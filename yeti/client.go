// Package yeti is the operations an agent uses on a Yeti server, shaped around
// tasks rather than API endpoints. The yeti command and its MCP server are
// thin front ends over this package.
package yeti

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/yeti-platform/goyeti/internal/api"
)

// Version is set at build time.
var Version = "dev"

// Config is how to reach Yeti.
type Config struct {
	// URL of the Yeti server, e.g. "https://yeti.example.com".
	URL string
	// APIKey of the Yeti user to act as.
	APIKey string
	// HTTPClient sends requests to Yeti. When set without APIKey, its
	// Transport is responsible for authentication; a *Error returned by
	// RoundTrip is passed through untouched.
	HTTPClient *http.Client
	// AuthHint is the recovery hint for 401 and 403 responses when set.
	AuthHint string
	// Timeout for each request to Yeti. Overrides HTTPClient.Timeout when
	// set; defaults to 30 seconds.
	Timeout time.Duration
	// AgentName identifies the agent in the User-Agent header, if set.
	AgentName string
}

// Client runs operations against one Yeti server. It is safe for concurrent
// use.
type Client struct {
	api *api.Client
	// agent names the agent in what it writes, e.g. context sources.
	agent     string
	hasAPIKey bool
	authHint  string
}

// New checks the configuration and returns a client. It doesn't contact Yeti.
func New(cfg Config) (*Client, error) {
	if cfg.URL == "" {
		return nil, &Error{Code: CodeConfig, Message: "No Yeti URL is configured.",
			Hint: "Set YETI_URL, e.g. https://yeti.example.com, or pass --url."}
	}
	parsed, err := url.Parse(cfg.URL)
	if err != nil || (parsed.Scheme != "http" && parsed.Scheme != "https") || parsed.Host == "" {
		return nil, &Error{Code: CodeConfig, Message: fmt.Sprintf("%q is not an http(s) URL.", cfg.URL),
			Hint: "Set YETI_URL to the server's address, e.g. https://yeti.example.com."}
	}
	if cfg.APIKey == "" && cfg.HTTPClient == nil {
		return nil, &Error{Code: CodeConfig, Message: "No Yeti API key is configured.",
			Hint: "Set YETI_API_KEY, or pass --api-key. Create a key in Yeti under your user's API keys."}
	}
	timeout := cfg.Timeout
	if timeout == 0 && cfg.HTTPClient != nil {
		timeout = cfg.HTTPClient.Timeout
	}
	if timeout == 0 {
		timeout = 30 * time.Second
	}
	userAgent := "goyeti/" + Version
	if cfg.AgentName != "" {
		userAgent += " (agent: " + cfg.AgentName + ")"
	}
	httpClient := &http.Client{Timeout: timeout}
	if cfg.HTTPClient != nil {
		c := *cfg.HTTPClient
		c.Timeout = timeout
		httpClient = &c
	}
	agent := cfg.AgentName
	if agent == "" {
		agent = "goyeti"
	}
	return &Client{
		api:       api.New(cfg.URL, cfg.APIKey, userAgent, httpClient),
		agent:     agent,
		hasAPIKey: cfg.APIKey != "",
		authHint:  cfg.AuthHint,
	}, nil
}

// Meta is what an operation's caller should know besides its data.
type Meta struct {
	// Truncated is set when results were left out.
	Truncated *Truncated `json:"truncated,omitempty"`
	// Warnings are problems that didn't stop the operation.
	Warnings []string `json:"warnings,omitempty"`
	// Failed counts the items of a batch that failed; each says why.
	Failed int `json:"failed,omitempty"`
}

// Truncated says how many results were left out and how to get them.
type Truncated struct {
	Omitted int    `json:"omitted"`
	Hint    string `json:"hint"`
}

// Error codes.
const (
	CodeUsage          = "usage"
	CodeConfig         = "config"
	CodeUnauthorized   = "unauthorized"
	CodeForbidden      = "forbidden"
	CodeNotFound       = "not_found"
	CodeInvalidRequest = "invalid_request"
	CodeRateLimited    = "rate_limited"
	CodeServerError    = "server_error"
	CodeUnreachable    = "unreachable"
	CodeTimeout        = "timeout"
)

// Error is a failed operation, worded for the agent that has to recover from
// it: what went wrong, what to do about it, and whether trying again can help.
type Error struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	Hint      string `json:"hint,omitempty"`
	Retryable bool   `json:"retryable"`
}

func (e *Error) Error() string {
	if e.Hint == "" {
		return e.Message
	}
	return e.Message + " " + e.Hint
}

func usage(message, hint string) *Error {
	return &Error{Code: CodeUsage, Message: message, Hint: hint}
}

// wrap turns an error from talking to Yeti into an *Error. notFoundHint is
// used for 404s.
func (c *Client) wrap(err error, notFoundHint string) error {
	if err == nil {
		return nil
	}
	var yetiErr *Error
	if errors.As(err, &yetiErr) {
		return yetiErr
	}
	if errors.Is(err, api.ErrUnauthorized) {
		return &Error{Code: CodeUnauthorized, Message: "Yeti refused the API key.",
			Hint: firstSet(c.authHint, "Check YETI_API_KEY: the key may be wrong, deleted, disabled or expired.")}
	}
	var status *api.StatusError
	if errors.As(err, &status) {
		switch {
		case status.Status == http.StatusUnauthorized:
			if c.hasAPIKey {
				return &Error{Code: CodeUnauthorized, Message: "Yeti refused the session for this API key.",
					Hint: firstSet(c.authHint, "Check YETI_API_KEY, and that its user is enabled.")}
			}
			return &Error{Code: CodeUnauthorized, Message: "Yeti refused the request's credentials.",
				Hint: c.authHint}
		case status.Status == http.StatusForbidden:
			defaultHint := "The authenticated Yeti user lacks the permission; ask a Yeti admin."
			if c.hasAPIKey {
				defaultHint = "The Yeti user behind this API key lacks the permission; ask a Yeti admin."
			}
			return &Error{Code: CodeForbidden, Message: status.Detail,
				Hint: firstSet(c.authHint, defaultHint)}
		case status.Status == http.StatusNotFound:
			return &Error{Code: CodeNotFound, Message: status.Detail, Hint: notFoundHint}
		case status.Status == http.StatusTooManyRequests:
			return &Error{Code: CodeRateLimited, Message: "Yeti is rate limiting requests.",
				Hint: "Wait, then try again.", Retryable: true}
		case status.Status >= 500:
			return &Error{Code: CodeServerError, Message: status.Error(),
				Hint: "Yeti failed to answer; try again, or check the Yeti server's logs.", Retryable: true}
		default:
			return &Error{Code: CodeInvalidRequest, Message: status.Detail}
		}
	}
	if errors.Is(err, context.DeadlineExceeded) || isTimeout(err) {
		return &Error{Code: CodeTimeout, Message: "Yeti didn't answer in time.",
			Hint: "Try again, ask for fewer results, or raise --timeout.", Retryable: true}
	}
	var netErr net.Error
	var opErr *net.OpError
	if errors.As(err, &opErr) || errors.As(err, &netErr) {
		return &Error{Code: CodeUnreachable, Message: "Can't reach Yeti: " + err.Error(),
			Hint: "Check YETI_URL and that the server is up.", Retryable: true}
	}
	var syntaxErr *json.SyntaxError
	if errors.As(err, &syntaxErr) {
		return &Error{Code: CodeServerError, Message: "The server's answer isn't the JSON Yeti sends: " + err.Error(),
			Hint: "Check that YETI_URL points at a Yeti server."}
	}
	return &Error{Code: CodeServerError, Message: err.Error()}
}

func isNotFound(err error) bool {
	var status *api.StatusError
	return errors.As(err, &status) && status.Status == http.StatusNotFound
}

func isTimeout(err error) bool {
	var netErr net.Error
	return errors.As(err, &netErr) && netErr.Timeout()
}

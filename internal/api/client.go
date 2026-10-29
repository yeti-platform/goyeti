// Package api talks HTTP to a Yeti server: authentication, requests and the
// endpoints the operations in package yeti use. Responses are decoded loosely,
// into the few fields the operations need, so new Yeti types don't break it.
package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
)

// Client is safe for concurrent use.
type Client struct {
	baseURL   string
	apiKey    string
	userAgent string
	http      *http.Client

	mu    sync.Mutex
	token string
}

// New returns a client for the Yeti server at baseURL, e.g.
// "https://yeti.example.com". The API key is exchanged for a session token on
// the first request.
func New(baseURL, apiKey, userAgent string, httpClient *http.Client) *Client {
	return &Client{
		baseURL:   strings.TrimRight(baseURL, "/"),
		apiKey:    apiKey,
		userAgent: userAgent,
		http:      httpClient,
	}
}

// StatusError is a response from Yeti with a status other than 2xx.
type StatusError struct {
	Status int
	Path   string
	// Detail is Yeti's explanation, from the response's "detail" field.
	Detail string
}

func (e *StatusError) Error() string {
	if e.Detail == "" {
		return fmt.Sprintf("%s returned HTTP %d", e.Path, e.Status)
	}
	return fmt.Sprintf("%s returned HTTP %d: %s", e.Path, e.Status, e.Detail)
}

// ErrUnauthorized means Yeti refused the API key itself.
var ErrUnauthorized = errors.New("Yeti refused the API key")

// Do sends a request to path (relative to /api/v2) and decodes the JSON
// response into out, unless out is nil. body, unless nil, is sent as JSON.
func (c *Client) Do(ctx context.Context, method, path string, body, out any) error {
	var token string
	if c.apiKey != "" {
		var err error
		if token, err = c.sessionToken(ctx, ""); err != nil {
			return err
		}
	}
	err := c.send(ctx, method, path, body, out, token)
	var status *StatusError
	if c.apiKey != "" && errors.As(err, &status) && status.Status == http.StatusUnauthorized {
		// The session token expired: exchange the key again, once.
		if token, err = c.sessionToken(ctx, token); err != nil {
			return err
		}
		err = c.send(ctx, method, path, body, out, token)
	}
	return err
}

// sessionToken returns the cached session token, exchanging the API key for a
// new one when there is none or when the cached one is stale.
func (c *Client) sessionToken(ctx context.Context, stale string) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.token != "" && c.token != stale {
		return c.token, nil
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/api/v2/auth/api-token", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("x-yeti-apikey", c.apiKey)
	var resp struct {
		AccessToken string `json:"access_token"`
	}
	if err := c.roundTrip(req, &resp); err != nil {
		var status *StatusError
		if errors.As(err, &status) && status.Status == http.StatusUnauthorized {
			return "", ErrUnauthorized
		}
		return "", err
	}
	if resp.AccessToken == "" {
		return "", errors.New("/auth/api-token returned no access token")
	}
	c.token = resp.AccessToken
	return c.token, nil
}

func (c *Client) send(ctx context.Context, method, path string, body, out any, token string) error {
	var reader io.Reader
	if body != nil {
		encoded, err := json.Marshal(body)
		if err != nil {
			return err
		}
		reader = bytes.NewReader(encoded)
	}
	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+"/api/v2"+path, reader)
	if err != nil {
		return err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	return c.roundTrip(req, out)
}

func (c *Client) roundTrip(req *http.Request, out any) error {
	req.Header.Set("Accept", "application/json")
	if c.userAgent != "" {
		req.Header.Set("User-Agent", c.userAgent)
	}
	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		path := strings.TrimPrefix(req.URL.Path, "/api/v2")
		return &StatusError{Status: resp.StatusCode, Path: path, Detail: detail(data)}
	}
	if out == nil {
		return nil
	}
	if err := json.Unmarshal(data, out); err != nil {
		return fmt.Errorf("decoding the response of %s: %w", req.URL.Path, err)
	}
	return nil
}

// detail extracts FastAPI's error explanation: a string, or a list of
// validation errors.
func detail(body []byte) string {
	var parsed struct {
		Detail json.RawMessage `json:"detail"`
	}
	if json.Unmarshal(body, &parsed) != nil || parsed.Detail == nil {
		return strings.TrimSpace(string(body))
	}
	var text string
	if json.Unmarshal(parsed.Detail, &text) == nil {
		return text
	}
	var validation []struct {
		Loc []any  `json:"loc"`
		Msg string `json:"msg"`
	}
	if json.Unmarshal(parsed.Detail, &validation) == nil && len(validation) > 0 {
		parts := make([]string, 0, len(validation))
		for _, v := range validation {
			loc := make([]string, 0, len(v.Loc))
			for _, l := range v.Loc {
				loc = append(loc, fmt.Sprint(l))
			}
			parts = append(parts, strings.Join(loc, ".")+": "+v.Msg)
		}
		return strings.Join(parts, "; ")
	}
	// Some errors are {"description": ..., "meta": {...}}.
	var described struct {
		Description string         `json:"description"`
		Meta        map[string]any `json:"meta"`
	}
	if json.Unmarshal(parsed.Detail, &described) == nil && described.Description != "" {
		if len(described.Meta) == 0 {
			return described.Description
		}
		meta, _ := json.Marshal(described.Meta)
		return described.Description + " " + string(meta)
	}
	return string(parsed.Detail)
}

package main

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/yeti-platform/goyeti/yeti"
)

// Exit codes.
const (
	exitOK     = 0
	exitFailed = 1
	exitUsage  = 2
)

// envelope is every command's output: data and what was left out, or an
// error.
type envelope struct {
	Data      any             `json:"data,omitempty"`
	Truncated *yeti.Truncated `json:"truncated,omitempty"`
	Warnings  []string        `json:"warnings,omitempty"`
	Failed    int             `json:"failed,omitempty"`
	Error     *yeti.Error     `json:"error,omitempty"`
}

func success(data any, meta yeti.Meta) envelope {
	return envelope{Data: data, Truncated: meta.Truncated, Warnings: meta.Warnings, Failed: meta.Failed}
}

// failure turns any error into an error envelope, and the exit code to use.
func failure(err error) (envelope, int) {
	var yetiErr *yeti.Error
	if !errors.As(err, &yetiErr) {
		yetiErr = &yeti.Error{Code: yeti.CodeUsage, Message: err.Error(), Hint: `Run "yeti --help" for usage.`}
	}
	if yetiErr.Code == yeti.CodeUsage {
		return envelope{Error: yetiErr}, exitUsage
	}
	return envelope{Error: yetiErr}, exitFailed
}

func write(w io.Writer, e envelope, pretty bool) {
	encoder := json.NewEncoder(w)
	// Output goes to terminals and agents, not into HTML: keep "->" readable.
	encoder.SetEscapeHTML(false)
	if pretty {
		encoder.SetIndent("", "  ")
	}
	encoder.Encode(e)
}

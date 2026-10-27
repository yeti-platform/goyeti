package main

import (
	"encoding/json"
	"errors"
	"fmt"
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
	Error     *yeti.Error     `json:"error,omitempty"`
}

func success(data any, meta yeti.Meta) envelope {
	return envelope{Data: data, Truncated: meta.Truncated, Warnings: meta.Warnings}
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
	var data []byte
	if pretty {
		data, _ = json.MarshalIndent(e, "", "  ")
	} else {
		data, _ = json.Marshal(e)
	}
	fmt.Fprintln(w, string(data))
}

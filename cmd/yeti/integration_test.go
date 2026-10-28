//go:build integration

package main

import (
	"context"
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/yeti-platform/goyeti/yeti"
)

// The operations are tested end to end in package yeti; these check that the
// command line and the MCP server reach a real Yeti.

func liveEnv(t *testing.T) map[string]string {
	t.Helper()
	env := map[string]string{"YETI_URL": os.Getenv("YETI_URL"), "YETI_API_KEY": os.Getenv("YETI_API_KEY")}
	if env["YETI_URL"] == "" || env["YETI_API_KEY"] == "" {
		t.Skip("set YETI_URL and YETI_API_KEY to run against a real Yeti")
	}
	return env
}

func TestLiveCommandLine(t *testing.T) {
	env := liveEnv(t)
	out, code := runYeti(t, env, "unknown.goyeti.invalid\n", "read", "lookup", "-")
	if code != exitOK {
		t.Fatalf("exit %d: %s", code, out)
	}
	results := decode(t, out).Data.([]any)
	if len(results) != 1 || results[0].(map[string]any)["known"] != false {
		t.Errorf("lookup: %s", out)
	}

	out, code = runYeti(t, env, "", "read", "get", "observables/goyeti-does-not-exist")
	if e := decode(t, out).Error; code != exitFailed || e == nil || e.Code != yeti.CodeNotFound {
		t.Errorf("get of a missing object, exit %d: %s", code, out)
	}

	env["YETI_API_KEY"] = "not-a-key"
	out, code = runYeti(t, env, "", "read", "types")
	if e := decode(t, out).Error; code != exitFailed || e == nil || e.Code != yeti.CodeUnauthorized {
		t.Errorf("a wrong key, exit %d: %s", code, out)
	}
}

func TestLiveMCP(t *testing.T) {
	env := liveEnv(t)
	client, err := yeti.New(yeti.Config{URL: env["YETI_URL"], APIKey: env["YETI_API_KEY"]})
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	clientTransport, serverTransport := mcp.NewInMemoryTransports()
	if _, err := newMCPServer(client, "read").Connect(ctx, serverTransport, nil); err != nil {
		t.Fatal(err)
	}
	session, err := mcp.NewClient(&mcp.Implementation{Name: "e2e"}, nil).Connect(ctx, clientTransport, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer session.Close()

	result, err := session.CallTool(ctx, &mcp.CallToolParams{Name: "types"})
	if err != nil || result.IsError {
		t.Fatalf("types: %v %+v", err, result)
	}
	structured, _ := json.Marshal(result.StructuredContent)
	if !strings.Contains(string(structured), `"ipv4"`) {
		t.Errorf("types returned %s", structured)
	}
}

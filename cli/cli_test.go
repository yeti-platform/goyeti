package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/yeti-platform/goyeti/yeti"
)

// fixtures maps requests to recorded Yeti responses.
var fixtures = map[string]string{
	"POST /api/v2/graph/match":            "match.json",
	"GET /api/v2/observables/22898861":    "get_observable.json",
	"POST /api/v2/indicators/yara/bundle": "yara_bundle.json",
	"GET /api/v2/system/types":            "types.json",
}

func fakeYeti(t *testing.T) string {
	t.Helper()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v2/auth/api-token" {
			w.Write([]byte(`{"access_token":"token"}`))
			return
		}
		fixture, ok := fixtures[r.Method+" "+r.URL.Path]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"detail":"Observable not found"}`))
			return
		}
		data, err := os.ReadFile("../internal/api/testdata/" + fixture)
		if err != nil {
			t.Error(err)
		}
		w.Write(data)
	}))
	t.Cleanup(server.Close)
	return server.URL
}

// runYeti runs the command line with stdin, and returns its output and exit
// code.
func runYeti(t *testing.T, env map[string]string, stdin string, args ...string) (string, int) {
	t.Helper()
	var out bytes.Buffer
	a := &app{stdin: strings.NewReader(stdin), stdout: &out, getenv: func(k string) string { return env[k] }}
	code := a.run(context.Background(), args)
	return out.String(), code
}

func decode(t *testing.T, out string) envelope {
	t.Helper()
	var e envelope
	if err := json.Unmarshal([]byte(out), &e); err != nil {
		t.Fatalf("output is not one JSON envelope: %v\n%s", err, out)
	}
	return e
}

func TestLookupReadsValuesFromStdin(t *testing.T) {
	env := map[string]string{"YETI_URL": fakeYeti(t), "YETI_API_KEY": "key"}
	out, code := runYeti(t, env, "203.0.113.7\n nothing.example.org\n", "read", "lookup")
	if code != exitOK {
		t.Fatalf("exit %d: %s", code, out)
	}
	if data := decode(t, out).Data.([]any); len(data) != 2 {
		t.Errorf("got %d results: %s", len(data), out)
	}
}

func TestExitCodes(t *testing.T) {
	env := map[string]string{"YETI_URL": fakeYeti(t), "YETI_API_KEY": "key"}
	for _, tc := range []struct {
		args []string
		env  map[string]string
		code int
		err  string
	}{
		{[]string{"read", "types"}, env, exitOK, ""},
		{[]string{"read", "get", "observables/404"}, env, exitFailed, yeti.CodeNotFound},
		{[]string{"read", "types"}, map[string]string{}, exitFailed, yeti.CodeConfig},
		{[]string{"read", "get", "not-a-reference"}, env, exitUsage, yeti.CodeUsage},
		{[]string{"read", "get"}, env, exitUsage, yeti.CodeUsage},
		{[]string{"read", "bogus"}, env, exitUsage, yeti.CodeUsage},
		{[]string{"mutate", "add", "bogus"}, env, exitUsage, yeti.CodeUsage},
		{[]string{"admin", "delete", "observables/22898861"}, env, exitUsage, yeti.CodeUsage},
		{[]string{"admin", "bogus"}, env, exitUsage, yeti.CodeUsage},
		{[]string{"admin", "delete"}, env, exitUsage, yeti.CodeUsage},
		{[]string{"mutate", "add-context", "observables/1", "no-equals-sign"}, env, exitUsage, yeti.CodeUsage},
		{[]string{"mcp", "--tier", "root"}, env, exitUsage, yeti.CodeUsage},
	} {
		t.Run(strings.Join(tc.args, " "), func(t *testing.T) {
			out, code := runYeti(t, tc.env, "", tc.args...)
			if code != tc.code {
				t.Errorf("exit %d, want %d: %s", code, tc.code, out)
			}
			e := decode(t, out)
			if tc.err == "" && e.Error != nil {
				t.Errorf("unexpected error %+v", e.Error)
			}
			if tc.err != "" && (e.Error == nil || e.Error.Code != tc.err) {
				t.Errorf("got %s, want error code %q", out, tc.err)
			}
		})
	}
}

func TestYaraBundleRaw(t *testing.T) {
	env := map[string]string{"YETI_URL": fakeYeti(t), "YETI_API_KEY": "key"}
	out, code := runYeti(t, env, "", "read", "yara-bundle", "indicators/22899145", "--raw")
	if code != exitOK || !strings.HasPrefix(out, "private rule emotet_strings") {
		t.Errorf("exit %d: %q", code, out)
	}
}

func TestNewClientOverride(t *testing.T) {
	url := fakeYeti(t)
	var out bytes.Buffer
	a := &app{
		stdin:  strings.NewReader(""),
		stdout: &out,
		getenv: func(string) string { return "" },
		newClient: func(ctx context.Context, cfg yeti.Config) (*yeti.Client, error) {
			if cfg.Timeout != 0 {
				t.Errorf("cfg.Timeout = %v when --timeout is omitted, want 0", cfg.Timeout)
			}
			cfg.URL = url
			cfg.HTTPClient = &http.Client{}
			return yeti.New(cfg)
		},
	}
	if code := a.run(context.Background(), []string{"read", "types"}); code != exitOK {
		t.Fatalf("exit %d: %s", code, out.String())
	}

	out.Reset()
	a.newClient = func(ctx context.Context, cfg yeti.Config) (*yeti.Client, error) {
		return nil, errors.New("sso credentials expired")
	}
	code := a.run(context.Background(), []string{"read", "types"})
	if e := decode(t, out.String()); code != exitFailed || e.Error == nil || e.Error.Code != yeti.CodeConfig {
		t.Errorf("plain error from NewClient: exit %d, envelope %+v", code, e)
	}
}

func TestMCPServesTheReadTier(t *testing.T) {
	client, err := yeti.New(yeti.Config{URL: fakeYeti(t), APIKey: "key"})
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	clientTransport, serverTransport := mcp.NewInMemoryTransports()
	if _, err := newMCPServer(client, "read").Connect(ctx, serverTransport, nil); err != nil {
		t.Fatal(err)
	}
	session, err := mcp.NewClient(&mcp.Implementation{Name: "test"}, nil).Connect(ctx, clientTransport, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer session.Close()

	tools, err := session.ListTools(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	var names []string
	for _, tool := range tools.Tools {
		names = append(names, tool.Name)
		if !tool.Annotations.ReadOnlyHint || tool.Description == "" {
			t.Errorf("tool %s: read-only %v, description %q", tool.Name, tool.Annotations.ReadOnlyHint, tool.Description)
		}
	}
	if want := "get lookup neighbors search tasks types yara_bundle"; strings.Join(names, " ") != want {
		t.Errorf("tools %v, want %s", names, want)
	}

	result, err := session.CallTool(ctx, &mcp.CallToolParams{Name: "lookup", Arguments: map[string]any{"values": []string{"203.0.113.7"}}})
	if err != nil || result.IsError {
		t.Fatalf("lookup: %v %+v", err, result)
	}
	structured, _ := json.Marshal(result.StructuredContent)
	if !strings.Contains(string(structured), `"ref":"observables/22898861"`) {
		t.Errorf("lookup returned %s", structured)
	}

	result, err = session.CallTool(ctx, &mcp.CallToolParams{Name: "get", Arguments: map[string]any{"ref": "203.0.113.7"}})
	if err != nil || !result.IsError {
		t.Fatalf("get with a bad reference should be a tool error: %v %+v", err, result)
	}
	text := result.Content[0].(*mcp.TextContent).Text
	if e := decode(t, text); e.Error == nil || e.Error.Code != yeti.CodeUsage || e.Error.Hint == "" {
		t.Errorf("error result %s", text)
	}
}

func TestFailedItemsExitOne(t *testing.T) {
	env := map[string]string{"YETI_URL": fakeYeti(t), "YETI_API_KEY": "key"}
	out, code := runYeti(t, env, "", "mutate", "tag", "dfiq/1", "--tag", "x")
	e := decode(t, out)
	if code != exitFailed || e.Failed != 1 || e.Data == nil {
		t.Errorf("exit %d: %s", code, out)
	}
}

func TestParseFields(t *testing.T) {
	fields, err := parseFields([]string{"note=seen in campaign X", "ports=[80,443]", "count=3", "empty=", "url=http://a.example/?q=1"})
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]any{"note": "seen in campaign X", "ports": []any{float64(80), float64(443)}, "count": float64(3),
		"empty": "", "url": "http://a.example/?q=1"}
	if !reflect.DeepEqual(fields, want) {
		t.Errorf("got %#v", fields)
	}
	if _, err := parseFields([]string{"=value"}); err == nil {
		t.Error("a field without a name should be refused")
	}
}

// mcpSession connects an MCP client to a server of a tier, over a fake Yeti.
func mcpSession(t *testing.T, tier string) *mcp.ClientSession {
	t.Helper()
	client, err := yeti.New(yeti.Config{URL: fakeYeti(t), APIKey: "key"})
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	clientTransport, serverTransport := mcp.NewInMemoryTransports()
	if _, err := newMCPServer(client, tier).Connect(ctx, serverTransport, nil); err != nil {
		t.Fatal(err)
	}
	session, err := mcp.NewClient(&mcp.Implementation{Name: "test"}, nil).Connect(ctx, clientTransport, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { session.Close() })
	return session
}

// writingTools maps the tools that aren't read-only to whether they're
// destructive.
func writingTools(t *testing.T, session *mcp.ClientSession) (map[string]bool, int) {
	t.Helper()
	tools, err := session.ListTools(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	writing := map[string]bool{}
	for _, tool := range tools.Tools {
		if !tool.Annotations.ReadOnlyHint {
			writing[tool.Name] = *tool.Annotations.DestructiveHint
		}
	}
	return writing, len(tools.Tools)
}

var mutateTools = map[string]bool{"add_observables": false, "add_entity": false, "add_indicator": false, "link": false,
	"tag": false, "untag": true, "add_context": false}

func TestMCPMutateTier(t *testing.T) {
	writing, count := writingTools(t, mcpSession(t, "mutate"))
	if count != 14 || !reflect.DeepEqual(writing, mutateTools) {
		t.Errorf("%d tools; writing tools and whether they're destructive: %v", count, writing)
	}
}

func TestMCPAdminTier(t *testing.T) {
	session := mcpSession(t, "admin")
	writing, count := writingTools(t, session)
	want := map[string]bool{"delete": true, "unlink": true, "run_task": false, "enable_task": false, "disable_task": false}
	for k, v := range mutateTools {
		want[k] = v
	}
	if count != 19 || !reflect.DeepEqual(writing, want) {
		t.Errorf("%d tools; writing tools and whether they're destructive: %v", count, writing)
	}

	result, err := session.CallTool(context.Background(), &mcp.CallToolParams{Name: "delete",
		Arguments: map[string]any{"refs": []string{"observables/22898861"}}})
	if err != nil || !result.IsError {
		t.Fatalf("delete without confirm or dry_run should be a tool error: %v %+v", err, result)
	}
	if e := decode(t, result.Content[0].(*mcp.TextContent).Text); e.Error == nil || e.Error.Code != yeti.CodeUsage {
		t.Errorf("delete without confirm returned %+v", e)
	}
}

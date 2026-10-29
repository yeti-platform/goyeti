package api

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync/atomic"
	"testing"
)

// fakeYeti issues session tokens for the API key "key" and serves handler
// behind them. expire makes it reject the current token once.
type fakeYeti struct {
	exchanges atomic.Int32
	expire    atomic.Bool
	handler   http.HandlerFunc
}

func (f *fakeYeti) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/api/v2/auth/api-token" {
		if r.Header.Get("x-yeti-apikey") != "key" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"detail":"Could not validate credentials"}`))
			return
		}
		n := f.exchanges.Add(1)
		w.Write([]byte(`{"access_token":"token-` + string(rune('0'+n)) + `","token_type":"bearer"}`))
		return
	}
	if !strings.HasPrefix(r.Header.Get("Authorization"), "Bearer token-") || f.expire.CompareAndSwap(true, false) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"detail":"Could not validate credentials"}`))
		return
	}
	f.handler(w, r)
}

func serve(t *testing.T, fixture string) (*Client, *fakeYeti) {
	t.Helper()
	data, err := os.ReadFile("testdata/" + fixture)
	if err != nil {
		t.Fatal(err)
	}
	fake := &fakeYeti{handler: func(w http.ResponseWriter, r *http.Request) { w.Write(data) }}
	server := httptest.NewServer(fake)
	t.Cleanup(server.Close)
	return New(server.URL, "key", "goyeti-test", server.Client()), fake
}

func TestTokenIsExchangedOnceAndReused(t *testing.T) {
	client, fake := serve(t, "types.json")
	for range 3 {
		if _, err := client.Types(context.Background()); err != nil {
			t.Fatal(err)
		}
	}
	if got := fake.exchanges.Load(); got != 1 {
		t.Errorf("API key exchanged %d times, want 1", got)
	}
}

func TestExpiredTokenIsExchangedAgain(t *testing.T) {
	client, fake := serve(t, "types.json")
	if _, err := client.Types(context.Background()); err != nil {
		t.Fatal(err)
	}
	fake.expire.Store(true)
	if _, err := client.Types(context.Background()); err != nil {
		t.Fatalf("request after expiry: %v", err)
	}
	if got := fake.exchanges.Load(); got != 2 {
		t.Errorf("API key exchanged %d times, want 2", got)
	}
}

func TestRejectedKey(t *testing.T) {
	client, _ := serve(t, "types.json")
	client.apiKey = "wrong"
	_, err := client.Types(context.Background())
	if !errors.Is(err, ErrUnauthorized) {
		t.Fatalf("got %v, want ErrUnauthorized", err)
	}
}

func TestCustomHTTPClientWithoutAPIKey(t *testing.T) {
	data, err := os.ReadFile("testdata/types.json")
	if err != nil {
		t.Fatal(err)
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v2/auth/api-token" {
			t.Errorf("unexpected call to %s", r.URL.Path)
		}
		if got := r.Header.Get("Authorization"); got != "" {
			t.Errorf("Authorization = %q, want empty", got)
		}
		w.Write(data)
	}))
	defer server.Close()

	client := New(server.URL, "", "goyeti-test", server.Client())
	if _, err := client.Types(context.Background()); err != nil {
		t.Fatal(err)
	}
}

func TestStatusErrorCarriesDetail(t *testing.T) {
	fake := &fakeYeti{handler: func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(`{"detail":[{"loc":["body","source"],"msg":"Field required"}]}`))
	}}
	server := httptest.NewServer(fake)
	defer server.Close()
	_, err := New(server.URL, "key", "", server.Client()).Get(context.Background(), "observables", "1")
	var status *StatusError
	if !errors.As(err, &status) {
		t.Fatalf("got %v, want a StatusError", err)
	}
	if status.Status != 422 || status.Detail != "body.source: Field required" || status.Path != "/observables/1" {
		t.Errorf("got %+v", status)
	}
}

func TestMatchDecodesPairs(t *testing.T) {
	client, _ := serve(t, "match.json")
	resp, err := client.Match(context.Background(), []string{"203.0.113.7"})
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Known) != 2 || len(resp.Entities) != 1 || len(resp.Observables) != 1 || len(resp.Matches) != 2 {
		t.Fatalf("got %d known, %d entities, %d observables, %d matches",
			len(resp.Known), len(resp.Entities), len(resp.Observables), len(resp.Matches))
	}
	if e := resp.Entities[0]; e.Edge.Type != "c2" || e.Object["name"] != "Emotet" {
		t.Errorf("entity pair decoded as %+v", e)
	}
	if m := resp.Matches[0]; m.Value == "" || m.Indicator["name"] != "evil_subdomains" {
		t.Errorf("match pair decoded as %+v", m)
	}
}

func TestGraphSearchDecodesPaths(t *testing.T) {
	client, _ := serve(t, "graph_search.json")
	resp, err := client.GraphSearch(context.Background(), GraphSearchRequest{Source: "entities/22899024"})
	if err != nil {
		t.Fatal(err)
	}
	if resp.Total != 3 || len(resp.Paths) != 3 || len(resp.Vertices) == 0 {
		t.Fatalf("got total %d, %d paths, %d vertices", resp.Total, len(resp.Paths), len(resp.Vertices))
	}
}

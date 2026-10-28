package yeti

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

// route is a recorded response for one request.
type route struct {
	status  int
	fixture string
	// then are the fixtures for the following calls, the last one repeating.
	then []string
}

// fakeYeti serves recorded Yeti responses, and keeps the request bodies.
type fakeYeti struct {
	routes map[string]route // "METHOD /path" -> response
	bodies map[string]map[string]any
	calls  map[string]int
}

func (f *fakeYeti) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/api/v2/auth/api-token" {
		if r.Header.Get("x-yeti-apikey") != "key" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.Write([]byte(`{"access_token":"token"}`))
		return
	}
	key := r.Method + " " + r.URL.Path
	if body, _ := io.ReadAll(r.Body); len(body) > 0 {
		var decoded map[string]any
		json.Unmarshal(body, &decoded)
		f.bodies[key] = decoded
	}
	rt, ok := f.routes[key]
	if withQuery, found := f.routes[key+"?"+r.URL.RawQuery]; found && r.URL.RawQuery != "" {
		rt, ok = withQuery, true
	}
	if f.calls[key]++; f.calls[key] > 1 && len(rt.then) > 0 {
		rt.fixture = rt.then[min(f.calls[key]-2, len(rt.then)-1)]
	}
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"detail":"Not Found"}`))
		return
	}
	if rt.status != 0 {
		w.WriteHeader(rt.status)
	}
	if rt.fixture != "" {
		data, err := os.ReadFile("../internal/api/testdata/" + rt.fixture)
		if err != nil {
			panic(err)
		}
		w.Write(data)
	}
}

var recorded = map[string]route{
	"POST /api/v2/graph/match":            {fixture: "match.json"},
	"POST /api/v2/search/":                {fixture: "search.json"},
	"POST /api/v2/search/semantic":        {fixture: "semantic_results.json"},
	"GET /api/v2/observables/22898861":    {fixture: "get_observable.json"},
	"GET /api/v2/observables/404":         {status: 404, fixture: "get_missing.json"},
	"POST /api/v2/graph/search":           {fixture: "graph_search.json"},
	"POST /api/v2/indicators/yara/bundle": {fixture: "yara_bundle.json"},
	"GET /api/v2/system/types":            {fixture: "types.json"},
}

func newTestClient(t *testing.T, overrides map[string]route) (*Client, *fakeYeti) {
	t.Helper()
	routes := map[string]route{}
	for k, v := range recorded {
		routes[k] = v
	}
	for k, v := range overrides {
		routes[k] = v
	}
	fake := &fakeYeti{routes: routes, bodies: map[string]map[string]any{}, calls: map[string]int{}}
	server := httptest.NewServer(fake)
	t.Cleanup(server.Close)
	client, err := New(Config{URL: server.URL, APIKey: "key"})
	if err != nil {
		t.Fatal(err)
	}
	return client, fake
}

func wantCode(t *testing.T, err error, code string) *Error {
	t.Helper()
	var yetiErr *Error
	if !errors.As(err, &yetiErr) || yetiErr.Code != code {
		t.Fatalf("got error %v, want code %q", err, code)
	}
	return yetiErr
}

func TestLookup(t *testing.T) {
	client, fake := newTestClient(t, nil)
	results, _, err := client.Lookup(context.Background(), []string{
		"203.0.113.7", " evil[.]example[.]com ", "sub.evil.example.com", "nothing.example.org", "203.0.113.7",
	}, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(results) != 4 {
		t.Fatalf("got %d results, want 4 (duplicates dropped)", len(results))
	}

	ip := results[0]
	if !ip.Known || ip.Ref != "observables/22898861" || ip.Type != "ipv4" {
		t.Errorf("ip: %+v", ip)
	}
	if !reflect.DeepEqual(ip.Tags, []string{"c2", "emotet"}) || !reflect.DeepEqual(ip.ContextSources, []string{"FeedA"}) {
		t.Errorf("ip tags %v, context sources %v", ip.Tags, ip.ContextSources)
	}
	if len(ip.Entities) != 1 || ip.Entities[0].Name != "Emotet" || ip.Entities[0].Link != "c2" || ip.Entities[0].Direction != "out" {
		t.Errorf("ip entities: %+v", ip.Entities)
	}

	host := results[1]
	if host.Value != "evil[.]example[.]com" || host.Normalized != "evil.example.com" || !host.Known {
		t.Errorf("defanged host: %+v", host)
	}
	if len(host.Observables) != 1 || host.Observables[0].Link != "hosted-on" || host.Observables[0].Direction != "in" {
		t.Errorf("host observables: %+v", host.Observables)
	}
	if len(host.Indicators) != 1 || host.Indicators[0].Name != "evil_subdomains" {
		t.Errorf("host indicators: %+v", host.Indicators)
	}

	if sub := results[2]; sub.Known || len(sub.Indicators) != 1 {
		t.Errorf("unknown value matching an indicator: %+v", sub)
	}
	if none := results[3]; none.Known || none.Indicators != nil || none.Entities != nil {
		t.Errorf("unknown value: %+v", none)
	}

	sent := fake.bodies["POST /api/v2/graph/match"]
	if sent["add_unknown"] != false {
		t.Errorf("lookup must never add unknown values, sent %v", sent)
	}
	if got := sent["observables"].([]any); len(got) != 4 || got[1] != "evil.example.com" {
		t.Errorf("sent values %v", got)
	}
}

func TestLookupNeedsValues(t *testing.T) {
	client, _ := newTestClient(t, nil)
	_, _, err := client.Lookup(context.Background(), []string{" ", ""}, 0)
	wantCode(t, err, CodeUsage)
}

func TestCapListCountsOmitted(t *testing.T) {
	r := &LookupResult{}
	kept := capList(r, "entities", []Linked{{}, {}, {}}, 2)
	if len(kept) != 2 || r.Omitted["entities"] != 1 {
		t.Errorf("kept %d, omitted %v", len(kept), r.Omitted)
	}
}

func TestSearchMergesSemanticAndKeyword(t *testing.T) {
	client, _ := newTestClient(t, nil)
	result, meta, err := client.Search(context.Background(), "emotet", 0)
	if err != nil {
		t.Fatal(err)
	}
	first := result.Hits[0]
	if first.Name != "Emotet" || first.Score != 0.82 || !reflect.DeepEqual(first.FoundBy, []string{"semantic", "keyword"}) {
		t.Errorf("first hit: %+v", first)
	}
	if result.Hits[1].Name != "TA542" {
		t.Errorf("second hit should be the next best semantic match: %+v", result.Hits[1])
	}
	if meta.Truncated != nil || meta.Warnings != nil {
		t.Errorf("meta: %+v", meta)
	}
}

func TestSearchTruncates(t *testing.T) {
	client, _ := newTestClient(t, nil)
	result, meta, err := client.Search(context.Background(), "emotet", 2)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Hits) != 2 || meta.Truncated == nil || meta.Truncated.Omitted == 0 {
		t.Errorf("got %d hits, truncated %+v", len(result.Hits), meta.Truncated)
	}
}

func TestSearchWithoutSemantic(t *testing.T) {
	client, _ := newTestClient(t, map[string]route{"POST /api/v2/search/semantic": {status: 503}})
	result, meta, err := client.Search(context.Background(), "emotet", 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(meta.Warnings) != 1 || len(result.Hits) == 0 {
		t.Errorf("got %d hits, warnings %v", len(result.Hits), meta.Warnings)
	}
	for _, hit := range result.Hits {
		if !reflect.DeepEqual(hit.FoundBy, []string{"keyword"}) {
			t.Errorf("hit %s found by %v", hit.Ref, hit.FoundBy)
		}
	}
}

func TestGet(t *testing.T) {
	client, _ := newTestClient(t, nil)
	obj, _, err := client.Get(context.Background(), "observables/22898861")
	if err != nil {
		t.Fatal(err)
	}
	if obj["ref"] != "observables/22898861" || obj["value"] != "203.0.113.7" {
		t.Errorf("got %v", obj)
	}
	if _, ok := obj["acls"]; ok {
		t.Error("acls should be left out")
	}

	_, _, err = client.Get(context.Background(), "observables/404")
	if e := wantCode(t, err, CodeNotFound); e.Hint == "" || e.Retryable {
		t.Errorf("not found: %+v", e)
	}
	for _, bad := range []string{"203.0.113.7", "observable/1", "observables/", "observables/1/2"} {
		_, _, err = client.Get(context.Background(), bad)
		wantCode(t, err, CodeUsage)
	}
}

func TestNeighbors(t *testing.T) {
	client, fake := newTestClient(t, nil)
	result, meta, err := client.Neighbors(context.Background(), "entities/22899024", NeighborsOptions{Direction: "in"})
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Neighbors) != 3 || meta.Truncated != nil {
		t.Fatalf("got %+v, %+v", result.Neighbors, meta)
	}
	for _, n := range result.Neighbors {
		if n.Direction != "in" || n.Via != nil {
			t.Errorf("neighbor %+v", n)
		}
	}
	if n := result.Neighbors[len(result.Neighbors)-1]; n.Name != "TA542" || n.Link != "uses" {
		t.Errorf("neighbors are sorted by link type: last is %+v", n)
	}
	sent := fake.bodies["POST /api/v2/graph/search"]
	if sent["direction"] != "inbound" || sent["max_hops"] != float64(1) || sent["include_original"] != false {
		t.Errorf("sent %v", sent)
	}

	_, _, err = client.Neighbors(context.Background(), "entities/1", NeighborsOptions{Direction: "sideways"})
	wantCode(t, err, CodeUsage)
	_, _, err = client.Neighbors(context.Background(), "entities/1", NeighborsOptions{Hops: 4})
	wantCode(t, err, CodeUsage)
}

func TestYaraBundle(t *testing.T) {
	client, fake := newTestClient(t, nil)
	bundle, meta, err := client.YaraBundle(context.Background(), []string{"indicators/22899145"}, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	if bundle.Rules != 2 || meta.Warnings != nil {
		t.Errorf("got %d rules, warnings %v", bundle.Rules, meta.Warnings)
	}
	if ids := fake.bodies["POST /api/v2/indicators/yara/bundle"]["ids"].([]any); len(ids) != 1 || ids[0] != "22899145" {
		t.Errorf("sent ids %v", ids)
	}

	_, _, err = client.YaraBundle(context.Background(), []string{"observables/1"}, nil, nil)
	wantCode(t, err, CodeUsage)
	_, _, err = client.YaraBundle(context.Background(), nil, nil, nil)
	wantCode(t, err, CodeUsage)
}

func TestTypes(t *testing.T) {
	client, _ := newTestClient(t, nil)
	types, _, err := client.Types(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(types["observables"]) != 33 || types["observables"][0] != "asn" {
		t.Errorf("got %v", types)
	}
}

func TestErrors(t *testing.T) {
	_, err := New(Config{APIKey: "key"})
	wantCode(t, err, CodeConfig)
	_, err = New(Config{URL: "yeti.example.com", APIKey: "key"})
	wantCode(t, err, CodeConfig)
	_, err = New(Config{URL: "https://yeti.example.com"})
	wantCode(t, err, CodeConfig)

	wrong, err := New(Config{URL: "http://127.0.0.1:1", APIKey: "key"})
	if err != nil {
		t.Fatal(err)
	}
	_, _, err = wrong.Types(context.Background())
	if e := wantCode(t, err, CodeUnreachable); !e.Retryable {
		t.Errorf("unreachable should be retryable: %+v", e)
	}

	server := httptest.NewServer(&fakeYeti{routes: recorded, bodies: map[string]map[string]any{}, calls: map[string]int{}})
	defer server.Close()
	badKey, _ := New(Config{URL: server.URL, APIKey: "wrong"})
	_, _, err = badKey.Types(context.Background())
	wantCode(t, err, CodeUnauthorized)

	failing, _ := newTestClient(t, map[string]route{"GET /api/v2/system/types": {status: 500}})
	_, _, err = failing.Types(context.Background())
	if e := wantCode(t, err, CodeServerError); !e.Retryable {
		t.Errorf("5xx should be retryable: %+v", e)
	}
}

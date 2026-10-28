//go:build integration

// End-to-end tests against a real Yeti, e.g. the yeti-docker production stack
// in CI or the dev stack locally:
//
//	YETI_URL=http://localhost YETI_API_KEY=... go test -tags integration ./...
//
// They seed a small investigation, with names unique to the run so that they
// can run again against the same server, and check every read operation
// against it. The API key's user needs write access, for the seeding.
package yeti

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"slices"
	"strings"
	"testing"
	"time"
)

func liveClient(t *testing.T) *Client {
	t.Helper()
	if os.Getenv("YETI_URL") == "" || os.Getenv("YETI_API_KEY") == "" {
		t.Skip("set YETI_URL and YETI_API_KEY to run against a real Yeti")
	}
	client, err := New(Config{URL: os.Getenv("YETI_URL"), APIKey: os.Getenv("YETI_API_KEY"), AgentName: "goyeti-e2e"})
	if err != nil {
		t.Fatal(err)
	}
	return client
}

// scenario is an IP and a hostname used by a malware, the actor behind it, a
// URL on the hostname, two YARA rules where one depends on the other, and a
// regex indicator for the hostname's subdomains.
type scenario struct {
	run                       string
	ip, host, url, subdomain  string
	ipRef, hostRef, urlRef    string
	malware, actor            string
	malwareRef, actorRef      string
	tag, baseRule, loaderRule string
	loaderRef                 string
}

func seed(t *testing.T, c *Client) *scenario {
	t.Helper()
	ctx := context.Background()
	now := time.Now().UnixNano()
	run := fmt.Sprintf("%d", now%1_000_000_000)
	s := &scenario{
		run:        run,
		ip:         fmt.Sprintf("10.%d.%d.%d", now%250+1, now/250%250+1, now/62500%250+1),
		host:       "e2e-" + run + ".example.com",
		malware:    "GoyetiMalware" + run,
		actor:      "GoyetiActor" + run,
		tag:        "goyeti-e2e-" + run,
		baseRule:   "goyeti_strings_" + run,
		loaderRule: "goyeti_loader_" + run,
	}
	s.url = "http://" + s.host + "/payload.bin"
	s.subdomain = "sub." + s.host

	post := func(path string, body map[string]any) map[string]any {
		t.Helper()
		var out map[string]any
		if err := c.api.Do(ctx, http.MethodPost, path, body, &out); err != nil {
			t.Fatalf("seeding %s: %v", path, err)
		}
		return out
	}
	newObservable := func(value, typ string) string {
		return "observables/" + str(post("/observables/", map[string]any{"value": value, "type": typ})["id"])
	}
	s.ipRef = newObservable(s.ip, "ipv4")
	s.hostRef = newObservable(s.host, "hostname")
	s.urlRef = newObservable(s.url, "url")
	post("/observables/tag", map[string]any{"ids": []string{refID(s.ipRef)}, "tags": []string{s.tag}})
	post("/observables/"+refID(s.ipRef)+"/context", map[string]any{"source": "GoyetiE2E", "context": map[string]any{"run": run}})

	s.malwareRef = "entities/" + str(post("/entities/", map[string]any{"entity": map[string]any{
		"type": "malware", "name": s.malware, "description": "Loader seeded by goyeti's end-to-end tests."}})["id"])
	s.actorRef = "entities/" + str(post("/entities/", map[string]any{"entity": map[string]any{
		"type": "threat-actor", "name": s.actor}})["id"])
	for _, link := range [][3]string{
		{s.ipRef, s.malwareRef, "c2"},
		{s.hostRef, s.malwareRef, "delivers"},
		{s.urlRef, s.hostRef, "hosted-on"},
		{s.actorRef, s.malwareRef, "uses"},
	} {
		post("/graph/add", map[string]any{"source": link[0], "target": link[1], "link_type": link[2], "description": ""})
	}

	post("/indicators/", map[string]any{"indicator": map[string]any{"type": "yara", "name": s.baseRule, "diamond": "capability",
		"pattern": fmt.Sprintf(`private rule %s { strings: $a = "goyeti" condition: $a }`, s.baseRule)}})
	loader := post("/indicators/", map[string]any{"indicator": map[string]any{"type": "yara", "name": s.loaderRule, "diamond": "capability",
		"pattern": fmt.Sprintf(`rule %s { condition: %s and filesize < 2MB }`, s.loaderRule, s.baseRule)}})
	s.loaderRef = "indicators/" + str(loader["id"])
	post("/indicators/tag", map[string]any{"ids": []string{refID(s.loaderRef)}, "tags": []string{s.tag}})
	post("/indicators/", map[string]any{"indicator": map[string]any{"type": "regex", "name": "goyeti_subdomains_" + run,
		"diamond": "infrastructure", "location": "network",
		"pattern": `\.` + strings.ReplaceAll(s.host, ".", `\.`) + `$`}})
	return s
}

// eventually retries check until it returns true, for 15 seconds. Searches
// and tag filters read ArangoSearch views, which index new objects about a
// second after they are saved.
func eventually(check func() bool) {
	for deadline := time.Now().Add(15 * time.Second); !check() && time.Now().Before(deadline); {
		time.Sleep(500 * time.Millisecond)
	}
}

func refID(ref string) string {
	_, id, _ := strings.Cut(ref, "/")
	return id
}

func refs[T interface{ ref() string }](items []T) []string {
	out := make([]string, 0, len(items))
	for _, item := range items {
		out = append(out, item.ref())
	}
	return out
}

// Linked, Neighbor and SearchHit get ref from the Summary they embed.
func (s Summary) ref() string { return s.Ref }

func find[T interface{ ref() string }](items []T, ref string) (T, bool) {
	for _, item := range items {
		if item.ref() == ref {
			return item, true
		}
	}
	var zero T
	return zero, false
}

func TestEndToEnd(t *testing.T) {
	c := liveClient(t)
	s := seed(t, c)
	ctx := context.Background()

	t.Run("lookup", func(t *testing.T) {
		unknown := "unknown-" + s.run + ".invalid"
		defanged := strings.ReplaceAll(s.host, ".", "[.]")
		results, _, err := c.Lookup(ctx, []string{s.ip, defanged, s.subdomain, unknown}, 0)
		if err != nil {
			t.Fatal(err)
		}
		ip, host, sub, none := results[0], results[1], results[2], results[3]

		if !ip.Known || ip.Ref != s.ipRef || ip.Type != "ipv4" {
			t.Errorf("ip: %+v", ip)
		}
		if !slices.Contains(ip.Tags, s.tag) || !slices.Contains(ip.ContextSources, "GoyetiE2E") {
			t.Errorf("ip tags %v, context sources %v", ip.Tags, ip.ContextSources)
		}
		if l, ok := find(ip.Entities, s.malwareRef); !ok || l.Link != "c2" || l.Direction != "out" || l.Name != s.malware {
			t.Errorf("ip entities %+v, want %s linked by c2", ip.Entities, s.malwareRef)
		}

		if !host.Known || host.Normalized != s.host || host.Ref != s.hostRef {
			t.Errorf("defanged host: %+v", host)
		}
		if l, ok := find(host.Observables, s.urlRef); !ok || l.Link != "hosted-on" || l.Direction != "in" {
			t.Errorf("host observables %+v, want %s hosted on it", host.Observables, s.urlRef)
		}
		// The host's link to the malware isn't asserted: /graph/match lists
		// each entity once per request, under the IP here (yeti#1463).

		if sub.Known || len(sub.Indicators) != 1 || !strings.HasPrefix(sub.Indicators[0].Name, "goyeti_subdomains_") {
			t.Errorf("unknown subdomain matching the regex: %+v", sub)
		}
		if none.Known {
			t.Errorf("unknown value: %+v", none)
		}
		again, _, err := c.Lookup(ctx, []string{unknown}, 0)
		if err != nil || again[0].Known {
			t.Errorf("lookup must never add a value: %+v, %v", again, err)
		}
	})

	t.Run("search", func(t *testing.T) {
		var hits []SearchHit
		eventually(func() bool {
			result, _, err := c.Search(ctx, s.malware, 0)
			if err != nil {
				t.Fatal(err)
			}
			hits = result.Hits
			return len(hits) > 0
		})
		if hit, ok := find(hits, s.malwareRef); !ok || !slices.Contains(hit.FoundBy, "keyword") {
			t.Errorf("hits %v, want %s found by keyword", refs(hits), s.malwareRef)
		}
	})

	t.Run("get", func(t *testing.T) {
		obj, _, err := c.Get(ctx, s.malwareRef)
		if err != nil {
			t.Fatal(err)
		}
		if obj["ref"] != s.malwareRef || obj["name"] != s.malware {
			t.Errorf("got %v", obj)
		}
		if _, ok := obj["acls"]; ok {
			t.Error("acls should be left out")
		}
		_, _, err = c.Get(ctx, "observables/goyeti-does-not-exist")
		var yetiErr *Error
		if !errors.As(err, &yetiErr) || yetiErr.Code != CodeNotFound {
			t.Errorf("get of a missing object: %v", err)
		}
	})

	t.Run("neighbors", func(t *testing.T) {
		result, _, err := c.Neighbors(ctx, s.actorRef, NeighborsOptions{Hops: 2})
		if err != nil {
			t.Fatal(err)
		}
		if n, ok := find(result.Neighbors, s.malwareRef); !ok || n.Link != "uses" || n.Direction != "out" || n.Via != nil {
			t.Errorf("neighbors %+v, want %s used directly", result.Neighbors, s.malwareRef)
		}
		if n, ok := find(result.Neighbors, s.ipRef); !ok || !slices.Equal(n.Via, []string{s.malwareRef}) {
			t.Errorf("neighbors %+v, want %s via %s", result.Neighbors, s.ipRef, s.malwareRef)
		}

		observables, _, err := c.Neighbors(ctx, s.malwareRef, NeighborsOptions{TargetTypes: []string{"observable"}})
		if err != nil {
			t.Fatal(err)
		}
		got := refs(observables.Neighbors)
		slices.Sort(got)
		want := []string{s.hostRef, s.ipRef}
		slices.Sort(want)
		if !slices.Equal(got, want) {
			t.Errorf("observable neighbors %v, want %v", got, want)
		}
	})

	t.Run("yara-bundle", func(t *testing.T) {
		for name, bundle := range map[string]func() (*YaraBundle, Meta, error){
			"by tag":       func() (*YaraBundle, Meta, error) { return c.YaraBundle(ctx, nil, []string{s.tag}, nil) },
			"by reference": func() (*YaraBundle, Meta, error) { return c.YaraBundle(ctx, []string{s.loaderRef}, nil, nil) },
		} {
			var got *YaraBundle
			eventually(func() bool {
				var err error
				if got, _, err = bundle(); err != nil {
					t.Fatalf("%s: %v", name, err)
				}
				return got.Rules > 0
			})
			if got.Rules != 2 || !strings.Contains(got.Bundle, s.baseRule) || !strings.Contains(got.Bundle, s.loaderRule) {
				t.Errorf("%s: %d rules, want the loader and its dependency:\n%s", name, got.Rules, got.Bundle)
			}
		}
	})

	t.Run("types", func(t *testing.T) {
		types, _, err := c.Types(ctx)
		if err != nil {
			t.Fatal(err)
		}
		for family, want := range map[string]string{"observables": "ipv4", "entities": "malware", "indicators": "yara"} {
			if !slices.Contains(types[family], want) {
				t.Errorf("%s types %v lack %s", family, types[family], want)
			}
		}
	})
}

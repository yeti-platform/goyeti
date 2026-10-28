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

func TestEndToEndMutate(t *testing.T) {
	c := liveClient(t)
	ctx := context.Background()
	now := time.Now().UnixNano()
	run := fmt.Sprintf("%d", now%1_000_000_000)
	ip := fmt.Sprintf("10.%d.%d.%d", now%250+1, now/250%250+1, now/62500%250+1)
	host := "mutate-" + run + ".example.com"
	tag := "goyeti-mutate-" + run

	t.Run("add observables", func(t *testing.T) {
		dry, _, err := c.AddObservables(ctx, []string{ip}, AddObservablesOptions{DryRun: true})
		if err != nil || !dry.DryRun || dry.Changes[0].Action != ActionCreated {
			t.Fatalf("dry run: %+v, %v", dry, err)
		}
		if looked, _, _ := c.Lookup(ctx, []string{ip}, 0); looked[0].Known {
			t.Fatal("a dry run added a value")
		}

		values := []string{ip, strings.ReplaceAll(host, ".", "[.]"), "not an observable " + run}
		added, meta, err := c.AddObservables(ctx, values, AddObservablesOptions{Tags: []string{tag}})
		if err != nil {
			t.Fatal(err)
		}
		if got := actions(added); !slices.Equal(got, []string{ActionCreated, ActionCreated, ActionFailed}) || meta.Failed != 1 {
			t.Fatalf("actions %v, meta %+v", got, meta)
		}
		if added.Changes[0].Type != "ipv4" || added.Changes[1].Type != "hostname" {
			t.Errorf("guessed types: %+v", added.Changes)
		}

		again, _, err := c.AddObservables(ctx, values[:2], AddObservablesOptions{})
		if err != nil || !slices.Equal(actions(again), []string{ActionExists, ActionExists}) {
			t.Errorf("adding again: %+v, %v", again, err)
		}
		looked, _, _ := c.Lookup(ctx, []string{ip}, 0)
		if !slices.Contains(looked[0].Tags, tag) {
			t.Errorf("tags %v lack %s", looked[0].Tags, tag)
		}
	})

	var entityRef string
	t.Run("add entity", func(t *testing.T) {
		name := "GoyetiMutate" + run
		created, _, err := c.AddEntity(ctx, "malware", name, AddObjectOptions{Fields: map[string]any{"description": "original"}})
		if err != nil || created.Changes[0].Action != ActionCreated {
			t.Fatalf("create: %+v, %v", created, err)
		}
		entityRef = created.Changes[0].Ref
		// Right away, before Yeti's search views have indexed it: an
		// existing entity must be found, not created again over itself.
		again, _, err := c.AddEntity(ctx, "malware", name, AddObjectOptions{Fields: map[string]any{"description": "overwritten"}, Tags: []string{tag}})
		if err != nil || again.Changes[0].Action != ActionExists || again.Changes[0].Ref != entityRef {
			t.Fatalf("adding again: %+v, %v", again, err)
		}
		obj, _, err := c.Get(ctx, entityRef)
		if err != nil || obj["description"] != "original" {
			t.Errorf("the existing entity changed: %v, %v", obj["description"], err)
		}
		if !slices.Contains(tagNames(obj["tags"]), tag) {
			t.Errorf("tags %v lack %s", tagNames(obj["tags"]), tag)
		}
	})

	t.Run("add indicator", func(t *testing.T) {
		name := "goyeti_mutate_" + run
		fields := map[string]any{"pattern": "goyeti-" + run, "diamond": "infrastructure", "location": "network"}
		created, _, err := c.AddIndicator(ctx, "regex", name, AddObjectOptions{Fields: fields, Tags: []string{tag}})
		if err != nil || created.Changes[0].Action != ActionCreated {
			t.Fatalf("create: %+v, %v", created, err)
		}
		again, _, err := c.AddIndicator(ctx, "regex", name, AddObjectOptions{Fields: fields})
		if err != nil || again.Changes[0].Action != ActionExists {
			t.Errorf("adding again: %+v, %v", again, err)
		}
		_, _, err = c.AddIndicator(ctx, "yara", "broken_"+run, AddObjectOptions{Fields: map[string]any{
			"pattern": "rule broken_" + run + " { condition: missing_" + run + " }", "diamond": "capability"}})
		var yetiErr *Error
		if !errors.As(err, &yetiErr) || yetiErr.Code != CodeInvalidRequest || !strings.Contains(yetiErr.Message, "missing_"+run) {
			t.Errorf("an invalid rule: %v", err)
		}
	})

	t.Run("link", func(t *testing.T) {
		looked, _, _ := c.Lookup(ctx, []string{ip}, 0)
		ipRef := looked[0].Ref
		linked, _, err := c.Link(ctx, ipRef, "c2", entityRef, "", false)
		if err != nil || linked.Changes[0].Action != ActionLinked {
			t.Fatalf("link: %+v, %v", linked, err)
		}
		again, _, err := c.Link(ctx, ipRef, "c2", entityRef, "", false)
		if err != nil || again.Changes[0].Action != ActionExists || again.Changes[0].Ref != linked.Changes[0].Ref {
			t.Errorf("linking again: %+v, %v", again, err)
		}
		around, _, _ := c.Neighbors(ctx, entityRef, NeighborsOptions{})
		if n, ok := find(around.Neighbors, ipRef); !ok || n.Link != "c2" {
			t.Errorf("neighbors %+v lack %s", around.Neighbors, ipRef)
		}
	})

	t.Run("tag and untag", func(t *testing.T) {
		looked, _, _ := c.Lookup(ctx, []string{ip}, 0)
		ipRef := looked[0].Ref
		tagged, _, err := c.Tag(ctx, []string{ipRef}, []string{"Second Tag"}, false)
		if err != nil || tagged.Changes[0].Action != ActionTagged || !slices.Contains(tagged.Changes[0].Tags, "second_tag") {
			t.Fatalf("tag: %+v, %v", tagged, err)
		}
		untagged, _, err := c.Untag(ctx, []string{ipRef}, []string{"second tag"}, false)
		if err != nil || untagged.Changes[0].Action != ActionUntagged {
			t.Fatalf("untag: %+v, %v", untagged, err)
		}
		looked, _, _ = c.Lookup(ctx, []string{ip}, 0)
		if slices.Contains(looked[0].Tags, "second_tag") || !slices.Contains(looked[0].Tags, tag) {
			t.Errorf("tags after untagging: %v", looked[0].Tags)
		}
	})

	t.Run("add context", func(t *testing.T) {
		added, _, err := c.AddContext(ctx, entityRef, map[string]any{"run": run}, false)
		if err != nil || added.Changes[0].Action != ActionAdded {
			t.Fatalf("add context: %+v, %v", added, err)
		}
		obj, _, _ := c.Get(ctx, entityRef)
		if sources := contextSources(obj["context"]); !slices.Contains(sources, "agent:goyeti-e2e") {
			t.Errorf("context sources %v", sources)
		}
	})
}

func TestEndToEndAdmin(t *testing.T) {
	c := liveClient(t)
	ctx := context.Background()
	run := fmt.Sprintf("%d", time.Now().UnixNano()%1_000_000_000)

	var actor, malware string
	for name, ref := range map[string]*string{"GoyetiAdminActor" + run: &actor, "GoyetiAdminMalware" + run: &malware} {
		added, _, err := c.AddEntity(ctx, "malware", name, AddObjectOptions{})
		if err != nil {
			t.Fatal(err)
		}
		*ref = added.Changes[0].Ref
	}
	if _, _, err := c.Link(ctx, actor, "uses", malware, "", false); err != nil {
		t.Fatal(err)
	}

	t.Run("unlink", func(t *testing.T) {
		_, _, err := c.Unlink(ctx, actor, "uses", malware, RemoveOptions{})
		wantCode(t, err, CodeUsage)
		dry, _, err := c.Unlink(ctx, actor, "uses", malware, RemoveOptions{DryRun: true})
		if err != nil || dry.Changes[0].Action != ActionUnlinked {
			t.Fatalf("dry run: %+v, %v", dry, err)
		}
		unlinked, _, err := c.Unlink(ctx, actor, "uses", malware, RemoveOptions{Confirm: true})
		if err != nil || unlinked.Changes[0].Action != ActionUnlinked {
			t.Fatalf("unlink: %+v, %v", unlinked, err)
		}
		again, _, err := c.Unlink(ctx, actor, "uses", malware, RemoveOptions{Confirm: true})
		if err != nil || again.Changes[0].Action != ActionAbsent {
			t.Errorf("unlinking again: %+v, %v", again, err)
		}
		if _, _, err := c.Link(ctx, actor, "uses", malware, "", false); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("delete", func(t *testing.T) {
		dry, _, err := c.Delete(ctx, []string{actor}, RemoveOptions{DryRun: true})
		if err != nil {
			t.Fatal(err)
		}
		if ch := dry.Changes[0]; ch.Action != ActionDeleted || ch.Name != "GoyetiAdminActor"+run || ch.Links != 1 {
			t.Errorf("dry run: %+v", ch)
		}
		if _, _, err := c.Get(ctx, actor); err != nil {
			t.Fatalf("a dry run deleted %s: %v", actor, err)
		}

		deleted, meta, err := c.Delete(ctx, []string{actor, malware}, RemoveOptions{Confirm: true})
		if err != nil || meta.Failed != 0 || !slices.Equal(actions(deleted), []string{ActionDeleted, ActionDeleted}) {
			t.Fatalf("delete: %+v, %+v, %v", deleted, meta, err)
		}
		_, _, err = c.Get(ctx, actor)
		wantCode(t, err, CodeNotFound)
		again, _, err := c.Delete(ctx, []string{actor}, RemoveOptions{Confirm: true})
		if err != nil || again.Changes[0].Action != ActionAbsent {
			t.Errorf("deleting again: %+v, %v", again, err)
		}
	})

	t.Run("tasks", func(t *testing.T) {
		// Oneshot tasks have no schedule, so enabling one runs nothing.
		var oneshot Task
		eventually(func() bool {
			tasks, _, err := c.Tasks(ctx, TasksOptions{Type: "oneshot"})
			if err == nil && len(tasks) > 0 {
				oneshot = tasks[0]
			}
			return oneshot.Name != ""
		})
		if oneshot.Name == "" {
			t.Fatal("Yeti lists no oneshot tasks")
		}
		listed, _, err := c.Tasks(ctx, TasksOptions{Names: []string{strings.ToLower(oneshot.Name)}})
		if _, found := findTask(listed, oneshot.Name); err != nil || !found {
			t.Errorf("listing by lowercase name: %+v, %v", listed, err)
		}
		t.Cleanup(func() {
			if oneshot.Enabled {
				c.EnableTasks(ctx, []string{oneshot.Name}, false)
			} else {
				c.DisableTasks(ctx, []string{oneshot.Name}, false)
			}
		})
		if _, _, err := c.DisableTasks(ctx, []string{oneshot.Name}, false); err != nil {
			t.Fatal(err)
		}

		_, _, err = c.RunTask(ctx, oneshot.Name, map[string]any{"value": "example.com"}, true)
		wantCode(t, err, CodeInvalidRequest)

		enabled, _, err := c.EnableTasks(ctx, []string{oneshot.Name}, false)
		if err != nil || enabled.Changes[0].Action != ActionEnabled {
			t.Fatalf("enable: %+v, %v", enabled, err)
		}
		again, _, err := c.EnableTasks(ctx, []string{oneshot.Name}, false)
		if err != nil || again.Changes[0].Action != ActionUnchanged {
			t.Errorf("enabling again: %+v, %v", again, err)
		}

		_, _, err = c.RunTask(ctx, oneshot.Name, nil, true)
		wantCode(t, err, CodeUsage)
		queued, _, err := c.RunTask(ctx, oneshot.Name, map[string]any{"value": "example.com"}, true)
		if err != nil || queued.Changes[0].Action != ActionQueued || !queued.DryRun {
			t.Errorf("run, dry: %+v, %v", queued, err)
		}
		_, _, err = c.RunTask(ctx, "NoSuchTask"+run, nil, true)
		wantCode(t, err, CodeNotFound)
	})
}

func findTask(tasks []Task, name string) (Task, bool) {
	for _, task := range tasks {
		if task.Name == name {
			return task, true
		}
	}
	return Task{}, false
}

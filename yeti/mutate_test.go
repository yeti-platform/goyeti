package yeti

import (
	"context"
	"reflect"
	"strings"
	"testing"
)

var writes = map[string]route{
	"POST /api/v2/graph/match":                       {fixture: "match_before_add.json", then: []string{"match_after_add.json"}},
	"POST /api/v2/observables/import/text":           {fixture: "import_text.json"},
	"POST /api/v2/observables/bulk":                  {fixture: "bulk.json"},
	"POST /api/v2/observables/tag":                   {fixture: "tag.json"},
	"GET /api/v2/entities/?name=Emotet&type=malware": {fixture: "get_entity.json"},
	"POST /api/v2/entities/":                         {fixture: "new_entity.json"},
	"POST /api/v2/entities/tag":                      {fixture: "tag.json"},
	"POST /api/v2/indicators/":                       {fixture: "new_indicator.json"},
	"POST /api/v2/indicators/tag":                    {fixture: "tag.json"},
	"POST /api/v2/graph/add":                         {fixture: "graph_add.json"},
	"POST /api/v2/observables/22898861/context":      {fixture: "add_context.json"},
}

func actions(c *Changes) []string {
	var out []string
	for _, change := range c.Changes {
		out = append(out, change.Action)
	}
	return out
}

func TestAddObservables(t *testing.T) {
	client, fake := newTestClient(t, writes)
	result, meta, err := client.AddObservables(context.Background(),
		[]string{"203.0.113.7", "198.51.100[.]23", "not an observable"}, AddObservablesOptions{Tags: []string{"campaign-x"}})
	if err != nil {
		t.Fatal(err)
	}
	if got := actions(result); !reflect.DeepEqual(got, []string{ActionExists, ActionCreated, ActionFailed}) {
		t.Errorf("actions %v", got)
	}
	if created := result.Changes[1]; created.Item != "198.51.100[.]23" || created.Ref != "observables/22919642" || created.Type != "ipv4" {
		t.Errorf("created %+v", created)
	}
	if failure := result.Changes[2]; failure.Error == nil || failure.Error.Hint == "" {
		t.Errorf("failure %+v", failure)
	}
	if meta.Failed != 1 || len(meta.Warnings) != 1 {
		t.Errorf("meta %+v", meta)
	}
	sent := fake.bodies["POST /api/v2/observables/import/text"]
	if sent["text"] != "198.51.100.23\nnot an observable" || !reflect.DeepEqual(sent["tags"], []any{"campaign-x"}) {
		t.Errorf("import sent %v", sent)
	}
	// The existing value is tagged too.
	if tagged := fake.bodies["POST /api/v2/observables/tag"]; !reflect.DeepEqual(tagged["ids"], []any{"22898861"}) {
		t.Errorf("tag sent %v", tagged)
	}
}

func TestAddObservablesWithAType(t *testing.T) {
	client, fake := newTestClient(t, writes)
	if _, _, err := client.AddObservables(context.Background(), []string{"198.51.100.23"}, AddObservablesOptions{Type: "ipv4"}); err != nil {
		t.Fatal(err)
	}
	sent := fake.bodies["POST /api/v2/observables/bulk"]["observables"].([]any)
	if first := sent[0].(map[string]any); first["type"] != "ipv4" || first["value"] != "198.51.100.23" {
		t.Errorf("bulk sent %v", sent)
	}
	if _, imported := fake.bodies["POST /api/v2/observables/import/text"]; imported {
		t.Error("an explicit type must not go through import, which guesses types")
	}
}

func TestDryRunsWriteNothing(t *testing.T) {
	client, fake := newTestClient(t, writes)
	ctx := context.Background()
	steps := map[string]func() (*Changes, Meta, error){
		"add observables": func() (*Changes, Meta, error) {
			return client.AddObservables(ctx, []string{"203.0.113.7", "198.51.100.23"}, AddObservablesOptions{Tags: []string{"x"}, DryRun: true})
		},
		"add entity": func() (*Changes, Meta, error) {
			return client.AddEntity(ctx, "malware", "Qakbot", AddObjectOptions{DryRun: true})
		},
		"link": func() (*Changes, Meta, error) {
			return client.Link(ctx, "entities/22899051", "uses", "observables/22898861", "", true)
		},
		"tag": func() (*Changes, Meta, error) {
			return client.Tag(ctx, []string{"observables/22898861"}, []string{"new"}, true)
		},
		"add context": func() (*Changes, Meta, error) {
			return client.AddContext(ctx, "observables/22898861", map[string]any{"k": "v"}, true)
		},
	}
	for name, step := range steps {
		result, _, err := step()
		if err != nil {
			t.Fatalf("%s: %v", name, err)
		}
		if !result.DryRun || len(result.Changes) == 0 {
			t.Errorf("%s: %+v", name, result)
		}
	}
	for request := range fake.bodies {
		if !strings.HasSuffix(request, "/graph/match") && !strings.HasSuffix(request, "/graph/search") {
			t.Errorf("a dry run sent %s", request)
		}
	}
}

func TestAddEntity(t *testing.T) {
	client, fake := newTestClient(t, writes)
	ctx := context.Background()

	existing, _, err := client.AddEntity(ctx, "malware", "Emotet", AddObjectOptions{Tags: []string{"loader"},
		Fields: map[string]any{"description": "would overwrite"}})
	if err != nil {
		t.Fatal(err)
	}
	if c := existing.Changes[0]; c.Action != ActionExists || c.Ref != "entities/22899024" {
		t.Errorf("existing entity: %+v", c)
	}
	if _, created := fake.bodies["POST /api/v2/entities/"]; created {
		t.Error("an existing entity must not be created again: Yeti would overwrite its fields")
	}
	if tagged := fake.bodies["POST /api/v2/entities/tag"]; !reflect.DeepEqual(tagged["ids"], []any{"22899024"}) {
		t.Errorf("tag sent %v", tagged)
	}

	created, _, err := client.AddEntity(ctx, "malware", "Qakbot", AddObjectOptions{Tags: []string{"qakbot"},
		Fields: map[string]any{"description": "Banking trojan."}})
	if err != nil {
		t.Fatal(err)
	}
	if c := created.Changes[0]; c.Action != ActionCreated || c.Ref != "entities/22919794" {
		t.Errorf("new entity: %+v", c)
	}
	sent := fake.bodies["POST /api/v2/entities/"]
	want := map[string]any{"type": "malware", "name": "Qakbot", "description": "Banking trojan."}
	if !reflect.DeepEqual(sent["entity"], want) || !reflect.DeepEqual(sent["tags"], []any{"qakbot"}) {
		t.Errorf("create sent %v", sent)
	}

	_, _, err = client.AddEntity(ctx, "malware", " ", AddObjectOptions{})
	wantCode(t, err, CodeUsage)
}

func TestAddIndicator(t *testing.T) {
	client, fake := newTestClient(t, writes)
	result, _, err := client.AddIndicator(context.Background(), "regex", "qakbot_paths", AddObjectOptions{
		Tags: []string{"qakbot"}, Fields: map[string]any{"pattern": `\\qbot\\`, "diamond": "capability", "location": "filesystem"}})
	if err != nil {
		t.Fatal(err)
	}
	if c := result.Changes[0]; c.Action != ActionCreated || !strings.HasPrefix(c.Ref, "indicators/") {
		t.Errorf("new indicator: %+v", c)
	}
	// Indicators are created without tags, then tagged.
	if _, ok := fake.bodies["POST /api/v2/indicators/"]["tags"]; ok {
		t.Error("the indicator request has no tags field")
	}
	if tagged := fake.bodies["POST /api/v2/indicators/tag"]; !reflect.DeepEqual(tagged["tags"], []any{"qakbot"}) {
		t.Errorf("tag sent %v", tagged)
	}

	invalid, _ := newTestClient(t, map[string]route{
		"POST /api/v2/indicators/": {status: 400, fixture: "new_indicator_invalid.json"},
	})
	_, _, err = invalid.AddIndicator(context.Background(), "yara", "broken", AddObjectOptions{})
	if e := wantCode(t, err, CodeInvalidRequest); !strings.Contains(e.Message, `"missing_dependencies":["nope"]`) {
		t.Errorf("error %+v", e)
	}
}

func TestLink(t *testing.T) {
	client, fake := newTestClient(t, writes)
	ctx := context.Background()
	existing, _, err := client.Link(ctx, "entities/22899051", "uses", "entities/22899024", "", false)
	if err != nil {
		t.Fatal(err)
	}
	if c := existing.Changes[0]; c.Action != ActionExists || c.Ref != "links/22899112" {
		t.Errorf("existing link: %+v", c)
	}
	if _, linked := fake.bodies["POST /api/v2/graph/add"]; linked {
		t.Error("an existing link must not be added again: Yeti would count it twice")
	}
	sent := fake.bodies["POST /api/v2/graph/search"]
	if sent["direction"] != "outbound" || !reflect.DeepEqual(sent["link_types"], []any{"uses"}) {
		t.Errorf("search sent %v", sent)
	}

	created, _, err := client.Link(ctx, "entities/22899051", "uses", "entities/22919794", "since 2024", false)
	if err != nil {
		t.Fatal(err)
	}
	if c := created.Changes[0]; c.Action != ActionLinked || c.Ref != "links/22919864" {
		t.Errorf("new link: %+v", c)
	}
	if sent := fake.bodies["POST /api/v2/graph/add"]; sent["description"] != "since 2024" || sent["link_type"] != "uses" {
		t.Errorf("add sent %v", sent)
	}

	_, _, err = client.Link(ctx, "entities/1", "", "entities/2", "", false)
	wantCode(t, err, CodeUsage)
	_, _, err = client.Link(ctx, "Emotet", "uses", "entities/2", "", false)
	wantCode(t, err, CodeUsage)
}

func TestTagAndUntag(t *testing.T) {
	client, fake := newTestClient(t, writes)
	ctx := context.Background()

	tagged, _, err := client.Tag(ctx, []string{"observables/22898861"}, []string{"Emotet", "New Tag"}, false)
	if err != nil {
		t.Fatal(err)
	}
	if c := tagged.Changes[0]; c.Action != ActionTagged || !reflect.DeepEqual(c.Tags, []string{"c2", "emotet", "new_tag"}) {
		t.Errorf("tag: %+v", c)
	}
	sent := fake.bodies["POST /api/v2/observables/tag"]
	if sent["strict"] != false || !reflect.DeepEqual(sent["tags"], []any{"emotet", "new_tag"}) {
		t.Errorf("tag sent %v", sent)
	}

	delete(fake.bodies, "POST /api/v2/observables/tag")
	unchanged, _, err := client.Tag(ctx, []string{"observables/22898861"}, []string{"c2"}, false)
	if err != nil {
		t.Fatal(err)
	}
	if c := unchanged.Changes[0]; c.Action != ActionUnchanged {
		t.Errorf("tag again: %+v", c)
	}
	if _, sent := fake.bodies["POST /api/v2/observables/tag"]; sent {
		t.Error("an object that already has the tags must not be tagged again")
	}

	untagged, _, err := client.Untag(ctx, []string{"observables/22898861"}, []string{"C2"}, false)
	if err != nil {
		t.Fatal(err)
	}
	if c := untagged.Changes[0]; c.Action != ActionUntagged || !reflect.DeepEqual(c.Tags, []string{"emotet"}) {
		t.Errorf("untag: %+v", c)
	}
	// Removing replaces the tags with the ones the object keeps.
	if sent := fake.bodies["POST /api/v2/observables/tag"]; sent["strict"] != true || !reflect.DeepEqual(sent["tags"], []any{"emotet"}) {
		t.Errorf("untag sent %v", sent)
	}

	mixed, meta, err := client.Tag(ctx, []string{"dfiq/1", "observables/404"}, []string{"x"}, false)
	if err != nil {
		t.Fatal(err)
	}
	if got := actions(mixed); !reflect.DeepEqual(got, []string{ActionFailed, ActionFailed}) || meta.Failed != 2 {
		t.Errorf("actions %v, meta %+v", got, meta)
	}
	if mixed.Changes[1].Error.Code != CodeNotFound {
		t.Errorf("missing object: %+v", mixed.Changes[1].Error)
	}

	_, _, err = client.Tag(ctx, []string{"observables/22898861"}, []string{" "}, false)
	wantCode(t, err, CodeUsage)
}

func TestAddContext(t *testing.T) {
	client, fake := newTestClient(t, writes)
	agent, err := New(Config{URL: "http://unused.invalid", APIKey: "key", AgentName: "triage-bot"})
	if err != nil {
		t.Fatal(err)
	}
	client.agent = agent.agent

	result, _, err := client.AddContext(context.Background(), "observables/22898861", map[string]any{"note": "seen in campaign X"}, false)
	if err != nil {
		t.Fatal(err)
	}
	if c := result.Changes[0]; c.Action != ActionAdded || c.Ref != "observables/22898861" {
		t.Errorf("change %+v", c)
	}
	sent := fake.bodies["POST /api/v2/observables/22898861/context"]
	if sent["source"] != "agent:triage-bot" || !reflect.DeepEqual(sent["context"], map[string]any{"note": "seen in campaign X"}) {
		t.Errorf("sent %v", sent)
	}

	_, _, err = client.AddContext(context.Background(), "dfiq/1", map[string]any{"k": "v"}, false)
	wantCode(t, err, CodeUsage)
	_, _, err = client.AddContext(context.Background(), "observables/1", nil, false)
	wantCode(t, err, CodeUsage)
}

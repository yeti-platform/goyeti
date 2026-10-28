package yeti

import (
	"context"
	"reflect"
	"testing"
)

var admin = map[string]route{
	"GET /api/v2/entities/22899024":       {fixture: "get_entity.json"},
	"GET /api/v2/entities/404":            {status: 404, fixture: "get_missing.json"},
	"DELETE /api/v2/observables/22898861": {fixture: "delete.json"},
	"DELETE /api/v2/entities/22899024":    {fixture: "delete.json"},
	"DELETE /api/v2/graph/22899112":       {fixture: "delete.json"},
	"POST /api/v2/tasks/search":           {fixture: "tasks_search.json"},
	"POST /api/v2/tasks/AbuseIPDB/toggle": {fixture: "task_toggle_enabled.json"},
	"POST /api/v2/tasks/AbuseIPDB/run":    {fixture: "task_run.json"},
}

func TestDelete(t *testing.T) {
	client, fake := newTestClient(t, admin)
	ctx := context.Background()
	refs := []string{"observables/22898861", "entities/22899024", "entities/404"}

	_, _, err := client.Delete(ctx, refs, RemoveOptions{})
	wantCode(t, err, CodeUsage)
	_, _, err = client.Delete(ctx, []string{"Emotet"}, RemoveOptions{Confirm: true})
	wantCode(t, err, CodeUsage)

	planned, _, err := client.Delete(ctx, refs, RemoveOptions{DryRun: true})
	if err != nil {
		t.Fatal(err)
	}
	if fake.calls["DELETE /api/v2/observables/22898861"]+fake.calls["DELETE /api/v2/entities/22899024"] > 0 {
		t.Error("a dry run deleted something")
	}
	if got := actions(planned); !reflect.DeepEqual(got, []string{ActionDeleted, ActionDeleted, ActionAbsent}) || !planned.DryRun {
		t.Errorf("dry run: %v", planned)
	}
	if c := planned.Changes[1]; c.Name != "Emotet" || c.Type != "malware" || c.Links != 3 {
		t.Errorf("an object to delete says what it is and how many links go with it: %+v", c)
	}
	if sent := fake.bodies["POST /api/v2/graph/search"]; sent["direction"] != "any" || sent["max_hops"] != 1.0 {
		t.Errorf("link count searched %v", sent)
	}

	deleted, meta, err := client.Delete(ctx, refs, RemoveOptions{Confirm: true})
	if err != nil || meta.Failed != 0 {
		t.Fatal(err, meta)
	}
	if got := actions(deleted); !reflect.DeepEqual(got, []string{ActionDeleted, ActionDeleted, ActionAbsent}) {
		t.Errorf("delete: %v", got)
	}
	if fake.calls["DELETE /api/v2/observables/22898861"] != 1 || fake.calls["DELETE /api/v2/entities/22899024"] != 1 {
		t.Errorf("delete calls: %v", fake.calls)
	}
}

func TestUnlink(t *testing.T) {
	client, fake := newTestClient(t, admin)
	ctx := context.Background()

	_, _, err := client.Unlink(ctx, "entities/22899051", "uses", "entities/22899024", RemoveOptions{})
	wantCode(t, err, CodeUsage)

	unlinked, _, err := client.Unlink(ctx, "entities/22899051", "uses", "entities/22899024", RemoveOptions{Confirm: true})
	if err != nil {
		t.Fatal(err)
	}
	if c := unlinked.Changes[0]; c.Action != ActionUnlinked || c.Ref != "links/22899112" {
		t.Errorf("unlink: %+v", c)
	}
	if fake.calls["DELETE /api/v2/graph/22899112"] != 1 {
		t.Errorf("delete calls: %v", fake.calls)
	}

	absent, _, err := client.Unlink(ctx, "entities/22899051", "uses", "entities/1", RemoveOptions{Confirm: true})
	if err != nil {
		t.Fatal(err)
	}
	if c := absent.Changes[0]; c.Action != ActionAbsent {
		t.Errorf("a link that doesn't exist: %+v", c)
	}
}

func TestTasks(t *testing.T) {
	client, fake := newTestClient(t, admin)
	tasks, _, err := client.Tasks(context.Background(), TasksOptions{Names: []string{"abuse"}, Type: "feed"})
	if err != nil {
		t.Fatal(err)
	}
	if len(tasks) != 2 || tasks[0].Name != "AbuseCHMalwareBazaaar" || tasks[1].Frequency != "PT5H" || tasks[1].Enabled {
		t.Errorf("tasks: %+v", tasks)
	}
	sent := fake.bodies["POST /api/v2/tasks/search"]
	if sent["type"] != "feed" || !reflect.DeepEqual(sent["query"], map[string]any{"name": "abuse"}) {
		t.Errorf("search sent %v", sent)
	}
}

func TestRunTask(t *testing.T) {
	client, _ := newTestClient(t, admin)
	ctx := context.Background()
	_, _, err := client.RunTask(ctx, "AbuseIPDB", nil, false)
	if e := wantCode(t, err, CodeInvalidRequest); e.Hint == "" {
		t.Error("running a disabled task says how to enable it")
	}

	client, fake := newTestClient(t, map[string]route{
		"POST /api/v2/tasks/search":        {fixture: "tasks_search_enabled.json"},
		"POST /api/v2/tasks/AbuseIPDB/run": {fixture: "task_run.json"},
	})
	queued, _, err := client.RunTask(ctx, "abuseipdb", map[string]any{"since": "1d"}, false)
	if err != nil {
		t.Fatal(err)
	}
	if c := queued.Changes[0]; c.Action != ActionQueued || c.Name != "AbuseIPDB" || c.Type != "feed" {
		t.Errorf("run: %+v", c)
	}
	if sent := fake.bodies["POST /api/v2/tasks/AbuseIPDB/run"]; !reflect.DeepEqual(sent, map[string]any{"params": map[string]any{"since": "1d"}}) {
		t.Errorf("run sent %v", sent)
	}

	_, _, err = client.RunTask(ctx, "NoSuchTask", nil, false)
	wantCode(t, err, CodeNotFound)
}

func TestEnableAndDisableTasks(t *testing.T) {
	client, fake := newTestClient(t, admin)
	ctx := context.Background()

	enabled, meta, err := client.EnableTasks(ctx, []string{"AbuseIPDB", "Abuse"}, false)
	if err != nil {
		t.Fatal(err)
	}
	if got := actions(enabled); !reflect.DeepEqual(got, []string{ActionEnabled, ActionFailed}) || meta.Failed != 1 {
		t.Errorf("enable: %v %+v", got, meta)
	}
	if e := enabled.Changes[1].Error; e == nil || e.Code != CodeNotFound || e.Hint == "" {
		t.Errorf("an ambiguous name suggests the tasks it could be: %+v", e)
	}
	if fake.calls["POST /api/v2/tasks/AbuseIPDB/toggle"] != 1 {
		t.Errorf("toggle calls: %v", fake.calls)
	}

	disabled, _, err := client.DisableTasks(ctx, []string{"AbuseIPDB"}, false)
	if err != nil {
		t.Fatal(err)
	}
	if got := actions(disabled); !reflect.DeepEqual(got, []string{ActionUnchanged}) {
		t.Errorf("disabling a disabled task: %v", got)
	}
	if fake.calls["POST /api/v2/tasks/AbuseIPDB/toggle"] != 1 {
		t.Error("a task already in the wanted state must not be toggled: Yeti only flips it")
	}
}

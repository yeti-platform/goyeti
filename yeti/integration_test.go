//go:build integration

// Run against a real Yeti, e.g. the dev stack from yeti-docker:
//
//	YETI_URL=http://localhost:8000 YETI_API_KEY=... go test -tags integration ./yeti/
package yeti

import (
	"context"
	"errors"
	"os"
	"testing"
)

func liveClient(t *testing.T) *Client {
	t.Helper()
	if os.Getenv("YETI_URL") == "" || os.Getenv("YETI_API_KEY") == "" {
		t.Skip("set YETI_URL and YETI_API_KEY to run against a real Yeti")
	}
	client, err := New(Config{URL: os.Getenv("YETI_URL"), APIKey: os.Getenv("YETI_API_KEY")})
	if err != nil {
		t.Fatal(err)
	}
	return client
}

func TestLiveRead(t *testing.T) {
	client := liveClient(t)
	ctx := context.Background()

	types, _, err := client.Types(ctx)
	if err != nil {
		t.Fatal(err)
	}
	for _, family := range []string{"observables", "entities", "indicators", "dfiq"} {
		if len(types[family]) == 0 {
			t.Errorf("no %s types", family)
		}
	}

	results, _, err := client.Lookup(ctx, []string{"goyeti-integration-test.invalid"}, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(results) != 1 || results[0].Known {
		t.Errorf("an unknown value was reported as %+v", results)
	}
	// Looking up must never add the value.
	if again, _, _ := client.Lookup(ctx, []string{"goyeti-integration-test.invalid"}, 0); again[0].Known {
		t.Error("lookup added an unknown value")
	}

	search, _, err := client.Search(ctx, "goyeti", 5)
	if err != nil {
		t.Fatal(err)
	}
	for _, hit := range search.Hits {
		if _, _, err := parseRef(hit.Ref); err != nil {
			t.Errorf("search hit %+v has a bad reference", hit)
			continue
		}
		if _, _, err := client.Get(ctx, hit.Ref); err != nil {
			t.Errorf("get %s: %v", hit.Ref, err)
		}
		if _, _, err := client.Neighbors(ctx, hit.Ref, NeighborsOptions{}); err != nil {
			t.Errorf("neighbors %s: %v", hit.Ref, err)
		}
	}

	_, _, err = client.Get(ctx, "observables/goyeti-does-not-exist")
	var yetiErr *Error
	if !errors.As(err, &yetiErr) || yetiErr.Code != CodeNotFound {
		t.Errorf("get of a missing object: %v", err)
	}
}

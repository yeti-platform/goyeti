package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Object is a Yeti object as the API returns it.
type Object map[string]any

// Edge is a relationship between two objects, referenced as
// "<collection>/<id>".
type Edge struct {
	ID          string `json:"id"`
	Source      string `json:"source"`
	Target      string `json:"target"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

// Linked is an object and the relationship it was reached through.
type Linked struct {
	Edge   Edge
	Object Object
}

// UnmarshalJSON reads Yeti's [relationship, object] pairs.
func (l *Linked) UnmarshalJSON(data []byte) error {
	var pair []json.RawMessage
	if err := json.Unmarshal(data, &pair); err != nil {
		return err
	}
	if len(pair) != 2 {
		return fmt.Errorf("expected a [relationship, object] pair, got %d items", len(pair))
	}
	if err := json.Unmarshal(pair[0], &l.Edge); err != nil {
		return err
	}
	return json.Unmarshal(pair[1], &l.Object)
}

// Match is a value an indicator matched.
type Match struct {
	Value     string
	Indicator Object
}

// UnmarshalJSON reads Yeti's [value, indicator] pairs.
func (m *Match) UnmarshalJSON(data []byte) error {
	var pair []json.RawMessage
	if err := json.Unmarshal(data, &pair); err != nil {
		return err
	}
	if len(pair) != 2 {
		return fmt.Errorf("expected a [value, indicator] pair, got %d items", len(pair))
	}
	if err := json.Unmarshal(pair[0], &m.Value); err != nil {
		return err
	}
	return json.Unmarshal(pair[1], &m.Indicator)
}

// MatchResponse is what /graph/match knows about a list of values.
type MatchResponse struct {
	Known       []Object `json:"known"`
	Unknown     []string `json:"unknown"`
	Entities    []Linked `json:"entities"`
	Observables []Linked `json:"observables"`
	Matches     []Match  `json:"matches"`
}

// Match looks values up, with their linked entities and observables and the
// indicators that match them. It never adds unknown values.
func (c *Client) Match(ctx context.Context, values []string) (*MatchResponse, error) {
	body := map[string]any{"observables": values, "fetch_neighbors": true, "add_unknown": false}
	var resp MatchResponse
	if err := c.Do(ctx, http.MethodPost, "/graph/match", body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Section is one family's results from a search.
type Section struct {
	Type    string   `json:"type"`
	Total   int      `json:"total"`
	Results []Object `json:"results"`
}

// Search finds objects of every family by keyword.
func (c *Client) Search(ctx context.Context, query string, countPerType int) ([]Section, error) {
	body := map[string]any{"query": query, "count_per_type": countPerType}
	var resp struct {
		Sections []Section `json:"sections"`
	}
	if err := c.Do(ctx, http.MethodPost, "/search/", body, &resp); err != nil {
		return nil, err
	}
	return resp.Sections, nil
}

// SemanticSummary is the summary Yeti returns with each semantic match.
type SemanticSummary struct {
	ID          string   `json:"id"`
	RootType    string   `json:"root_type"`
	Type        string   `json:"type"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

// SemanticResult is one semantic match, with its 0-1 similarity score.
type SemanticResult struct {
	Score   float64         `json:"score"`
	Summary SemanticSummary `json:"object_summary"`
}

// SemanticSection is one family's semantic matches.
type SemanticSection struct {
	Type    string           `json:"type"`
	Total   int              `json:"total"`
	Results []SemanticResult `json:"results"`
}

// SemanticSearch finds entities, indicators and DFIQ by meaning.
func (c *Client) SemanticSearch(ctx context.Context, query string, count int) ([]SemanticSection, error) {
	body := map[string]any{"query": query, "count": count}
	var resp struct {
		Sections []SemanticSection `json:"sections"`
	}
	if err := c.Do(ctx, http.MethodPost, "/search/semantic", body, &resp); err != nil {
		return nil, err
	}
	return resp.Sections, nil
}

// Get fetches one object from collection: "observables", "entities",
// "indicators" or "dfiq".
func (c *Client) Get(ctx context.Context, collection, id string) (Object, error) {
	var obj Object
	path := "/" + collection + "/" + url.PathEscape(id)
	if err := c.Do(ctx, http.MethodGet, path, nil, &obj); err != nil {
		return nil, err
	}
	return obj, nil
}

// GraphSearchRequest bounds a walk of the graph from Source.
type GraphSearchRequest struct {
	Source      string   `json:"source"`
	Graph       string   `json:"graph"`
	Direction   string   `json:"direction"`
	MinHops     int      `json:"min_hops"`
	MaxHops     int      `json:"max_hops"`
	LinkTypes   []string `json:"link_types"`
	TargetTypes []string `json:"target_types"`
	Count       int      `json:"count"`
	Page        int      `json:"page"`
	// Always false: the source is the caller's own reference.
	IncludeOriginal bool `json:"include_original"`
}

// GraphResponse is the paths found from the source, and their vertices.
type GraphResponse struct {
	Vertices map[string]Object `json:"vertices"`
	Paths    [][]Edge          `json:"paths"`
	Total    int               `json:"total"`
}

// GraphSearch walks the graph from a source object.
func (c *Client) GraphSearch(ctx context.Context, req GraphSearchRequest) (*GraphResponse, error) {
	var resp GraphResponse
	if err := c.Do(ctx, http.MethodPost, "/graph/search", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// YaraBundle compiles YARA rules, given by id or tag, with their dependencies
// into one bundle.
func (c *Client) YaraBundle(ctx context.Context, ids, tags, excludeTags []string) (string, error) {
	body := map[string]any{"ids": nonNil(ids), "tags": nonNil(tags), "exclude_tags": nonNil(excludeTags)}
	var resp struct {
		Bundle string `json:"bundle"`
	}
	if err := c.Do(ctx, http.MethodPost, "/indicators/yara/bundle", body, &resp); err != nil {
		return "", err
	}
	return resp.Bundle, nil
}

// Types lists the object types the server knows, by family.
func (c *Client) Types(ctx context.Context) (map[string][]string, error) {
	var resp map[string][]struct {
		Type string `json:"type"`
	}
	if err := c.Do(ctx, http.MethodGet, "/system/types", nil, &resp); err != nil {
		return nil, err
	}
	types := make(map[string][]string, len(resp))
	for family, entries := range resp {
		for _, entry := range entries {
			types[family] = append(types[family], entry.Type)
		}
	}
	return types, nil
}

func nonNil(values []string) []string {
	if values == nil {
		return []string{}
	}
	return values
}

package api

import (
	"context"
	"encoding/json"
	"errors"
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

// AddResponse is what Yeti added from a list of observables, and the lines it
// couldn't read as one.
type AddResponse struct {
	Added  []Object `json:"added"`
	Failed []string `json:"failed"`
}

// ImportText adds the observables in text, one per line, guessing their
// types. Existing observables are returned as added, unchanged.
func (c *Client) ImportText(ctx context.Context, text string, tags []string) (*AddResponse, error) {
	body := map[string]any{"text": text, "tags": nonNil(tags)}
	var resp AddResponse
	if err := c.Do(ctx, http.MethodPost, "/observables/import/text", body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// BulkAdd adds observables of one type. Yeti doesn't check that the values
// are of that type.
func (c *Client) BulkAdd(ctx context.Context, values []string, typ string, tags []string) (*AddResponse, error) {
	observables := make([]map[string]any, 0, len(values))
	for _, value := range values {
		observables = append(observables, map[string]any{"value": value, "type": typ, "tags": nonNil(tags)})
	}
	var resp AddResponse
	if err := c.Do(ctx, http.MethodPost, "/observables/bulk", map[string]any{"observables": observables}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// FindByName returns the entity or indicator ("entities" or "indicators") of
// a type and name, or nil when there is none. It reads the collection itself,
// so it sees objects as soon as they're saved, unlike searches.
func (c *Client) FindByName(ctx context.Context, collection, name, typ string) (Object, error) {
	query := url.Values{"name": {name}, "type": {typ}}
	var obj Object
	err := c.Do(ctx, http.MethodGet, "/"+collection+"/?"+query.Encode(), nil, &obj)
	var status *StatusError
	if errors.As(err, &status) && status.Status == http.StatusNotFound {
		return nil, nil
	}
	return obj, err
}

// Create creates an entity or an indicator: body is the request, e.g.
// {"entity": {...}, "tags": [...]}.
func (c *Client) Create(ctx context.Context, collection string, body map[string]any) (Object, error) {
	var obj Object
	if err := c.Do(ctx, http.MethodPost, "/"+collection+"/", body, &obj); err != nil {
		return nil, err
	}
	return obj, nil
}

// Tag adds tags to objects of one collection or, with strict, replaces their
// tags.
func (c *Client) Tag(ctx context.Context, collection string, ids, tags []string, strict bool) error {
	body := map[string]any{"ids": ids, "tags": nonNil(tags), "strict": strict}
	return c.Do(ctx, http.MethodPost, "/"+collection+"/tag", body, nil)
}

// Link links source to target, both "<collection>/<id>". Linking again only
// counts the link once more.
func (c *Client) Link(ctx context.Context, source, linkType, target, description string) (*Edge, error) {
	body := map[string]any{"source": source, "target": target, "link_type": linkType, "description": description}
	var edge Edge
	if err := c.Do(ctx, http.MethodPost, "/graph/add", body, &edge); err != nil {
		return nil, err
	}
	return &edge, nil
}

// AddContext adds context from source to an object. Yeti keeps one copy of
// identical context from the same source.
func (c *Client) AddContext(ctx context.Context, collection, id, source string, context map[string]any) (Object, error) {
	var obj Object
	body := map[string]any{"source": source, "context": context}
	if err := c.Do(ctx, http.MethodPost, "/"+collection+"/"+url.PathEscape(id)+"/context", body, &obj); err != nil {
		return nil, err
	}
	return obj, nil
}

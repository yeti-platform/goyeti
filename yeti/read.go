package yeti

import (
	"context"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/yeti-platform/goyeti/internal/api"
)

// Linked is an object linked to another, and how.
type Linked struct {
	Summary
	// Link is the relationship's type, e.g. "c2" or "uses".
	Link string `json:"link"`
	// Direction is "out" when the relationship points from the object it
	// was found from, "in" when it points to it.
	Direction string `json:"direction"`
}

// LookupResult is what Yeti knows about one value.
type LookupResult struct {
	Value string `json:"value"`
	// Normalized is the value that was looked up, when it differs: defanged
	// values are refanged.
	Normalized     string    `json:"normalized,omitempty"`
	Known          bool      `json:"known"`
	Ref            string    `json:"ref,omitempty"`
	Type           string    `json:"type,omitempty"`
	Tags           []string  `json:"tags,omitempty"`
	ContextSources []string  `json:"context_sources,omitempty"`
	Entities       []Linked  `json:"entities,omitempty"`
	Observables    []Linked  `json:"observables,omitempty"`
	Indicators     []Summary `json:"indicators,omitempty"`
	// Omitted counts, per list, the items left out by the limit.
	Omitted map[string]int `json:"omitted,omitempty"`
}

var refang = strings.NewReplacer(
	"[.]", ".", "(.)", ".", "{.}", ".", "[dot]", ".", "(dot)", ".",
	"[:]", ":", "[://]", "://", "hxxp", "http", "hXXp", "http", "HXXP", "HTTP",
)

// Lookup says, for each value, whether Yeti knows it, its tags and context
// sources, the entities and observables linked to it, and the indicators that
// match it. Each list keeps at most limit items. Unknown values are reported,
// never added.
func (c *Client) Lookup(ctx context.Context, values []string, limit int) ([]LookupResult, Meta, error) {
	results := []LookupResult{}
	byValue := map[string]int{} // normalized value -> index in results
	var query []string
	for _, value := range values {
		value = strings.TrimSpace(value)
		normalized := refang.Replace(value)
		if value == "" {
			continue
		}
		if _, seen := byValue[normalized]; seen {
			continue
		}
		byValue[normalized] = len(results)
		result := LookupResult{Value: value}
		if normalized != value {
			result.Normalized = normalized
		}
		results = append(results, result)
		query = append(query, normalized)
	}
	if len(query) == 0 {
		return nil, Meta{}, usage("No values to look up.", "Pass one or more values, e.g. IPs, hostnames, URLs or hashes.")
	}
	limit = clamp(limit, 10, 100)

	resp, err := c.api.Match(ctx, query)
	if err != nil {
		return nil, Meta{}, wrap(err, "")
	}

	byRef := map[string]int{} // reference of a known observable -> index in results
	for _, obj := range resp.Known {
		i, ok := byValue[str(obj["value"])]
		if !ok {
			continue
		}
		r := &results[i]
		r.Known = true
		r.Ref = ref(obj)
		r.Type = str(obj["type"])
		r.Tags = tagNames(obj["tags"])
		r.ContextSources = contextSources(obj["context"])
		byRef[r.Ref] = i
	}
	attach := func(linked []api.Linked, list func(*LookupResult) *[]Linked) {
		for _, l := range linked {
			from, direction := l.Edge.Source, "out"
			if _, ok := byRef[from]; !ok {
				from, direction = l.Edge.Target, "in"
			}
			i, ok := byRef[from]
			if !ok {
				continue
			}
			*list(&results[i]) = append(*list(&results[i]), Linked{Summary: summarize(l.Object), Link: l.Edge.Type, Direction: direction})
		}
	}
	attach(resp.Entities, func(r *LookupResult) *[]Linked { return &r.Entities })
	attach(resp.Observables, func(r *LookupResult) *[]Linked { return &r.Observables })
	for _, m := range resp.Matches {
		if i, ok := byValue[m.Value]; ok {
			results[i].Indicators = append(results[i].Indicators, summarize(m.Indicator))
		}
	}

	for i := range results {
		r := &results[i]
		sortLinked(r.Entities)
		sortLinked(r.Observables)
		sort.Slice(r.Indicators, func(a, b int) bool { return r.Indicators[a].Ref < r.Indicators[b].Ref })
		r.Entities = capList(r, "entities", r.Entities, limit)
		r.Observables = capList(r, "observables", r.Observables, limit)
		r.Indicators = capList(r, "indicators", r.Indicators, limit)
	}
	return results, Meta{}, nil
}

func sortLinked(items []Linked) {
	sort.Slice(items, func(a, b int) bool {
		if items[a].Link != items[b].Link {
			return items[a].Link < items[b].Link
		}
		return items[a].Ref < items[b].Ref
	})
}

func capList[T any](r *LookupResult, name string, items []T, limit int) []T {
	if len(items) <= limit {
		return items
	}
	if r.Omitted == nil {
		r.Omitted = map[string]int{}
	}
	r.Omitted[name] = len(items) - limit
	return items[:limit]
}

// SearchHit is an object a search found.
type SearchHit struct {
	Summary
	// FoundBy is "keyword", "semantic" or both.
	FoundBy []string `json:"found_by"`
	// Score is the semantic similarity, from 0 to 1, when found by meaning.
	Score float64 `json:"score,omitempty"`
}

// SearchResult is what a search found.
type SearchResult struct {
	Query string      `json:"query"`
	Hits  []SearchHit `json:"hits"`
}

// Search finds objects by keyword, across every family, and entities,
// indicators and DFIQ by meaning. Hits found by meaning come first, best
// first. At most limit hits are returned.
func (c *Client) Search(ctx context.Context, query string, limit int) (*SearchResult, Meta, error) {
	query = strings.TrimSpace(query)
	if query == "" {
		return nil, Meta{}, usage("The search query is empty.", "Pass a word, a name or a phrase to search for.")
	}
	limit = clamp(limit, 20, 50)
	var meta Meta

	keyword, keywordErr := c.api.Search(ctx, query, limit)
	semantic, semanticErr := c.api.SemanticSearch(ctx, query, limit)
	if keywordErr != nil && semanticErr != nil {
		return nil, Meta{}, wrap(keywordErr, "")
	}
	if keywordErr != nil {
		meta.Warnings = append(meta.Warnings, "Keyword search failed, so only semantic matches are shown: "+wrap(keywordErr, "").Error())
	}
	if semanticErr != nil {
		meta.Warnings = append(meta.Warnings, "Semantic search is unavailable on this server, so only keyword matches are shown: "+wrap(semanticErr, "").Error())
	}

	hits := []SearchHit{}
	index := map[string]int{}
	var semanticHits []SearchHit
	for _, section := range semantic {
		for _, r := range section.Results {
			s := r.Summary
			semanticHits = append(semanticHits, SearchHit{
				Summary: Summary{Ref: collections[s.RootType] + "/" + s.ID, Type: s.Type, Name: s.Name,
					Tags: sortedCopy(s.Tags), Description: shorten(s.Description)},
				FoundBy: []string{"semantic"},
				Score:   r.Score,
			})
		}
	}
	sort.SliceStable(semanticHits, func(a, b int) bool { return semanticHits[a].Score > semanticHits[b].Score })
	for _, hit := range semanticHits {
		if _, seen := index[hit.Ref]; !seen {
			index[hit.Ref] = len(hits)
			hits = append(hits, hit)
		}
	}
	omitted := 0
	for _, section := range keyword {
		omitted += section.Total - len(section.Results)
		for _, obj := range section.Results {
			summary := summarize(obj)
			if i, seen := index[summary.Ref]; seen {
				hits[i].FoundBy = append(hits[i].FoundBy, "keyword")
				continue
			}
			index[summary.Ref] = len(hits)
			hits = append(hits, SearchHit{Summary: summary, FoundBy: []string{"keyword"}})
		}
	}
	if len(hits) > limit {
		omitted += len(hits) - limit
		hits = hits[:limit]
	}
	if omitted > 0 {
		meta.Truncated = &Truncated{Omitted: omitted,
			Hint: "Raise the limit, or search for something more specific."}
	}
	return &SearchResult{Query: query, Hits: hits}, meta, nil
}

// Get returns one object, with every field Yeti stores, and its reference.
func (c *Client) Get(ctx context.Context, reference string) (map[string]any, Meta, error) {
	collection, id, err := parseRef(reference)
	if err != nil {
		return nil, Meta{}, err
	}
	obj, err := c.api.Get(ctx, collection, id)
	if err != nil {
		return nil, Meta{}, wrap(err, fmt.Sprintf("Nothing exists at %s. %s", reference, refHint))
	}
	delete(obj, "acls")
	obj["ref"] = ref(obj)
	return obj, Meta{}, nil
}

// NeighborsOptions bounds a walk of the graph.
type NeighborsOptions struct {
	// Hops is how far to walk, from 1 (direct links, the default) to 3.
	Hops int
	// Direction is "any" (the default), "out" or "in".
	Direction string
	// LinkTypes keeps only relationships of these types, if set.
	LinkTypes []string
	// TargetTypes keeps only neighbors of these object types or families, if
	// set.
	TargetTypes []string
	// Limit is the most paths to return.
	Limit int
}

// Neighbor is an object reached from the source.
type Neighbor struct {
	Linked
	// Via lists the objects in between, for neighbors more than one hop away.
	Via []string `json:"via,omitempty"`
}

// NeighborsResult is what was found around one object.
type NeighborsResult struct {
	Source    string     `json:"source"`
	Neighbors []Neighbor `json:"neighbors"`
}

var directions = map[string]string{"any": "any", "out": "outbound", "in": "inbound"}

// Neighbors walks the graph from one object and returns the objects it
// reaches, each once.
func (c *Client) Neighbors(ctx context.Context, reference string, opts NeighborsOptions) (*NeighborsResult, Meta, error) {
	if _, _, err := parseRef(reference); err != nil {
		return nil, Meta{}, err
	}
	source := strings.TrimSpace(reference)
	if opts.Direction == "" {
		opts.Direction = "any"
	}
	direction, ok := directions[opts.Direction]
	if !ok {
		return nil, Meta{}, usage(fmt.Sprintf("%q is not a direction.", opts.Direction), `Use "any", "out" or "in".`)
	}
	if opts.Hops == 0 {
		opts.Hops = 1
	}
	if opts.Hops < 1 || opts.Hops > 3 {
		return nil, Meta{}, usage(fmt.Sprintf("%d hops is out of range.", opts.Hops), "Walk 1 to 3 hops; start with 1.")
	}
	limit := clamp(opts.Limit, 20, 100)

	resp, err := c.api.GraphSearch(ctx, api.GraphSearchRequest{
		Source: source, Graph: "links", Direction: direction,
		MinHops: 1, MaxHops: opts.Hops,
		LinkTypes: nonNil(opts.LinkTypes), TargetTypes: nonNil(opts.TargetTypes),
		Count: limit,
	})
	if err != nil {
		return nil, Meta{}, wrap(err, fmt.Sprintf("Nothing exists at %s. %s", reference, refHint))
	}

	result := &NeighborsResult{Source: source, Neighbors: []Neighbor{}}
	seen := map[string]bool{source: true}
	for _, path := range resp.Paths {
		current, linkDirection := source, ""
		var via []string
		for i, edge := range path {
			next := edge.Target
			linkDirection = "out"
			if edge.Source != current {
				next, linkDirection = edge.Source, "in"
			}
			if i < len(path)-1 {
				via = append(via, next)
			}
			current = next
		}
		if len(path) == 0 || seen[current] {
			continue
		}
		seen[current] = true
		summary := Summary{Ref: current}
		if obj, ok := resp.Vertices[current]; ok {
			summary = summarize(obj)
		}
		last := path[len(path)-1]
		result.Neighbors = append(result.Neighbors, Neighbor{
			Linked: Linked{Summary: summary, Link: last.Type, Direction: linkDirection},
			Via:    via,
		})
	}
	sort.SliceStable(result.Neighbors, func(a, b int) bool {
		x, y := result.Neighbors[a], result.Neighbors[b]
		if len(x.Via) != len(y.Via) {
			return len(x.Via) < len(y.Via)
		}
		if x.Link != y.Link {
			return x.Link < y.Link
		}
		return x.Ref < y.Ref
	})
	var meta Meta
	if omitted := resp.Total - len(resp.Paths); omitted > 0 {
		meta.Truncated = &Truncated{Omitted: omitted,
			Hint: "Paths were left out: raise the limit, or narrow by link type, target type or direction."}
	}
	return result, meta, nil
}

// YaraBundle is a set of YARA rules ready to compile.
type YaraBundle struct {
	// Rules is how many rules the bundle holds, dependencies included.
	Rules  int    `json:"rules"`
	Bundle string `json:"bundle"`
}

var yaraRule = regexp.MustCompile(`(?m)^\s*(?:(?:private|global)\s+)*rule\s`)

// YaraBundle builds one YARA source from rules given by reference
// ("indicators/<id>") or by tag, with every rule they depend on. Rules with
// an excluded tag are left out.
func (c *Client) YaraBundle(ctx context.Context, refs, tags, excludeTags []string) (*YaraBundle, Meta, error) {
	if len(refs) == 0 && len(tags) == 0 {
		return nil, Meta{}, usage("No rules were asked for.", "Pass YARA rule references (indicators/<id>), tags, or both.")
	}
	ids := make([]string, 0, len(refs))
	for _, r := range refs {
		collection, id, err := parseRef(r)
		if err != nil {
			return nil, Meta{}, err
		}
		if collection != "indicators" {
			return nil, Meta{}, usage(fmt.Sprintf("%s is not an indicator.", r), "YARA rules are indicators: pass references like indicators/123.")
		}
		ids = append(ids, id)
	}
	bundle, err := c.api.YaraBundle(ctx, ids, tags, excludeTags)
	if err != nil {
		return nil, Meta{}, wrap(err, "A rule doesn't exist, or isn't a YARA rule. Find YARA rules with \"yeti read search\".")
	}
	result := &YaraBundle{Rules: len(yaraRule.FindAllStringIndex(bundle, -1)), Bundle: bundle}
	var meta Meta
	if result.Rules == 0 {
		meta.Warnings = append(meta.Warnings, "No rules matched: check the references and tags.")
	}
	return result, meta, nil
}

// Types lists the object types the server knows, by family: "observables",
// "entities", "indicators" and "dfiq".
func (c *Client) Types(ctx context.Context) (map[string][]string, Meta, error) {
	types, err := c.api.Types(ctx)
	if err != nil {
		return nil, Meta{}, wrap(err, "")
	}
	for _, names := range types {
		sort.Strings(names)
	}
	return types, Meta{}, nil
}

// clamp returns n, or fallback when n isn't set, capped at most.
func clamp(n, fallback, most int) int {
	if n <= 0 {
		return fallback
	}
	return min(n, most)
}

func nonNil(values []string) []string {
	if values == nil {
		return []string{}
	}
	return values
}

func sortedCopy(values []string) []string {
	if len(values) == 0 {
		return nil
	}
	sorted := append([]string(nil), values...)
	sort.Strings(sorted)
	return sorted
}

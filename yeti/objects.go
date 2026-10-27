package yeti

import (
	"fmt"
	"sort"
	"strings"

	"github.com/yeti-platform/goyeti/internal/api"
)

// collections maps Yeti's root types to the collections in references.
var collections = map[string]string{
	"observable": "observables",
	"entity":     "entities",
	"indicator":  "indicators",
	"dfiq":       "dfiq",
}

const refHint = `References look like "observables/123", "entities/45", "indicators/67" or "dfiq/89"; find them with "yeti read lookup" or "yeti read search".`

// parseRef splits "<collection>/<id>".
func parseRef(ref string) (collection, id string, err error) {
	collection, id, ok := strings.Cut(strings.TrimSpace(ref), "/")
	valid := false
	for _, c := range collections {
		valid = valid || c == collection
	}
	if !ok || !valid || id == "" || strings.Contains(id, "/") {
		return "", "", usage(fmt.Sprintf("%q is not a Yeti reference.", ref), refHint)
	}
	return collection, id, nil
}

// Summary is the short form of an object in list results. Get returns the
// whole object.
type Summary struct {
	Ref  string `json:"ref"`
	Type string `json:"type"`
	// Value is set for observables, Name for other objects.
	Value       string   `json:"value,omitempty"`
	Name        string   `json:"name,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Description string   `json:"description,omitempty"`
}

const descriptionLength = 160

func summarize(obj api.Object) Summary {
	return Summary{
		Ref:         ref(obj),
		Type:        str(obj["type"]),
		Value:       str(obj["value"]),
		Name:        str(obj["name"]),
		Tags:        tagNames(obj["tags"]),
		Description: shorten(str(obj["description"])),
	}
}

func ref(obj api.Object) string {
	collection := collections[str(obj["root_type"])]
	if collection == "" {
		collection = str(obj["root_type"])
	}
	return collection + "/" + str(obj["id"])
}

func str(v any) string {
	s, _ := v.(string)
	return s
}

// tagNames reads tag names from objects' embedded tag instances, or from a
// list of names.
func tagNames(v any) []string {
	items, _ := v.([]any)
	names := make([]string, 0, len(items))
	for _, item := range items {
		switch tag := item.(type) {
		case string:
			names = append(names, tag)
		case map[string]any:
			if name := str(tag["name"]); name != "" {
				names = append(names, name)
			}
		}
	}
	sort.Strings(names)
	if len(names) == 0 {
		return nil
	}
	return names
}

// contextSources lists the distinct sources of an object's context.
func contextSources(v any) []string {
	items, _ := v.([]any)
	seen := map[string]bool{}
	var sources []string
	for _, item := range items {
		entry, _ := item.(map[string]any)
		if source := str(entry["source"]); source != "" && !seen[source] {
			seen[source] = true
			sources = append(sources, source)
		}
	}
	sort.Strings(sources)
	return sources
}

// shorten keeps a description's first line, cut at descriptionLength.
func shorten(text string) string {
	line, _, cut := strings.Cut(strings.TrimSpace(text), "\n")
	runes := []rune(line)
	if len(runes) > descriptionLength {
		return string(runes[:descriptionLength-1]) + "…"
	}
	if cut {
		return line + " …"
	}
	return line
}

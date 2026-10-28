package yeti

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/yeti-platform/goyeti/internal/api"
)

// Actions a mutate operation reports for each item. In a dry run they say what
// would happen.
const (
	ActionCreated   = "created"
	ActionExists    = "exists"
	ActionLinked    = "linked"
	ActionTagged    = "tagged"
	ActionUntagged  = "untagged"
	ActionUnchanged = "unchanged"
	ActionAdded     = "context_added"
	ActionFailed    = "failed"
)

// Change is what happened, or would happen, to one item.
type Change struct {
	// Item is the value, name or reference the caller passed.
	Item   string `json:"item"`
	Ref    string `json:"ref,omitempty"`
	Type   string `json:"type,omitempty"`
	Action string `json:"action"`
	// Tags are the object's tags after the change, for tag operations.
	Tags  []string `json:"tags,omitempty"`
	Error *Error   `json:"error,omitempty"`
}

// Changes is a mutate operation's outcome, item by item.
type Changes struct {
	// DryRun is set when nothing was written: the actions say what would
	// happen.
	DryRun  bool     `json:"dry_run,omitempty"`
	Changes []Change `json:"changes"`
}

func (c *Changes) add(change Change) { c.Changes = append(c.Changes, change) }

// meta counts failed changes.
func (c *Changes) meta() Meta {
	var meta Meta
	for _, change := range c.Changes {
		if change.Action == ActionFailed {
			meta.Failed++
		}
	}
	if meta.Failed > 0 {
		meta.Warnings = append(meta.Warnings, fmt.Sprintf("%d of %d items failed; each says why.", meta.Failed, len(c.Changes)))
	}
	return meta
}

func failed(item string, err error) Change {
	var yetiErr *Error
	errors.As(wrap(err, ""), &yetiErr)
	return Change{Item: item, Action: ActionFailed, Error: yetiErr}
}

// AddObservablesOptions describes observables to add.
type AddObservablesOptions struct {
	// Type of every value. Empty lets Yeti guess each value's type, which
	// also rejects values it doesn't recognize; Yeti doesn't check values
	// against an explicit type.
	Type string
	// Tags to add to every value, new or existing.
	Tags   []string
	DryRun bool
}

// AddObservables creates the values Yeti doesn't know yet, leaves the others
// as they are, and tags them all.
func (c *Client) AddObservables(ctx context.Context, values []string, opts AddObservablesOptions) (*Changes, Meta, error) {
	var query []string
	items := map[string]string{} // normalized value -> value as passed
	for _, value := range values {
		value = strings.TrimSpace(value)
		normalized := refang.Replace(value)
		if value == "" || items[normalized] != "" {
			continue
		}
		items[normalized] = value
		query = append(query, normalized)
	}
	if len(query) == 0 {
		return nil, Meta{}, usage("No values to add.", "Pass one or more observable values.")
	}

	known, err := c.knownObservables(ctx, query)
	if err != nil {
		return nil, Meta{}, err
	}
	var unknown []string
	for _, value := range query {
		if _, ok := known[value]; !ok {
			unknown = append(unknown, value)
		}
	}

	result := &Changes{DryRun: opts.DryRun}
	if opts.DryRun {
		for _, value := range query {
			if obj, ok := known[value]; ok {
				result.add(Change{Item: items[value], Ref: ref(obj), Type: str(obj["type"]), Action: ActionExists})
			} else {
				result.add(Change{Item: items[value], Type: opts.Type, Action: ActionCreated})
			}
		}
		return result, result.meta(), nil
	}

	if len(unknown) > 0 {
		if opts.Type == "" {
			_, err = c.api.ImportText(ctx, strings.Join(unknown, "\n"), opts.Tags)
		} else {
			_, err = c.api.BulkAdd(ctx, unknown, opts.Type, opts.Tags)
		}
		if err != nil {
			return nil, Meta{}, wrap(err, "")
		}
	}
	// Existing values get the tags too.
	if len(opts.Tags) > 0 && len(known) > 0 {
		var ids []string
		for _, obj := range known {
			ids = append(ids, str(obj["id"]))
		}
		slices.Sort(ids)
		if err := c.api.Tag(ctx, "observables", ids, opts.Tags, false); err != nil {
			return nil, Meta{}, wrap(err, "")
		}
	}
	// Ask Yeti again what it knows: the add responses don't tell new values
	// from existing ones, and leave out values they couldn't read.
	now, err := c.knownObservables(ctx, query)
	if err != nil {
		return nil, Meta{}, err
	}
	for _, value := range query {
		obj, isKnown := now[value]
		_, wasKnown := known[value]
		switch {
		case wasKnown:
			result.add(Change{Item: items[value], Ref: ref(obj), Type: str(obj["type"]), Action: ActionExists})
		case isKnown:
			result.add(Change{Item: items[value], Ref: ref(obj), Type: str(obj["type"]), Action: ActionCreated})
		default:
			result.add(Change{Item: items[value], Action: ActionFailed, Error: &Error{Code: CodeInvalidRequest,
				Message: fmt.Sprintf("Yeti didn't recognize %q as an observable.", items[value]),
				Hint:    `Check the value, or pass its type, e.g. --type path; "yeti read types" lists them.`}})
		}
	}
	return result, result.meta(), nil
}

func (c *Client) knownObservables(ctx context.Context, values []string) (map[string]api.Object, error) {
	resp, err := c.api.Match(ctx, values)
	if err != nil {
		return nil, wrap(err, "")
	}
	known := make(map[string]api.Object, len(resp.Known))
	for _, obj := range resp.Known {
		known[str(obj["value"])] = obj
	}
	return known, nil
}

// AddObjectOptions describes an entity or indicator to add.
type AddObjectOptions struct {
	// Fields are the object's other fields, e.g. "description", "aliases",
	// or an indicator's "pattern", "diamond" and "location".
	Fields map[string]any
	// Tags to add, whether the object is new or existing.
	Tags   []string
	DryRun bool
}

// AddEntity creates an entity of a type and name, unless one exists, and tags
// it. An existing entity's fields are left as they are.
func (c *Client) AddEntity(ctx context.Context, typ, name string, opts AddObjectOptions) (*Changes, Meta, error) {
	return c.addObject(ctx, "entities", "entity", typ, name, opts)
}

// AddIndicator creates an indicator of a type and name, unless one exists, and
// tags it. Indicators need a "pattern" and a "diamond" field, and some types
// a "location". A YARA rule's name is the rule's own name.
func (c *Client) AddIndicator(ctx context.Context, typ, name string, opts AddObjectOptions) (*Changes, Meta, error) {
	return c.addObject(ctx, "indicators", "indicator", typ, name, opts)
}

func (c *Client) addObject(ctx context.Context, collection, family, typ, name string, opts AddObjectOptions) (*Changes, Meta, error) {
	typ, name = strings.TrimSpace(typ), strings.TrimSpace(name)
	if typ == "" || name == "" {
		return nil, Meta{}, usage(fmt.Sprintf("An %s needs a type and a name.", family),
			fmt.Sprintf(`"yeti read types" lists the %s types.`, family))
	}
	// Creating an existing object would overwrite its fields, so look for it
	// first.
	existing, err := c.api.FindByName(ctx, collection, name, typ)
	if err != nil {
		return nil, Meta{}, wrap(err, "")
	}
	result := &Changes{DryRun: opts.DryRun}
	if existing != nil {
		if !opts.DryRun && len(opts.Tags) > 0 {
			if err := c.api.Tag(ctx, collection, []string{str(existing["id"])}, opts.Tags, false); err != nil {
				return nil, Meta{}, wrap(err, "")
			}
		}
		result.add(Change{Item: name, Ref: ref(existing), Type: typ, Action: ActionExists})
		return result, result.meta(), nil
	}
	if opts.DryRun {
		result.add(Change{Item: name, Type: typ, Action: ActionCreated})
		return result, result.meta(), nil
	}

	fields := map[string]any{}
	for k, v := range opts.Fields {
		fields[k] = v
	}
	fields["type"], fields["name"] = typ, name
	body := map[string]any{family: fields}
	if family == "entity" {
		body["tags"] = nonNil(opts.Tags)
	}
	obj, err := c.api.Create(ctx, collection, body)
	if err != nil {
		return nil, Meta{}, wrap(err, "")
	}
	if family == "indicator" && len(opts.Tags) > 0 {
		if err := c.api.Tag(ctx, collection, []string{str(obj["id"])}, opts.Tags, false); err != nil {
			return nil, Meta{}, wrap(err, "")
		}
	}
	result.add(Change{Item: name, Ref: ref(obj), Type: typ, Action: ActionCreated})
	if created := str(obj["name"]); created != name {
		result.Changes[0].Item = name + " (saved as " + created + ")"
	}
	return result, result.meta(), nil
}

// Link links source to target with a relationship type, e.g.
// "observables/1" -c2-> "entities/2", unless they are already linked that
// way.
func (c *Client) Link(ctx context.Context, source, linkType, target, description string, dryRun bool) (*Changes, Meta, error) {
	linkType = strings.TrimSpace(linkType)
	if linkType == "" {
		return nil, Meta{}, usage("The link has no type.", `Name the relationship, e.g. "uses", "c2" or "delivers".`)
	}
	for _, r := range []string{source, target} {
		if _, _, err := parseRef(r); err != nil {
			return nil, Meta{}, err
		}
	}
	source, target = strings.TrimSpace(source), strings.TrimSpace(target)
	item := source + " -" + linkType + "-> " + target

	paths, err := c.api.GraphSearch(ctx, api.GraphSearchRequest{
		Source: source, Graph: "links", Direction: "outbound", MinHops: 1, MaxHops: 1,
		LinkTypes: []string{linkType}, TargetTypes: []string{}, Count: 1000,
	})
	if err != nil {
		return nil, Meta{}, wrap(err, fmt.Sprintf("Nothing exists at %s. %s", source, refHint))
	}
	result := &Changes{DryRun: dryRun}
	for _, path := range paths.Paths {
		if len(path) == 1 && path[0].Target == target {
			result.add(Change{Item: item, Ref: "links/" + path[0].ID, Type: linkType, Action: ActionExists})
			return result, result.meta(), nil
		}
	}
	if dryRun {
		if _, _, err := c.Get(ctx, target); err != nil {
			return nil, Meta{}, err
		}
		result.add(Change{Item: item, Type: linkType, Action: ActionLinked})
		return result, result.meta(), nil
	}
	edge, err := c.api.Link(ctx, source, linkType, target, description)
	if err != nil {
		return nil, Meta{}, wrap(err, refHint)
	}
	result.add(Change{Item: item, Ref: "links/" + edge.ID, Type: linkType, Action: ActionLinked})
	return result, result.meta(), nil
}

// taggable are the collections whose objects have tags and context.
var taggable = map[string]bool{"observables": true, "entities": true, "indicators": true}

var tagSpaces = regexp.MustCompile(`\s+`)

// tagName approximates Yeti's tag normalization, to compare tags with an
// object's.
func tagName(tag string) string {
	return tagSpaces.ReplaceAllString(strings.ToLower(strings.TrimSpace(tag)), "_")
}

// Tag adds tags to objects. Objects that already have them are left
// unchanged.
func (c *Client) Tag(ctx context.Context, refs, tags []string, dryRun bool) (*Changes, Meta, error) {
	return c.retag(ctx, refs, tags, dryRun, false)
}

// Untag removes tags from objects.
func (c *Client) Untag(ctx context.Context, refs, tags []string, dryRun bool) (*Changes, Meta, error) {
	return c.retag(ctx, refs, tags, dryRun, true)
}

func (c *Client) retag(ctx context.Context, refs, tags []string, dryRun, remove bool) (*Changes, Meta, error) {
	if len(refs) == 0 {
		return nil, Meta{}, usage("No objects to tag.", "Pass the references of the objects, e.g. observables/123.")
	}
	var names []string
	for _, tag := range tags {
		if name := tagName(tag); name != "" && !slices.Contains(names, name) {
			names = append(names, name)
		}
	}
	if len(names) == 0 {
		return nil, Meta{}, usage("No tags were given.", "Pass one or more tags.")
	}
	result := &Changes{DryRun: dryRun}
	for _, r := range refs {
		collection, id, err := parseRef(r)
		if err != nil {
			return nil, Meta{}, err
		}
		if !taggable[collection] {
			result.add(Change{Item: r, Action: ActionFailed, Error: usage(fmt.Sprintf("%s can't be tagged.", collection),
				"Only observables, entities and indicators have tags.")})
			continue
		}
		obj, _, err := c.Get(ctx, r)
		if err != nil {
			result.add(failed(r, err))
			continue
		}
		current := tagNames(obj["tags"])
		var after []string
		action := ActionUnchanged
		if remove {
			for _, tag := range current {
				if slices.Contains(names, tag) {
					action = ActionUntagged
				} else {
					after = append(after, tag)
				}
			}
		} else {
			after = slices.Clone(current)
			for _, name := range names {
				if !slices.Contains(after, name) {
					after = append(after, name)
					action = ActionTagged
				}
			}
			slices.Sort(after)
		}
		if action != ActionUnchanged && !dryRun {
			// Removing replaces the object's tags with the ones it keeps.
			tagsToSend, strict := names, false
			if remove {
				tagsToSend, strict = after, true
			}
			if err := c.api.Tag(ctx, collection, []string{id}, tagsToSend, strict); err != nil {
				result.add(failed(r, err))
				continue
			}
		}
		result.add(Change{Item: r, Ref: r, Type: str(obj["type"]), Action: action, Tags: after})
	}
	return result, result.meta(), nil
}

// AddContext adds context to an object, from the source "agent:<name>", the
// agent's name being Config.AgentName. Yeti keeps one copy of identical
// context from the same source.
func (c *Client) AddContext(ctx context.Context, reference string, context map[string]any, dryRun bool) (*Changes, Meta, error) {
	collection, id, err := parseRef(reference)
	if err != nil {
		return nil, Meta{}, err
	}
	if !taggable[collection] {
		return nil, Meta{}, usage(fmt.Sprintf("%s have no context.", collection), "Only observables, entities and indicators have context.")
	}
	if len(context) == 0 {
		return nil, Meta{}, usage("The context is empty.", `Pass fields, e.g. note="seen in phishing campaign X".`)
	}
	reference = strings.TrimSpace(reference)
	source := "agent:" + c.agent
	result := &Changes{DryRun: dryRun}
	if dryRun {
		obj, _, err := c.Get(ctx, reference)
		if err != nil {
			return nil, Meta{}, err
		}
		result.add(Change{Item: reference, Ref: reference, Type: str(obj["type"]), Action: ActionAdded})
		return result, result.meta(), nil
	}
	obj, err := c.api.AddContext(ctx, collection, id, source, context)
	if err != nil {
		return nil, Meta{}, wrap(err, fmt.Sprintf("Nothing exists at %s. %s", reference, refHint))
	}
	result.add(Change{Item: reference, Ref: reference, Type: str(obj["type"]), Action: ActionAdded})
	return result, result.meta(), nil
}

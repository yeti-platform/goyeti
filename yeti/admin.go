package yeti

import (
	"context"
	"fmt"
	"strings"

	"github.com/yeti-platform/goyeti/internal/api"
)

// RemoveOptions guard the operations that remove objects or links.
type RemoveOptions struct {
	// Confirm must be set to remove anything, unless DryRun is.
	Confirm bool
	DryRun  bool
}

func (o RemoveOptions) check(what string) error {
	if o.Confirm || o.DryRun {
		return nil
	}
	return usage(fmt.Sprintf("Deleting %s is permanent, and needs confirmation.", what),
		"Confirm with --yes (confirm: true over MCP), or do a dry run first to see what would be deleted.")
}

// Delete deletes objects, given by reference, with their links. Objects that
// don't exist are reported absent.
func (c *Client) Delete(ctx context.Context, refs []string, opts RemoveOptions) (*Changes, Meta, error) {
	if len(refs) == 0 {
		return nil, Meta{}, usage("No objects to delete.", "Pass the references of the objects, e.g. observables/123.")
	}
	for _, r := range refs {
		if _, _, err := parseRef(r); err != nil {
			return nil, Meta{}, err
		}
	}
	if err := opts.check("objects"); err != nil {
		return nil, Meta{}, err
	}
	result := &Changes{DryRun: opts.DryRun}
	for _, r := range refs {
		r = strings.TrimSpace(r)
		collection, id, _ := parseRef(r)
		obj, err := c.api.Get(ctx, collection, id)
		if isNotFound(err) {
			result.add(Change{Item: r, Ref: r, Action: ActionAbsent})
			continue
		}
		if err != nil {
			result.add(failed(r, err))
			continue
		}
		links, err := c.api.GraphSearch(ctx, api.GraphSearchRequest{
			Source: r, Graph: "links", Direction: "any", MinHops: 1, MaxHops: 1,
			LinkTypes: []string{}, TargetTypes: []string{}, Count: 1,
		})
		if err != nil {
			result.add(failed(r, err))
			continue
		}
		if !opts.DryRun {
			if err := c.api.Delete(ctx, collection, id); err != nil {
				result.add(failed(r, err))
				continue
			}
		}
		summary := summarize(obj)
		result.add(Change{Item: r, Ref: r, Type: summary.Type, Name: firstSet(summary.Value, summary.Name),
			Action: ActionDeleted, Links: links.Total})
	}
	return result, result.meta(), nil
}

// Unlink deletes the link of a type from source to target. A link that
// doesn't exist is reported absent.
func (c *Client) Unlink(ctx context.Context, source, linkType, target string, opts RemoveOptions) (*Changes, Meta, error) {
	linkType = strings.TrimSpace(linkType)
	if linkType == "" {
		return nil, Meta{}, usage("The link has no type.", `Name the relationship, e.g. "uses", "c2" or "delivers".`)
	}
	for _, r := range []string{source, target} {
		if _, _, err := parseRef(r); err != nil {
			return nil, Meta{}, err
		}
	}
	if err := opts.check("a link"); err != nil {
		return nil, Meta{}, err
	}
	source, target = strings.TrimSpace(source), strings.TrimSpace(target)
	item := source + " -" + linkType + "-> " + target
	existing, err := c.findLink(ctx, source, linkType, target)
	if err != nil {
		return nil, Meta{}, err
	}
	result := &Changes{DryRun: opts.DryRun}
	if existing == nil {
		result.add(Change{Item: item, Type: linkType, Action: ActionAbsent})
		return result, result.meta(), nil
	}
	if !opts.DryRun {
		if err := c.api.DeleteLink(ctx, existing.ID); err != nil {
			return nil, Meta{}, wrap(err, "")
		}
	}
	result.add(Change{Item: item, Ref: "links/" + existing.ID, Type: linkType, Action: ActionUnlinked})
	return result, result.meta(), nil
}

func firstSet(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}

package yeti

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/yeti-platform/goyeti/internal/api"
)

// Task is a feed, analytics, export or other task, and how its last run went.
type Task struct {
	Name          string `json:"name"`
	Type          string `json:"type"`
	Enabled       bool   `json:"enabled"`
	Status        string `json:"status"`
	StatusMessage string `json:"status_message,omitempty"`
	LastRun       string `json:"last_run,omitempty"`
	// Frequency is an ISO 8601 duration, e.g. "PT1H", for scheduled tasks.
	Frequency   string `json:"frequency,omitempty"`
	Description string `json:"description,omitempty"`
}

func taskOf(t api.Task) Task {
	return Task{Name: t.Name, Type: t.Type, Enabled: t.Enabled, Status: t.Status, StatusMessage: t.StatusMessage,
		LastRun: t.LastRun, Frequency: t.Frequency, Description: shorten(t.Description)}
}

// TasksOptions narrows the tasks to list.
type TasksOptions struct {
	// Names keeps the tasks whose name contains one of these, ignoring case.
	Names []string
	// Type keeps the tasks of one type: feed, analytics, export, oneshot,
	// event or log.
	Type string
}

// Tasks lists tasks by name.
func (c *Client) Tasks(ctx context.Context, opts TasksOptions) ([]Task, Meta, error) {
	names := opts.Names
	if len(names) == 0 {
		names = []string{""}
	}
	seen := map[string]bool{}
	tasks := []Task{}
	for _, name := range names {
		found, err := c.api.Tasks(ctx, strings.TrimSpace(name), opts.Type)
		if err != nil {
			return nil, Meta{}, wrap(err, "")
		}
		for _, t := range found {
			if !seen[t.Name] {
				seen[t.Name] = true
				tasks = append(tasks, taskOf(t))
			}
		}
	}
	slices.SortFunc(tasks, func(a, b Task) int { return strings.Compare(a.Name, b.Name) })
	return tasks, Meta{}, nil
}

// findTask returns the task of a name, matched exactly or else ignoring case.
func (c *Client) findTask(ctx context.Context, name string) (*api.Task, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, usage("No task was named.", `"yeti read tasks" lists the tasks.`)
	}
	found, err := c.api.Tasks(ctx, name, "")
	if err != nil {
		return nil, wrap(err, "")
	}
	var folded []api.Task
	for _, t := range found {
		if t.Name == name {
			return &t, nil
		}
		if strings.EqualFold(t.Name, name) {
			folded = append(folded, t)
		}
	}
	if len(folded) == 1 {
		return &folded[0], nil
	}
	hint := `"yeti read tasks" lists the tasks.`
	if len(found) > 0 {
		var similar []string
		for _, t := range found[:min(len(found), 5)] {
			similar = append(similar, t.Name)
		}
		hint = "Did you mean " + strings.Join(similar, ", ") + "? " + hint
	}
	return nil, &Error{Code: CodeNotFound, Message: fmt.Sprintf("No task is named %q.", name), Hint: hint}
}

// exclusive are the task types Yeti runs one at a time: a run queued while
// one is going is skipped.
var exclusive = map[string]bool{"feed": true, "analytics": true, "export": true}

// RunTask queues a run of an enabled task, with optional parameters for the
// task, e.g. {"value": "example.com"} for a oneshot analytics. Yeti runs it in
// the background: Tasks shows its status and last run.
func (c *Client) RunTask(ctx context.Context, name string, params map[string]any, dryRun bool) (*Changes, Meta, error) {
	task, err := c.findTask(ctx, name)
	if err != nil {
		return nil, Meta{}, err
	}
	if !task.Enabled {
		return nil, Meta{}, &Error{Code: CodeInvalidRequest, Message: fmt.Sprintf("%s is disabled, so Yeti won't run it.", task.Name),
			Hint: fmt.Sprintf(`Enable it first: "yeti admin enable-task %s".`, task.Name)}
	}
	if _, ok := params["value"]; task.Type == "oneshot" && !ok {
		return nil, Meta{}, usage(fmt.Sprintf("%s acts on one observable, and needs its value.", task.Name),
			`Pass the observable's value as the "value" parameter, e.g. --param value=example.com.`)
	}
	var meta Meta
	if exclusive[task.Type] && task.Status == "running" {
		meta.Warnings = append(meta.Warnings, fmt.Sprintf("%s is running already; Yeti skips a run queued meanwhile.", task.Name))
	}
	if !dryRun {
		if err := c.api.RunTask(ctx, task.Name, params); err != nil {
			return nil, Meta{}, wrap(err, "")
		}
	}
	result := &Changes{DryRun: dryRun, Changes: []Change{{Item: name, Name: task.Name, Type: task.Type, Action: ActionQueued}}}
	return result, meta, nil
}

// EnableTasks enables tasks, so that Yeti runs them on schedule and on demand.
func (c *Client) EnableTasks(ctx context.Context, names []string, dryRun bool) (*Changes, Meta, error) {
	return c.setTasksEnabled(ctx, names, true, dryRun)
}

// DisableTasks disables tasks.
func (c *Client) DisableTasks(ctx context.Context, names []string, dryRun bool) (*Changes, Meta, error) {
	return c.setTasksEnabled(ctx, names, false, dryRun)
}

func (c *Client) setTasksEnabled(ctx context.Context, names []string, enabled, dryRun bool) (*Changes, Meta, error) {
	if len(names) == 0 {
		return nil, Meta{}, usage("No task was named.", `"yeti read tasks" lists the tasks.`)
	}
	action := ActionDisabled
	if enabled {
		action = ActionEnabled
	}
	result := &Changes{DryRun: dryRun}
	for _, name := range names {
		task, err := c.findTask(ctx, name)
		if err != nil {
			result.add(failed(name, err))
			continue
		}
		change := Change{Item: name, Name: task.Name, Type: task.Type, Action: action}
		if task.Enabled == enabled {
			change.Action = ActionUnchanged
		} else if !dryRun {
			// Yeti only toggles, so a concurrent toggle could undo this one.
			toggled, err := c.api.ToggleTask(ctx, task.Name)
			if err != nil {
				result.add(failed(name, err))
				continue
			}
			if toggled.Enabled != enabled {
				result.add(Change{Item: name, Name: task.Name, Type: task.Type, Action: ActionFailed, Error: &Error{
					Code: CodeServerError, Message: fmt.Sprintf("%s was toggled at the same time, and is %s again.", task.Name, map[bool]string{true: "enabled", false: "disabled"}[toggled.Enabled]),
					Hint: "Try again.", Retryable: true}})
				continue
			}
		}
		result.add(change)
	}
	return result, result.meta(), nil
}

package cli

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/spf13/cobra"
	"github.com/yeti-platform/goyeti/yeti"
)

func (a *app) mcpCommand() *cobra.Command {
	var tier string
	cmd := &cobra.Command{
		Use:   "mcp",
		Short: "Serve the operations to an MCP client over stdio",
		Long: `Serves the operations of a tier, and of the tiers below it, as MCP tools over
standard input and output, for MCP hosts such as Claude Desktop or an ADK
agent. --tier read, the default, serves the read operations only; --tier mutate
adds the operations that add objects, links, tags and context; --tier admin
adds deleting objects and links, and running and toggling tasks.`,
		Example: `  claude mcp add yeti -e YETI_URL=https://yeti.example.com -e YETI_API_KEY=... -- yeti mcp`,
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if tier != "read" && tier != "mutate" && tier != "admin" {
				return a.emit(nil, yeti.Meta{}, &yeti.Error{Code: yeti.CodeUsage,
					Message: fmt.Sprintf("%q is not a tier.", tier), Hint: "Use --tier read, mutate or admin."})
			}
			client, err := a.client()
			if err != nil {
				return a.emit(nil, yeti.Meta{}, err)
			}
			return newMCPServer(client, tier).Run(cmd.Context(), &mcp.StdioTransport{})
		},
	}
	cmd.Flags().StringVar(&tier, "tier", "read", "the most powerful tier of operations to serve")
	return cmd
}

type lookupInput struct {
	Values []string `json:"values" jsonschema:"the values to look up: IPs, hostnames, URLs, hashes…, defanged or not"`
	Limit  int      `json:"limit,omitempty" jsonschema:"most linked objects per list for each value; default 10, at most 100"`
}

type searchInput struct {
	Query string `json:"query" jsonschema:"a name, word or phrase"`
	Limit int    `json:"limit,omitempty" jsonschema:"most hits to return; default 20, at most 50"`
}

type getInput struct {
	Ref string `json:"ref" jsonschema:"the object's reference, e.g. observables/123 or entities/45"`
}

type neighborsInput struct {
	Ref         string   `json:"ref" jsonschema:"the reference of the object to start from, e.g. entities/45"`
	Hops        int      `json:"hops,omitempty" jsonschema:"how far to walk, 1 to 3; default 1"`
	Direction   string   `json:"direction,omitempty" jsonschema:"any (default), out or in"`
	LinkTypes   []string `json:"link_types,omitempty" jsonschema:"keep only these relationship types, e.g. uses or c2"`
	TargetTypes []string `json:"target_types,omitempty" jsonschema:"keep only these object types or families, e.g. malware or observable"`
	Limit       int      `json:"limit,omitempty" jsonschema:"most paths to return; default 20, at most 100"`
}

type yaraBundleInput struct {
	Refs        []string `json:"refs,omitempty" jsonschema:"references of YARA rules, e.g. indicators/67"`
	Tags        []string `json:"tags,omitempty" jsonschema:"include the YARA rules with these tags"`
	ExcludeTags []string `json:"exclude_tags,omitempty" jsonschema:"leave out the rules with these tags"`
}

type noInput struct{}

type tasksInput struct {
	Names []string `json:"names,omitempty" jsonschema:"keep the tasks whose name contains one of these, ignoring case"`
	Type  string   `json:"type,omitempty" jsonschema:"keep the tasks of one type: feed, analytics, export, oneshot, event or log"`
}

type addObservablesInput struct {
	Values []string `json:"values" jsonschema:"the values to add: IPs, hostnames, URLs, hashes…, defanged or not"`
	Type   string   `json:"type,omitempty" jsonschema:"the type of every value; leave it out to let Yeti guess each value's type"`
	Tags   []string `json:"tags,omitempty" jsonschema:"tags for every value, new or existing"`
	DryRun bool     `json:"dry_run,omitempty" jsonschema:"say what would change, and change nothing"`
}

type addEntityInput struct {
	Type        string         `json:"type" jsonschema:"the entity type, e.g. malware, threat-actor or campaign"`
	Name        string         `json:"name" jsonschema:"the entity's name"`
	Description string         `json:"description,omitempty" jsonschema:"the description"`
	Fields      map[string]any `json:"fields,omitempty" jsonschema:"other fields of the type, e.g. aliases"`
	Tags        []string       `json:"tags,omitempty" jsonschema:"tags, whether the entity is new or existing"`
	DryRun      bool           `json:"dry_run,omitempty" jsonschema:"say what would change, and change nothing"`
}

type addIndicatorInput struct {
	Type        string         `json:"type" jsonschema:"the indicator type, e.g. yara, sigma, regex or query"`
	Name        string         `json:"name" jsonschema:"the indicator's name; for YARA, the rule's name"`
	Pattern     string         `json:"pattern" jsonschema:"the rule, regex or query itself"`
	Diamond     string         `json:"diamond" jsonschema:"the diamond model vertex: adversary, capability, infrastructure or victim"`
	Location    string         `json:"location,omitempty" jsonschema:"where the indicator applies, e.g. filesystem or network; needed by regex and query"`
	Description string         `json:"description,omitempty" jsonschema:"the description"`
	Fields      map[string]any `json:"fields,omitempty" jsonschema:"other fields of the type"`
	Tags        []string       `json:"tags,omitempty" jsonschema:"tags, whether the indicator is new or existing"`
	DryRun      bool           `json:"dry_run,omitempty" jsonschema:"say what would change, and change nothing"`
}

type linkInput struct {
	Source      string `json:"source" jsonschema:"the source's reference, e.g. observables/123"`
	LinkType    string `json:"link_type" jsonschema:"the relationship, e.g. c2, uses or delivers"`
	Target      string `json:"target" jsonschema:"the target's reference, e.g. entities/45"`
	Description string `json:"description,omitempty" jsonschema:"what the link means"`
	DryRun      bool   `json:"dry_run,omitempty" jsonschema:"say what would change, and change nothing"`
}

type tagInput struct {
	Refs   []string `json:"refs" jsonschema:"references of the objects, e.g. observables/123"`
	Tags   []string `json:"tags" jsonschema:"the tags"`
	DryRun bool     `json:"dry_run,omitempty" jsonschema:"say what would change, and change nothing"`
}

type addContextInput struct {
	Ref     string         `json:"ref" jsonschema:"the object's reference, e.g. observables/123"`
	Context map[string]any `json:"context" jsonschema:"the facts to add, as name/value pairs"`
	DryRun  bool           `json:"dry_run,omitempty" jsonschema:"say what would change, and change nothing"`
}

type deleteInput struct {
	Refs    []string `json:"refs" jsonschema:"references of the objects, e.g. observables/123"`
	Confirm bool     `json:"confirm,omitempty" jsonschema:"delete for good; needed unless dry_run is set"`
	DryRun  bool     `json:"dry_run,omitempty" jsonschema:"say what would be deleted, and delete nothing"`
}

type unlinkInput struct {
	Source   string `json:"source" jsonschema:"the source's reference, e.g. observables/123"`
	LinkType string `json:"link_type" jsonschema:"the relationship, e.g. c2, uses or delivers"`
	Target   string `json:"target" jsonschema:"the target's reference, e.g. entities/45"`
	Confirm  bool   `json:"confirm,omitempty" jsonschema:"delete the link; needed unless dry_run is set"`
	DryRun   bool   `json:"dry_run,omitempty" jsonschema:"say what would be deleted, and delete nothing"`
}

type runTaskInput struct {
	Name   string         `json:"name" jsonschema:"the task's exact name, as tasks lists it"`
	Params map[string]any `json:"params,omitempty" jsonschema:"parameters for the task, e.g. {\"value\": \"example.com\"} for a oneshot analytics"`
	DryRun bool           `json:"dry_run,omitempty" jsonschema:"say what would run, and run nothing"`
}

type tasksByNameInput struct {
	Names  []string `json:"names" jsonschema:"the tasks' exact names"`
	DryRun bool     `json:"dry_run,omitempty" jsonschema:"say what would change, and change nothing"`
}

// readOnly annotates tools that never change anything in Yeti.
func readOnly(title string) *mcp.ToolAnnotations {
	closed := false
	return &mcp.ToolAnnotations{Title: title, ReadOnlyHint: true, IdempotentHint: true, OpenWorldHint: &closed}
}

// writes annotates tools that change Yeti: destructive ones remove
// something, idempotent ones change nothing more when called again.
func writes(title string, destructive, idempotent bool) *mcp.ToolAnnotations {
	closed := false
	return &mcp.ToolAnnotations{Title: title, DestructiveHint: &destructive, IdempotentHint: idempotent, OpenWorldHint: &closed}
}

func newMCPServer(c *yeti.Client, tier string) *mcp.Server {
	server := mcp.NewServer(&mcp.Implementation{Name: "yeti", Version: yeti.Version}, &mcp.ServerOptions{
		Instructions: "Tools for the Yeti threat intelligence platform. Start with lookup for observable " +
			"values (IPs, domains, hashes) and search for names and topics; follow references " +
			"(observables/123, entities/45) with get and neighbors.",
	})
	mcp.AddTool(server, &mcp.Tool{Name: "lookup", Description: lookupDescription, Annotations: readOnly("Look up observables")},
		func(ctx context.Context, _ *mcp.CallToolRequest, in lookupInput) (*mcp.CallToolResult, any, error) {
			return reply(c.Lookup(ctx, in.Values, in.Limit))
		})
	mcp.AddTool(server, &mcp.Tool{Name: "search", Description: searchDescription, Annotations: readOnly("Search")},
		func(ctx context.Context, _ *mcp.CallToolRequest, in searchInput) (*mcp.CallToolResult, any, error) {
			return reply(c.Search(ctx, in.Query, in.Limit))
		})
	mcp.AddTool(server, &mcp.Tool{Name: "get", Description: getDescription, Annotations: readOnly("Get an object")},
		func(ctx context.Context, _ *mcp.CallToolRequest, in getInput) (*mcp.CallToolResult, any, error) {
			return reply(c.Get(ctx, in.Ref))
		})
	mcp.AddTool(server, &mcp.Tool{Name: "neighbors", Description: neighborsDescription, Annotations: readOnly("Linked objects")},
		func(ctx context.Context, _ *mcp.CallToolRequest, in neighborsInput) (*mcp.CallToolResult, any, error) {
			return reply(c.Neighbors(ctx, in.Ref, yeti.NeighborsOptions{Hops: in.Hops, Direction: in.Direction,
				LinkTypes: in.LinkTypes, TargetTypes: in.TargetTypes, Limit: in.Limit}))
		})
	mcp.AddTool(server, &mcp.Tool{Name: "yara_bundle", Description: yaraBundleDescription, Annotations: readOnly("YARA bundle")},
		func(ctx context.Context, _ *mcp.CallToolRequest, in yaraBundleInput) (*mcp.CallToolResult, any, error) {
			return reply(c.YaraBundle(ctx, in.Refs, in.Tags, in.ExcludeTags))
		})
	mcp.AddTool(server, &mcp.Tool{Name: "types", Description: typesDescription, Annotations: readOnly("Object types")},
		func(ctx context.Context, _ *mcp.CallToolRequest, _ noInput) (*mcp.CallToolResult, any, error) {
			return reply(c.Types(ctx))
		})
	mcp.AddTool(server, &mcp.Tool{Name: "tasks", Description: tasksDescription, Annotations: readOnly("Tasks")},
		func(ctx context.Context, _ *mcp.CallToolRequest, in tasksInput) (*mcp.CallToolResult, any, error) {
			return reply(c.Tasks(ctx, yeti.TasksOptions{Names: in.Names, Type: in.Type}))
		})
	if tier == "mutate" || tier == "admin" {
		addMutateTools(server, c)
	}
	if tier == "admin" {
		addAdminTools(server, c)
	}
	return server
}

func addMutateTools(server *mcp.Server, c *yeti.Client) {
	mcp.AddTool(server, &mcp.Tool{Name: "add_observables", Description: addObservablesDescription, Annotations: writes("Add observables", false, true)},
		func(ctx context.Context, _ *mcp.CallToolRequest, in addObservablesInput) (*mcp.CallToolResult, any, error) {
			return reply(c.AddObservables(ctx, in.Values, yeti.AddObservablesOptions{Type: in.Type, Tags: in.Tags, DryRun: in.DryRun}))
		})
	mcp.AddTool(server, &mcp.Tool{Name: "add_entity", Description: addEntityDescription, Annotations: writes("Add an entity", false, true)},
		func(ctx context.Context, _ *mcp.CallToolRequest, in addEntityInput) (*mcp.CallToolResult, any, error) {
			fields := withFields(in.Fields, map[string]string{"description": in.Description})
			return reply(c.AddEntity(ctx, in.Type, in.Name, yeti.AddObjectOptions{Fields: fields, Tags: in.Tags, DryRun: in.DryRun}))
		})
	mcp.AddTool(server, &mcp.Tool{Name: "add_indicator", Description: addIndicatorDescription, Annotations: writes("Add an indicator", false, true)},
		func(ctx context.Context, _ *mcp.CallToolRequest, in addIndicatorInput) (*mcp.CallToolResult, any, error) {
			fields := withFields(in.Fields, map[string]string{"pattern": in.Pattern, "diamond": in.Diamond,
				"location": in.Location, "description": in.Description})
			return reply(c.AddIndicator(ctx, in.Type, in.Name, yeti.AddObjectOptions{Fields: fields, Tags: in.Tags, DryRun: in.DryRun}))
		})
	mcp.AddTool(server, &mcp.Tool{Name: "link", Description: linkDescription, Annotations: writes("Link objects", false, true)},
		func(ctx context.Context, _ *mcp.CallToolRequest, in linkInput) (*mcp.CallToolResult, any, error) {
			return reply(c.Link(ctx, in.Source, in.LinkType, in.Target, in.Description, in.DryRun))
		})
	mcp.AddTool(server, &mcp.Tool{Name: "tag", Description: tagDescription, Annotations: writes("Tag objects", false, true)},
		func(ctx context.Context, _ *mcp.CallToolRequest, in tagInput) (*mcp.CallToolResult, any, error) {
			return reply(c.Tag(ctx, in.Refs, in.Tags, in.DryRun))
		})
	mcp.AddTool(server, &mcp.Tool{Name: "untag", Description: untagDescription, Annotations: writes("Untag objects", true, true)},
		func(ctx context.Context, _ *mcp.CallToolRequest, in tagInput) (*mcp.CallToolResult, any, error) {
			return reply(c.Untag(ctx, in.Refs, in.Tags, in.DryRun))
		})
	mcp.AddTool(server, &mcp.Tool{Name: "add_context", Description: addContextDescription, Annotations: writes("Add context", false, true)},
		func(ctx context.Context, _ *mcp.CallToolRequest, in addContextInput) (*mcp.CallToolResult, any, error) {
			return reply(c.AddContext(ctx, in.Ref, in.Context, in.DryRun))
		})
}

func addAdminTools(server *mcp.Server, c *yeti.Client) {
	mcp.AddTool(server, &mcp.Tool{Name: "delete", Description: deleteDescription, Annotations: writes("Delete objects", true, true)},
		func(ctx context.Context, _ *mcp.CallToolRequest, in deleteInput) (*mcp.CallToolResult, any, error) {
			return reply(c.Delete(ctx, in.Refs, yeti.RemoveOptions{Confirm: in.Confirm, DryRun: in.DryRun}))
		})
	mcp.AddTool(server, &mcp.Tool{Name: "unlink", Description: unlinkDescription, Annotations: writes("Unlink objects", true, true)},
		func(ctx context.Context, _ *mcp.CallToolRequest, in unlinkInput) (*mcp.CallToolResult, any, error) {
			return reply(c.Unlink(ctx, in.Source, in.LinkType, in.Target, yeti.RemoveOptions{Confirm: in.Confirm, DryRun: in.DryRun}))
		})
	mcp.AddTool(server, &mcp.Tool{Name: "run_task", Description: runTaskDescription, Annotations: writes("Run a task", false, false)},
		func(ctx context.Context, _ *mcp.CallToolRequest, in runTaskInput) (*mcp.CallToolResult, any, error) {
			return reply(c.RunTask(ctx, in.Name, in.Params, in.DryRun))
		})
	mcp.AddTool(server, &mcp.Tool{Name: "enable_task", Description: enableTaskDescription, Annotations: writes("Enable tasks", false, true)},
		func(ctx context.Context, _ *mcp.CallToolRequest, in tasksByNameInput) (*mcp.CallToolResult, any, error) {
			return reply(c.EnableTasks(ctx, in.Names, in.DryRun))
		})
	mcp.AddTool(server, &mcp.Tool{Name: "disable_task", Description: disableTaskDescription, Annotations: writes("Disable tasks", false, true)},
		func(ctx context.Context, _ *mcp.CallToolRequest, in tasksByNameInput) (*mcp.CallToolResult, any, error) {
			return reply(c.DisableTasks(ctx, in.Names, in.DryRun))
		})
}

// withFields merges the set named fields into fields.
func withFields(fields map[string]any, named map[string]string) map[string]any {
	merged := map[string]any{}
	for k, v := range fields {
		merged[k] = v
	}
	for k, v := range named {
		if v != "" {
			merged[k] = v
		}
	}
	return merged
}

// reply returns an operation's outcome as a tool result: the success envelope
// as structured content, or the error envelope as an error result.
func reply[T any](data T, meta yeti.Meta, err error) (*mcp.CallToolResult, any, error) {
	if err != nil {
		e, _ := failure(err)
		text, _ := json.Marshal(e)
		return &mcp.CallToolResult{IsError: true, Content: []mcp.Content{&mcp.TextContent{Text: string(text)}}}, nil, nil
	}
	return nil, success(data, meta), nil
}

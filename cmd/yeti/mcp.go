package main

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
		Long: `Serves the operations of a tier as MCP tools over standard input and output,
for MCP hosts such as Claude Desktop or an ADK agent. The read tier, the only
one so far, is the default.`,
		Example: `  claude mcp add yeti -e YETI_URL=https://yeti.example.com -e YETI_API_KEY=... -- yeti mcp`,
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if tier != "read" {
				return a.emit(nil, yeti.Meta{}, &yeti.Error{Code: yeti.CodeUsage,
					Message: fmt.Sprintf("The %q tier isn't available yet.", tier), Hint: "Use --tier read."})
			}
			client, err := a.client()
			if err != nil {
				return a.emit(nil, yeti.Meta{}, err)
			}
			return newMCPServer(client).Run(cmd.Context(), &mcp.StdioTransport{})
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

// readOnly annotates tools that never change anything in Yeti.
func readOnly(title string) *mcp.ToolAnnotations {
	closed := false
	return &mcp.ToolAnnotations{Title: title, ReadOnlyHint: true, IdempotentHint: true, OpenWorldHint: &closed}
}

func newMCPServer(c *yeti.Client) *mcp.Server {
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
	return server
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

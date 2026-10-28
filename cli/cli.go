// Package cli implements the yeti command line and MCP server (yeti mcp) over
// the operations in package yeti.
package cli

import (
	"bufio"
	"context"
	"io"
	"os"
	"os/signal"
	"runtime/debug"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/yeti-platform/goyeti/yeti"
)

// NewClient constructs the Yeti client used by CLI commands and the MCP server.
var NewClient = func(ctx context.Context, cfg yeti.Config) (*yeti.Client, error) {
	return yeti.New(cfg)
}

// Execute runs the yeti command line and exits with its status code.
func Execute() {
	// Release builds set the version at link time; go install records it in
	// the build information instead.
	if info, ok := debug.ReadBuildInfo(); ok && yeti.Version == "dev" && strings.HasPrefix(info.Main.Version, "v") {
		yeti.Version = strings.TrimPrefix(info.Main.Version, "v")
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	stdinIsTerminal := false
	if info, err := os.Stdin.Stat(); err == nil {
		stdinIsTerminal = info.Mode()&os.ModeCharDevice != 0
	}
	a := &app{stdin: os.Stdin, stdinIsTerminal: stdinIsTerminal, stdout: os.Stdout, getenv: os.Getenv}
	os.Exit(a.run(ctx, os.Args[1:]))
}

type app struct {
	ctx             context.Context
	stdin           io.Reader
	stdinIsTerminal bool
	stdout          io.Writer
	getenv          func(string) string

	url, apiKey, agentName string
	timeout                time.Duration
	pretty                 bool
	exit                   int
}

func (a *app) run(ctx context.Context, args []string) int {
	a.ctx = ctx
	root := a.rootCommand()
	root.SetArgs(args)
	root.SetIn(a.stdin)
	root.SetOut(a.stdout)
	root.SetErr(a.stdout)
	if err := root.ExecuteContext(ctx); err != nil {
		e, code := failure(err)
		write(a.stdout, e, a.pretty)
		return code
	}
	return a.exit
}

func (a *app) client() (*yeti.Client, error) {
	return NewClient(a.ctx, yeti.Config{
		URL:       firstSet(a.url, a.getenv("YETI_URL")),
		APIKey:    firstSet(a.apiKey, a.getenv("YETI_API_KEY")),
		Timeout:   a.timeout,
		AgentName: firstSet(a.agentName, a.getenv("YETI_AGENT_NAME")),
	})
}

// emit writes an operation's outcome, and sets the exit code.
func (a *app) emit(data any, meta yeti.Meta, err error) error {
	if err != nil {
		e, code := failure(err)
		write(a.stdout, e, a.pretty)
		a.exit = code
		return nil
	}
	write(a.stdout, success(data, meta), a.pretty)
	if meta.Failed > 0 {
		a.exit = exitFailed
	}
	return nil
}

// withClient runs an operation once the configuration is known to be valid.
func (a *app) withClient(op func(*yeti.Client) error) error {
	client, err := a.client()
	if err != nil {
		return a.emit(nil, yeti.Meta{}, err)
	}
	return op(client)
}

func (a *app) rootCommand() *cobra.Command {
	root := &cobra.Command{
		Use:   "yeti",
		Short: "A client for Yeti, built for AI agents",
		Long: `yeti queries a Yeti server. The first word of every command is its tier:

  yeti read ...    never changes anything
  yeti mutate ...  adds objects, links, tags and context; --dry-run shows what
                   it would do
  yeti admin ...   deletes objects and links (with --yes), runs and toggles
                   tasks
  yeti mcp         serves the operations to MCP clients over stdio

Output is JSON: {"data": ...} on success, with "truncated", "warnings" and
"failed" when they apply, or {"error": {"code", "message", "hint",
"retryable"}}. Exit codes: 0 success, 1 failed operation or failed items, 2 bad
usage.

Configure with YETI_URL and YETI_API_KEY, and optionally YETI_AGENT_NAME.`,
		SilenceUsage:  true,
		SilenceErrors: true,
		Version:       yeti.Version,
	}
	root.SetVersionTemplate(`{"data":{"version":"{{.Version}}"}}` + "\n")
	flags := root.PersistentFlags()
	flags.StringVar(&a.url, "url", "", "Yeti server URL (default $YETI_URL)")
	flags.StringVar(&a.apiKey, "api-key", "", "Yeti API key (default $YETI_API_KEY; prefer the variable, which other users can't see)")
	flags.StringVar(&a.agentName, "agent-name", "", "name of the agent, sent to Yeti (default $YETI_AGENT_NAME)")
	flags.DurationVar(&a.timeout, "timeout", 30*time.Second, "timeout for each request to Yeti")
	flags.BoolVar(&a.pretty, "pretty", false, "indent the JSON output")

	read := &cobra.Command{
		Use:   "read",
		Short: "Operations that never change anything in Yeti",
		// Without Args and RunE, cobra answers an unknown operation with the
		// help text and exit code 0.
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error { return cmd.Help() },
	}
	read.AddCommand(a.lookupCommand(), a.searchCommand(), a.getCommand(), a.neighborsCommand(),
		a.yaraBundleCommand(), a.typesCommand(), a.tasksCommand())
	root.AddCommand(read, a.mutateCommand(), a.adminCommand(), a.mcpCommand())
	return root
}

func (a *app) lookupCommand() *cobra.Command {
	var limit int
	cmd := &cobra.Command{
		Use:   "lookup [value...]",
		Short: "Look up observables: what Yeti knows about each value",
		Long: lookupDescription + `

Values come from the arguments, or from standard input (whitespace-separated)
when there are none or the only argument is "-".`,
		Example: "  yeti read lookup 203.0.113.7 evil[.]example[.]com\n  cut -f1 iocs.tsv | yeti read lookup",
		RunE: func(cmd *cobra.Command, args []string) error {
			values := args
			if len(args) == 0 || (len(args) == 1 && args[0] == "-") {
				if a.stdinIsTerminal {
					return a.emit(nil, yeti.Meta{}, &yeti.Error{Code: yeti.CodeUsage, Message: "No values to look up.",
						Hint: "Pass values as arguments, or pipe them in."})
				}
				var err error
				if values, err = readValues(a.stdin); err != nil {
					return err
				}
			}
			return a.withClient(func(c *yeti.Client) error {
				return a.emit(c.Lookup(cmd.Context(), values, limit))
			})
		},
	}
	cmd.Flags().IntVar(&limit, "limit", 0, "most linked objects per list for each value (default 10, at most 100)")
	return cmd
}

func readValues(r io.Reader) ([]string, error) {
	var values []string
	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		values = append(values, scanner.Text())
	}
	return values, scanner.Err()
}

func (a *app) searchCommand() *cobra.Command {
	var limit int
	cmd := &cobra.Command{
		Use:     "search <query>",
		Short:   "Find objects by keyword and by meaning",
		Long:    searchDescription,
		Example: `  yeti read search "cobalt strike"`,
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return a.withClient(func(c *yeti.Client) error {
				return a.emit(c.Search(cmd.Context(), strings.Join(args, " "), limit))
			})
		},
	}
	cmd.Flags().IntVar(&limit, "limit", 0, "most hits to return (default 20, at most 50)")
	return cmd
}

func (a *app) getCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "get <reference>",
		Short:   "Get one object in full",
		Long:    getDescription,
		Example: "  yeti read get entities/45",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return a.withClient(func(c *yeti.Client) error {
				return a.emit(c.Get(cmd.Context(), args[0]))
			})
		},
	}
}

func (a *app) neighborsCommand() *cobra.Command {
	var opts yeti.NeighborsOptions
	cmd := &cobra.Command{
		Use:     "neighbors <reference>",
		Short:   "List the objects linked to one object",
		Long:    neighborsDescription,
		Example: "  yeti read neighbors entities/45 --target-type observable --hops 2",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return a.withClient(func(c *yeti.Client) error {
				return a.emit(c.Neighbors(cmd.Context(), args[0], opts))
			})
		},
	}
	flags := cmd.Flags()
	flags.IntVar(&opts.Hops, "hops", 1, "how far to walk, 1 to 3")
	flags.StringVar(&opts.Direction, "direction", "any", `"any", "out" or "in"`)
	flags.StringSliceVar(&opts.LinkTypes, "link-type", nil, "keep only these relationship types (repeat or comma-separate)")
	flags.StringSliceVar(&opts.TargetTypes, "target-type", nil, "keep only these object types or families, e.g. malware or observable")
	flags.IntVar(&opts.Limit, "limit", 0, "most paths to return (default 20, at most 100)")
	return cmd
}

func (a *app) yaraBundleCommand() *cobra.Command {
	var tags, excludeTags []string
	var raw bool
	cmd := &cobra.Command{
		Use:     "yara-bundle [reference...]",
		Short:   "Build a YARA bundle from rules and their dependencies",
		Long:    yaraBundleDescription + "\n\nWith --raw, prints only the YARA source.",
		Example: "  yeti read yara-bundle indicators/67\n  yeti read yara-bundle --tag emotet --raw > emotet.yar",
		RunE: func(cmd *cobra.Command, args []string) error {
			return a.withClient(func(c *yeti.Client) error {
				bundle, meta, err := c.YaraBundle(cmd.Context(), args, tags, excludeTags)
				if raw && err == nil {
					_, err := io.WriteString(a.stdout, bundle.Bundle)
					return err
				}
				return a.emit(bundle, meta, err)
			})
		},
	}
	flags := cmd.Flags()
	flags.StringSliceVar(&tags, "tag", nil, "include the YARA rules with these tags")
	flags.StringSliceVar(&excludeTags, "exclude-tag", nil, "leave out the rules with these tags")
	flags.BoolVar(&raw, "raw", false, "print only the YARA source")
	return cmd
}

func (a *app) typesCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "types",
		Short: "List the object types the server knows",
		Long:  typesDescription,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return a.withClient(func(c *yeti.Client) error {
				return a.emit(c.Types(cmd.Context()))
			})
		},
	}
}

func (a *app) tasksCommand() *cobra.Command {
	var opts yeti.TasksOptions
	cmd := &cobra.Command{
		Use:     "tasks [name...]",
		Short:   "List tasks: feeds, analytics, exports…",
		Long:    tasksDescription,
		Example: "  yeti read tasks abuse\n  yeti read tasks --type feed",
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Names = args
			return a.withClient(func(c *yeti.Client) error {
				return a.emit(c.Tasks(cmd.Context(), opts))
			})
		},
	}
	cmd.Flags().StringVar(&opts.Type, "type", "", "keep the tasks of one type: feed, analytics, export, oneshot, event or log")
	return cmd
}

func firstSet(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}

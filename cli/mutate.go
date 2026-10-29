package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yeti-platform/goyeti/yeti"
)

func (a *app) mutateCommand() *cobra.Command {
	var dryRun bool
	mutate := &cobra.Command{
		Use:   "mutate",
		Short: "Operations that add to Yeti: objects, links, tags and context",
		Long: `Operations that add to Yeti. They create what's missing and leave what exists:
running one twice changes nothing more. With --dry-run, they only say what they
would do.`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error { return cmd.Help() },
	}
	mutate.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "say what would change, and change nothing")

	add := &cobra.Command{
		Use:   "add",
		Short: "Add observables, an entity or an indicator",
		Args:  cobra.NoArgs,
		RunE:  func(cmd *cobra.Command, args []string) error { return cmd.Help() },
	}
	add.AddCommand(a.addObservablesCommand(&dryRun), a.addEntityCommand(&dryRun), a.addIndicatorCommand(&dryRun))
	mutate.AddCommand(add, a.linkCommand(&dryRun), a.tagCommand(&dryRun, false), a.tagCommand(&dryRun, true),
		a.addContextCommand(&dryRun))
	return mutate
}

func (a *app) addObservablesCommand(dryRun *bool) *cobra.Command {
	var opts yeti.AddObservablesOptions
	cmd := &cobra.Command{
		Use:   "observables [value...]",
		Short: "Add observables Yeti doesn't know yet",
		Long: addObservablesDescription + `

Values come from the arguments, or from standard input (one per line) when
there are none or the only argument is "-".`,
		Example: "  yeti mutate add observables 203.0.113.7 evil[.]example[.]com --tag campaign-x\n" +
			"  yeti mutate add observables 'C:\\Windows\\Temp\\x.exe' --type path",
		RunE: func(cmd *cobra.Command, args []string) error {
			values := args
			if len(args) == 0 || (len(args) == 1 && args[0] == "-") {
				if a.stdinIsTerminal {
					return a.emit(nil, yeti.Meta{}, &yeti.Error{Code: yeti.CodeUsage, Message: "No values to add.",
						Hint: "Pass values as arguments, or pipe them in, one per line."})
				}
				var err error
				if values, err = readLines(a.stdin); err != nil {
					return err
				}
			}
			opts.DryRun = *dryRun
			return a.withClient(func(c *yeti.Client) error {
				return a.emit(c.AddObservables(cmd.Context(), values, opts))
			})
		},
	}
	cmd.Flags().StringVar(&opts.Type, "type", "", "the type of every value (default: Yeti guesses each value's type)")
	cmd.Flags().StringSliceVar(&opts.Tags, "tag", nil, "tags for every value, new or existing")
	return cmd
}

// readLines reads one value per line, so that values may contain spaces.
func readLines(r io.Reader) ([]string, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	var values []string
	for _, line := range strings.Split(string(data), "\n") {
		if line = strings.TrimSpace(line); line != "" {
			values = append(values, line)
		}
	}
	return values, nil
}

// objectFlags are the flags entities and indicators share.
type objectFlags struct {
	description string
	fields      []string
	tags        []string
}

func (f *objectFlags) register(cmd *cobra.Command) {
	cmd.Flags().StringVar(&f.description, "description", "", "the description")
	cmd.Flags().StringArrayVar(&f.fields, "field", nil, "another field, as name=value; a JSON value is decoded (repeat)")
	cmd.Flags().StringSliceVar(&f.tags, "tag", nil, "tags, whether the object is new or existing")
}

func (f *objectFlags) options(dryRun bool, extra map[string]any) (yeti.AddObjectOptions, error) {
	fields, err := parseFields(f.fields)
	if err != nil {
		return yeti.AddObjectOptions{}, err
	}
	if f.description != "" {
		fields["description"] = f.description
	}
	for k, v := range extra {
		if v != "" {
			fields[k] = v
		}
	}
	return yeti.AddObjectOptions{Fields: fields, Tags: f.tags, DryRun: dryRun}, nil
}

// parseFields reads name=value pairs. A value that is valid JSON, such as a
// number, a list or an object, is decoded; any other value is a string.
func parseFields(pairs []string) (map[string]any, error) {
	fields := map[string]any{}
	for _, pair := range pairs {
		name, value, ok := strings.Cut(pair, "=")
		if !ok || strings.TrimSpace(name) == "" {
			return nil, &yeti.Error{Code: yeti.CodeUsage, Message: fmt.Sprintf("%q is not name=value.", pair),
				Hint: `Write fields as name=value, e.g. note="seen in campaign X" or ports=[80,443].`}
		}
		var decoded any
		if json.Unmarshal([]byte(value), &decoded) == nil {
			fields[strings.TrimSpace(name)] = decoded
		} else {
			fields[strings.TrimSpace(name)] = value
		}
	}
	return fields, nil
}

func (a *app) addEntityCommand(dryRun *bool) *cobra.Command {
	var flags objectFlags
	cmd := &cobra.Command{
		Use:     "entity <type> <name>",
		Short:   "Add an entity, unless one of that type and name exists",
		Long:    addEntityDescription,
		Example: `  yeti mutate add entity malware Qakbot --description "Banking trojan" --field aliases='["QBot"]' --tag loader`,
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts, err := flags.options(*dryRun, nil)
			if err != nil {
				return a.emit(nil, yeti.Meta{}, err)
			}
			return a.withClient(func(c *yeti.Client) error {
				return a.emit(c.AddEntity(cmd.Context(), args[0], args[1], opts))
			})
		},
	}
	flags.register(cmd)
	return cmd
}

func (a *app) addIndicatorCommand(dryRun *bool) *cobra.Command {
	var flags objectFlags
	var pattern, diamond, location string
	cmd := &cobra.Command{
		Use:   "indicator <type> <name>",
		Short: "Add an indicator, unless one of that type and name exists",
		Long:  addIndicatorDescription + "\n\nPass --pattern @file to read the pattern from a file.",
		Example: "  yeti mutate add indicator yara qakbot_loader --pattern @qakbot.yar --diamond capability\n" +
			`  yeti mutate add indicator regex qakbot_paths --pattern '\\qbot\\' --diamond capability --location filesystem`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			if path, ok := strings.CutPrefix(pattern, "@"); ok {
				data, err := os.ReadFile(path)
				if err != nil {
					return a.emit(nil, yeti.Meta{}, &yeti.Error{Code: yeti.CodeUsage, Message: err.Error(),
						Hint: "--pattern @file reads the pattern from a file."})
				}
				pattern = string(data)
			}
			opts, err := flags.options(*dryRun, map[string]any{"pattern": pattern, "diamond": diamond, "location": location})
			if err != nil {
				return a.emit(nil, yeti.Meta{}, err)
			}
			return a.withClient(func(c *yeti.Client) error {
				return a.emit(c.AddIndicator(cmd.Context(), args[0], args[1], opts))
			})
		},
	}
	flags.register(cmd)
	cmd.Flags().StringVar(&pattern, "pattern", "", "the pattern: the rule, regex or query itself, or @file")
	cmd.Flags().StringVar(&diamond, "diamond", "", "the diamond model vertex: adversary, capability, infrastructure or victim")
	cmd.Flags().StringVar(&location, "location", "", "where the indicator applies, e.g. filesystem or network, for types that need it")
	return cmd
}

func (a *app) linkCommand(dryRun *bool) *cobra.Command {
	var description string
	cmd := &cobra.Command{
		Use:     "link <source> <link-type> <target>",
		Short:   "Link two objects, unless they're already linked that way",
		Long:    linkDescription,
		Example: "  yeti mutate link observables/123 c2 entities/45\n  yeti mutate link entities/46 uses entities/45",
		Args:    cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			return a.withClient(func(c *yeti.Client) error {
				return a.emit(c.Link(cmd.Context(), args[0], args[1], args[2], description, *dryRun))
			})
		},
	}
	cmd.Flags().StringVar(&description, "description", "", "what the link means")
	return cmd
}

func (a *app) tagCommand(dryRun *bool, remove bool) *cobra.Command {
	var tags []string
	use, short, long, example := "tag", "Add tags to objects", tagDescription, "  yeti mutate tag observables/123 entities/45 --tag campaign-x"
	if remove {
		use, short, long, example = "untag", "Remove tags from objects", untagDescription, "  yeti mutate untag observables/123 --tag false-positive"
	}
	cmd := &cobra.Command{
		Use:     use + " <reference...> --tag <tag>",
		Short:   short,
		Long:    long,
		Example: example,
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return a.withClient(func(c *yeti.Client) error {
				if remove {
					return a.emit(c.Untag(cmd.Context(), args, tags, *dryRun))
				}
				return a.emit(c.Tag(cmd.Context(), args, tags, *dryRun))
			})
		},
	}
	cmd.Flags().StringSliceVar(&tags, "tag", nil, "the tags (repeat or comma-separate)")
	return cmd
}

func (a *app) addContextCommand(dryRun *bool) *cobra.Command {
	var raw string
	cmd := &cobra.Command{
		Use:   "add-context <reference> [name=value...]",
		Short: "Add context to an object",
		Long:  addContextDescription,
		Example: `  yeti mutate add-context observables/123 note="beaconing every 60s" first_seen=2026-10-01` + "\n" +
			`  yeti mutate add-context entities/45 --json '{"report": "https://example.com/report"}'`,
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			context, err := parseFields(args[1:])
			if err != nil {
				return a.emit(nil, yeti.Meta{}, err)
			}
			if raw != "" {
				var fromJSON map[string]any
				if err := json.Unmarshal([]byte(raw), &fromJSON); err != nil {
					return a.emit(nil, yeti.Meta{}, &yeti.Error{Code: yeti.CodeUsage,
						Message: "--json isn't a JSON object: " + err.Error(), Hint: `Pass an object, e.g. --json '{"note": "..."}'.`})
				}
				for k, v := range fromJSON {
					context[k] = v
				}
			}
			return a.withClient(func(c *yeti.Client) error {
				return a.emit(c.AddContext(cmd.Context(), args[0], context, *dryRun))
			})
		},
	}
	cmd.Flags().StringVar(&raw, "json", "", "the context as a JSON object, merged with name=value arguments")
	return cmd
}

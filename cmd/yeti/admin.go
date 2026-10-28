package main

import (
	"github.com/spf13/cobra"
	"github.com/yeti-platform/goyeti/yeti"
)

func (a *app) adminCommand() *cobra.Command {
	var dryRun bool
	admin := &cobra.Command{
		Use:   "admin",
		Short: "Operations that remove from Yeti, and run its tasks",
		Long: `Operations that delete objects and links, and run, enable and disable tasks.
Deleting needs --yes; with --dry-run, every operation only says what it would
do.`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error { return cmd.Help() },
	}
	admin.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "say what would change, and change nothing")
	admin.AddCommand(a.deleteCommand(&dryRun), a.unlinkCommand(&dryRun), a.runTaskCommand(&dryRun),
		a.setTasksCommand(&dryRun, true), a.setTasksCommand(&dryRun, false))
	return admin
}

func (a *app) deleteCommand(dryRun *bool) *cobra.Command {
	var yes bool
	cmd := &cobra.Command{
		Use:     "delete <reference...>",
		Short:   "Delete objects and their links",
		Long:    deleteDescription,
		Example: "  yeti admin delete observables/123 entities/45 --dry-run\n  yeti admin delete observables/123 entities/45 --yes",
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return a.withClient(func(c *yeti.Client) error {
				return a.emit(c.Delete(cmd.Context(), args, yeti.RemoveOptions{Confirm: yes, DryRun: *dryRun}))
			})
		},
	}
	cmd.Flags().BoolVar(&yes, "yes", false, "delete: the objects and their links are gone for good")
	return cmd
}

func (a *app) unlinkCommand(dryRun *bool) *cobra.Command {
	var yes bool
	cmd := &cobra.Command{
		Use:     "unlink <source> <link-type> <target>",
		Short:   "Delete a link between two objects",
		Long:    unlinkDescription,
		Example: "  yeti admin unlink observables/123 c2 entities/45 --yes",
		Args:    cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			return a.withClient(func(c *yeti.Client) error {
				return a.emit(c.Unlink(cmd.Context(), args[0], args[1], args[2], yeti.RemoveOptions{Confirm: yes, DryRun: *dryRun}))
			})
		},
	}
	cmd.Flags().BoolVar(&yes, "yes", false, "delete the link")
	return cmd
}

func (a *app) runTaskCommand(dryRun *bool) *cobra.Command {
	var params []string
	cmd := &cobra.Command{
		Use:   "run-task <name>",
		Short: "Queue a run of a task",
		Long:  runTaskDescription,
		Example: "  yeti admin run-task AbuseIPDB\n" +
			"  yeti admin run-task VTDomainReport --param value=example.com",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			fields, err := parseFields(params)
			if err != nil {
				return a.emit(nil, yeti.Meta{}, err)
			}
			return a.withClient(func(c *yeti.Client) error {
				return a.emit(c.RunTask(cmd.Context(), args[0], fields, *dryRun))
			})
		},
	}
	cmd.Flags().StringArrayVar(&params, "param", nil, "a parameter for the task, as name=value; a JSON value is decoded (repeat)")
	return cmd
}

func (a *app) setTasksCommand(dryRun *bool, enable bool) *cobra.Command {
	use, short, long := "disable-task", "Disable tasks", disableTaskDescription
	if enable {
		use, short, long = "enable-task", "Enable tasks", enableTaskDescription
	}
	return &cobra.Command{
		Use:     use + " <name...>",
		Short:   short,
		Long:    long,
		Example: "  yeti admin " + use + " AbuseIPDB AbuseCHMalwareBazaar",
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return a.withClient(func(c *yeti.Client) error {
				if enable {
					return a.emit(c.EnableTasks(cmd.Context(), args, *dryRun))
				}
				return a.emit(c.DisableTasks(cmd.Context(), args, *dryRun))
			})
		},
	}
}

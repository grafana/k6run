// Package cmd contains run cobra command factory function.
package cmd

import (
	_ "embed"

	"github.com/spf13/cobra"
)

type options struct{}

//go:embed help.md
var help string

// New creates new cobra command for deps command.
func New() *cobra.Command {
	opts := new(options)

	cmd := &cobra.Command{
		Use:   "run [flags] [script-file]",
		Short: "Launching k6 with extensions",
		Long:  help,
		Args:  cobra.MaximumNArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			return run(opts, args)
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	return cmd
}

func run(_ *options, _ []string) error {
	return nil
}

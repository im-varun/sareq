package cmd

import (
	"github.com/im-varun/sareq/cmd/flags"
	"github.com/spf13/cobra"
)

var patchCmd = &cobra.Command{
	Use:     "patch <url>",
	Aliases: []string{"PATCH"},
	Short:   "Send HTTP PATCH request to the specified URL",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := runRequest(cmd.Name(), args[0])

		return err
	},
}

func init() {
	flags.RegisterRequestFlags(patchCmd)
	rootCmd.AddCommand(patchCmd)
}

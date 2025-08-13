package cmd

import (
	"github.com/im-varun/sareq/cli/cliutil/httpprinter"
	"github.com/im-varun/sareq/cli/cliutil/httprunner"
	"github.com/im-varun/sareq/cli/flags"
	"github.com/spf13/cobra"
)

var patchCmd = &cobra.Command{
	Use:     "patch <url>",
	Aliases: []string{"PATCH"},
	Short:   "Send HTTP PATCH request to the specified URL",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := httprunner.Run(cmd.Name(), args[0])
		if err != nil {
			return err
		}

		httpprinter.PrintResponse(resp)

		return nil
	},
}

func init() {
	flags.RegisterRequestFlags(patchCmd)
	rootCmd.AddCommand(patchCmd)
}

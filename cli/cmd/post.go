package cmd

import (
	"github.com/im-varun/sareq/cli/cliutil/httpprinter"
	"github.com/im-varun/sareq/cli/cliutil/httprunner"
	"github.com/im-varun/sareq/cli/flags"
	"github.com/spf13/cobra"
)

var postCmd = &cobra.Command{
	Use:     "post <url>",
	Aliases: []string{"POST"},
	Short:   "Send HTTP POST request to the specified URL",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := httprunner.Run(cmd.Name(), args[0])
		if err != nil {
			return err
		}

		httpprinter.PrintResponse(resp, flags.RespNoColor)

		return nil
	},
}

func init() {
	flags.RegisterRequestFlags(postCmd)
	flags.RegisterResponseFormattingFlags(postCmd)

	rootCmd.AddCommand(postCmd)
}

package cmd

import (
	"github.com/im-varun/sareq/cli/cliutil/httpprinter"
	"github.com/im-varun/sareq/cli/cliutil/httprunner"
	"github.com/im-varun/sareq/cli/flags"
	"github.com/spf13/cobra"
)

var putCmd = &cobra.Command{
	Use:     "put <url>",
	Aliases: []string{"PUT"},
	Short:   "Send HTTP PUT request to the specified URL",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := httprunner.Run(cmd.Name(), args[0], flags.ReqBody, flags.ReqHeader, flags.ReqTimeout)
		if err != nil {
			return err
		}

		httpprinter.PrintResponse(resp, flags.RespNoColor, flags.RespNoPrettify)

		return nil
	},
}

func init() {
	flags.RegisterRequestFlags(putCmd)
	flags.RegisterResponseFormattingFlags(putCmd)

	rootCmd.AddCommand(putCmd)
}

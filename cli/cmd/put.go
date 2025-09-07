package cmd

import (
	"github.com/im-varun/sareq/cli/cliutil/httpprinter"
	"github.com/im-varun/sareq/cli/cliutil/httprunner"
	"github.com/im-varun/sareq/cli/flags"
	"github.com/spf13/cobra"
)

var putCmd = &cobra.Command{
	Use:     "put URL --body BODY",
	Aliases: []string{"PUT"},
	Short:   "Send HTTP PUT request to the specified URL",
	Example: putCmdExample,
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := httprunner.Run(cmd.Name(), args[0], flags.RequestBody(), flags.RequestHeader(), flags.RequestTimeout())
		if err != nil {
			return err
		}

		httpprinter.Print(resp, flags.ResponseNoColor(), flags.ResponseNoPrettify())

		return nil
	},
}

func init() {
	flags.RegisterRequestFlags(putCmd)
	flags.RegisterResponseFormattingFlags(putCmd)
	flags.DisableSorting(putCmd)

	rootCmd.AddCommand(putCmd)
}

const putCmdExample string = `
# basic PUT request
sareq put https://api.example.com/users/user123 --body '{"name": "John Doe"}'

# PUT request with header
sareq put https://api.example.com/users/user123 --header "Authorization=abc123" --body '{"name": "John Doe"}'`

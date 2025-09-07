package cmd

import (
	"github.com/im-varun/sareq/cli/cliutil/httpprinter"
	"github.com/im-varun/sareq/cli/cliutil/httprunner"
	"github.com/im-varun/sareq/cli/flags"
	"github.com/spf13/cobra"
)

var postCmd = &cobra.Command{
	Use:     "post URL --body BODY",
	Aliases: []string{"POST"},
	Short:   "Send HTTP POST request to the specified URL",
	GroupID: "http request",
	Example: postCmdExample,
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
	flags.RegisterRequestFlags(postCmd)
	flags.RegisterResponseFormattingFlags(postCmd)
	flags.DisableSorting(postCmd)

	rootCmd.AddCommand(postCmd)
}

const postCmdExample string = `
# basic POST request
sareq post https://api.example.com/users --body '{"name": "John Doe"}'

# POST request with timeout of 5 seconds
sareq post https://api.example.com/users --body '{"name": "John Doe"}' --timeout 5

# POST request with headers
sareq post https://api.example.com/users --header "Authorization=abc123" --header "Content-Type=application/json" --body '{"name": "John Doe"}'`

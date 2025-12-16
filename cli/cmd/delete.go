package cmd

import (
	"github.com/im-varun/sareq/cli/cliutil/httpprinter"
	"github.com/im-varun/sareq/cli/cliutil/httprunner"
	"github.com/im-varun/sareq/cli/flags"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete URL",
	Aliases: []string{"DELETE"},
	Short:   "Send HTTP DELETE request to the specified URL",
	GroupID: "http request",
	Example: deleteCmdExample,
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
	flags.RegisterRequestFlags(deleteCmd)
	flags.RegisterResponseFormattingFlags(deleteCmd)
	flags.DisableSorting(deleteCmd)
	rootCmd.AddCommand(deleteCmd)
}

const deleteCmdExample string = `# basic DELETE request
sareq delete https://api.example.com/users/user123

# DELETE request with header
sareq delete https://api.example.com/users/user123 --header "Authorization=abc123"`

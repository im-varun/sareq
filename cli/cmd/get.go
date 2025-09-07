package cmd

import (
	"github.com/im-varun/sareq/cli/cliutil/httpprinter"
	"github.com/im-varun/sareq/cli/cliutil/httprunner"
	"github.com/im-varun/sareq/cli/flags"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:     "get URL",
	Aliases: []string{"GET"},
	Short:   "Send HTTP GET request to the specified URL",
	GroupID: "http request",
	Example: getCmdExample,
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
	flags.RegisterRequestFlags(getCmd)
	flags.RegisterResponseFormattingFlags(getCmd)
	flags.DisableSorting(getCmd)

	rootCmd.AddCommand(getCmd)
}

const getCmdExample string = `
# basic GET request
sareq get https://api.example.com/users/user123

# GET request with timeout of 5 seconds
sareq get https://api.example.com/users/user123 --timeout 5

# GET request with header
sareq get https://api.example.com/users/user123 --header "Authorization=abc123"`

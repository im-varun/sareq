package cmd

import (
	"github.com/im-varun/sareq/cli/cliutil/httpprinter"
	"github.com/im-varun/sareq/cli/cliutil/httprunner"
	"github.com/im-varun/sareq/cli/flags"
	"github.com/spf13/cobra"
)

var patchCmd = &cobra.Command{
	Use:     "patch URL --body BODY",
	Aliases: []string{"PATCH"},
	Short:   "Send HTTP PATCH request to the specified URL",
	GroupID: "http request",
	Example: patchCmdExample,
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
	flags.RegisterRequestFlags(patchCmd)
	flags.RegisterResponseFormattingFlags(patchCmd)
	flags.DisableSorting(patchCmd)
	rootCmd.AddCommand(patchCmd)
}

const patchCmdExample string = `# basic PATCH request
sareq patch https://api.example.com/users/user123 --body '{"email": "john@example.com"}'

# PATCH request with header
sareq patch https://api.example.com/users/user123 --header "Authorization=abc123" --body '{"email": "john@example.com"}'`

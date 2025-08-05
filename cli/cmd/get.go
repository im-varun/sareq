package cmd

import (
	"github.com/im-varun/sareq/cli/cliutil/httprunner"
	"github.com/im-varun/sareq/cli/flags"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:     "get <url>",
	Aliases: []string{"GET"},
	Short:   "Send HTTP GET request to the specified URL",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := httprunner.Run(cmd.Name(), args[0])
		return err
	},
}

func init() {
	flags.RegisterRequestFlags(getCmd)
	rootCmd.AddCommand(getCmd)
}

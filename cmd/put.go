package cmd

import (
	"github.com/im-varun/sareq/cmd/flags"
	"github.com/spf13/cobra"
)

var putCmd = &cobra.Command{
	Use:     "put <url>",
	Aliases: []string{"PUT"},
	Short:   "Send HTTP PUT request to the specified URL",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := runRequest(cmd.Name(), args[0])
		return err
	},
}

func init() {
	flags.RegisterRequestFlags(putCmd)
	rootCmd.AddCommand(putCmd)
}

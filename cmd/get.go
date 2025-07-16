package cmd

import (
	"github.com/im-varun/sareq/internal/httpclient"
	"github.com/spf13/cobra"
)

var timeoutFlag int

var getCmd = &cobra.Command{
	Use:     "get <url>",
	Aliases: []string{"GET"},
	Short:   "Send HTTP GET request to the specified URL",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		url, err := httpclient.ValidateRequestURL(args[0])
		if err != nil {
			return err
		}

		client := httpclient.NewClient(timeoutFlag)
		err = client.Do(cmd.Name(), url, "")
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	getCmd.Flags().IntVarP(&timeoutFlag, "timeout", "t", httpclient.DefaultTimeout, "sets the timeout for HTTP GET request")

	rootCmd.AddCommand(getCmd)
}

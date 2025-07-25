package cmd

import (
	"github.com/im-varun/sareq/cmd/flags"
	"github.com/im-varun/sareq/internal/httpclient"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:     "get <url>",
	Aliases: []string{"GET"},
	Short:   "Send HTTP GET request to the specified URL",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		reqURL, err := httpclient.ValidateRequestURL(args[0])
		if err != nil {
			return err
		}

		if flags.ReqBody != "" {
			err = httpclient.ValidateRequestBody(flags.ReqBody)
			if err != nil {
				return err
			}
		}

		client := httpclient.NewClient(flags.ReqTimeout)

		err = client.Do(cmd.Name(), reqURL, flags.ReqBody)

		return err
	},
}

func init() {
	flags.RegisterRequestFlags(getCmd)
	rootCmd.AddCommand(getCmd)
}

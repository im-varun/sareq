package flags

import (
	"github.com/im-varun/sareq/internal/httpclient"
	"github.com/spf13/cobra"
)

var ReqTimeout int
var ReqBody string
var ReqHeader map[string]string

func RegisterRequestFlags(reqCmd *cobra.Command) {
	reqCmd.Flags().IntVarP(&ReqTimeout, "timeout", "t", httpclient.DefaultTimeoutSeconds, "set timeout for HTTP request")
	reqCmd.Flags().StringVarP(&ReqBody, "body", "B", "", "set body to send with HTTP request")
	reqCmd.Flags().StringToStringVarP(&ReqHeader, "header", "H", nil, "set header to send with HTTP request")

	commandName := reqCmd.Name()
	if commandName == "post" || commandName == "put" || commandName == "patch" {
		reqCmd.MarkFlagRequired("body")
	}
}

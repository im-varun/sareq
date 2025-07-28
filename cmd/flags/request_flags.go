package flags

import (
	"github.com/im-varun/sareq/internal/httpclient"
	"github.com/spf13/cobra"
)

const (
	reqTimeoutLong    string = "timeout"
	reqTimeoutShort   string = "t"
	reqTimeoutDefault int    = httpclient.DefaultTimeoutSeconds
	reqTimeoutUsage   string = "set timeout for HTTP request"
)

const (
	reqBodyLong    string = "body"
	reqBodyShort   string = "B"
	reqBodyDefault string = ""
	reqBodyUsage   string = "set body to send with HTTP request"
)

var ReqTimeout int
var ReqBody string

func RegisterRequestFlags(reqCmd *cobra.Command) {
	reqCmd.Flags().IntVarP(&ReqTimeout, reqTimeoutLong, reqTimeoutShort, reqTimeoutDefault, reqTimeoutUsage)
	reqCmd.Flags().StringVarP(&ReqBody, reqBodyLong, reqBodyShort, reqBodyDefault, reqBodyUsage)

	commandName := reqCmd.Name()
	if commandName == "post" || commandName == "put" || commandName == "patch" {
		reqCmd.MarkFlagRequired(reqBodyLong)
	}
}

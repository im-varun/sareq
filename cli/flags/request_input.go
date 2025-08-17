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

const (
	reqHeaderLong  string = "header"
	reqHeaderShort string = "H"
	reqHeaderUsage string = "set header to sent with HTTP request"
)

var ReqTimeout int
var ReqBody string
var ReqHeader map[string]string

func RegisterRequestFlags(reqCmd *cobra.Command) {
	reqCmd.Flags().IntVarP(&ReqTimeout, reqTimeoutLong, reqTimeoutShort, reqTimeoutDefault, reqTimeoutUsage)
	reqCmd.Flags().StringVarP(&ReqBody, reqBodyLong, reqBodyShort, reqBodyDefault, reqBodyUsage)
	reqCmd.Flags().StringToStringVarP(&ReqHeader, reqHeaderLong, reqHeaderShort, nil, reqHeaderUsage)

	commandName := reqCmd.Name()
	if commandName == "post" || commandName == "put" || commandName == "patch" {
		reqCmd.MarkFlagRequired(reqBodyLong)
	}
}

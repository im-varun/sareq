package flags

import "github.com/spf13/cobra"

const (
	reqTimeoutLong    string = "timeout"
	reqTimeoutShort   string = "t"
	reqTimeoutDefault int    = 10
	reqTimeoutUsage   string = "set timeout for HTTP request"
)

const (
	reqBodyLong    string = "body"
	reqBodyShort   string = "B"
	reqBodyDefault string = ""
	reqBodyUsage   string = "initialize body to send with HTTP request"
)

var ReqTimeout int
var ReqBody string

func RegisterRequestFlags(cmd *cobra.Command) {
	cmd.Flags().IntVarP(&ReqTimeout, reqTimeoutLong, reqTimeoutShort, reqTimeoutDefault, reqTimeoutUsage)
	cmd.Flags().StringVarP(&ReqBody, reqBodyLong, reqBodyShort, reqBodyDefault, reqBodyUsage)
}

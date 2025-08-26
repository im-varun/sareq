package flags

import (
	"maps"

	"github.com/im-varun/sareq/internal/httpclient"
	"github.com/spf13/cobra"
)

var reqTimeout int
var reqBody string
var reqHeader map[string]string

func RegisterRequestFlags(reqCmd *cobra.Command) {
	reqCmd.Flags().StringVarP(&reqBody, "body", "B", "", "set body to send with HTTP request (e.g '{\"key1\": int, \"key2\": \"string\"}')")
	reqCmd.Flags().StringToStringVarP(&reqHeader, "header", "H", nil, "set header to send with HTTP request (e.g key=value or \"key=value\")")
	reqCmd.Flags().IntVarP(&reqTimeout, "timeout", "t", httpclient.DefaultTimeoutSeconds, "set timeout for HTTP request")

	commandName := reqCmd.Name()
	if commandName == "post" || commandName == "put" || commandName == "patch" {
		reqCmd.MarkFlagRequired("body")

		bodyFlag := reqCmd.Flags().Lookup("body")
		bodyFlag.Usage = bodyFlag.Usage + " (required)"
	}
}

func RequestTimeout() int {
	return reqTimeout
}

func RequestBody() string {
	return reqBody
}

func RequestHeader() map[string]string {
	headerCopy := make(map[string]string, len(reqHeader))

	maps.Copy(headerCopy, reqHeader)

	return headerCopy
}

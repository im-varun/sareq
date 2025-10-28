package flags

import (
	"maps"

	"github.com/im-varun/sareq/internal/httpclient"
	"github.com/spf13/cobra"
)

// reqTimeout stores the value passed to the --timeout (or -t) flag
var reqTimeout int

// reqBody stores the value passed to the --body (or -B) flag
var reqBody string

// reqHeader stores value passed to the --header (or -H) flag
var reqHeader map[string]string

// RegisterRequestFlags registers HTTP request related flags with the specified request command.
func RegisterRequestFlags(reqCmd *cobra.Command) {
	reqCmd.Flags().StringVarP(&reqBody, "body", "B", "", "define the body to send with HTTP request (e.g. '{\"key1\": 1, \"key2\": \"abc\"}')")
	reqCmd.Flags().StringToStringVarP(&reqHeader, "header", "H", nil, "add a header to include with HTTP request (e.g. \"key=value\")")
	reqCmd.Flags().IntVar(&reqTimeout, "timeout", httpclient.DefaultTimeoutSeconds, "specify timeout to use with HTTP request")

	commandName := reqCmd.Name()
	if commandName == "post" || commandName == "put" || commandName == "patch" {
		reqCmd.MarkFlagRequired("body")

		bodyFlag := reqCmd.Flags().Lookup("body")
		bodyFlag.Usage = bodyFlag.Usage + " (required)"
	}
}

// RequestTimeout returns the value stored in the request timeout flag (--timeout or -t).
func RequestTimeout() int {
	return reqTimeout
}

// RequestBody returns the value stored in the request body flag (--body or -B).
func RequestBody() string {
	return reqBody
}

// RequestHeader returns the value stored in the request header flag (--header or -H).
func RequestHeader() map[string]string {
	headerCopy := make(map[string]string, len(reqHeader))

	maps.Copy(headerCopy, reqHeader)

	return headerCopy
}

package flags

import "github.com/spf13/cobra"

const (
	respNoColorLong    string = "no-color"
	respNoColorDefault bool   = false
	respNoColorUsage   string = "disable coloring for HTTP response"
)

const (
	respNoPrettifyLong    string = "no-prettify"
	respNoPrettifyDefault bool   = false
	respNoPrettifyUsage   string = "disable prettification for HTTP response"
)

var RespNoColor bool
var RespNoPrettify bool

func RegisterResponseFormattingFlags(reqCmd *cobra.Command) {
	reqCmd.Flags().BoolVar(&RespNoColor, respNoColorLong, respNoColorDefault, respNoColorUsage)
	reqCmd.Flags().BoolVar(&RespNoPrettify, respNoPrettifyLong, respNoPrettifyDefault, respNoPrettifyUsage)
}

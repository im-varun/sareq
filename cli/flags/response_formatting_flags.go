package flags

import "github.com/spf13/cobra"

const (
	respNoColorLong    string = "no-color"
	respNoColorDefault bool   = false
	respNoColorUsage   string = "disable coloring for HTTP response"
)

var RespNoColor bool

func RegisterResponseFormattingFlags(reqCmd *cobra.Command) {
	reqCmd.Flags().BoolVar(&RespNoColor, respNoColorLong, respNoColorDefault, respNoColorUsage)
}

package flags

import "github.com/spf13/cobra"

var RespNoColor bool
var RespNoPrettify bool

func RegisterResponseFormattingFlags(reqCmd *cobra.Command) {
	reqCmd.Flags().BoolVar(&RespNoColor, "no-color", false, "disable coloring for HTTP response")
	reqCmd.Flags().BoolVar(&RespNoPrettify, "no-prettify", false, "disable prettification for HTTP response")
}

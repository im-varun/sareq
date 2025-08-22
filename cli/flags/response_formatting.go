package flags

import "github.com/spf13/cobra"

var respNoColor bool
var respNoPrettify bool

func RegisterResponseFormattingFlags(reqCmd *cobra.Command) {
	reqCmd.Flags().BoolVar(&respNoColor, "no-color", false, "disable coloring for HTTP response")
	reqCmd.Flags().BoolVar(&respNoPrettify, "no-prettify", false, "disable prettification for HTTP response")
}

func ResponseNoColor() bool {
	return respNoColor
}

func ResponseNoPrettify() bool {
	return respNoPrettify
}

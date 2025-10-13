package flags

import "github.com/spf13/cobra"

// respNoColor stores the value passed to the --no-color flag
var respNoColor bool

// respNoPrettify stores the value passed to the --no-prettify flag
var respNoPrettify bool

// RegisterResponseFormattingFlags registers HTTP response formatting related flags with the
// specified request command.
func RegisterResponseFormattingFlags(reqCmd *cobra.Command) {
	reqCmd.Flags().BoolVar(&respNoColor, "no-color", false, "disable coloring in HTTP response")
	reqCmd.Flags().BoolVar(&respNoPrettify, "no-prettify", false, "disable prettification in HTTP response")
}

// ResponseNoColor returns the value stored in the response no-color flag (--no-color).
func ResponseNoColor() bool {
	return respNoColor
}

// ResponseNoPrettify returns the value stored in the response no-prettify flag (--no-prettify).
func ResponseNoPrettify() bool {
	return respNoPrettify
}

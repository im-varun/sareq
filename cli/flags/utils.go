package flags

import "github.com/spf13/cobra"

// DisableSorting disables the automatic sorting of the specified command's flags during the
// display in help messages.
func DisableSorting(cmd *cobra.Command) {
	cmd.Flags().SortFlags = false
}

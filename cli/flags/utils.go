package flags

import "github.com/spf13/cobra"

func DisableSorting(cmd *cobra.Command) {
	cmd.Flags().SortFlags = false
}

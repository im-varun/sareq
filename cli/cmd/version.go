package cmd

import (
	"fmt"

	"github.com/im-varun/sareq/internal/release"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of SAReq",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(release.Version())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

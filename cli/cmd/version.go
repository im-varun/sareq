package cmd

import (
	"fmt"

	"github.com/im-varun/sareq/internal/release"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Print the version of SAReq",
	Example: versionCmdExample,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sareq version", release.Version())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

const versionCmdExample string = `# print the version of sareq
sareq version
# output: sareq version vX.Y.Z`

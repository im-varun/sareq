package cmd

import (
	"fmt"
	"os"

	"github.com/im-varun/sareq/internal/utils"
	"github.com/spf13/cobra"
)

var versionFlag bool

var rootCmd = &cobra.Command{
	Use:   "sareq",
	Short: "SAReq is a CLI-based HTTP client for modern developers",
	Run: func(cmd *cobra.Command, args []string) {
		if versionFlag {
			fmt.Println(utils.VersionString())
		} else {
			cmd.Help()
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&versionFlag, "version", "v", false, "prints the version of SAReq in use")
}

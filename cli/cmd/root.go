package cmd

import (
	"fmt"
	"os"

	"github.com/im-varun/sareq/internal/release"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "sareq",
	Short:         "SAReq is a CLI-based HTTP client for modern developers",
	Version:       release.Version(),
	SilenceErrors: true,
	SilenceUsage:  true,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	cobra.EnableCommandSorting = false
}

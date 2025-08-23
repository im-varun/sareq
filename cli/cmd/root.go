package cmd

import (
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
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

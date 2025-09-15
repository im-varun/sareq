package main

import (
	"fmt"
	"os"

	"github.com/im-varun/sareq/cli/cmd"
)

func main() {
	cliDocsDir := "./docs/cli"

	err := os.MkdirAll(cliDocsDir, 0o755)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not create directory for cli docs: %v", err)
		os.Exit(1)
	}

	root := cmd.Root()
	root.DisableAutoGenTag = true

	err = genMarkdownTree(root, cliDocsDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not generate cli docs: %v", err)
		os.Exit(1)
	}

	fmt.Println("successfully generated cli docs")
}

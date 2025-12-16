package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

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

	prep := func(filename string) string {
		base := filepath.Base(filename)
		name := strings.TrimSuffix(base, filepath.Ext(base))
		title := strings.ReplaceAll(name, "_", " ")
		return fmt.Sprintf("---\ntitle: %s\n---\n\n", title)
	}
	link := func(name string) string {
		return strings.ToLower(name)
	}
	err = genMarkdownTreeCustom(root, cliDocsDir, prep, link)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not generate cli docs: %v", err)
		os.Exit(1)
	}

	fmt.Println("successfully generated cli docs")
}

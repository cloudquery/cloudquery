package cmd

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// docCmd represents the doc markdown generation command
// This is an internal command to generate our documentation
var docCmd = &cobra.Command{
	Use:    "doc [directory_path]",
	Short:  "Generate CLI documentation markdown files",
	Args:   cobra.ExactValidArgs(1),
	Hidden: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return doc.GenMarkdownTreeCustom(rootCmd, args[0], filePrepender, linkHandler)
	},
}

func linkHandler(s string) string { return s }

func filePrepender(filename string) string {
	const fmTemplate = `---
id: "%s"
hide_title: true
sidebar_label: "%s"
---
`
	name := filepath.Base(filename)
	base := strings.TrimSuffix(name, path.Ext(name))
	id := strings.TrimPrefix(base, "cloudquery_")
	sidebarLabel := strings.ReplaceAll(id, "_", " ")
	return fmt.Sprintf(fmTemplate, id, sidebarLabel)
}

func init() {
	rootCmd.AddCommand(docCmd)
}

package cmd

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

const (
	docShort = "Generate CLI documentation markdown files"
)

func newCmdDoc() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "doc [directory_path]",
		Short:  docShort,
		Args:   cobra.ExactValidArgs(1),
		Hidden: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			// This is no danger of infinite recursion here as it just goes through the docs
			// and not running the doc command
			// nolint:revive
			return doc.GenMarkdownTreeCustom(newCmdDoc(), args[0], filePrepender, linkHandler)
		},
	}
	return cmd
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

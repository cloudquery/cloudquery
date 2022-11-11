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
		Args:   cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Hidden: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return doc.GenMarkdownTreeCustom(cmd.Parent(), args[0], filePrepender, linkHandler)
		},
	}
	return cmd
}

func linkHandler(s string) string {
	if strings.HasSuffix(s, ".md") {
		fileName := strings.TrimSuffix(s, ".md")
		fullPath := "/docs/reference/cli/" + fileName

		return fullPath
	}

	return s
}

func filePrepender(filename string) string {
	const fmTemplate = `---
title: "%s"
---
`
	name := filepath.Base(filename)
	base := strings.TrimSuffix(name, path.Ext(name))
	id := strings.TrimPrefix(base, "cloudquery_")
	return fmt.Sprintf(fmTemplate, id)
}

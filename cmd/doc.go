package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// NewCmdDoc creates the doc generation command
func NewCmdDoc() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "doc [directory_path]",
		Short:  "Generate CLI documentation markdown files",
		Args:   cobra.ExactValidArgs(1),
		Hidden: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return doc.GenMarkdownTree(newRootCmd(), args[0])
		},
	}
	return cmd
}

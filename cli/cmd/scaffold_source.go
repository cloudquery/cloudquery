package cmd

import (
	"embed"
	"fmt"
	"go/format"
	"os"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

//go:embed templates/source/*
var sourceFS embed.FS

const (
	scaffoldSourceShort = "Create an empty source plugin project"
)

func newCmdScaffoldSource() *cobra.Command {
	var outputDir string
	cmd := &cobra.Command{
		Use:   "source [org] [name]",
		Short: scaffoldSourceShort,
		Args:  cobra.ExactValidArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			if outputDir == "" {
				outputDir = "cq-source-" + args[1]
			}
			return runScaffoldSource(outputDir, args[0], args[1])
		},
	}
	cmd.Flags().StringVar(&outputDir, "output", "", "output directory")
	return cmd
}

type SourceData struct {
	Org  string
	Name string
}

var sourceTemplates = map[string]string{
	"main.go.tpl":   "main.go",
	"README.md.tpl": "README.md",
	"client.go.tpl": "client/client.go",
	"plugin.go.tpl": "plugin/plugin.go",
	"table.go.tpl": "resources/table.go",
}


func runScaffoldSource(outputDir, org, name string) error {
	var dirs = []string{"client", "plugin", "resources"}
	for _, dir := range dirs {
		if err := os.MkdirAll(outputDir+"/" + dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}
	tpl, err := template.New("source").ParseFS(sourceFS,
		"templates/source/*.tpl",
		"templates/source/plugin/*.tpl",
		"templates/source/client/*.tpl",
		"templates/source/resources/*.tpl",
	)
	if err != nil {
		return fmt.Errorf("failed to parse templates: %w", err)
	}
	data := SourceData{
		Org:  org,
		Name: name,
	}
	for templatePath, filePath := range sourceTemplates {
		var sb strings.Builder
		if err := tpl.ExecuteTemplate(&sb, templatePath, data); err != nil {
			return fmt.Errorf("failed to execute template: %w", err)
		}
		content := []byte(sb.String())
		if strings.HasSuffix(filePath, ".go") {
			formattedContent, err := format.Source(content)
			if err != nil {
				// we still write the file even if it's not formatted for easy debugging
				os.WriteFile(outputDir+"/"+filePath, content, 0644)
				return fmt.Errorf("failed to format source %s: %w", filePath, err)
			}
			content = formattedContent
		}
		if err := os.WriteFile(outputDir+"/"+filePath, content, 0644); err != nil {
			return fmt.Errorf("failed to write file: %w", err)
		}
	}
	return nil
}

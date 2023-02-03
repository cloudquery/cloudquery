package cmd

import (
	"embed"
	"fmt"
	"go/format"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

const (
	scaffoldSourceShort = "Create an empty source plugin project"
)

func newCmdScaffoldSource() *cobra.Command {
	var outputDir string
	cmd := &cobra.Command{
		Use:   "source [org] [name]",
		Short: scaffoldSourceShort,
		Args:  cobra.MatchAll(cobra.ExactArgs(2)),
		RunE: func(cmd *cobra.Command, args []string) error {
			if outputDir == "" {
				outputDir = "cq-source-" + args[1]
			}
			return runScaffoldSource(args[0], args[1], outputDir)
		},
	}
	cmd.Flags().StringVar(&outputDir, "output", "", "output directory")
	return cmd
}

var scaffoldTemplates = map[string]string{
	"release.yaml.tpl":     ".github/workflows/release.yaml",
	"test.yaml.tpl":        ".github/workflows/test.yaml",
	".goreleaser.yaml.tpl": ".goreleaser.yaml",
	"go.mod.tpl":           "go.mod",
	"main.go.tpl":          "main.go",
	"Makefile.tpl":         "Makefile",
	"README.md.tpl":        "README.md",
	"client.go.tpl":        "client/client.go",
	"spec.go.tpl":          "client/spec.go",
	"plugin.go.tpl":        "plugin/plugin.go",
	"table.go.tpl":         "resources/table.go",
	".gitignore.tpl":       ".gitignore",
}

//go:embed templates/source/*
var sourceFS embed.FS

type scaffoldData struct {
	Org  string
	Name string
}

func runScaffoldSource(org string, name string, outputDir string) error {
	tpl, err := template.New("source").ParseFS(sourceFS,
		"templates/source/.github/workflows/*.tpl",
		"templates/source/*.tpl",
		"templates/source/plugin/*.tpl",
		"templates/source/client/*.tpl",
		"templates/source/resources/*.tpl",
	)
	if err != nil {
		return fmt.Errorf("failed to parse templates: %w", err)
	}

	data := scaffoldData{
		Org:  org,
		Name: name,
	}
	for templatePath, filePath := range scaffoldTemplates {
		var sb strings.Builder
		if err := tpl.ExecuteTemplate(&sb, templatePath, data); err != nil {
			return fmt.Errorf("failed to execute template: %w", err)
		}
		content := []byte(sb.String())
		fullPath := outputDir + "/" + filePath
		baseDir := path.Dir(fullPath)
		if err := os.MkdirAll(baseDir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", baseDir, err)
		}
		if strings.HasSuffix(filePath, ".go") {
			formattedContent, err := format.Source(content)
			if err != nil {
				// we still write the file even if it's not formatted for easy debugging
				_ = os.WriteFile(outputDir+"/"+filePath, content, 0644)
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

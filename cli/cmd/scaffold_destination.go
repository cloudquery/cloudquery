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

//go:embed templates/*
var destinationFS embed.FS

const (
	scaffoldDestinationShort = "Create an empty destination plugin project"
)

func newCmdScaffoldDestination() *cobra.Command {
	var outputDir string
	cmd := &cobra.Command{
		Use:   "destination [org] [name]",
		Short: scaffoldDestinationShort,
		Args:  cobra.ExactValidArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			if outputDir == "" {
				outputDir = "cq-destination-" + args[1]
			}
			return runScaffoldDestination(outputDir, args[0], args[1])
		},
	}
	cmd.Flags().StringVar(&outputDir, "output", "", "output directory")
	return cmd
}

type DestinationData struct {
	Org  string
	Name string
}

var destinationTemplates = map[string]string{
	"main.go.tpl":   "main.go",
	"README.md.tpl": "README.md",
	"plugin.go.tpl": "plugin/plugin.go",
}

func runScaffoldDestination(outputDir, org, name string) error {
	if err := os.MkdirAll(outputDir+"/plugin", 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	tpl, err := template.New("destination").ParseFS(destinationFS, "templates/destination/*.tpl", "templates/destination/plugin/*.tpl")
	if err != nil {
		return fmt.Errorf("failed to parse templates: %w", err)
	}
	data := DestinationData{
		Org:  org,
		Name: name,
	}
	for templatePath, filePath := range destinationTemplates {
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
				return fmt.Errorf("failed to format source: %w", err)
			}
			content = formattedContent
		}
		if err := os.WriteFile(outputDir+"/"+filePath, content, 0644); err != nil {
			return fmt.Errorf("failed to write file: %w", err)
		}
	}
	return nil
}

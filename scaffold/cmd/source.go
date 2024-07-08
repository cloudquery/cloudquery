package cmd

import (
	"embed"
	"fmt"
	"go/format"
	"io/fs"
	"os"
	"path"
	"path/filepath"
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

//go:embed templates/source/*
var sourceFS embed.FS

type scaffoldData struct {
	Org  string
	Name string
}

func runScaffoldSource(org string, name string, outputDir string) error {
	data := scaffoldData{
		Org:  org,
		Name: name,
	}

	err := fs.WalkDir(sourceFS, "templates/source", func(templatePath string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("failed to walk directory: %w", err)
		}
		if d.IsDir() {
			return nil
		}
		if strings.HasSuffix(templatePath, ".tpl") {
			tpl, err := template.New(filepath.Base(templatePath)).ParseFS(sourceFS, templatePath)
			if err != nil {
				return fmt.Errorf("failed to parse template %s: %w", templatePath, err)
			}
			outputPath := strings.TrimSuffix(strings.TrimPrefix(templatePath, "templates/source"), ".tpl")
			var sb strings.Builder
			if err := tpl.Execute(&sb, data); err != nil {
				return fmt.Errorf("failed to execute template: %w", err)
			}
			content := []byte(sb.String())
			fullPath := outputDir + "/" + outputPath
			baseDir := path.Dir(fullPath)
			if err := os.MkdirAll(baseDir, 0755); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", baseDir, err)
			}
			if strings.HasSuffix(outputPath, ".go") {
				formattedContent, err := format.Source(content)
				if err != nil {
					// we still write the file even if it's not formatted for easy debugging
					_ = os.WriteFile(outputDir+"/"+outputPath, content, 0644)
					return fmt.Errorf("failed to format source %s: %w", outputPath, err)
				}
				content = formattedContent
			}
			if err := os.WriteFile(outputDir+"/"+outputPath, content, 0644); err != nil {
				return fmt.Errorf("failed to write file: %w", err)
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	
	n := len(data.Name)
	fmt.Println("------------------------------------------------")
	fmt.Printf("Successfully created new plugin under %s 🎉\n\n", outputDir)
	fmt.Printf("Next steps:\n")
	fmt.Printf("1. cd %s\n", outputDir)
	fmt.Printf("2. go mod tidy             %s# fetch dependencies\n", strings.Repeat(" ", n))
	fmt.Printf("3. go build .              %s# build the plugin\n", strings.Repeat(" ", n))
	fmt.Printf("4. ./cq-source-%s serve      # run the plugin as a gRPC server\n\n", data.Name)
	fmt.Printf("------------------------------------------------\n\n")
	fmt.Printf("For more information, see the README.md in the plugin directory.\n\n")
	fmt.Println("Developer guide: https://cql.ink/go-source-plugin-developer-guide")

	return nil
}

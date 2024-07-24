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

// Note: The long list here is deliberate, to make sure that node_modules does not accidentally
// get included in the binary.
//
//go:embed templates/source/*
//go:embed templates/cloud-config-ui/public/*
//go:embed templates/cloud-config-ui/src/*
//go:embed templates/cloud-config-ui/.eslintrc.json.tpl
//go:embed templates/cloud-config-ui/.prettierrc.tpl
//go:embed templates/cloud-config-ui/.gitignore.tpl
//go:embed templates/cloud-config-ui/.nvmrc.tpl
//go:embed templates/cloud-config-ui/package.json.tpl
//go:embed templates/cloud-config-ui/README.md.tpl
//go:embed templates/cloud-config-ui/tsconfig.json.tpl
var sourceFS embed.FS

type scaffoldData struct {
	Org  string
	Name string
	Kind string
}

func runScaffoldSource(org string, name string, outputDir string) error {
	data := scaffoldData{
		Org:  org,
		Name: name,
		Kind: "source",
	}
	if err := copyGoFiles(data, outputDir); err != nil {
		return fmt.Errorf("failed to copy go files: %w", err)
	}
	if err := copyConfigUIFiles(data, outputDir); err != nil {
		return fmt.Errorf("failed to copy config ui files: %w", err)
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

func copyGoFiles(data scaffoldData, outputDir string) error {
	return fs.WalkDir(sourceFS, "templates/source", func(fpath string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("failed to walk directory: %w", err)
		}
		if d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(fpath, ".tpl") {
			return nil
		}
		outputPath := strings.TrimSuffix(strings.TrimPrefix(fpath, "templates/source"), ".tpl")
		fullPath := outputDir + "/" + outputPath
		err = writeTemplate(data, fpath, fullPath)
		if err != nil {
			return fmt.Errorf("failed to write template: %w", err)
		}
		return nil
	})
}

func copyConfigUIFiles(data scaffoldData, outputDir string) error {
	return fs.WalkDir(sourceFS, "templates/cloud-config-ui", func(fpath string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("failed to walk directory: %w", err)
		}
		if d.IsDir() {
			return nil
		}
		if strings.HasSuffix(fpath, ".tpl") {
			outputPath := strings.TrimSuffix(strings.TrimPrefix(fpath, "templates/"), ".tpl")
			fullPath := outputDir + "/" + outputPath
			err = writeTemplate(data, fpath, fullPath)
			if err != nil {
				return fmt.Errorf("failed to write template: %w", err)
			}
			return nil
		}
		fullPath := outputDir + "/" + strings.TrimPrefix(fpath, "templates")
		return copyFile(fpath, fullPath)
	})
}

func copyFile(inputPath string, outputPath string) error {
	b, err := fs.ReadFile(sourceFS, inputPath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}
	baseDir := path.Dir(outputPath)
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", baseDir, err)
	}
	if err := os.WriteFile(outputPath, b, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	return nil
}

func writeTemplate(data scaffoldData, tplPath string, outputPath string) error {
	tpl, err := template.New(filepath.Base(tplPath)).ParseFS(sourceFS, tplPath)
	if err != nil {
		return fmt.Errorf("failed to parse template %s: %w", tplPath, err)
	}
	var sb strings.Builder
	if err := tpl.Execute(&sb, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}
	content := []byte(sb.String())
	baseDir := path.Dir(outputPath)
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", baseDir, err)
	}
	if strings.HasSuffix(outputPath, ".go") {
		formattedContent, err := format.Source(content)
		if err != nil {
			// we still write the file even if it's not formatted for easy debugging
			_ = os.WriteFile(outputPath, content, 0644)
			return fmt.Errorf("failed to format source %s: %w", outputPath, err)
		}
		content = formattedContent
	}
	if err := os.WriteFile(outputPath, content, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	return nil
}

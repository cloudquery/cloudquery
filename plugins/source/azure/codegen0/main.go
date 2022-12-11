package main

import (
	"bytes"
	"embed"
	"fmt"
	"go/format"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
	"text/template"

	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen0/internal/azparser"
	"github.com/iancoleman/strcase"
)

//go:embed templates/*.go.tpl
var templateFS embed.FS

var (
	currentFilename string
	currentDir      string
)

// Module is a struct that contains all the information needed to generate
// cloudquery code for a given Azure SDK package.
type Recipe struct {
	Tables   []*azparser.Table
	Import   string
	BaseName string
}

func main() {
	var ok bool
	_, currentFilename, _, ok = runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	currentDir = path.Dir(currentFilename)
	var updateGoMod bool
	if len(os.Args) > 1 && os.Args[1] == "--update-go-mod" {
		updateGoMod = true
	}

	if updateGoMod {
		packagesToGoGet, err := azparser.DiscoverSubpackages()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("go getting %d packages\n", len(packagesToGoGet))
		args := []string{"get", "-u"}
		args = append(args, packagesToGoGet...)
		cmd := exec.Command("go", args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}

	armModules, err := azparser.GetArmModules(path.Join(currentDir, "../go.mod"))
	if err != nil {
		log.Fatal(err)
	}
	namespaces := make(map[string]string, 0)
	for _, armModule := range armModules {
		tables, err := azparser.CreateTablesFromPackage(armModule)
		if err != nil {
			log.Fatal(err)
		}
		if len(tables) == 0 {
			continue
		}
		importPath := strings.Split(armModule, "@")[0]
		mod := &Recipe{
			Tables:   tables,
			Import:   importPath,
			BaseName: path.Base(importPath),
		}
		if err := generatePackage(armModule, mod); err != nil {
			log.Fatal(err)
		}
		for _, table := range tables {
			if table.Namespace == "" {
				panic(fmt.Sprintf("table %s has no namespace %s %s", table.NewFuncName, table.URL, importPath))
			}
			namespaces[strings.ReplaceAll(table.Namespace, ".", "_")] = table.Namespace
		}
	}
	if err := generateNamespaces(namespaces); err != nil {
		log.Fatal(err)
	}
}

func generatePackage(pkg string, mod *Recipe) error {
	tpl, err := template.New("recipe.go.tpl").Funcs(template.FuncMap{
		"ToCamel": strcase.ToCamel,
	}).ParseFS(templateFS, "templates/recipe.go.tpl")
	if err != nil {
		return fmt.Errorf("failed to parse recipe.go.tpl: %w", err)
	}

	var buff bytes.Buffer
	if err := tpl.Execute(&buff, mod); err != nil {
		return fmt.Errorf("failed to execute recipe template: %w", err)
	}
	basename := strings.TrimPrefix(mod.BaseName, "arm")
	filePath := path.Join(currentDir, "../codegen1/recipes", basename+".go")
	content := buff.Bytes()
	formattedContent, err := format.Source(content)
	if err != nil {
		fmt.Printf("failed to format code for %s: %v\n", filePath, err)
	} else {
		content = formattedContent
	}
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}
	return nil
}

func generateNamespaces(namespaces map[string]string) error {
	tpl, err := template.New("namespaces.go.tpl").Funcs(template.FuncMap{
		"ToCamel": strcase.ToCamel,
	}).ParseFS(templateFS, "templates/namespaces.go.tpl")
	if err != nil {
		return fmt.Errorf("failed to parse recipe.go.tpl: %w", err)
	}

	var buff bytes.Buffer
	if err := tpl.Execute(&buff, namespaces); err != nil {
		return fmt.Errorf("failed to execute recipe template: %w", err)
	}
	filePath := path.Join(currentDir, "../client/namespaces.go")
	content := buff.Bytes()
	formattedContent, err := format.Source(content)
	if err != nil {
		fmt.Printf("failed to format code for %s: %v\n", filePath, err)
	} else {
		content = formattedContent
	}
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}
	return nil
}

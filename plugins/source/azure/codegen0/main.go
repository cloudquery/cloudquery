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

var templateFuncs = template.FuncMap{
	"ToCamel": strcase.ToCamel,
	"ToLower": strings.ToLower,
}

//go:embed templates/*.go.tpl
var templateFS embed.FS

var (
	currentFilename string
	currentDir      string
)

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
	var allTables []*azparser.Table
	for _, armModule := range armModules {
		tables, err := azparser.CreateTablesFromPackage(armModule)
		if err != nil {
			log.Fatal(err)
		}
		if len(tables) == 0 {
			continue
		}
		importPath := strings.Split(armModule, "@")[0]
		for _, table := range tables {
			if table.Namespace == "" {
				panic(fmt.Sprintf("table %s has no namespace %s %s", table.NewFuncName, table.URL, importPath))
			}
			namespaces[strings.ReplaceAll(table.Namespace, ".", "_")] = table.Namespace
		}
		for _, table := range tables {
			table.Name = strings.TrimPrefix(table.NewFuncName, "New")
			table.Name = strcase.ToSnake(strings.TrimSuffix(table.Name, "Client"))
			table.BaseImportPath = path.Base(importPath)
			table.PackageName = strings.TrimPrefix(path.Base(importPath), "arm")
			table.ImportPath = importPath
			table.NamespaceConst = strings.ReplaceAll(table.Namespace, ".", "_")
			if len(table.NewClientParams) == 3 {
				table.NewFuncHasSubscriptionId = true
			}
			if err := generateTable(table); err != nil {
				log.Fatal(err)
			}
		}
		allTables = append(allTables, tables...)
	}
	if err := generateTables(allTables); err != nil {
		log.Fatal(err)
	}
	if err := generateNamespaces(namespaces); err != nil {
		log.Fatal(err)
	}
}

func generateNamespaces(namespaces map[string]string) error {
	tpl, err := template.New("namespaces.go.tpl").Funcs(template.FuncMap{
		"ToCamel": strcase.ToCamel,
	}).ParseFS(templateFS, "templates/namespaces.go.tpl")
	if err != nil {
		return fmt.Errorf("failed to parse namespaces.go.tpl: %w", err)
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

func generateTable(r *azparser.Table) error {
	tpl, err := template.New("list.go.tpl").Funcs(templateFuncs).ParseFS(templateFS, "templates/*.go.tpl")
	if err != nil {
		return fmt.Errorf("failed to parse templates: %w", err)
	}
	var buff bytes.Buffer
	if err := tpl.Execute(&buff, r); err != nil {
		return fmt.Errorf("failed to execute template for  %w", err)
	}

	filePath := path.Join(currentDir, "../resources/services", r.PackageName)
	if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
		return err
	}
	name := r.Name
	if strings.HasSuffix(name, "_test") {
		name = name + "_not"
	}
	filePath = path.Join(filePath, name+".go")

	content := buff.Bytes()
	formattedContent, err := format.Source(buff.Bytes())
	if err != nil {
		fmt.Printf("failed to format code for %s: %v\n", filePath, err)
	} else {
		content = formattedContent
	}
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}
	if err := generateTableMock(r); err != nil {
		return err
	}
	return nil
}

func generateTableMock(r *azparser.Table) error {
	tpl, err := template.New("list_mock.go.tpl").Funcs(templateFuncs).ParseFS(templateFS, "templates/*.go.tpl")
	if err != nil {
		return fmt.Errorf("failed to parse templates: %w", err)
	}
	var buff bytes.Buffer
	if err := tpl.Execute(&buff, r); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	filePath := path.Join(currentDir, "../resources/services", r.PackageName)
	if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
		return err
	}
	filePath = path.Join(filePath, r.Name+"_mock_test.go")

	content := buff.Bytes()
	formattedContent, err := format.Source(buff.Bytes())
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

func generateTables(rr []*azparser.Table) error {
	tpl, err := template.New("tables.go.tpl").Funcs(templateFuncs).ParseFS(templateFS, "templates/tables.go.tpl")
	if err != nil {
		return fmt.Errorf("failed to parse services.go.tpl: %w", err)
	}

	var buff bytes.Buffer
	if err := tpl.Execute(&buff, rr); err != nil {
		return fmt.Errorf("failed to execute services template: %w", err)
	}

	filePath := path.Join(currentDir, "../resources/plugin/tables.go")
	content := buff.Bytes()
	formattedContent, err := format.Source(buff.Bytes())
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

package main

import (
	"bytes"
	"embed"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
	"text/template"

	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen0/internal/azparser"
)

//go:embed templates/*.go.tpl
var templateFS embed.FS

var (
	currentFilename string
	currentDir string
)

// Module is a struct that contains all the information needed to generate
// cloudquery code for a given Azure SDK package.
type Module struct {
	Tables []*azparser.Table
	Import string
	BaseName string
}

func main() {
	var ok bool
	_, currentFilename, _, ok = runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	currentDir = path.Dir(currentFilename)
	
	armModules, err := azparser.GetArmModules(path.Join(currentDir, "../go.mod"))
	if err != nil {
		log.Fatal(err)
	}
	for _, armModule := range armModules {
		tables, err := azparser.CreateTablesFromPackage(armModule)
		if err != nil {
			log.Fatal(err)
		}
		if len(tables) == 0 {
			continue
		}
		importPath := strings.Split(armModule, "@")[0]
		mod := &Module{
			Tables: tables,
			Import: importPath,
			BaseName: path.Base(importPath),
		}
		if err := generatePackage(armModule, mod); err != nil {
			log.Fatal(err)
		}
	}
}

func generatePackage(pkg string, mod *Module) error {
	tpl, err := template.New("package.go.tpl").Funcs(template.FuncMap{
		"ToCamel": strings.Title,
	}).ParseFS(templateFS, "templates/package.go.tpl")
	if err != nil {
		return fmt.Errorf("failed to parse package.go.tpl: %w", err)
	}

	var buff bytes.Buffer
	if err := tpl.Execute(&buff, mod); err != nil {
		return fmt.Errorf("failed to execute package template: %w", err)
	}
	filePath := path.Join(currentDir, "../codegen1/packages", mod.BaseName+".go")
	if err := os.WriteFile(filePath, buff.Bytes(), 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}
	return nil
}




package main

import (
	"bytes"
	"embed"
	"fmt"
	"go/format"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
	"text/template"

	"github.com/cloudquery/plugins/source/gcp/codegen/client"
	"github.com/iancoleman/strcase"
)

//go:embed templates/*.go.tpl
var gcpTemplatesFS embed.FS

func main() {
	genServices()
}

func genServices() {
	fmt.Println("Generating services")
	serviceList, err := client.Discover()
	if err != nil {
		log.Fatal(err)
	}
	generateTemplate("services.go.tpl", "client/services.go", serviceList)
}

func generateTemplate(name string, output string, data any) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	dir := path.Dir(filename)
	tpl, err := template.New(name).Funcs(template.FuncMap{
		"ToCamel": strcase.ToCamel,
		"ToLower": strings.ToLower,
	}).ParseFS(gcpTemplatesFS, "templates/"+name)

	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse %s: %w", name, err))
	}

	var buff bytes.Buffer
	if err := tpl.Execute(&buff, data); err != nil {
		log.Fatal(fmt.Errorf("failed to execute template: %w", err))
	}

	filePath := path.Join(dir, "..", output)
	if err := os.MkdirAll(path.Dir(filePath), os.ModePerm); err != nil {
		log.Fatal(fmt.Errorf("failed to create directory %s: %w", filePath, err))
	}
	content, err := format.Source(buff.Bytes())
	if err != nil {
		log.Fatal(fmt.Errorf("failed to format code for %s: %w", filePath, err))
	}
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		log.Fatal(fmt.Errorf("failed to write file %s: %w", filePath, err))
	}
}

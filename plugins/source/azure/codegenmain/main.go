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

	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen"
	sdkgen "github.com/cloudquery/plugin-sdk/codegen"

	"github.com/iancoleman/strcase"
)

//go:embed templates/*.go.tpl
var azureTemplatesFS embed.FS

func main() {
	var resources = []codegen.Resource{}
	resources = append(resources, codegen.NetworkResources()...)
	for _, r := range resources {
		generateResource(r)
	}
}

func getFilename() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	return filename
}

func initTemplate(name string) *template.Template {
	templateNmae := name + ".go.tpl"
	tpl, err := template.New(templateNmae).Funcs(template.FuncMap{"ToCamel": strcase.ToCamel}).ParseFS(azureTemplatesFS, "templates/"+templateNmae)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse azure templates: %w", err))
	}

	tpl, err = tpl.ParseFS(sdkgen.TemplatesFS, "templates/*.go.tpl")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse codegen template: %w", err))
	}

	return tpl
}

func getDestination(dir string, t string, r codegen.Resource) string {
	filePath := path.Join(dir, "../resources/servicesv2", r.AzureService)
	if strings.HasSuffix(t, "_mock_test") {
		filePath = path.Join(filePath, r.AzureSubService+"_mock_test.go")
	} else {
		filePath = path.Join(filePath, r.AzureSubService+".go")
	}

	return filePath
}

func getContent(t string, filePath string, r codegen.Resource) []byte {
	tpl := initTemplate(t)
	var buff bytes.Buffer
	if err := tpl.Execute(&buff, r); err != nil {
		log.Fatal(fmt.Errorf("failed to execute template: %w", err))
	}
	content, err := format.Source(buff.Bytes())
	if err != nil {
		fmt.Println(buff.String())
		log.Fatal(fmt.Errorf("failed to format code for %s: %w", filePath, err))
	}
	return content
}

func writeContent(filePath string, content []byte) {
	if err := os.MkdirAll(path.Dir(filePath), 0755); err != nil {
		log.Fatal(fmt.Errorf("failed to create directory for file %s: %w", filePath, err))
	}
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		log.Fatal(fmt.Errorf("failed to write file %s: %w", filePath, err))
	}
}

func generateResource(r codegen.Resource) {
	filename := getFilename()
	dir := path.Dir(filename)
	for _, t := range r.Templates {
		filePath := getDestination(dir, t, r)
		content := getContent(t, filePath, r)
		writeContent(filePath, content)
	}

}

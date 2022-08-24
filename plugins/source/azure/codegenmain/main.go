package main

import (
	"bytes"
	"embed"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"
	"text/template"

	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen"
	sdkgen "github.com/cloudquery/plugin-sdk/codegen"
)

//go:embed templates/*.go.tpl
var azureTemplatesFS embed.FS

func main() {
	var resources = []codegen.Resource{}
	resources = append(resources, codegen.AuthorizationResources()...)
	resources = append(resources, codegen.BatchResources()...)
	resources = append(resources, codegen.CDNResources()...)
	resources = append(resources, codegen.ComputeResources()...)
	resources = append(resources, codegen.NetworkResources()...)
	resources = append(resources, codegen.KeyValueResources()...)
	resources = append(resources, codegen.LogicResources()...)
	resources = append(resources, codegen.MariaDbResources()...)
	resources = append(resources, codegen.MonitorResources()...)
	resources = append(resources, codegen.MySQLResources()...)
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

func initTemplate(templateName string) *template.Template {
	tpl, err := template.New(templateName).ParseFS(azureTemplatesFS, "templates/*.go.tpl")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse azure templates: %w", err))
	}

	tpl, err = tpl.ParseFS(sdkgen.TemplatesFS, "templates/*.go.tpl")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse codegen template: %w", err))
	}

	return tpl
}

func getContent(t codegen.Template, destination string, r codegen.Resource) []byte {
	tpl := initTemplate(t.Source)
	var buff bytes.Buffer
	if err := tpl.Execute(&buff, r); err != nil {
		log.Fatal(fmt.Errorf("failed to execute template: %w", err))
	}
	return buff.Bytes()
}

func writeContent(destination string, content []byte) {
	if err := os.MkdirAll(path.Dir(destination), 0755); err != nil {
		log.Fatal(fmt.Errorf("failed to create directory for file %s: %w", destination, err))
	}
	if err := os.WriteFile(destination, content, 0644); err != nil {
		log.Fatal(fmt.Errorf("failed to write file %s: %w", destination, err))
	}
}

func generateResource(r codegen.Resource) {
	filename := getFilename()
	dir := path.Dir(filename)
	destination := path.Join(dir, "../resources/servicesv2", r.Template.Destination)
	content := getContent(r.Template, destination, r)
	writeContent(destination, content)

	exec.Command("goimports", "-w", destination).Run()
}

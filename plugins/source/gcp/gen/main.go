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
	"text/template"

	sdkgen "github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugins/source/gcp/codegen"
)

//go:embed templates/*.go.tpl
var gcpTemplatesFS embed.FS

var resources = []codegen.Resource{}

func main() {

	resources = append(resources, codegen.ComputeResources...)
	resources = append(resources, codegen.CloudFunctionsResources...)
	resources = append(resources, codegen.CloudRunResources...)

	var err error
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	dir := path.Dir(filename)
	for _, r := range resources {
		r.Table, err = sdkgen.NewTableFromStruct("gcp_cloudfunctions_functions", r.Struct)
		if err != nil {
			log.Fatal(err)
		}
		r.Table.Multiplex = "client.ProjectMultiplex"
		r.Table.Resolver = "fetch" + r.TableFunctionName
		tpl, err := template.New(r.Template).ParseFS(gcpTemplatesFS, "templates/"+r.Template)
		if err != nil {
			log.Fatal(fmt.Errorf("failed to parse gcp templates: %w", err))
		}
		tpl, err = tpl.ParseFS(sdkgen.TemplatesFS, "templates/*.go.tpl")
		if err != nil {
			log.Fatal(fmt.Errorf("failed to parse codegen template: %w", err))
		}
		var buff bytes.Buffer
		if err := tpl.Execute(&buff, r); err != nil {
			log.Fatal(fmt.Errorf("failed to execute template: %w", err))
		}
		filePath := path.Join(dir, "../resources/servicesv2", r.PackageName, r.FileName)
		content, err := format.Source(buff.Bytes())
		if err != nil {
			log.Fatal(fmt.Errorf("failed to format code: %w", err))
		}
		if err := os.WriteFile(filePath, content, 0644); err != nil {
			log.Fatal(fmt.Errorf("failed to write file %s: %w", filePath, err))
		}
	}
}

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
	"github.com/iancoleman/strcase"
)

//go:embed templates/*.go.tpl
var gcpTemplatesFS embed.FS

var resources = []*codegen.Resource{}

func main() {
	resources = append(resources, codegen.ComputeResources()...)
	// resources = append(resources, codegen.DnsResources()...)
	// resources = append(resources, codegen.DomainsResources()...)
	// resources = append(resources, codegen.IamResources()...)
	// resources = append(resources, codegen.KmsResources()...)
	// resources = append(resources, codegen.KubernetesResources()...)
	// resources = append(resources, codegen.LoggingResources()...)
	// resources = append(resources, codegen.RedisResources()...)
	// resources = append(resources, codegen.MonitoringResources()...)
	// resources = append(resources, codegen.SecretManagerResources()...)
	// resources = append(resources, codegen.ServiceusageResources()...)
	// resources = append(resources, codegen.SqlResources()...)
	// resources = append(resources, codegen.StorageResources()...)
	// resources = append(resources, codegen.CloudFunctionsResources()...)
	resources = append(resources, codegen.BigqueryResources()...)
	// resources = append(resources, codegen.CloudBillingResources()...)
	// resources = append(resources, codegen.CloudRunResources...)

	for _, r := range resources {
		generateResource(*r, false)
		generateResource(*r, true)
	}
}

func generateResource(r codegen.Resource, mock bool) {
	var err error
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	dir := path.Dir(filename)
	r.Table, err = sdkgen.NewTableFromStruct(
		fmt.Sprintf("gcp_%s_%s", r.Service, r.SubService),
		r.Struct,
		sdkgen.WithSkipFields(r.SkipFields),
		sdkgen.WithOverrideColumns(r.OverrideColumns))
	if err != nil {
		log.Fatal(err)
	}
	r.Table.Columns = append(r.DefaultColumns, r.Table.Columns...)
	r.Table.Multiplex = "client.ProjectMultiplex"
	r.Table.Resolver = "fetch" + strcase.ToCamel(r.SubService)
	mainTemplate := r.Template + ".go.tpl"
	if mock {
		mainTemplate = r.Template + "_mock_test.go.tpl"
	}
	tpl, err := template.New(mainTemplate).Funcs(template.FuncMap{
		"ToCamel": strcase.ToCamel,
	}).ParseFS(gcpTemplatesFS, "templates/"+mainTemplate)
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
	filePath := path.Join(dir, "../resources/servicesv2", r.Service)
	if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	if mock {
		filePath = path.Join(filePath, r.SubService+"_mock_test.go")
	} else {
		filePath = path.Join(filePath, r.SubService+".go")
	}

	content, err := format.Source(buff.Bytes())
	if err != nil {
		fmt.Println(buff.String())
		log.Fatal(fmt.Errorf("failed to format code for %s: %w", filePath, err))
	}
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		log.Fatal(fmt.Errorf("failed to write file %s: %w", filePath, err))
	}
}

package main

import (
	"bytes"
	"embed"
	"fmt"
	"go/format"
	"log"
	"os"
	"path"
	"reflect"
	"runtime"
	"strings"
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
	resources = append(resources, codegen.DnsResources()...)
	resources = append(resources, codegen.DomainsResources()...)
	resources = append(resources, codegen.IamResources()...)
	resources = append(resources, codegen.KmsResources()...)
	resources = append(resources, codegen.KubernetesResources()...)
	resources = append(resources, codegen.LoggingResources()...)
	resources = append(resources, codegen.RedisResources()...)
	resources = append(resources, codegen.MonitoringResources()...)
	resources = append(resources, codegen.SecretManagerResources()...)
	resources = append(resources, codegen.ServiceusageResources()...)
	resources = append(resources, codegen.SqlResources()...)
	resources = append(resources, codegen.StorageResources()...)
	resources = append(resources, codegen.CloudFunctionsResources()...)
	resources = append(resources, codegen.BigqueryResources()...)
	resources = append(resources, codegen.CloudBillingResources()...)
	resources = append(resources, codegen.CloudResourceManagerResources()...)
	// resources = append(resources, codegen.RunResources()...)

	for _, r := range resources {
		generateResource(*r, false)
		if !r.SkipMock {
			generateResource(*r, true)
		}
	}
	generatePlugin(resources)
}

func generatePlugin(rr []*codegen.Resource) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	dir := path.Dir(filename)
	tpl, err := template.New("autogen_tables.go.tpl").Funcs(template.FuncMap{
		"ToCamel": strcase.ToCamel,
		"ToLower": strings.ToLower,
	}).ParseFS(gcpTemplatesFS, "templates/autogen_tables.go.tpl")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse autogen_tables.go.tpl: %w", err))
	}

	var buff bytes.Buffer
	if err := tpl.Execute(&buff, rr); err != nil {
		log.Fatal(fmt.Errorf("failed to execute template: %w", err))
	}

	filePath := path.Join(dir, "../resources/plugin/autogen_tables.go")
	content, err := format.Source(buff.Bytes())
	if err != nil {
		fmt.Println(buff.String())
		log.Fatal(fmt.Errorf("failed to format code for %s: %w", filePath, err))
	}
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		log.Fatal(fmt.Errorf("failed to write file %s: %w", filePath, err))
	}
}

func generateResource(r codegen.Resource, mock bool) {
	var err error
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	dir := path.Dir(filename)
	if r.SkipFields == nil {
		r.SkipFields = []string{"ServerResponse", "NullFields", "ForceSendFields"}
	}
	if r.OutputField == "" {
		r.OutputField = "Items"
	}
	if r.DefaultColumns == nil {
		r.DefaultColumns = []sdkgen.ColumnDefinition{codegen.ProjectIdColumn}
	}
	if r.StructName == "" {
		r.StructName = reflect.TypeOf(r.Struct).Elem().Name()
	}
	if r.MockListStruct == "" {
		r.MockListStruct = strcase.ToCamel(r.StructName)
	}

	if r.MockImports == nil {
		r.MockImports = []string{reflect.TypeOf(r.Struct).Elem().PkgPath()}
	}

	r.Table, err = sdkgen.NewTableFromStruct(
		fmt.Sprintf("gcp_%s_%s", r.Service, r.SubService),
		r.Struct,
		sdkgen.WithSkipFields(r.SkipFields),
		sdkgen.WithOverrideColumns(r.OverrideColumns),
		sdkgen.WithExtraColumns(r.DefaultColumns),
	)
	if err != nil {
		log.Fatal(err)
	}
	if r.Multiplex == nil {
		r.Table.Multiplex = "client.ProjectMultiplex"
	} else {
		r.Table.Multiplex = *r.Multiplex
	}
	r.Table.Resolver = "fetch" + strcase.ToCamel(r.SubService)
	if r.GetFunction != "" {
		r.Table.PreResourceResolver = "get" + strcase.ToCamel(r.StructName)
	}
	if r.Relations != nil {
		r.Table.Relations = r.Relations
	}
	mainTemplate := r.Template + ".go.tpl"
	if mock {
		mainTemplate = r.Template + "_mock.go.tpl"
	}
	tpl, err := template.New(mainTemplate).Funcs(template.FuncMap{
		"ToCamel": strcase.ToCamel,
		"ToLower": strings.ToLower,
	}).ParseFS(gcpTemplatesFS, "templates/"+mainTemplate)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse gcp templates: %w", err))
	}
	tpl, err = tpl.ParseFS(sdkgen.TemplatesFS, "templates/*.go.tpl")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse sdk template: %w", err))
	}
	var buff bytes.Buffer
	if err := tpl.Execute(&buff, r); err != nil {
		log.Fatal(fmt.Errorf("failed to execute template: %w", err))
	}
	filePath := path.Join(dir, "../resources/services", r.Service)
	if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	if mock {
		filePath = path.Join(filePath, r.SubService+"_mock_test.go")
	} else {
		filePath = path.Join(filePath, r.SubService+".go")
	}

	content := buff.Bytes()
	formattedContent, err := format.Source(buff.Bytes())
	if err != nil {
		log.Printf("failed to format %s\n", filePath)
	} else {
		content = formattedContent
	}
	// if err != nil {
	// 	fmt.Println(buff.String())
	// 	log.Fatal(fmt.Errorf("failed to format code for %s: %w", filePath, err))
	// }
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		log.Fatal(fmt.Errorf("failed to write file %s: %w", filePath, err))
	}
}

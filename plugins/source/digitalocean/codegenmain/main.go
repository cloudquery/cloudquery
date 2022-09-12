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

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/codegen"
	sdkgen "github.com/cloudquery/plugin-sdk/codegen"
	"github.com/iancoleman/strcase"
)

//go:embed templates/*.go.tpl
var gcpTemplatesFS embed.FS

var resources = []*codegen.Resource{}

func main() {
	resources = append(resources, codegen.Resources...)

	for _, r := range resources {
		configureResource(r)
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
		log.Fatal(fmt.Errorf("failed to format code for %s: %w", filePath, err))
	}
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		log.Fatal(fmt.Errorf("failed to write file %s: %w", filePath, err))
	}
}

func configureResource(r *codegen.Resource) {
	if r.MockStruct == nil {
		r.MockStruct = r.Struct
	}
	r.StructName, r.IsStructPointer = getElemNameAndPointer(r.Struct)
	r.ParentStructName, r.IsParentPointer = getElemNameAndPointer(r.ParentStruct)
	r.MockStructName, _ = getElemNameAndPointer(r.MockStruct)

	r.SubServiceName = r.SubService
	if r.SubServiceName == "" {
		r.SubServiceName = r.Service
	}
	if r.Args == "" {
		r.Args = ",p.ID"
	}
}

func generateResource(r codegen.Resource, mock bool) {
	var err error
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	dir := path.Dir(filename)

	tableName := fmt.Sprintf("digitalocean_%s_%s", r.Service, r.SubService)
	if r.SubService == "" {
		tableName = fmt.Sprintf("digitalocean_%s", r.Service)
	}

	r.Table, err = sdkgen.NewTableFromStruct(
		tableName,
		r.Struct,
		sdkgen.WithSkipFields(r.SkipFields),
		sdkgen.WithExtraColumns(r.ExtraColumns),
	)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to create table for %s: %w", r.StructName, err))
	}
	if r.Multiplex == nil {
		r.Table.Multiplex = "client.ProjectMultiplex"
	} else {
		r.Table.Multiplex = *r.Multiplex
	}
	r.Table.Resolver = "fetch" + strcase.ToCamel(r.SubServiceName)

	if r.Relations != nil {
		relations := make([]string, 0, len(r.Relations))
		for _, r := range r.Relations {
			f := fmt.Sprintf("%s()", r)
			relations = append(relations, f)
		}
		r.Table.Relations = []string{strings.Join(relations, ", ")}
	}
	mainTemplate := r.Template + ".go.tpl"
	if mock {
		if r.MockTemplate == "" {
			mainTemplate = r.Template + "_mock.go.tpl"
		} else {
			mainTemplate = r.MockTemplate + ".go.tpl"
		}
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
		filePath = path.Join(filePath, r.SubServiceName+"_mock_test.go")
	} else {
		filePath = path.Join(filePath, r.SubServiceName+".go")
	}

	content := buff.Bytes()
	formattedContent, err := format.Source(buff.Bytes())
	if err != nil {
		log.Printf("failed to format %s\n", filePath)
	} else {
		content = formattedContent
	}
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		log.Fatal(fmt.Errorf("failed to write file %s: %w", filePath, err))
	}
}

func getElemNameAndPointer(i interface{}) (string, bool) {
	if i == nil {
		return "", false
	}
	isPointer := reflect.ValueOf(i).Type().Kind() == reflect.Pointer
	name := reflect.TypeOf(i).Name()
	if isPointer {
		name = reflect.TypeOf(i).Elem().Name()
	}
	return name, isPointer
}

package recipes

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/iancoleman/strcase"
	"go/format"
	"os"
	"path"
	"runtime"
	"strings"
	"text/template"
)

//go:embed templates/*.go.tpl
var templatesFS embed.FS

type Resource struct {
	// Table is the table definition that will be used to generate the cloudquery table
	Table *codegen.TableDefinition
	// Struct that will be used to generate the cloudquery table
	Struct interface{}
	// MockStruct that will be used to generate the mock
	IsStructPointer bool
	MockStruct      interface{}
	// ParentStruct that will be used to generate the mock
	ParentStruct interface{}
	// ParentStructName that will be used to generate the mock
	IsParentPointer bool
	// Args for get list function
	Args             string
	ParentStructName string
	// MockWrapper
	MockWrapper bool
	// MockStructName is the name of the Struct because it can't be inferred by reflection
	MockStructName string
	// StructName is the name of the Struct because it can't be inferred by reflection
	StructName string
	// Service is the name of the gcp service the struct/api is residing
	Service string
	// SubService is the name of the subservice
	SubService string
	// SubService is the name of the subservice
	SubServiceName string
	Relations      []string
	// Template is the template to use to generate the resource (some services has different template as some services were generated using different original codegen)
	Template string
	// imports to add for this resource
	Imports []string
	// Multiplex
	Multiplex string
	// ChildTable
	ChildTable bool
	// Pass to MockTemplate
	MockTemplate string
	// MockFieldName is a name of a struct that is used for mocking
	ResponsePath string
	// SkipFields fields in go struct to skip when generating the table from the go struct
	SkipFields []string
	// Columns override, override generated columns
	ExtraColumns []codegen.ColumnDefinition
}

func (r *Resource) Generate() error {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("failed to get caller information")
	}
	dir := path.Dir(filename)

	var err error
	opts := []codegen.TableOptions{
		codegen.WithSkipFields(r.SkipFields),
		codegen.WithExtraColumns(r.ExtraColumns),
	}
	//if r.UnwrapEmbeddedStructs {
	//	opts = append(opts, codegen.WithUnwrapAllEmbeddedStructs())
	//}

	tableName := fmt.Sprintf("digitalocean_%s_%s", r.Service, r.SubService)
	if r.SubService == "" {
		r.SubService = r.Service
		tableName = fmt.Sprintf("digitalocean_%s", r.Service)
	}

	r.Table, err = codegen.NewTableFromStruct(
		tableName,
		r.Struct,
		opts...,
	)
	if err != nil {
		return err
	}
	r.Table.Resolver = "fetch" + strcase.ToCamel(r.Service) + strcase.ToCamel(r.SubService)
	if r.Multiplex != "" {
		r.Table.Multiplex = r.Multiplex
	}

	if r.Relations != nil {
		r.Table.Relations = r.Relations
	}

	tpl, err := template.New("resource.go.tpl").Funcs(template.FuncMap{
		"ToCamel": strcase.ToCamel,
		"ToLower": strings.ToLower,
	}).ParseFS(templatesFS, "templates/resource.go.tpl")
	if err != nil {
		return fmt.Errorf("failed to parse gcp templates: %w", err)
	}
	tpl, err = tpl.ParseFS(codegen.TemplatesFS, "templates/*.go.tpl")
	if err != nil {
		return fmt.Errorf("failed to parse sdk template: %w", err)
	}

	var buff bytes.Buffer
	if err := tpl.Execute(&buff, r); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}
	dir = path.Join(dir, "../../resources/services", r.Service)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	filePath := path.Join(dir, r.SubService+".go")
	content := buff.Bytes()
	formattedContent, err := format.Source(buff.Bytes())
	if err != nil {
		fmt.Printf("failed to format source: %s: %w\n", filePath, err)
	} else {
		content = formattedContent
	}
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}
	return nil
}

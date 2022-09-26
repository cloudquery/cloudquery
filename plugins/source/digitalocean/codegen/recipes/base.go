package recipes

import (
	"bytes"
	"embed"
	"fmt"
	"go/format"
	"os"
	"path"
	"runtime"
	"strings"
	"text/template"

	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

//go:embed templates/*.go.tpl
var templatesFS embed.FS

type Resource struct {
	// Table is the table definition that will be used to generate the cloudquery table
	Table *codegen.TableDefinition
	// Struct that will be used to generate the cloudquery table
	Struct interface{}
	// Service is the name of the digitalocean service the struct/api is residing
	Service string
	// SubService is the name of the subservice
	SubService string
	// Relations list of resources that are relations of current resource
	Relations []string
	// imports to add for this resource
	Imports []string
	// Multiplex
	Multiplex string
	// SkipFields fields in go struct to skip when generating the table from the go struct
	SkipFields []string
	// Columns override, override generated columns
	ExtraColumns []codegen.ColumnDefinition
	// PostResolver name of post resolver function
	PostResolver string
	// TableName name of the table
	TableName string
}

func (r *Resource) Generate() error {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("failed to get caller information")
	}
	dir := path.Dir(filename)

	var err error
	opts := []codegen.TableOption{
		codegen.WithSkipFields(r.SkipFields),
		codegen.WithExtraColumns(r.ExtraColumns),
	}

	plural := pluralize.NewClient()
	if r.TableName == "" {
		r.TableName = fmt.Sprintf("digitalocean_%s_%s", plural.Singular(r.Service), r.SubService)
		if r.SubService == "" {
			r.SubService = r.Service
			r.TableName = fmt.Sprintf("digitalocean_%s", r.Service)
		}
	}

	r.Table, err = codegen.NewTableFromStruct(
		r.TableName,
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

	if r.PostResolver != "" {
		r.Table.PostResourceResolver = r.PostResolver
	}

	tpl, err := template.New("resource.go.tpl").Funcs(template.FuncMap{
		"ToCamel": strcase.ToCamel,
		"ToLower": strings.ToLower,
	}).ParseFS(templatesFS, "templates/resource.go.tpl")
	if err != nil {
		return fmt.Errorf("failed to parse templates: %w", err)
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
		return fmt.Errorf("failed to format source: %s: %w\n", filePath, err)
	} else {
		content = formattedContent
	}
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}
	return nil
}

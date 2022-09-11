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
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/iancoleman/strcase"
)

type Resource struct {
	Service string
	SubService string
	Struct interface{}
	SkipFields []string
	ExtraColumns []codegen.ColumnDefinition
	Table *codegen.TableDefinition
	Multiplex string
}

//go:embed templates/*.go.tpl
var templatesFS embed.FS

var defaultRegionalColumns = []codegen.ColumnDefinition{
	{
		Name: "account_id",
		Type: schema.TypeString,
		Resolver: "client.ResolveAWSAccount",
	},
	{
		Name: "region",
		Type: schema.TypeString,
		Resolver: "client.ResolveAWSRegion",
	},
}

func (r *Resource) Generate() error {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("failed to get caller information")
	}
	dir := path.Dir(filename)
	
	var err error
	r.Table, err = codegen.NewTableFromStruct(
		fmt.Sprintf("aws_%s_%s", r.Service, r.SubService),
		r.Struct,
		codegen.WithSkipFields(r.SkipFields),
		codegen.WithExtraColumns(r.ExtraColumns),
	)
	if err != nil {
		return err
	}
	r.Table.Resolver = "fetch" + strcase.ToCamel(r.Service) + strcase.ToCamel(r.SubService)
	if r.Multiplex != "" {
		r.Table.Multiplex = r.Multiplex
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

	filePath := path.Join(dir, r.SubService + ".go")
	content, err := format.Source(buff.Bytes())
	if err != nil {
		return fmt.Errorf("failed to format code for %s: %w", filePath, err)
	}
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}
	return nil
}
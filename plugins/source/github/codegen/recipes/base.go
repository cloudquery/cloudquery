package recipes

import (
	"bytes"
	"embed"
	"fmt"
	"go/format"
	"os"
	"path"
	"reflect"
	"runtime"
	"strings"
	"text/template"

	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/go-github/v48/github"
	"github.com/iancoleman/strcase"
)

type Resource struct {
	Service              string
	SubService           string
	Struct               interface{}
	SkipFields           []string
	ExtraColumns         []codegen.ColumnDefinition
	Table                *codegen.TableDefinition
	TableName            string
	Multiplex            string
	PreResourceResolver  string
	PostResourceResolver string
	Relations            []string
}

//go:embed templates/*.go.tpl
var templatesFS embed.FS

func timestampTransformer(field reflect.StructField) (schema.ValueType, error) {
	timestamp := github.Timestamp{}
	switch field.Type {
	case reflect.TypeOf(timestamp), reflect.TypeOf(&timestamp):
		return schema.TypeTimestamp, nil
	default:
		return schema.TypeInvalid, nil
	}
}

func (r *Resource) Generate() error {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("failed to get caller information")
	}
	dir := path.Dir(filename)

	var err error
	if r.TableName == "" {
		r.TableName = r.Service + "_" + r.SubService
	}
	r.TableName = `github_` + r.TableName
	r.Table, err = codegen.NewTableFromStruct(r.TableName, r.Struct,
		codegen.WithSkipFields(r.SkipFields),
		codegen.WithExtraColumns(r.ExtraColumns),
		codegen.WithTypeTransformer(timestampTransformer),
	)
	if err != nil {
		return err
	}
	r.Table.Resolver = "fetch" + strcase.ToCamel(r.SubService)
	if r.Multiplex != "" {
		r.Table.Multiplex = r.Multiplex
	}
	if r.PreResourceResolver != "" {
		r.Table.PreResourceResolver = r.PreResourceResolver
	}
	if r.PostResourceResolver != "" {
		r.Table.PostResourceResolver = r.PostResourceResolver
	}
	if r.Relations != nil {
		r.Table.Relations = r.Relations
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
		fmt.Printf("failed to format source: %s: %s\n", filePath, err.Error())
	} else {
		content = formattedContent
	}
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}
	return nil
}

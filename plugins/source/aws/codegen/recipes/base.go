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

	"github.com/cloudquery/plugin-sdk/caser"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/iancoleman/strcase"
)

type Resource struct {
	// Name overrides the table name: used only in rare cases for backwards-compatibility.
	Name                  string
	Service               string
	SubService            string
	Struct                interface{}
	SkipFields            []string
	Description           string
	ExtraColumns          []codegen.ColumnDefinition
	Table                 *codegen.TableDefinition
	Multiplex             string
	PreResourceResolver   string
	PostResourceResolver  string
	Relations             []string
	UnwrapEmbeddedStructs bool
}

//go:embed templates/*.go.tpl
var templatesFS embed.FS

var defaultAccountColumns = []codegen.ColumnDefinition{
	{
		Name:     "account_id",
		Type:     schema.TypeString,
		Resolver: "client.ResolveAWSAccount",
	},
}

var defaultRegionalColumns = []codegen.ColumnDefinition{
	{
		Name:     "account_id",
		Type:     schema.TypeString,
		Resolver: "client.ResolveAWSAccount",
	},
	{
		Name:     "region",
		Type:     schema.TypeString,
		Resolver: "client.ResolveAWSRegion",
	},
}

var defaultRegionalColumnsPK = []codegen.ColumnDefinition{
	{
		Name:     "account_id",
		Type:     schema.TypeString,
		Resolver: "client.ResolveAWSAccount",
		Options:  schema.ColumnCreationOptions{PrimaryKey: true},
	},
	{
		Name:     "region",
		Type:     schema.TypeString,
		Resolver: "client.ResolveAWSRegion",
		Options:  schema.ColumnCreationOptions{PrimaryKey: true},
	},
}

func awsNameTransformer(f reflect.StructField) (string, error) {
	c := caser.New(caser.WithCustomInitialisms(map[string]bool{
		"EC2": true,
		"VPC": true,
	}))
	return c.ToSnake(f.Name), nil
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
		codegen.WithNameTransformer(awsNameTransformer),
	}
	if r.UnwrapEmbeddedStructs {
		opts = append(opts, codegen.WithUnwrapAllEmbeddedStructs())
	}
	name := fmt.Sprintf("aws_%s_%s", r.Service, r.SubService)
	if r.Name != "" {
		name = r.Name
	}
	r.Table, err = codegen.NewTableFromStruct(
		name,
		r.Struct,
		opts...,
	)
	r.Table.Description = r.Description
	if err != nil {
		return err
	}
	r.Table.Resolver = "fetch" + strcase.ToCamel(r.Service) + strcase.ToCamel(r.SubService)
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
		fmt.Printf("failed to format source: %s: %v\n", filePath, err)
	} else {
		content = formattedContent
	}
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}
	return nil
}

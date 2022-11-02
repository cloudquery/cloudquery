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

	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
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

	// Used for generating the resolver and mock tests.
	// --------------------------------
	ShouldGenerateResolverAndMockTest bool
	ResolverAndMockTestTemplate       string

	// Needed because it's usually capitalised differently than 'Service'.
	// Used for accessing 'client.Services().{{.CloudqueryServiceName}}'.
	CloudqueryServiceName string
}

//go:embed templates/resolver_and_mock_test/*/*.go.tpl
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

func awsResolverTransformer(f reflect.StructField, path string) (string, error) {
	if f.Type.String() == "[]types.Tag" {
		if path == "Tags" {
			return "client.ResolveTags", nil
		}
		return `client.ResolveTagField("` + path + `")`, nil
	}

	if path == "Tags" || path == "TagSet" {
		switch f.Type {
		case reflect.TypeOf(map[string]string{}), reflect.TypeOf(map[string]*string{}), reflect.TypeOf(map[string]interface{}{}), reflect.TypeOf([]types.TagDescription{}):
			// valid tag types
		default:
			return "", fmt.Errorf("%q field is not of type []types.Tag or acceptable map: %s", path, f.Type.String())
		}
	}

	return codegen.DefaultResolverTransformer(f, path)
}

func (r *Resource) Generate() error {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("failed to get caller information")
	}

	dir := path.Dir(filename)
	dir = path.Join(dir, "../../resources/services", r.Service)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	var err error
	opts := []codegen.TableOption{
		codegen.WithSkipFields(r.SkipFields),
		codegen.WithExtraColumns(r.ExtraColumns),
		codegen.WithNameTransformer(awsNameTransformer),
		codegen.WithResolverTransformer(awsResolverTransformer),
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
	if err != nil {
		return fmt.Errorf("error generating %s: %w", name, err)
	}
	r.Table.Description = r.Description
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

	if err := r.generateSchema(dir); err != nil {
		return err
	}

	if r.ShouldGenerateResolverAndMockTest {
		if err := r.generateResolver(dir); err != nil {
			return err
		}
		if err := r.generateMockTest(dir); err != nil {
			return err
		}
	}

	return nil
}

func (r *Resource) generateSchema(dir string) error {
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

func (r *Resource) generateResolver(dir string) error {
	tpl, err := template.New("fetch.go.tpl").Funcs(template.FuncMap{
		"ToCamel": strcase.ToCamel,
		"ToLower": strings.ToLower,
	}).ParseFS(templatesFS,
		fmt.Sprintf("templates/resolver_and_mock_test/%s/fetch.go.tpl", r.ResolverAndMockTestTemplate))
	if err != nil {
		return fmt.Errorf("failed to parse gcp templates: %w", err)
	}

	var buff bytes.Buffer
	if err := tpl.Execute(&buff, r); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	filePath := path.Join(dir, r.SubService+"_fetch.go")
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

func (r *Resource) generateMockTest(dir string) error {
	tpl, err := template.New("mock_test.go.tpl").Funcs(template.FuncMap{
		"ToCamel": strcase.ToCamel,
		"ToLower": strings.ToLower,
	}).ParseFS(templatesFS,
		fmt.Sprintf("templates/resolver_and_mock_test/%s/mock_test.go.tpl", r.ResolverAndMockTestTemplate))
	if err != nil {
		return fmt.Errorf("failed to parse gcp templates: %w", err)
	}

	var buff bytes.Buffer
	if err := tpl.Execute(&buff, r); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	filePath := path.Join(dir, r.SubService+"_mock_test.go")
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

// These methods are called from the template.
// Because of this, we use a value receiver.
// -------------------------------------------------------------------------------

// Because usually the 'Struct' field contains a pointer, we need to dereference with '.Elem()'.
func (resource Resource) StructName() string {
	return reflect.TypeOf(resource.Struct).Elem().Name()
}

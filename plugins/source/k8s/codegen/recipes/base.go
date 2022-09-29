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
	"github.com/iancoleman/strcase"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//go:embed templates/*.go.tpl
var templatesFS embed.FS

type Resource struct {
	Service      string
	SubService   string
	Struct       interface{}
	Multiplex    string // By default, Multiplex is `client.ContextMultiplex`
	Table        *codegen.TableDefinition
	ExtraColumns []codegen.ColumnDefinition
	SkipFields   []string
}

func (resource *Resource) Generate() error {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("failed to get caller information")
	}
	dir := path.Dir(filename)

	var err error

	skipFields := []string{
		"GenerateName",
		"SelfLink",
		"CreationTimestamp",
		"DeletionTimestamp",
		"ZZZ_DeprecatedClusterName",
		"ManagedFields",
		"UID", // Skip UID - but only to re-add it with primary-key options
	}
	skipFields = append(skipFields, resource.SkipFields...)

	extraColumns := []codegen.ColumnDefinition{
		{
			Name:     "context",
			Type:     schema.TypeString,
			Resolver: `client.ResolveContext`,
		},
		{
			Name:     "uid",
			Type:     schema.TypeString,
			Resolver: `schema.PathResolver("UID")`,
			Options:  schema.ColumnCreationOptions{PrimaryKey: true},
		},
	}

	extraColumns = append(extraColumns, resource.ExtraColumns...)

	resource.Table, err = codegen.NewTableFromStruct(
		fmt.Sprintf("k8s_%s_%s", resource.Service, resource.SubService),
		resource.Struct,
		codegen.WithUnwrapAllEmbeddedStructs(),
		codegen.WithSkipFields(skipFields),
		codegen.WithExtraColumns(extraColumns),
		codegen.WithUnwrapStructFields([]string{"Spec", "Status"}),
		codegen.WithTypeTransformer(typeTransformerForK8s),
	)

	if err != nil {
		return err
	}

	resource.Table.Resolver = "fetch" + strcase.ToCamel(resource.Service) + strcase.ToCamel(resource.SubService)
	if resource.Multiplex != "" {
		resource.Table.Multiplex = resource.Multiplex
	} else {
		resource.Table.Multiplex = "client.ContextMultiplex"
	}

	tpl, err := template.New("resource.go.tpl").Funcs(template.FuncMap{
		"ToCamel": strcase.ToCamel,
		"ToLower": strings.ToLower,
	}).ParseFS(templatesFS, "templates/*.go.tpl")
	if err != nil {
		return fmt.Errorf("failed to parse k8s templates: %w", err)
	}
	tpl, err = tpl.ParseFS(codegen.TemplatesFS, "templates/*.go.tpl")
	if err != nil {
		return fmt.Errorf("failed to parse sdk template: %w", err)
	}

	var buff bytes.Buffer
	if err := tpl.Execute(&buff, resource); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}
	dir = path.Join(dir, "../../resources/services", resource.Service)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	filePath := path.Join(dir, resource.SubService+".go")
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

func typeTransformerForK8s(field reflect.StructField) (schema.ValueType, error) {
	if isK8sTimeStruct(field.Type) {
		return schema.TypeTimestamp, nil
	}

	return codegen.DefaultTypeTransformer(field)
}

// isK8sTimeStruct returns true if the given type is a metav1.Time struct or a pointer to it.
func isK8sTimeStruct(fieldType reflect.Type) bool {
	fieldKind := fieldType.Kind()

	if fieldKind == reflect.Ptr {
		return isK8sTimeStruct(fieldType.Elem())
	}

	if fieldKind == reflect.Struct && fieldType == reflect.TypeOf(metav1.Time{}) {
		return true
	}

	return false
}

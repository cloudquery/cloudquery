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
	Service             string
	SubService          string
	ServicePath         string
	GlobalResource      bool
	ServiceFunc         any
	ResourceFunc        any
	ServiceFuncName     string
	ResourceFuncName    string
	ResourcePath        string
	ImportPath          string
	SubServiceInterface any
	ResourceInterface   any
	Struct              any
	StructName          string
	Multiplex           string // By default, Multiplex is `client.ContextMultiplex`
	Table               *codegen.TableDefinition
	ExtraColumns        []codegen.ColumnDefinition
	SkipFields          []string
	FakerOverride       string
	MockImports         []string
}

func getFunctionName(i any) string {
	s := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	return s[strings.LastIndex(s, ".")+1:]
}

func getPackagePath(myvar any) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return t.Elem().PkgPath()
	} else {
		return t.PkgPath()
	}
}

func getType(myvar any) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}

func (resource *Resource) generate(mock bool) error {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("failed to get caller information")
	}
	dir := path.Dir(filename)

	tplName := "resource.go.tpl"
	if mock {
		tplName = "mock.go.tpl"
	}
	tpl, err := template.New(tplName).Funcs(template.FuncMap{
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

	filePath := path.Join(dir, resource.SubService)
	if mock {
		filePath = filePath + "_test.go"
	} else {
		filePath = filePath + ".go"
	}
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

func (resource *Resource) Generate() error {
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

	resource.Table.Resolver = "fetch" + strcase.ToCamel(resource.SubService)
	if resource.Multiplex != "" {
		resource.Table.Multiplex = resource.Multiplex
	} else {
		resource.Table.Multiplex = "client.ContextMultiplex"
	}

	resource.StructName = getType(resource.Struct)
	resource.ImportPath = strings.TrimPrefix(getPackagePath(resource.Struct), "k8s.io/api/")
	resource.ServiceFuncName = getFunctionName(resource.ServiceFunc)
	resource.ResourceFuncName = getFunctionName(resource.ResourceFunc)

	if err := resource.generate(false); err != nil {
		return err
	}

	if err := resource.generate(true); err != nil {
		return err
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

func GeneratePlugin(resources []*Resource) error {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("failed to get caller information")
	}
	dir := path.Dir(filename)

	tpl, err := template.New("plugin.go.tpl").Funcs(template.FuncMap{
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
	if err := tpl.Execute(&buff, resources); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}
	dir = path.Join(dir, "../../resources/plugin")
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	filePath := path.Join(dir, "plugin.go")
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


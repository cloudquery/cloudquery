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
	"text/template"

	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/uuid"
	"github.com/iancoleman/strcase"
)

//go:embed templates/*.go.tpl
var templatesFS embed.FS
var Resources []*Resource

type Resource struct {
	Service    string
	SubService string
	Struct     any
	Table      *codegen.TableDefinition
}

func writeTemplateContentToFile(dir string, filePath string, buff bytes.Buffer) error {
	outputPath := path.Join(dir, "../..", filePath)
	outputDir := path.Dir(outputPath)
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", outputDir, err)
	}

	content := buff.Bytes()
	formattedContent, err := format.Source(content)
	if err != nil {
		fmt.Printf("failed to format source: %s: %v\n", filePath, err)
	} else {
		content = formattedContent
	}

	if err := os.WriteFile(outputPath, content, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}
	return nil
}

func renderTemplate(name string, filePath string, data any) error {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("failed to get caller information")
	}
	dir := path.Dir(filename)

	tpl, err := template.New(name).Funcs(template.FuncMap{"ToCamel": strcase.ToCamel}).ParseFS(templatesFS, "templates/*.go.tpl")
	if err != nil {
		return fmt.Errorf("failed to parse azureDevops templates: %w", err)
	}
	tpl, err = tpl.ParseFS(codegen.TemplatesFS, "templates/*.go.tpl")
	if err != nil {
		return fmt.Errorf("failed to parse sdk template: %w", err)
	}

	var buff bytes.Buffer
	if err := tpl.Execute(&buff, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	return writeTemplateContentToFile(dir, filePath, buff)
}

func (resource *Resource) generate() error {
	return renderTemplate("resource.go.tpl", path.Join("resources", "services", resource.Service, resource.SubService+".go"), resource)
}

func isUUID(fieldType reflect.Type) bool {
	fieldKind := fieldType.Kind()

	if fieldKind == reflect.Ptr {
		return isUUID(fieldType.Elem())
	}

	return fieldType == reflect.TypeOf(uuid.UUID{})
}

func typeTransformer(field reflect.StructField) (schema.ValueType, error) {
	if isUUID(field.Type) {
		return schema.TypeUUID, nil
	}

	return codegen.DefaultTypeTransformer(field)
}

func (resource *Resource) Generate() error {
	var err error
	resource.Table, err = codegen.NewTableFromStruct(
		fmt.Sprintf("azuredevops_%s_%s", resource.Service, resource.SubService),
		resource.Struct,
		codegen.WithTypeTransformer(typeTransformer),
	)

	if err != nil {
		return err
	}

	resource.Table.Resolver = "fetch" + strcase.ToCamel(resource.SubService)
	if err := resource.generate(); err != nil {
		return err
	}

	return nil
}

func GenerateTablesList(resources []*Resource) error {
	return renderTemplate("tables.go.tpl", path.Join("resources", "plugin", "tables.go"), resources)
}

package recipes

import (
	"bytes"
	"embed"
	"errors"
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
	"golang.org/x/exp/slices"
)

type Resource struct {
	SubService            string
	Struct                any
	SkipFields            []string
	Description           string
	ExtraColumns          []codegen.ColumnDefinition
	PKColumns             []string
	Table                 *codegen.TableDefinition
	Multiplex             string
	PreResourceResolver   string
	PostResourceResolver  string
	Relations             []string
	UnwrapEmbeddedStructs bool

	SkipFetchGeneration bool
	SkipMockGeneration  bool

	// The sub-directory in the `templates` directory where the `resource.go.tpl` and `fetch.go.tpl` file are located.
	Template string

	ListFunctionNameOverride      string
	ListOptionsStructNameOverride string
	ResponseStructOverride        string
	ResponseFieldOverride         string
	RestPathOverride              string
	ParentIsPointer               bool

	// used for generating resolver and mock tests, but set automatically
	Parent   *Resource
	Children []*Resource
}

//go:embed templates/*.go.tpl
//go:embed templates/*/*.go.tpl
var templatesFS embed.FS

func (r *Resource) Generate() error {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("failed to get caller information")
	}

	dir := path.Dir(filename)
	if r.Parent == nil {
		dir = path.Join(dir, "../../resources/services/", r.SubService)
	} else {
		dir = path.Join(dir, "../../resources/services", r.Parent.SubService)
	}

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	opts := []codegen.TableOption{
		codegen.WithSkipFields(r.SkipFields),
		codegen.WithExtraColumns(r.ExtraColumns),
		codegen.WithPKColumns(r.PKColumns...),
		codegen.WithUnwrapAllEmbeddedStructs(), // Unwrap the `APIObject` embedded-struct, that contains the `ID` PK.
		codegen.WithTypeTransformer(pagerDutyTypeTransformer),
	}

	name := fmt.Sprintf("pagerduty_%s", r.SubService)

	var err error
	r.Table, err = codegen.NewTableFromStruct(
		name,
		r.Struct,
		opts...,
	)
	if err != nil {
		return fmt.Errorf("error generating %s: %w", name, err)
	}

	r.Table.Description = r.Description
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
	if r.Template == "" {
		r.Template = "basic"
	}

	if err := r.generateSchema(dir); err != nil {
		return err
	}

	if !r.SkipFetchGeneration {
		if err := r.generateFetch(dir); err != nil {
			return err
		}
	}

	if r.Parent == nil && !r.SkipMockGeneration {
		if err := r.generateMockTest(dir); err != nil {
			return err
		}
	}

	return nil
}

func (r *Resource) generateSchema(dir string) error {
	var tpl *template.Template
	var err error

	tpl, err = template.New(path.Join("resource.go.tpl")).Funcs(template.FuncMap{
		"ToCamel": strcase.ToCamel,
		"ToLower": strings.ToLower,
	}).ParseFS(templatesFS, path.Join("templates", r.Template, "resource.go.tpl"))
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

	filePath := path.Join(dir, r.SubService+".go")
	filePath = MaybeReplaceUnderscoreWindowsSuffix(filePath)

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

func (r *Resource) generateFetch(dir string) error {
	var tpl *template.Template
	var err error

	tpl, err = template.New("fetch.go.tpl").Funcs(template.FuncMap{
		"ToCamel": strcase.ToCamel,
		"ToLower": strings.ToLower,
	}).ParseFS(templatesFS, path.Join("templates", r.Template, "fetch.go.tpl"))
	if err != nil {
		return fmt.Errorf("failed to parse templates: %w", err)
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
	var tpl *template.Template
	var err error

	tpl, err = template.New("mock_test.go.tpl").Funcs(template.FuncMap{
		"ToCamel": strcase.ToCamel,
		"ToLower": strings.ToLower,
		"ToSnake": strcase.ToSnake,
	}).ParseFS(templatesFS, "templates/mock_test.go.tpl")
	if err != nil {
		return fmt.Errorf("failed to parse templates: %w", err)
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

// SetParentChildRelationships calculates and sets the parent and children fields on resources.
func SetParentChildRelationships(resources []*Resource) error {
	manualRelations := []string{"incident_log_entries"}

	m := map[string]*Resource{}
	for _, r := range resources {
		m[r.SubService] = r
	}
	csr := caser.New()
	for _, r := range resources {
		for _, ch := range r.Relations {
			name := csr.ToSnake(strings.TrimSuffix(ch, "()"))
			if slices.Contains(manualRelations, name) {
				continue
			}
			v, ok := m[name]
			if !ok {
				return errors.New("child not found for " + r.SubService + " : " + name)
			}
			r.Children = append(r.Children, v)
			v.Parent = r
		}
	}
	return nil
}

func GenerateAllTablesList(resources []*Resource) error {
	tpl, err := template.New("all_tables.go.tpl").Funcs(template.FuncMap{
		"ToCamel": strcase.ToCamel,
		"ToLower": strings.ToLower,
	}).ParseFS(templatesFS, "templates/all_tables.go.tpl")
	if err != nil {
		return fmt.Errorf("failed to parse templates: %w", err)
	}

	var buff bytes.Buffer
	if err := tpl.Execute(&buff, filterRootResources(resources)); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	_, runtime_filename, _, ok := runtime.Caller(0) // Make codegen work independently of `pwd`.
	if !ok {
		return fmt.Errorf("failed to get caller information")
	}

	filePath := path.Join(path.Dir(runtime_filename), "../../", "resources/plugin", "all_tables.go")
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

// filterRootResources returns only root resources from the input list
func filterRootResources(resources []*Resource) []*Resource {
	result := make([]*Resource, 0)

	for _, r := range resources {
		if r.Parent == nil {
			result = append(result, r)
		}
	}

	return result
}

// Golang hates files with `_windows.go` extension...
// This function replaces such extension with `_window.go`.
func MaybeReplaceUnderscoreWindowsSuffix(str string) string {
	return strings.ReplaceAll(str, "_windows.go", "_window.go")
}

// The default type-transfomer, but replace `created_at`, `deleted_at` fields with `TypeTimestamp` instead
// of `TypeString`.
func pagerDutyTypeTransformer(field reflect.StructField) (schema.ValueType, error) {
	if !isStringField(field.Type) {
		return codegen.DefaultTypeTransformer(field)
	}

	timestampFieldNames := []string{
		"CreateAt", "CreatedAt", "DeletedAt", "LastStatusChangeAt", "StartTime", "EndTime", "LastIncidentTimestamp",
	}

	if slices.Contains(timestampFieldNames, field.Name) {
		return schema.TypeTimestamp, nil
	}
	return schema.TypeString, nil
}

func isStringField(fieldType reflect.Type) bool {
	fieldKind := fieldType.Kind()

	if fieldKind == reflect.Ptr {
		return isStringField(fieldType.Elem())
	}

	return fieldKind == reflect.String
}

// These methods are called from the template.
// Because of this, we use a value receiver.
// -------------------------------------------------------------------------------

// StructName returns the name of the resource's Struct field
func (r Resource) StructName() string {
	if reflect.TypeOf(r.Struct).Kind() == reflect.Ptr {
		return reflect.TypeOf(r.Struct).Elem().Name()
	}
	return reflect.TypeOf(r.Struct).Name()
}

// Returns the last word in `r.SubService`.
// i.e. for `IncidentAlerts`, access is done with `Alerts`.
func (r Resource) DefaultResponseStructFieldName() string {
	splitByUnderscore := strings.Split(r.SubService, "_")
	return strcase.ToCamel(splitByUnderscore[len(splitByUnderscore)-1])
}

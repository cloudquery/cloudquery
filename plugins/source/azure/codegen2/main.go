package main

import (
	"bytes"
	"embed"
	"fmt"
	"go/format"
	"log"
	"os"
	"path"
	"reflect"
	"runtime"
	"strings"
	"text/template"

	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen2/recipes"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

//go:embed templates/*.go.tpl
var templateFS embed.FS


var (
	currentFilename string
	currentDir string
)

var SubscriptionIdColumn = codegen.ColumnDefinition{
	Name:     "subscription_id",
	Type:     schema.TypeString,
	Resolver: "client.ResolveAzureSubscription",
}

func main() {
	var ok bool
	_, currentFilename, _, ok = runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	currentDir = path.Dir(currentFilename)

	for _, r := range recipes.Tables {
		initResource(&r)
		generateResource(r, false)
	}
	// generateResource()
}


func initResource(r *recipes.Table) {
	var err error
	if r.Client != nil {
		r.ClientName = reflect.TypeOf(r.Client).Elem().Name()
	}
	if r.NewFunc != nil {
		path := strings.Split(runtime.FuncForPC(reflect.ValueOf(r.NewFunc).Pointer()).Name(), ".")
		r.NewFuncName = path[len(path)-1]
	}
	if r.ListFunc != nil {
		path := strings.Split(runtime.FuncForPC(reflect.ValueOf(r.ListFunc).Pointer()).Name(), ".")
		r.ListFuncName = path[len(path)-1]
		r.ListFuncName = strings.TrimSuffix(r.ListFuncName, "-fm")
	}
	if r.ResponseStruct != nil {
		r.ResponseStructName = reflect.TypeOf(r.ResponseStruct).Elem().Name()
	}
	if r.StructName == "" {
		r.StructName = reflect.TypeOf(r.Struct).Elem().Name()
	}
	if r.MockListStruct == "" {
		r.MockListStruct = strcase.ToCamel(r.StructName)
	}

	if r.MockImports == nil {
		r.MockImports = []string{reflect.TypeOf(r.Struct).Elem().PkgPath()}
	}

	for _, f := range r.ExtraColumns {
		r.SkipFields = append(r.SkipFields, strcase.ToCamel(f.Name))
	}

	extraColumns := r.ExtraColumns
	// extraColumns = append([]codegen.ColumnDefinition{SubscriptionIdColumn}, extraColumns...)
	// if needsProjectIDColumn(r) {
	// 	extraColumns = append([]codegen.ColumnDefinition{recipes.ProjectIdColumn}, extraColumns...)
	// }

	opts := []codegen.TableOption{
		codegen.WithSkipFields(r.SkipFields),
		codegen.WithExtraColumns(extraColumns),
		codegen.WithTypeTransformer(func(field reflect.StructField) (schema.ValueType, error) {
			switch reflect.New(field.Type).Elem().Interface().(type) {
			case *timestamppb.Timestamp,
				timestamppb.Timestamp:
				return schema.TypeTimestamp, nil
			case *durationpb.Duration,
				durationpb.Duration:
				return schema.TypeInt, nil
			case protoreflect.Enum:
				return schema.TypeString, nil
			default:
				return schema.TypeInvalid, nil
			}
		}),
		codegen.WithResolverTransformer(func(field reflect.StructField, path string) (string, error) {
			switch reflect.New(field.Type).Elem().Interface().(type) {
			case *timestamppb.Timestamp,
				timestamppb.Timestamp:
				return `client.ResolveProtoTimestamp("` + path + `")`, nil
			case *durationpb.Duration,
				durationpb.Duration:
				return `client.ResolveProtoDuration("` + path + `")`, nil
			case protoreflect.Enum:
				return `client.ResolveProtoEnum("` + path + `")`, nil
			default:
				return "", nil
			}
		}),
	}

	if r.NameTransformer != nil {
		opts = append(opts, codegen.WithNameTransformer(r.NameTransformer))
	}

	r.Table, err = codegen.NewTableFromStruct(
		fmt.Sprintf("azure_%s_%s", r.Service, r.Name),
		r.Struct,
		opts...,
	)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to create table for %s: %w", r.StructName, err))
	}

	if r.Multiplex == nil {
		if !r.ChildTable {
			r.Table.Multiplex = "client.SubscriptionMultiplex"
		}
	} else {
		r.Table.Multiplex = *r.Multiplex
	}
	r.Table.Resolver = "fetch" + strcase.ToCamel(r.Name)
	if r.PreResourceResolver != "" {
		r.Table.PreResourceResolver = r.PreResourceResolver
	}
	if r.Relations != nil {
		r.Table.Relations = r.Relations
	}

}

func generateResource(r recipes.Table, mock bool) {
	mainTemplate := r.Template + ".go.tpl"
	if mock {
		if r.MockTemplate == "" {
			mainTemplate = r.Template + "_mock.go.tpl"
		} else {
			mainTemplate = r.MockTemplate + ".go.tpl"
		}
	}
	tpl, err := template.New(mainTemplate).Funcs(template.FuncMap{
		"ToCamel": strcase.ToCamel,
		"ToLower": strings.ToLower,
	}).ParseFS(templateFS, "templates/"+mainTemplate)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse templates: %w", err))
	}
	tpl, err = tpl.ParseFS(codegen.TemplatesFS, "templates/*.go.tpl")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse sdk template: %w", err))
	}
	var buff bytes.Buffer
	if err := tpl.Execute(&buff, r); err != nil {
		log.Fatal(fmt.Errorf("failed to execute template: %w", err))
	}
	filePath := path.Join(currentDir, "../resources/services", r.Service)
	if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	if mock {
		filePath = path.Join(filePath, r.Name+"_mock_test.go")
	} else {
		filePath = path.Join(filePath, r.Name+".go")
	}

	content := buff.Bytes()
	formattedContent, err := format.Source(buff.Bytes())
	if err != nil {
		log.Printf("failed to format %s\n", filePath)
	} else {
		content = formattedContent
	}
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		log.Fatal(fmt.Errorf("failed to write file %s: %w", filePath, err))
	}
}
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

	"github.com/cloudquery/cloudquery/plugins/source/azure/codegenv2/recipes"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

//go:embed templates/*.go.tpl
var templateFS embed.FS

var filename string
var rootDir string


func initResource(r *recipes.Resource) {
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
	extraColumns = append([]codegen.ColumnDefinition{recipes.SubscriptionIdColumn}, extraColumns...)
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
		fmt.Sprintf("azure_%s_%s", r.Service, r.SubService),
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
	r.Table.Resolver = "fetch" + strcase.ToCamel(r.SubService)
	if r.PreResourceResolver != "" {
		r.Table.PreResourceResolver = r.PreResourceResolver
	}
	if r.Relations != nil {
		r.Table.Relations = r.Relations
	}

}



func main() {
	var ok bool
	_, filename, _, ok = runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	rootDir = path.Dir(filename)
	var resources []*recipes.Resource
	// resources = append(resources, recipes.AadResources()...)
	// resources = append(resources, recipes.AddonsResources()...)
	resources = append(resources, recipes.AdvisorResources()...)
	resources = append(resources, recipes.BatchResources()...)
	resources = append(resources, recipes.CDNResources()...)
	resources = append(resources, recipes.ComputeResources()...)
	resources = append(resources, recipes.ContainerResources()...)
	resources = append(resources, recipes.ContainerServiceResources()...)
	resources = append(resources, recipes.CosmosDBResources()...)
	resources = append(resources, recipes.DataLakeStoreResources()...)
	resources = append(resources, recipes.DataLakeAnalyticsResources()...)
	resources = append(resources, recipes.EventHubResources()...)
	resources = append(resources, recipes.FrontDoorResources()...)
	resources = append(resources, recipes.KeyVaultResources()...)
	resources = append(resources, recipes.LogicResources()...)
	resources = append(resources, recipes.MariaDBResources()...)
	resources = append(resources, recipes.MySQLResources()...)
	resources = append(resources, recipes.NetworkResources()...)
	resources = append(resources, recipes.PostgreSQLResources()...)
	resources = append(resources, recipes.RedisResources()...)
	// resources = append(resources, recipes.IotSecurityResources()...)
	// resources = append(resources, recipes.WorkloadsResources()...)

	for _, r := range resources {
		initResource(r)
		generateResource(r, false)
		if !r.SkipMock {
			generateResource(r, true)
		}
	}
	generatePlugin(resources)
	generateServices(resources)
}

func generateServices(rr []*recipes.Resource) {
	tpl, err := template.New("services.go.tpl").Funcs(template.FuncMap{
		"ToCamel": strcase.ToCamel,
		"ToLower": strings.ToLower,
	}).ParseFS(templateFS, "templates/services.go.tpl")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse services.go.tpl: %w", err))
	}

	var buff bytes.Buffer
	if err := tpl.Execute(&buff, rr); err != nil {
		log.Fatal(fmt.Errorf("failed to execute services template: %w", err))
	}

	filePath := path.Join(rootDir, "../client/services.go")
	content := buff.Bytes()
	formattedContent, err := format.Source(buff.Bytes())
	if err != nil {
		fmt.Printf("failed to format code for %s: %v\n", filePath, err)
	} else {
		content = formattedContent
	}
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		log.Fatal(fmt.Errorf("failed to write file %s: %w", filePath, err))
	}
}

func generatePlugin(rr []*recipes.Resource) {
	dir := path.Dir(filename)
	tpl, err := template.New("autogen_tables.go.tpl").Funcs(template.FuncMap{
		"ToCamel": strcase.ToCamel,
		"ToLower": strings.ToLower,
	}).ParseFS(templateFS, "templates/autogen_tables.go.tpl")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse autogen_tables.go.tpl: %w", err))
	}

	var buff bytes.Buffer
	if err := tpl.Execute(&buff, rr); err != nil {
		log.Fatal(fmt.Errorf("failed to execute template: %w", err))
	}

	filePath := path.Join(dir, "../resources/plugin/autogen_tables.go")
	content, err := format.Source(buff.Bytes())
	if err != nil {
		log.Fatal(fmt.Errorf("failed to format code for %s: %w", filePath, err))
	}
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		log.Fatal(fmt.Errorf("failed to write file %s: %w", filePath, err))
	}
}

func generateResource(r *recipes.Resource, mock bool) {
	var err error
	dir := path.Dir(filename)

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
	filePath := path.Join(dir, "../resources/services", r.Service)
	if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	if mock {
		filePath = path.Join(filePath, r.SubService+"_mock_test.go")
	} else {
		filePath = path.Join(filePath, r.SubService+".go")
	}

	content := buff.Bytes()
	formattedContent, err := format.Source(buff.Bytes())
	if err != nil {
		log.Printf("failed to format %s\n", filePath)
	} else {
		content = formattedContent
	}
	// if err != nil {
	// 	fmt.Println(buff.String())
	// 	log.Fatal(fmt.Errorf("failed to format code for %s: %w", filePath, err))
	// }
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		log.Fatal(fmt.Errorf("failed to write file %s: %w", filePath, err))
	}
}
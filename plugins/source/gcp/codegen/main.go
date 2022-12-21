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

	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugins/source/gcp/codegen/recipes"
	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

//go:embed templates/*.go.tpl
var gcpTemplatesFS embed.FS

func main() {

	for _, r := range recipes.Resources {
		generateResource(*r, false)
		if !r.SkipMock {
			generateResource(*r, true)
		}
	}
	generatePlugin(recipes.Resources)
}

func generatePlugin(rr []*recipes.Resource) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	dir := path.Dir(filename)
	tpl, err := template.New("autogen_tables.go.tpl").Funcs(template.FuncMap{
		"ToCamel": strcase.ToCamel,
		"ToLower": strings.ToLower,
	}).ParseFS(gcpTemplatesFS, "templates/autogen_tables.go.tpl")
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

func needsProjectIDColumn(r recipes.Resource) bool {
	return r.Multiplex != &recipes.OrgMultiplex
}

func InferListFunction(r *recipes.Resource) {
	// We can infer the List function by matching a list function that return the same struct as r.Struct
	server := reflect.TypeOf(r.RegisterServer).In(1)
	for i := 0; i < server.NumMethod(); i++ {
		method := server.Method(i)
		if strings.HasPrefix(method.Name, "List") {
			response := method.Type.Out(0).Elem()
			for j := 0; j < response.NumField(); j++ {
				field := response.Field(j)
				kind := field.Type.Kind()
				if kind != reflect.Slice {
					continue
				}
				if field.Type.Elem() == reflect.TypeOf(r.Struct) {
					r.ListFunctionName = method.Name
					r.RequestStructName = method.Type.In(1).Elem().Name()
					r.ResponseStructName = method.Type.Out(0).Elem().Name()
					return
				}
			}
		}
	}
}

func getColumn(columns codegen.ColumnDefinitions, name string) *codegen.ColumnDefinition {
	for i := range columns {
		if columns[i].Name == name {
			return &columns[i]
		}
	}
	return nil
}

func listMethodSignature(method reflect.Method) string {
	t := method.Type
	buf := strings.Builder{}
	buf.WriteString(method.Name + "(")
	buf.WriteString(t.In(0).String())
	buf.WriteString(", ")
	buf.WriteString(t.In(1).String())
	buf.WriteString(")")
	buf.WriteString(" (")
	buf.WriteString(t.Out(0).String())
	buf.WriteString(", ")
	buf.WriteString(t.Out(1).String())
	buf.WriteString(")")
	return buf.String()
}

func isValidListMethod(method reflect.Method) bool {
	if !strings.HasPrefix(method.Name, "List") {
		return false
	}

	_, responseHasNextPage := method.Type.Out(0).Elem().FieldByName("NextPageToken")
	return responseHasNextPage
}

func generateMockTestData(r *recipes.Resource) {
	registerServerPath := runtime.FuncForPC(reflect.ValueOf(r.RelationsTestData.RegisterServer).Pointer()).Name()
	serverName := strings.Split(registerServerPath, "/")[4]
	pbName := strings.Split(serverName, ".")[0]
	pbPath := strings.Split(registerServerPath, pbName)[0]
	r.RelationsTestData.ProtobufImport = fmt.Sprintf("%s%s", pbPath, pbName)
	r.RelationsTestData.RegisterServerName = serverName
	r.RelationsTestData.UnimplementedServerName = strings.Replace(serverName, "Register", "Unimplemented", 1)
	server := reflect.TypeOf(r.RelationsTestData.RegisterServer).In(1)
	for i := 0; i < server.NumMethod(); i++ {
		method := server.Method(i)
		if isValidListMethod(method) {
			sig := listMethodSignature(method)
			responseType := method.Type.Out(0).Elem().String()
			r.RelationsTestData.ListFunctions = append(r.RelationsTestData.ListFunctions, recipes.ListFunctions{Signature: sig, ResponseStructName: responseType})
		}
	}
}

func generateResource(r recipes.Resource, mock bool) {
	var err error
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	dir := path.Dir(filename)

	if r.NewFunction != nil {
		path := strings.Split(runtime.FuncForPC(reflect.ValueOf(r.NewFunction).Pointer()).Name(), ".")
		r.NewFunctionName = path[len(path)-1]
	}

	if r.RegisterServer != nil {
		path := strings.Split(runtime.FuncForPC(reflect.ValueOf(r.RegisterServer).Pointer()).Name(), ".")
		r.RegisterServerName = path[len(path)-1]
		r.UnimplementedServerName = strings.Replace(r.RegisterServerName, "Register", "Unimplemented", 1)

	}

	if r.ClientName == "" && r.NewFunctionName != "" {
		n := strings.Split(fmt.Sprintf("%v", reflect.TypeOf(r.NewFunction).Out(0)), ".")[1]
		r.ClientName = n
	}

	if r.StructName == "" {
		r.StructName = reflect.TypeOf(r.Struct).Elem().Name()
	}

	if !r.SkipFetch {
		if r.ListFunction == nil {
			InferListFunction(&r)
		} else {
			path := strings.Split(runtime.FuncForPC(reflect.ValueOf(r.ListFunction).Pointer()).Name(), ".")
			r.ListFunctionName = path[len(path)-1]
			// https://stackoverflow.com/questions/32925344/why-is-there-a-fm-suffix-when-getting-a-functions-name-in-go
			r.ListFunctionName = strings.Split(r.ListFunctionName, "-")[0]
			r.RequestStructName = reflect.TypeOf(r.ListFunction).In(1).Elem().Name()
			switch {
			case r.RegisterServer != nil:
				server := reflect.TypeOf(r.RegisterServer).In(1)
				method, _ := server.MethodByName(r.ListFunctionName)
				r.ResponseStructName = method.Type.Out(0).Elem().Name()
			case r.ListFunctionName == "Get":
				r.ResponseStructName = r.StructName
			default:
				r.ResponseStructName = r.StructName + r.ListFunctionName
			}
		}
	}

	if mock && r.RelationsTestData.RegisterServer != nil {
		generateMockTestData(&r)
	}

	if r.ResponseStruct != nil {
		r.ResponseStructName = reflect.TypeOf(r.ResponseStruct).Elem().Name()
	}

	if r.MockListStruct == "" {
		r.MockListStruct = strcase.ToCamel(r.StructName)
	}

	if r.MockImports == nil {
		r.MockImports = []string{reflect.TypeOf(r.Struct).Elem().PkgPath()}
	}

	extraColumns := r.ExtraColumns
	if needsProjectIDColumn(r) {
		extraColumns = append([]codegen.ColumnDefinition{recipes.ProjectIdColumn}, r.ExtraColumns...)
	}

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
		fmt.Sprintf("gcp_%s_%s", r.Service, r.SubService),
		r.Struct,
		opts...,
	)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to create table for %s: %w", r.StructName, err))
	}
	if r.Multiplex == nil {
		if r.ServiceDNS == "" {
			r.ServiceDNS = r.Service + ".googleapis.com"
		}
		if _, ok := client.GcpServices[r.ServiceDNS]; !ok {
			panic("unknown service DNS: " + r.ServiceDNS)
		}
		r.Table.Multiplex = "client.ProjectMultiplexEnabledServices(\"" + r.ServiceDNS + "\")"
	} else {
		r.Table.Multiplex = *r.Multiplex
	}

	for _, f := range r.PrimaryKeys {
		column := getColumn(r.Table.Columns, f)
		if column != nil {
			column.Options.PrimaryKey = true
		}
	}

	for _, f := range r.IgnoreInTestsColumns {
		column := getColumn(r.Table.Columns, f)
		if column != nil {
			column.IgnoreInTests = true
		}
	}

	r.Table.Description = r.Description

	r.Table.Resolver = "fetch" + strcase.ToCamel(r.SubService)
	if r.PreResourceResolver != "" {
		r.Table.PreResourceResolver = r.PreResourceResolver
	}
	if r.Relations != nil {
		r.Table.Relations = r.Relations
	}
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
	}).ParseFS(gcpTemplatesFS, "templates/"+mainTemplate)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse gcp templates: %w", err))
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

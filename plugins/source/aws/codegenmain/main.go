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

	"github.com/cloudquery/cloudquery/plugins/source/aws/codegenmain/recipes"
	sdkgen "github.com/cloudquery/plugin-sdk/codegen"
	pluginschema "github.com/cloudquery/plugin-sdk/schema"
	"github.com/iancoleman/strcase"
)

//go:embed templates/*.go.tpl
var awsTemplatesFS embed.FS

var resources []*recipes.Resource

func main() {
	resources = append(resources, recipes.ACMResources...)
	resources = append(resources, recipes.APIGatewayv2Resources...)

	for _, r := range resources {
		generateResource(r, false)
		generateResource(r, true)
	}
}

func generateResource(r *recipes.Resource, mock bool) {
	var err error
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	dir := path.Dir(filename)

	tableNameFromSubService := strcase.ToSnake(r.AWSSubService)
	if r.Parent != nil && !strings.HasPrefix(tableNameFromSubService, strcase.ToSnake(r.Parent.ItemName)+"_") {
		tableNameFromSubService = strcase.ToSnake(r.Parent.ItemName) + "_" + tableNameFromSubService
	}
	if r.Parent != nil && r.Parent.Parent != nil && !strings.HasPrefix(tableNameFromSubService, strcase.ToSnake(r.Parent.Parent.ItemName)+"_") {
		tableNameFromSubService = strcase.ToSnake(r.Parent.Parent.ItemName) + "_" + tableNameFromSubService
	}

	r.Table, err = sdkgen.NewTableFromStruct(
		fmt.Sprintf("aws_%s_%s", strings.ToLower(r.AWSService), tableNameFromSubService),
		r.AWSStruct,
		sdkgen.WithSkipFields(append(r.SkipFields, "noSmithyDocumentSerde")),
	)
	if err != nil {
		log.Fatal(err)
	}
	r.Table.Columns = append(r.DefaultColumns, r.Table.Columns...)
	if r.ColumnOverrides != nil {
		for i, c := range r.Table.Columns {
			override, ok := r.ColumnOverrides[c.Name]
			if !ok {
				continue
			}
			if override.Name != "" {
				r.Table.Columns[i].Name = override.Name
			}
			if override.Resolver != "" {
				r.Table.Columns[i].Resolver = override.Resolver
			}
			if override.Description != "" {
				r.Table.Columns[i].Description = override.Description
			}
			delete(r.ColumnOverrides, c.Name)
		}
		// remaining, unmatched columns are added to the end of the table. Difference from DefaultColumns? none for now
		for k, c := range r.ColumnOverrides {
			if c.Type == pluginschema.TypeInvalid {
				if !mock {
					fmt.Println("Not adding unmatched column with unspecified type", k, c)
				}
				continue
			}
			if c.Name == "" {
				c.Name = k
			}
			r.Table.Columns = append(r.Table.Columns, c)
		}
	}

	for i := range r.Table.Columns {
		if len(r.Table.Options.PrimaryKeys) == 0 && r.Table.Columns[i].Name == "arn" {
			//	r.Table.Columns[i].Options.PrimaryKey = true
			r.Table.Options.PrimaryKeys = []string{"arn"}
		}
		if r.Table.Columns[i].Name == "tags" {
			r.HasTags = true
		}
	}
	r.Table.Multiplex = `client.ServiceAccountRegionMultiplexer("` + strings.ToLower(r.AWSService) + `")`

	if r.Parent == nil {
		r.Table.Resolver = "fetch" + r.AWSService + r.AWSSubService
	} else {
		if !strings.HasPrefix(r.AWSSubService, r.Parent.ItemName) {
			r.Table.Resolver = "fetch" + r.AWSService + r.Parent.ItemName + r.AWSSubService
		} else {
			r.Table.Resolver = "fetch" + r.AWSService + r.AWSSubService
		}
	}

	r.TableFuncName = strings.TrimPrefix(r.Table.Resolver, "fetch")
	if mock {
		r.MockFuncName = "build" + r.TableFuncName
		r.TestFuncName = "Test" + r.TableFuncName
	}

	t := reflect.TypeOf(r.AWSStruct).Elem()
	if r.AWSStructName == "" {
		r.AWSStructName = t.Name()
	}

	if r.ItemName == "" {
		r.ItemName = r.AWSStructName
	}

	if r.ListFunctionName == "" {
		r.ListFunctionName = "List" + r.AWSSubService
	}
	if r.DescribeFunctionName == "" {
		if r.ItemName != "" {
			r.DescribeFunctionName = "Describe" + r.ItemName
		} else {
			r.DescribeFunctionName = "Describe" + r.AWSSubService
		}
	}

	if sp := t.PkgPath(); strings.HasSuffix(sp, "/types") {
		if (r.HasTags || r.Parent != nil) && (!r.SkipTypesImport || mock) {
			r.Imports = append(r.Imports, sp)
		}
		r.Imports = append(r.Imports, strings.TrimSuffix(sp, "/types")) // auto-import main pkg (not "types")
	}

	mainTemplate := r.Template + stringSwitch(mock, "_mock_test", "") + ".go.tpl"
	tpl, err := template.New(mainTemplate).Funcs(template.FuncMap{
		"ToCamel": strcase.ToCamel,
		"ToLower": strings.ToLower,
	}).ParseFS(awsTemplatesFS, "templates/*.go.tpl")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse aws templates: %w", err))
	}
	tpl, err = tpl.ParseFS(sdkgen.TemplatesFS, "templates/*.go.tpl")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse codegen template: %w", err))
	}
	var buff bytes.Buffer
	if err := tpl.Execute(&buff, r); err != nil {
		log.Fatal(fmt.Errorf("failed to execute template: %w", err))
	}
	filePath := path.Join(dir, "../codegen")
	if err := os.MkdirAll(filePath, 0755); err != nil {
		log.Fatal(fmt.Errorf("failed to create directory: %w", err))
	}

	fileSuffix := stringSwitch(mock, "_mock_test.go", ".go")
	if r.Parent == nil {
		filePath = path.Join(filePath, strings.ToLower(r.AWSService)+"_"+strcase.ToSnake(r.AWSSubService)+fileSuffix)
	} else {
		parentPrefix := strcase.ToSnake(r.Parent.ItemName) + "_"
		ss := strcase.ToSnake(r.AWSSubService)
		if !strings.HasPrefix(ss, parentPrefix) {
			// prevent "apigatewayv2_integration_integration_responses" type names (repeated prefix)
			ss = parentPrefix + ss
		}
		filePath = path.Join(filePath, strings.ToLower(r.AWSService)+"_"+ss+fileSuffix)
	}
	content, err := format.Source(buff.Bytes())
	if err != nil {
		fmt.Println(buff.String())
		log.Fatal(fmt.Errorf("failed to format code for %s: %w", filePath, err))
	}
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		log.Fatal(fmt.Errorf("failed to write file %s: %w", filePath, err))
	}
}

func stringSwitch(b bool, ifTrue, ifFalse string) string {
	if b {
		return ifTrue
	}
	return ifFalse
}

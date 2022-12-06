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
)

//go:embed templates/*.go.tpl
var templateFS embed.FS

var templateFuncs = template.FuncMap{
	"ToCamel": strcase.ToCamel,
	"ToLower": strings.ToLower,
}

var (
	currentFilename string
	currentDir      string
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

	for i := range recipes.Tables {
		if err := initResource(&recipes.Tables[i]); err != nil {
			log.Fatal(err)
		}
		if err := generateResource(recipes.Tables[i], false); err != nil {
			log.Fatal(err)
		}
		if err := generateResource(recipes.Tables[i], true); err != nil {
			log.Fatal(err)
		}
	}
	if err := generateTables(recipes.Tables); err != nil {
		log.Fatal(err)
	}
}

func generateTables(rr []recipes.Table) error {
	tpl, err := template.New("tables.go.tpl").Funcs(templateFuncs).ParseFS(templateFS, "templates/tables.go.tpl")
	if err != nil {
		return fmt.Errorf("failed to parse services.go.tpl: %w", err)
	}

	var buff bytes.Buffer
	if err := tpl.Execute(&buff, rr); err != nil {
		return fmt.Errorf("failed to execute services template: %w", err)
	}

	filePath := path.Join(currentDir, "../resources/plugin/tables.go")
	content := buff.Bytes()
	formattedContent, err := format.Source(buff.Bytes())
	if err != nil {
		fmt.Printf("failed to format code for %s: %v\n", filePath, err)
	} else {
		content = formattedContent
	}
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}
	return nil
}

func initResource(r *recipes.Table) error {
	var err error
	if r.Client != nil {
		r.ClientName = reflect.TypeOf(r.Client).Elem().Name()
	}
	r.PackageName = strings.TrimPrefix(r.Service, "arm")
	if r.NewFunc != nil {
		path := strings.Split(runtime.FuncForPC(reflect.ValueOf(r.NewFunc).Pointer()).Name(), ".")
		r.NewFuncName = path[len(path)-1]
		if reflect.TypeOf(r.NewFunc).In(0).Name() == "string" {
			r.NewFuncHasSubscriptionId = true
		}
	}
	if r.ListFunc != nil {
		path := strings.Split(runtime.FuncForPC(reflect.ValueOf(r.ListFunc).Pointer()).Name(), ".")
		r.ListFuncName = path[len(path)-1]
		r.ListFuncName = strings.TrimSuffix(r.ListFuncName, "-fm")
		if reflect.TypeOf(r.ListFunc).In(0).Name() == "string" {
			r.ListFuncHasSubscriptionId = true
		}
	}
	if r.ResponseStruct != nil {
		r.ResponseStructName = reflect.TypeOf(r.ResponseStruct).Elem().Name()
		_, r.ResponspeStructNextLink = reflect.TypeOf(r.ResponseStruct).Elem().FieldByName("NextLink")
	}
	if r.StructName == "" {
		r.StructName = reflect.TypeOf(r.Struct).Elem().Name()
	}
	if r.MockListStruct == "" {
		r.MockListStruct = strcase.ToCamel(r.StructName)
	}

	if r.ImportPath == "" {
		r.ImportPath = reflect.TypeOf(r.Struct).Elem().PkgPath()
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
	}

	r.Table, err = codegen.NewTableFromStruct(
		fmt.Sprintf("azure_%s_%s", r.PackageName, r.Name),
		r.Struct,
		opts...,
	)
	if err != nil {
		return fmt.Errorf("failed to create table for %s: %w", r.StructName, err)
	}

	if r.Multiplex == nil {
		if !r.ChildTable {
			if r.ListFuncHasSubscriptionId {
				r.Table.Multiplex = "client.SubscriptionResourceGroupMultiplex"
			} else {
				r.Table.Multiplex = "client.SubscriptionMultiplex"
			}

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
	return nil
}

func generateResource(r recipes.Table, mock bool) error {
	mainTemplate := r.Template + ".go.tpl"
	if mock {
		if r.MockTemplate == "" {
			mainTemplate = r.Template + "_mock.go.tpl"
		} else {
			mainTemplate = r.MockTemplate + ".go.tpl"
		}
	}
	tpl, err := template.New(mainTemplate).Funcs(templateFuncs).ParseFS(templateFS, "templates/"+mainTemplate)
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

	filePath := path.Join(currentDir, "../resources/services", r.PackageName)
	if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
		return err
	}
	if mock {
		filePath = path.Join(filePath, r.Name+"_mock_test.go")
	} else {
		name := r.Name
		if strings.HasSuffix(name, "_test") {
			name = name + "_not"
		}
		filePath = path.Join(filePath, name+".go")
	}

	content := buff.Bytes()
	formattedContent, err := format.Source(buff.Bytes())
	if err != nil {
		fmt.Printf("failed to format code for %s: %v\n", filePath, err)
	} else {
		content = formattedContent
	}
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}
	return nil
}

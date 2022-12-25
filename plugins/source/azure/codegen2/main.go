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
		if err := initTable(nil, &recipes.Tables[i]); err != nil {
			log.Fatal(err)
		}
		if err := generateTable(nil, &recipes.Tables[i]); err != nil {
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

func initTable(parent *recipes.Table, r *recipes.Table) error {
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
	}
	if r.ResponseStruct != nil {
		r.ResponseStructName = reflect.TypeOf(r.ResponseStruct).Elem().Name()
		_, r.ResponspeStructNextLink = reflect.TypeOf(r.ResponseStruct).Elem().FieldByName("NextLink")
	}
	if r.StructName == "" {
		r.StructName = reflect.TypeOf(r.Struct).Elem().Name()
	}

	if r.ImportPath == "" {
		r.ImportPath = reflect.TypeOf(r.Struct).Elem().PkgPath()
	}

	opts := []codegen.TableOption{
		codegen.WithSkipFields(r.SkipFields),
		codegen.WithExtraColumns(r.ExtraColumns),
		codegen.WithPKColumns("id"),
		codegen.WithNameTransformer(func(f reflect.StructField) (string, error) {
			if f.Name == "ETag" {
				return "etag", nil
			}
			return codegen.DefaultNameTransformer(f)
		}),
	}
	tableName := fmt.Sprintf("azure_%s_%s", r.PackageName, r.Name)
	if len(tableName) > 63 {
		panic(fmt.Sprintf("table name %s is too long", tableName))
	}
	r.Table, err = codegen.NewTableFromStruct(
		tableName,
		r.Struct,
		opts...,
	)
	if err != nil {
		return fmt.Errorf("failed to create table for %s: %w", r.StructName, err)
	}
	r.Table.Multiplex = r.Multiplex
	r.Table.Resolver = "fetch" + strcase.ToCamel(r.Name)
	if r.PreResourceResolver != "" {
		r.Table.PreResourceResolver = r.PreResourceResolver
	}
	if r.Relations != nil {
		for _, relation := range r.Relations {
			r.Table.Relations = append(r.Table.Relations, relation.Name+"()")
			if err := initTable(r, relation); err != nil {
				return err
			}
		}
	}
	return nil
}

func generateTable(parent *recipes.Table, r *recipes.Table) error {
	tpl, err := template.New("list.go.tpl").Funcs(templateFuncs).ParseFS(templateFS, "templates/*.go.tpl")
	if err != nil {
		return fmt.Errorf("failed to parse templates: %w", err)
	}
	tpl, err = tpl.ParseFS(codegen.TemplatesFS, "templates/*.go.tpl")
	if err != nil {
		return fmt.Errorf("failed to parse sdk template: %w", err)
	}
	var buff bytes.Buffer
	if err := tpl.Execute(&buff, r); err != nil {
		return fmt.Errorf("failed to execute template for  %w", err)
	}

	filePath := path.Join(currentDir, "../resources/services", r.PackageName)
	if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
		return err
	}
	name := r.Name
	if strings.HasSuffix(name, "_test") {
		name = name + "_not"
	}
	filePath = path.Join(filePath, name+".go")

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
	if !r.SkipMock {
		if err := generateTableMock(r); err != nil {
			return err
		}
	}
	for _, relation := range r.Relations {
		relation.ChildTable = true
		if err := generateTable(r, relation); err != nil {
			return err
		}
	}
	return nil
}

func generateTableMock(r *recipes.Table) error {
	tpl, err := template.New("list_mock.go.tpl").Funcs(templateFuncs).ParseFS(templateFS, "templates/*.go.tpl")
	if err != nil {
		return fmt.Errorf("failed to parse templates: %w", err)
	}
	var buff bytes.Buffer
	if err := tpl.Execute(&buff, r); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	filePath := path.Join(currentDir, "../resources/services", r.PackageName)
	if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
		return err
	}
	filePath = path.Join(filePath, r.Name+"_mock_test.go")

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

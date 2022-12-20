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

	"github.com/cloudquery/cloudquery/plugins/source/slack/codegen/recipes"
	"github.com/cloudquery/cloudquery/plugins/source/slack/codegen/services"
	"github.com/cloudquery/cloudquery/plugins/source/slack/codegen/tables"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/slack-go/slack"
)

//go:embed templates/*.go.tpl
var templatesFS embed.FS

func main() {
	err := services.Generate()
	if err != nil {
		log.Fatal(err)
	}

	var resources []*recipes.Resource
	resources = append(resources, recipes.ConversationResources()...)
	resources = append(resources, recipes.FileResources()...)
	resources = append(resources, recipes.TeamResources()...)
	resources = append(resources, recipes.UserResources()...)

	for _, r := range resources {
		r.Infer()
	}
	for _, r := range resources {
		r.GenerateNames()

		_, filename, _, ok := runtime.Caller(0)
		if !ok {
			log.Fatal("Failed to get caller information")
		}
		codegenDir := path.Join(path.Dir(filename), "..", "resources", "services")

		generateTable(codegenDir, *r)
	}

	err = tables.Generate(resources)
	if err != nil {
		log.Fatal(err)
	}
}

func typeTransformer(f reflect.StructField) (schema.ValueType, error) {
	jsonTimeType := reflect.TypeOf(slack.JSONTime(0))
	isPointerOrInterface := f.Type.Kind() == reflect.Ptr || f.Type.Kind() == reflect.Interface
	if f.Type == jsonTimeType || (isPointerOrInterface && f.Type.Elem() == jsonTimeType) {
		// f is of type slack.JSONTime
		return schema.TypeTimestamp, nil
	}
	return codegen.DefaultTypeTransformer(f)
}

func resolverTransformer(f reflect.StructField, path string) (string, error) {
	jsonTimeType := reflect.TypeOf(slack.JSONTime(0))
	isPointerOrInterface := f.Type.Kind() == reflect.Ptr || f.Type.Kind() == reflect.Interface
	if f.Type == jsonTimeType || (isPointerOrInterface && f.Type.Elem() == jsonTimeType) {
		// f is of type slack.JSONTime
		return fmt.Sprintf(`client.JSONTimeResolver("%s")`, f.Name), nil
	}
	return codegen.DefaultResolverTransformer(f, path)
}

func generateTable(basedir string, r recipes.Resource) {
	var err error

	r.TableName = "slack_" + r.TableName

	log.Println("Generating table", r.TableName)
	opts := []codegen.TableOption{
		codegen.WithSkipFields(r.SkipFields),
		codegen.WithExtraColumns(r.ExtraColumns),
		codegen.WithPKColumns(r.PKColumns...),
		codegen.WithTypeTransformer(typeTransformer),
		codegen.WithResolverTransformer(resolverTransformer),
	}
	if r.UnwrapEmbeddedStructs {
		opts = append(opts, codegen.WithUnwrapAllEmbeddedStructs())
	}
	r.Table, err = codegen.NewTableFromStruct(r.TableName, r.DataStruct, opts...)

	if err != nil {
		log.Fatal(err)
	}

	r.Table.Description = r.Description
	r.Table.Resolver = r.ResolverFuncName
	r.Table.Multiplex = r.Multiplex
	r.ImportClient = strings.HasPrefix(r.Multiplex, "client.")
	r.Table.Relations = r.Relations
	r.Table.PreResourceResolver = r.PreResourceResolver
	r.Table.PostResourceResolver = r.PostResourceResolver

	for _, c := range r.Table.Columns {
		if strings.HasPrefix(c.Resolver, "client.") {
			r.ImportClient = true
		}
	}

	mainTemplate := r.Template + ".go.tpl"
	tpl, err := template.New(mainTemplate).ParseFS(templatesFS, "templates/"+mainTemplate)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse slack templates: %w", err))
	}
	tpl, err = tpl.ParseFS(codegen.TemplatesFS, "templates/*.go.tpl")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse recipes template: %w", err))
	}
	var buff bytes.Buffer
	if err := tpl.Execute(&buff, r); err != nil {
		log.Fatal(fmt.Errorf("failed to execute template: %w", err))
	}

	pkgPath := path.Join(basedir, r.Service)
	if err := os.Mkdir(pkgPath, 0o755); err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	filePath := path.Join(pkgPath, r.Filename)
	content, err := format.Source(buff.Bytes())
	if err != nil {
		fmt.Println(buff.String())
		log.Fatal(fmt.Errorf("failed to format code for %s: %w", filePath, err))
	}
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		log.Fatal(fmt.Errorf("failed to write file %s: %w", filePath, err))
	}
}

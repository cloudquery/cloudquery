package main

import (
	"bytes"
	"embed"
	"fmt"
	"go/format"
	"log"
	"os"
	"path"
	"runtime"
	"text/template"

	"github.com/cloudquery/cloudquery/plugins/source/heroku/codegen/recipes"
	sdkgen "github.com/cloudquery/plugin-sdk/codegen"
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
)

//go:embed templates/*.go.tpl
var templatesFS embed.FS

func main() {
	clearServicesDirectory()

	resources := recipes.All()
	for _, r := range resources {
		generateResource(r, false)
		generateResource(r, true)
	}
}

func clearServicesDirectory() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	dir := path.Dir(filename)
	filePath := path.Join(dir, "../resources/services")
	err := clearDirectory(filePath)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to clear services directory: %w", err))
	}
}

func clearDirectory(dir string) error {
	entry, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, e := range entry {
		if e.IsDir() {
			continue
		}
		err = os.Remove(path.Join(dir, e.Name()))
		if err != nil {
			return err
		}
	}
	return nil
}

func generateResource(r recipes.Resource, mock bool) {
	var err error
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	dir := path.Dir(filename)

	tableName := fmt.Sprintf("heroku_%s", strcase.ToSnake(inflection.Plural(r.HerokuStructName)))
	if r.TableName != "" {
		tableName = r.TableName
	}
	r.Table, err = sdkgen.NewTableFromStruct(tableName, r.HerokuStruct)
	if err != nil {
		log.Fatal(err)
	}
	for i := range r.Table.Columns {
		if r.Table.Columns[i].Name == r.PrimaryKey {
			r.Table.Columns[i].Options.PrimaryKey = true
		}
	}

	r.Table.Resolver = "fetch" + inflection.Plural(r.HerokuStructName)
	r.Table.Description = fmt.Sprintf("https://devcenter.heroku.com/articles/platform-api-reference#%s", strcase.ToKebab(r.HerokuStructName))
	mainTemplate := r.Template + ".go.tpl"
	if mock {
		mainTemplate = r.Template + "_mock_test.go.tpl"
	}
	tpl, err := template.New(mainTemplate).Funcs(template.FuncMap{
		"ToCamel":   strcase.ToCamel,
		"ToSnake":   strcase.ToSnake,
		"Pluralize": inflection.Plural,
	}).ParseFS(templatesFS, "templates/"+mainTemplate)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse gcp templates: %w", err))
	}
	tpl, err = tpl.ParseFS(sdkgen.TemplatesFS, "templates/*.go.tpl")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse recipes template: %w", err))
	}
	var buff bytes.Buffer
	if err := tpl.Execute(&buff, r); err != nil {
		log.Fatal(fmt.Errorf("failed to execute template: %w", err))
	}
	filePath := path.Join(dir, "../resources/services")
	fileName := strcase.ToSnake(r.HerokuStructName)
	if mock {
		filePath = path.Join(filePath, fileName+"_mock_test.go")
	} else {
		filePath = path.Join(filePath, fileName+".go")
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

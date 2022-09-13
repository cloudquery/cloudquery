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
	"strings"
	"text/template"

	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/codegenmain/recipes"
	sdkgen "github.com/cloudquery/plugin-sdk/codegen"
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
)

//go:embed templates/*.go.tpl
var templatesFS embed.FS

func main() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	codegenDir := path.Join(path.Dir(filename), "..", "codegen")

	if err := clearDirectory(codegenDir); err != nil {
		log.Fatal(fmt.Errorf("failed to clear codegen directory: %w", err))
	}

	resources := recipes.All()
	for _, r := range resources {
		generateTable(r)
		if r.Parent != nil && r.TableFuncName != "" {
			r.Parent.Table.Relations = append(r.Parent.Table.Relations, r.TableFuncName+"()")
		}
	}
	for _, r := range resources {
		writeResource(codegenDir, r)
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

func generateTable(r *recipes.Resource) {
	var err error

	tableName := fmt.Sprintf("cloudflare_%s", strcase.ToSnake(inflection.Plural(r.CFStructName)))
	if r.TableName != "" {
		tableName = r.TableName
	}
	log.Println("Generating table", tableName)
	r.Table, err = sdkgen.NewTableFromStruct(tableName, r.CFStruct, sdkgen.WithSkipFields(r.SkipFields))
	if err != nil {
		log.Fatal(err)
	}

	r.Table.Resolver = "services.Fetch" + inflection.Plural(r.CFStructName)
	r.Table.Multiplex = r.Multiplex
	r.ImportClient = strings.HasPrefix(r.Multiplex, "client.")

	r.Table.Columns = append(append(r.DefaultColumns, r.Table.Columns...), r.ExtraColumns...)
	for i := range r.Table.Columns {
		if r.RenameColumns != nil && r.RenameColumns[r.Table.Columns[i].Name] != "" {
			r.Table.Columns[i].Name = r.RenameColumns[r.Table.Columns[i].Name]
		}

		if r.Table.Columns[i].Name == r.PrimaryKey {
			r.Table.Columns[i].Options.PrimaryKey = true
		}
		if strings.HasPrefix(r.Table.Resolver, "client.") {
			r.ImportClient = true
		}
	}

}

func writeResource(basedir string, r *recipes.Resource) {
	mainTemplate := r.Template + ".go.tpl"
	tpl, err := template.New(mainTemplate).Funcs(template.FuncMap{
		"ToCamel":   strcase.ToCamel,
		"ToSnake":   strcase.ToSnake,
		"Pluralize": inflection.Plural,
	}).ParseFS(templatesFS, "templates/"+mainTemplate)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse cf templates: %w", err))
	}
	tpl, err = tpl.ParseFS(sdkgen.TemplatesFS, "templates/*.go.tpl")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse recipes template: %w", err))
	}
	var buff bytes.Buffer
	if err := tpl.Execute(&buff, r); err != nil {
		log.Fatal(fmt.Errorf("failed to execute template: %w", err))
	}
	filePath := path.Join(basedir, r.Filename+".go")
	content, err := format.Source(buff.Bytes())
	if err != nil {
		fmt.Println(buff.String())
		log.Fatal(fmt.Errorf("failed to format code for %s: %w", filePath, err))
	}
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		log.Fatal(fmt.Errorf("failed to write file %s: %w", filePath, err))
	}
}

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

	"github.com/cloudquery/cloudquery/plugins/source/stripe/codegen/recipes"
	"github.com/cloudquery/plugin-sdk/caser"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/gertd/go-pluralize"
)

//go:embed templates/*.go.tpl
var templatesFS embed.FS

func main() {
	for _, r := range recipes.AllResources {
		r.SkipFields = append(r.SkipFields, "APIResource")
		r.Infer()
	}
	if err := recipes.SetParentChildRelationships(recipes.AllResources); err != nil {
		log.Fatal(err)
	}
	for _, r := range recipes.AllResources {
		r.GenerateNames()

		_, filename, _, ok := runtime.Caller(0)
		if !ok {
			log.Fatal("Failed to get caller information")
		}
		codegenDir := path.Join(path.Dir(filename), "..", "resources", "services")

		generateTable(codegenDir, *r)
	}
}

func generateTable(basedir string, r recipes.Resource) {
	var err error

	r.TableName = "stripe_" + r.TableName

	log.Println("Generating table", r.TableName)
	opts := []codegen.TableOption{
		codegen.WithSkipFields(r.SkipFields),
		codegen.WithExtraColumns(r.ExtraColumns),
		codegen.WithPKColumns(r.PKColumns...),
	}
	if r.UnwrapEmbeddedStructs {
		opts = append(opts, codegen.WithUnwrapAllEmbeddedStructs())
	}
	r.Table, err = codegen.NewTableFromStruct(r.TableName, r.DataStruct, opts...)

	if err != nil {
		log.Fatal(err)
	}

	r.Table.Resolver = r.ResolverFuncName
	r.Table.Multiplex = r.Multiplex
	r.ImportClient = strings.HasPrefix(r.Multiplex, "client.") || r.GenerateResolver
	r.Table.Relations = r.Relations
	r.Table.PreResourceResolver = r.PreResourceResolver
	r.Table.PostResourceResolver = r.PostResourceResolver

	for _, c := range r.Table.Columns {
		if strings.HasPrefix(c.Resolver, "client.") {
			r.ImportClient = true
		}
	}

	csr := caser.New()
	pl := pluralize.NewClient()

	templates := []string{
		r.Template + ".go.tpl",
	}
	if !r.SkipMocks {
		templates = append(templates, r.Template+"_test.go.tpl")
	}
	for idx, templateName := range templates {
		generatingMocks := idx == 1

		tpl, err := template.New(templateName).Funcs(template.FuncMap{
			"ToPascal":  csr.ToPascal,
			"Pluralize": pl.Plural,
		}).ParseFS(templatesFS, "templates/"+templateName)

		if err != nil {
			log.Fatal(fmt.Errorf("failed to parse templates: %w", err))
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
		if err := os.Mkdir(pkgPath, 0755); err != nil && !os.IsExist(err) {
			log.Fatal(err)
		}

		var filePath string
		if generatingMocks {
			filePath = path.Join(pkgPath, strings.TrimSuffix(r.Filename, ".go")+"_test.go")
		} else {
			filePath = path.Join(pkgPath, r.Filename)
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
}

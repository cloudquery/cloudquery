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
	"strconv"
	"text/template"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/codegen/recipes"
	"github.com/cloudquery/plugin-sdk/caser"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/gertd/go-pluralize"
)

//go:embed templates/*.go.tpl
var templatesFS embed.FS

func main() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	codegenDir := path.Join(path.Dir(filename), "..", "resources", "services")

	for _, r := range recipes.AllResources {
		r.Infer()
		generateTable(codegenDir, *r)
	}
}

func generateTable(basedir string, r recipes.Resource) {
	log.Println("Generating table", r.TableName)

	csr := caser.New()
	pl := pluralize.NewClient()

	templates := []string{
		"resource.go.tpl",
	}
	if !r.SkipMocks {
		templates = append(templates, "resource_test.go.tpl")
	}
	for idx, templateName := range templates {
		generatingMocks := idx == 1

		tpl, err := template.New(templateName).Funcs(template.FuncMap{
			"ToPascal":    csr.ToPascal,
			"Pluralize":   pl.Plural,
			"Singularize": pl.Singular,
			"QuoteJoin": func(s []string) string {
				var buf bytes.Buffer
				for i, v := range s {
					buf.WriteString(strconv.Quote(v))
					if i != len(s)-1 {
						buf.WriteString(", ")
					}
				}
				return buf.String()
			},
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
			filePath = path.Join(pkgPath, r.TableName+"_test.go")
		} else {
			filePath = path.Join(pkgPath, r.TableName+".go")
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

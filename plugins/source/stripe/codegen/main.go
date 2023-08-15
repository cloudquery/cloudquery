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
	"github.com/cloudquery/plugin-sdk/v4/caser"
	"github.com/gertd/go-pluralize"
)

//go:embed templates/*.go.tpl
var templatesFS embed.FS

func main() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	servicesDir := path.Join(path.Dir(filename), "..", "resources", "services")

	for _, r := range recipes.AllResources {
		r.Infer(nil)
	}

	generateResources(servicesDir, recipes.AllResources)

	if err := generateTables(servicesDir, recipes.AllResources); err != nil {
		log.Fatal(err)
	}
}

func generateResources(servicesDir string, rr []*recipes.Resource) {
	for _, r := range rr {
		generateResource(servicesDir, *r)
		generateResources(servicesDir, r.Children)
	}
}

func generateResource(servicesDir string, r recipes.Resource) {
	log.Println("Generating table", r.TableName)

	csr := caser.New()
	pl := pluralize.NewClient()

	templates := []string{
		r.FetchTemplate + ".go.tpl",
	}
	if !r.SkipMocks {
		templates = append(templates, "resource_test.go.tpl")
	}
	for idx, templateName := range templates {
		generatingMocks := idx == 1

		tpl, err := template.New(templateName).Funcs(template.FuncMap{
			"ToSnake":     csr.ToSnake,
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

		var buff bytes.Buffer
		if err := tpl.Execute(&buff, r); err != nil {
			log.Fatal(fmt.Errorf("failed to execute template: %w", err))
		}

		pkgPath := path.Join(servicesDir, r.Service)
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

func generateTables(servicesDir string, rr []*recipes.Resource) error {
	csr := caser.New()
	pl := pluralize.NewClient()

	tpl, err := template.New("tables.go.tpl").Funcs(template.FuncMap{
		"ToPascal":    csr.ToPascal,
		"Pluralize":   pl.Plural,
		"Singularize": pl.Singular,
	}).ParseFS(templatesFS, "templates/tables.go.tpl")

	var buff bytes.Buffer
	if err := tpl.Execute(&buff, rr); err != nil {
		return fmt.Errorf("failed to execute tables template: %w", err)
	}

	filePath := path.Join(servicesDir, "../plugin/tables.go")
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

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

	"github.com/cloudquery/cloudquery/plugins/source/gandi/codegen/recipes"
	"github.com/cloudquery/plugin-sdk/caser"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/gertd/go-pluralize"
)

//go:embed templates/*.go.tpl
var templatesFS embed.FS

func main() {
	var resources []*recipes.Resource
	resources = append(resources, recipes.CertificateResources()...)
	resources = append(resources, recipes.DomainResources()...)
	resources = append(resources, recipes.LiveDNSResources()...)
	resources = append(resources, recipes.SimpleHostingResources()...)

	pluralizeClient := pluralize.NewClient()
	for _, s := range []string{"livedns", "simplehosting"} {
		pluralizeClient.AddUncountableRule(s)
	}

	csr := caser.New(
		caser.WithCustomInitialisms(map[string]bool{
			"DNS":     true,
			"DNSSec":  true,
			"LiveDNS": true,
			//"Simplehosting": true,
			//"Vhost":         true,
		}),
		caser.WithCustomExceptions(map[string]string{
			"dnssec":  "DNSSec",
			"livedns": "LiveDNS",
			"vhost":   "Vhost",
		}),
	)

	for _, r := range resources {
		// Set defaults and/or infer fields
		if r.Template == "" {
			r.Template = "resource"
		}

		ds := reflect.TypeOf(r.DataStruct)
		if ds.Kind() == reflect.Ptr {
			ds = ds.Elem()
		}
		basepkg := strings.ToLower(path.Base(ds.PkgPath()))

		if r.Package == "" {
			if !pluralizeClient.IsPlural(basepkg) {
				basepkg = pluralizeClient.Plural(basepkg)
			}
			r.Package = basepkg
		}

		if r.TableName == "" {
			// TODO include parent table name in child table name
			n := pluralizeClient.Singular(basepkg) + "_" + csr.ToSnake(ds.Name())
			if !pluralizeClient.IsPlural(n) {
				n = pluralizeClient.Plural(n)
			}
			r.TableName = n
		}

		r.Filename = csr.ToSnake(r.TableName) + ".go"
		r.TableFuncName = csr.ToPascal(r.TableName)
		r.ResolverFuncName = "fetch" + r.TableFuncName

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

	r.TableName = "gandi_" + r.TableName

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
		log.Fatal(fmt.Errorf("failed to parse gandi templates: %w", err))
	}
	tpl, err = tpl.ParseFS(codegen.TemplatesFS, "templates/*.go.tpl")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse recipes template: %w", err))
	}
	var buff bytes.Buffer
	if err := tpl.Execute(&buff, r); err != nil {
		log.Fatal(fmt.Errorf("failed to execute template: %w", err))
	}

	pkgPath := path.Join(basedir, r.Package)
	if err := os.Mkdir(pkgPath, 0755); err != nil && !os.IsExist(err) {
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

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

	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/codegen/recipes"
	sdkgen "github.com/cloudquery/plugin-sdk/codegen"
)

//go:embed templates/*.go.tpl
var templatesFS embed.FS

func main() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	codegenDir := path.Join(path.Dir(filename), "..", "resources", "services")

	var resources []recipes.Resource
	resources = append(resources, recipes.AccessGroupResources()...)
	resources = append(resources, recipes.AccountResources()...)
	resources = append(resources, recipes.CertificatePackResources()...)
	resources = append(resources, recipes.DNSRecordResources()...)
	resources = append(resources, recipes.ImageResources()...)
	resources = append(resources, recipes.WAFOverrideResources()...)
	resources = append(resources, recipes.WAFPackageResources()...)
	resources = append(resources, recipes.WorkerMetaDataResources()...)
	resources = append(resources, recipes.WorkerRouteResources()...)
	resources = append(resources, recipes.ZoneResources()...)

	for _, r := range resources {
		generateTable(codegenDir, r)
	}
}

func generateTable(basedir string, r recipes.Resource) {
	var err error

	log.Println("Generating table", r.TableName)
	r.Table, err = sdkgen.NewTableFromStruct(r.TableName, r.CFStruct, sdkgen.WithSkipFields(r.SkipFields))
	if err != nil {
		log.Fatal(err)
	}

	r.Table.Resolver = r.ResolverFuncName
	r.Table.Multiplex = r.Multiplex
	r.ImportClient = strings.HasPrefix(r.Multiplex, "client.")
	r.Table.Relations = r.Relations

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

	mainTemplate := r.Template + ".go.tpl"
	tpl, err := template.New(mainTemplate).ParseFS(templatesFS, "templates/"+mainTemplate)
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

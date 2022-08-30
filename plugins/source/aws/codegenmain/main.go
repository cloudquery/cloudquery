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
	"strconv"
	"strings"
	"text/template"

	"github.com/cloudquery/cloudquery/plugins/source/aws/codegenmain/helpers"
	"github.com/cloudquery/cloudquery/plugins/source/aws/codegenmain/recipes"
	sdkgen "github.com/cloudquery/plugin-sdk/codegen"
	pluginschema "github.com/cloudquery/plugin-sdk/schema"
	"github.com/iancoleman/strcase"
	"golang.org/x/exp/slices"
)

//go:embed templates/*.go.tpl
var awsTemplatesFS embed.FS

var resources []*recipes.Resource

func main() {
	resources = append(resources, recipes.ACMResources...)
	resources = append(resources, recipes.APIGatewayv2Resources...)
	resources = append(resources, recipes.ApplicationautoscalingResources...)
	resources = append(resources, recipes.AppsyncResources...)
	resources = append(resources, recipes.AthenaResources...)
	resources = append(resources, recipes.AutoscalingResources...)

	for _, r := range resources {
		generateResource(r, false)
		generateResource(r, true)
	}
}

func generateResource(r *recipes.Resource, mock bool) {
	var err error
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	dir := path.Dir(filename)

	tableNameFromSubService, fetcherNameFromSubService := helpers.TableAndFetcherNames(r)

	r.Table, err = sdkgen.NewTableFromStruct(
		fmt.Sprintf("aws_%s_%s", strings.ToLower(r.AWSService), tableNameFromSubService),
		r.AWSStruct,
		sdkgen.WithSkipFields(append(r.SkipFields, "noSmithyDocumentSerde")),
	)
	if err != nil {
		log.Fatal(err)
	}

	if r.TrimPrefix != "" {
		for i := range r.Table.Columns {
			r.Table.Columns[i].Name = strings.TrimPrefix(r.Table.Columns[i].Name, r.TrimPrefix)
		}
	}

	r.Table.Columns = append(r.DefaultColumns, r.Table.Columns...)
	if r.ColumnOverrides != nil {
		for i, c := range r.Table.Columns {
			override, ok := r.ColumnOverrides[c.Name]
			if !ok {
				continue
			}
			r.Table.Columns[i].Name = helpers.Coalesce(override.Name, r.Table.Columns[i].Name)
			r.Table.Columns[i].Resolver = helpers.Coalesce(override.Resolver, r.Table.Columns[i].Resolver)
			r.Table.Columns[i].Description = helpers.Coalesce(override.Description, r.Table.Columns[i].Description)

			delete(r.ColumnOverrides, c.Name)
		}
		// remaining, unmatched columns are added to the end of the table. Difference from DefaultColumns? none for now
		for k, c := range r.ColumnOverrides {
			if c.Type == pluginschema.TypeInvalid {
				if !mock {
					fmt.Println("Not adding unmatched column with unspecified type", k, c)
				}
				continue
			}
			c.Name = helpers.Coalesce(c.Name, k)
			r.Table.Columns = append(r.Table.Columns, c)
		}
	}

	hasReferenceToResolvers := false

	for i := range r.Table.Columns {
		if len(r.Table.Options.PrimaryKeys) == 0 && r.Table.Columns[i].Name == "arn" {
			//	r.Table.Columns[i].Options.PrimaryKey = true
			r.Table.Options.PrimaryKeys = []string{"arn"}
		}
		if r.Table.Columns[i].Name == "tags" {
			r.HasTags = true

			if r.Table.Columns[i].Resolver == recipes.ResolverAuto {
				r.Table.Columns[i].Resolver = "resolve" + r.AWSService + r.AWSSubService + "Tags"
			}
		}
		if strings.HasPrefix(r.Table.Columns[i].Resolver, "resolvers.") {
			hasReferenceToResolvers = true
		}
	}
	r.Table.Multiplex = `client.ServiceAccountRegionMultiplexer("` + helpers.Coalesce(r.MultiplexerServiceOverride, strings.ToLower(r.AWSService)) + `")`

	r.Table.Resolver = "fetch" + r.AWSService + fetcherNameFromSubService
	r.TableFuncName = strings.TrimPrefix(r.Table.Resolver, "fetch")

	if mock {
		r.MockFuncName = "build" + r.TableFuncName
		r.TestFuncName = "Test" + r.TableFuncName
	}

	t := reflect.TypeOf(r.AWSStruct).Elem()
	r.AWSStructName = helpers.Coalesce(r.AWSStructName, t.Name())
	r.ItemName = helpers.Coalesce(r.ItemName, r.AWSStructName)

	if !mock {
		r.Imports = quoteImports(r.Imports)
	} else {
		r.MockImports = quoteImports(r.MockImports)
	}

	r.TypesImport = ""
	if sp := t.PkgPath(); strings.HasSuffix(sp, "/types") {
		r.TypesImport = sp
		if r.AddTypesImport {
			if !mock {
				r.Imports = append(r.Imports, strconv.Quote(sp))
			} else {
				r.MockImports = append(r.MockImports, strconv.Quote(sp))
			}
		}
		r.Imports = append(r.Imports, strconv.Quote(strings.TrimSuffix(sp, "/types")))         // auto-import main pkg (not "types")
		r.MockImports = append(r.MockImports, strconv.Quote(strings.TrimSuffix(sp, "/types"))) // auto-import main pkg (not "types")
	}

	if hasReferenceToResolvers && !mock {
		res := "resolvers " + strconv.Quote(`github.com/cloudquery/cloudquery/plugins/source/aws/codegenmain/resolvers/`+strings.ToLower(r.AWSService))
		if !slices.Contains(r.Imports, res) {
			r.Imports = append(r.Imports, res)
		}
	}

	mainTemplate := r.Template + helpers.StringSwitch(mock, "_mock_test", "") + ".go.tpl"
	tpl, err := template.New(mainTemplate).Funcs(template.FuncMap{
		"ToCamel":  strcase.ToCamel,
		"ToLower":  strings.ToLower,
		"Coalesce": func(a1, a2 string) string { return helpers.Coalesce(a2, a1) }, // go templates argument order is backwards
	}).ParseFS(awsTemplatesFS, "templates/*.go.tpl")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse aws templates: %w", err))
	}
	tpl, err = tpl.ParseFS(sdkgen.TemplatesFS, "templates/*.go.tpl")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse codegen template: %w", err))
	}
	var buff bytes.Buffer
	if err := tpl.Execute(&buff, r); err != nil {
		log.Fatal(fmt.Errorf("failed to execute template: %w", err))
	}
	filePath := path.Join(dir, "../codegen", strings.ToLower(r.AWSService))
	if err := os.MkdirAll(filePath, 0755); err != nil {
		log.Fatal(fmt.Errorf("failed to create directory: %w", err))
	}

	fileSuffix := helpers.StringSwitch(mock, "_mock_test.go", ".go")
	filePath = path.Join(filePath, strings.ToLower(r.AWSService)+"_"+tableNameFromSubService+fileSuffix)
	content, err := format.Source(buff.Bytes())
	if err != nil {
		fmt.Println(buff.String())
		log.Fatal(fmt.Errorf("failed to format code for %s: %w", filePath, err))
	}
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		log.Fatal(fmt.Errorf("failed to write file %s: %w", filePath, err))
	}
}

func quoteImports(imports []string) []string {
	for i := range imports {
		if !strings.HasSuffix(imports[i], `"`) {
			imports[i] = strconv.Quote(imports[i])
		}
	}
	return imports
}

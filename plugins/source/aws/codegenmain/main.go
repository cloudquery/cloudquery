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

	"github.com/cloudquery/cloudquery/plugins/source/aws/codegenmain/recipes"
	sdkgen "github.com/cloudquery/plugin-sdk/codegen"
	pluginschema "github.com/cloudquery/plugin-sdk/schema"
	"github.com/iancoleman/strcase"
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

	cqSubservice := coalesce(r.CQSubserviceOverride, r.AWSSubService)

	tableNameFromSubService := strcase.ToSnake(cqSubservice)
	fetcherNameFromSubService := strcase.ToCamel(cqSubservice)
	{
		// Generate table and fetcher names using parent info

		prev := tableNameFromSubService
		var (
			preTableNames   []string
			preFetcherNames []string
		)
		rp := r.Parent
		for rp != nil {
			if rp.CQSubserviceOverride != "" {
				preTableNames = append(preTableNames, rp.CQSubserviceOverride)
				preFetcherNames = append(preFetcherNames, strcase.ToCamel(rp.CQSubserviceOverride))
			} else {
				ins := strcase.ToSnake(rp.ItemName)
				if !strings.HasPrefix(prev, ins) {
					preTableNames = append(preTableNames, ins)
					preFetcherNames = append(preFetcherNames, strcase.ToCamel(rp.ItemName))
					prev = ins
				}
			}
			rp = rp.Parent
		}
		if len(preTableNames) > 0 {
			tableNameFromSubService = strings.Join(reverseStringSlice(preTableNames), "_") + "_" + tableNameFromSubService
			fetcherNameFromSubService = strings.Join(reverseStringSlice(preFetcherNames), "") + fetcherNameFromSubService
		}
	}

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
			if override.Name != "" {
				r.Table.Columns[i].Name = override.Name
			}
			if override.Resolver != "" {
				r.Table.Columns[i].Resolver = override.Resolver
			}
			if override.Description != "" {
				r.Table.Columns[i].Description = override.Description
			}
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
			if c.Name == "" {
				c.Name = k
			}
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
	r.Table.Multiplex = `client.ServiceAccountRegionMultiplexer("` + coalesce(r.MultiplexerServiceOverride, strings.ToLower(r.AWSService)) + `")`

	r.Table.Resolver = "fetch" + r.AWSService + fetcherNameFromSubService
	r.TableFuncName = strings.TrimPrefix(r.Table.Resolver, "fetch")

	if mock {
		r.MockFuncName = "build" + r.TableFuncName
		r.TestFuncName = "Test" + r.TableFuncName
	}

	t := reflect.TypeOf(r.AWSStruct).Elem()
	if r.AWSStructName == "" {
		r.AWSStructName = t.Name()
	}

	if r.ItemName == "" {
		r.ItemName = r.AWSStructName
	}

	if !mock {
		for i := range r.Imports {
			if !strings.HasSuffix(r.Imports[i], `"`) {
				r.Imports[i] = strconv.Quote(r.Imports[i])
			}
		}
	} else {
		for i := range r.MockImports {
			if !strings.HasSuffix(r.MockImports[i], `"`) {
				r.MockImports[i] = strconv.Quote(r.MockImports[i])
			}
		}
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
		found := false
		for i := range r.Imports {
			if r.Imports[i] == res {
				found = true
				break
			}
		}
		if !found {
			r.Imports = append(r.Imports, res)
		}
	}

	mainTemplate := r.Template + stringSwitch(mock, "_mock_test", "") + ".go.tpl"
	tpl, err := template.New(mainTemplate).Funcs(template.FuncMap{
		"ToCamel":  strcase.ToCamel,
		"ToLower":  strings.ToLower,
		"Coalesce": func(a1, a2 string) string { return coalesce(a2, a1) }, // go templates argument order is backwards
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
	filePath := path.Join(dir, "../codegen")
	if err := os.MkdirAll(filePath, 0755); err != nil {
		log.Fatal(fmt.Errorf("failed to create directory: %w", err))
	}

	fileSuffix := stringSwitch(mock, "_mock_test.go", ".go")
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

func stringSwitch(b bool, ifTrue, ifFalse string) string {
	if b {
		return ifTrue
	}
	return ifFalse
}

func coalesce(input, defValue string) string {
	if input == "" {
		return defValue
	}
	return input
}

func reverseStringSlice(input []string) []string {
	ret := make([]string, 0, len(input))
	for i := len(input) - 1; i >= 0; i-- {
		ret = append(ret, input[i])
	}
	return ret
}

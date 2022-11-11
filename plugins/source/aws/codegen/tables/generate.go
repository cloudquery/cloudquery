package tables

import (
	"bytes"
	"embed"
	"fmt"
	"go/format"
	"os"
	"path"
	"runtime"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"

	"github.com/cloudquery/cloudquery/plugins/source/aws/codegen/recipes"
)

//go:embed templates/*.go.tpl
var templatesFS embed.FS

func Generate(resources []*recipes.Resource) error {
	tpl, err := template.New("tables.go.tpl").Funcs(template.FuncMap{
		"ToCamel": strcase.ToCamel,
	}).ParseFS(templatesFS, "templates/tables.go.tpl")
	if err != nil {
		return err
	}

	resources = removeChildResources(resources)
	var buff bytes.Buffer
	if err := tpl.Execute(&buff, resources); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("failed to get caller information")
	}

	filePath := path.Join(path.Dir(filename), "../../resources/plugin/tables.go")
	content := buff.Bytes()
	formattedContent, err := format.Source(buff.Bytes())
	if err != nil {
		fmt.Printf("failed to format source: %s: %v\n", filePath, err)
	} else {
		content = formattedContent
	}
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}

	return nil
}

func removeChildResources(resources []*recipes.Resource) []*recipes.Resource {
	filtered := make([]*recipes.Resource, 0)
	relations := map[string]bool{}
	for _, r := range resources {
		for _, rel := range r.Relations {
			relations[r.Service+"."+strings.TrimSuffix(rel, "()")] = true
		}
	}
	for _, r := range resources {
		funcName := r.Service + "." + strcase.ToCamel(r.SubService)
		if relations[funcName] {
			continue
		}
		filtered = append(filtered, r)
	}
	return filtered
}

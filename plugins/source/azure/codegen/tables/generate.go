package tables

import (
	"bytes"
	"embed"
	"fmt"
	"go/format"
	"os"
	"path"
	"runtime"
	"text/template"

	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/resource"
	"github.com/iancoleman/strcase"
)

//go:embed templates/*.go.tpl
var templatesFS embed.FS

func Generate(resources []*resource.Resource) error {
	tpl, err := template.New("tables.go.tpl").Funcs(template.FuncMap{
		"ToCamel": strcase.ToCamel,
	}).ParseFS(templatesFS, "templates/tables.go.tpl")
	if err != nil {
		return err
	}

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

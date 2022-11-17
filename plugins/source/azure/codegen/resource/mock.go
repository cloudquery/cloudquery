package resource

import (
	"bytes"
	"fmt"
	"path"
	"strings"
	"text/template"

	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/util"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

func (r *Resource) generateMockTest(dir string) error {
	templateName := "mock_test.go.tpl"
	if r.parent != nil {
		templateName = "mock_test_relation.go.tpl"
	}
	tpl, err := template.New(templateName).Funcs(template.FuncMap{
		"Singular":     pluralize.NewClient().Singular,
		"ToCamel":      strcase.ToCamel,
		"ToLower":      strings.ToLower,
		"ToLowerCamel": strcase.ToLowerCamel,
	}).ParseFS(templatesFS, "templates/*.go.tpl")
	if err != nil {
		return fmt.Errorf("failed to parse templates: %w", err)
	}
	tpl, err = tpl.ParseFS(codegen.TemplatesFS, "templates/*.go.tpl")
	if err != nil {
		return fmt.Errorf("failed to parse sdk template: %w", err)
	}

	var buff bytes.Buffer
	if err := tpl.Execute(&buff, r); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	filePath := path.Join(dir, r.SubService+"_mock_test.go")
	return util.WriteAndFormat(filePath, buff.Bytes())
}

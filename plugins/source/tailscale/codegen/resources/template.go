package resources

import (
	"embed"
	"strings"
	"text/template"

	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/iancoleman/strcase"
)

//go:embed templates/*.go.tpl
var templatesFS embed.FS

func parse(name string) (*template.Template, error) {
	tpl, err := template.New(name+".go.tpl").Funcs(template.FuncMap{
		"ToCamel":      strcase.ToCamel,
		"ToLower":      strings.ToLower,
		"ToLowerCamel": strcase.ToLowerCamel,
	}).ParseFS(templatesFS, "templates/*.go.tpl")
	if err != nil {
		return nil, err
	}
	return tpl.ParseFS(codegen.TemplatesFS, "templates/*.go.tpl")
}

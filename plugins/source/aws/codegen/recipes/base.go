package recipes

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/iancoleman/strcase"
)

type Resource struct {
	Struct interface{}
}

func (r *Resource) Generate() error {
	r.Table, err = codegen.NewTableFromStruct(
		fmt.Sprintf("gcp_%s_%s", r.Service, r.SubService),
		r.Struct,
		codegen.WithSkipFields(r.SkipFields),
		codegen.WithOverrideColumns(r.OverrideColumns),
		codegen.WithExtraColumns(r.DefaultColumns),
	)

	tpl, err := template.New(mainTemplate).Funcs(template.FuncMap{
		"ToCamel": strcase.ToCamel,
		"ToLower": strings.ToLower,
	}).ParseFS(gcpTemplatesFS, "templates/"+mainTemplate)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse gcp templates: %w", err))
	}
	tpl, err = tpl.ParseFS(sdkgen.TemplatesFS, "templates/*.go.tpl")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse sdk template: %w", err))
	}

	var buff bytes.Buffer
	if err := tpl.Execute(&buff, r); err != nil {
		log.Fatal(fmt.Errorf("failed to execute template: %w", err))
	}
	filePath := path.Join(dir, "../resources/services", r.Service)
	if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}
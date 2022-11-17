package resource

import (
	"bytes"
	"fmt"
	"path"
	"reflect"
	"strings"
	"text/template"

	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/util"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

func (r *Resource) generateSchema(dir string) error {
	tpl, err := template.New("resource.go.tpl").Funcs(template.FuncMap{
		"Singular":     pluralize.NewClient().Singular,
		"ToCamel":      strcase.ToCamel,
		"ToLower":      strings.ToLower,
		"ToLowerCamel": strcase.ToLowerCamel,
	}).ParseFS(templatesFS, "templates/resource.go.tpl")
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

	filePath := path.Join(dir, r.SubService+".go")
	return util.WriteAndFormat(filePath, buff.Bytes())
}

func fixStringArray(field reflect.StructField) (schema.ValueType, error) {
	typ := field.Type
	if typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}
	if typ.Kind() == reflect.Slice {
		typ = typ.Elem()
		if typ.Kind() == reflect.Pointer {
			typ = typ.Elem()
		}
		if typ.Kind() == reflect.String {
			return schema.TypeStringArray, nil
		}
	}
	// pass through to default
	return schema.TypeInvalid, nil
}

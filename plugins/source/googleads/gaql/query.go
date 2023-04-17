package gaql

import (
	"reflect"
	"strings"

	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func Query(a any, parent *schema.Resource, o ...*Options) string {
	var options *Options
	if len(o) > 0 {
		options = o[0]
	}

	tableName := FieldName(a)
	query := "SELECT\n\t" +
		strings.Join(selectFields(a, options, tableName+"."), ",\n\t") +
		"\nFROM " + tableName

	if parent == nil {
		return query
	}
	return query + "\n" + "WHERE " + tableName + "." + where(parent)
}

func selectFields(a any, o *Options, prefix string) []string {
	ref := reflect.TypeOf(a)
	if ref.Kind() == reflect.Pointer {
		ref = ref.Elem()
	}

	visible := reflect.VisibleFields(ref)
	fields := make([]string, 0, len(visible)) // might grow past this, though
	for _, fld := range visible {
		if o.skip(fld.Name) {
			continue
		}

		tag, ok := jsonTag(fld)
		if !ok {
			continue
		}

		name := prefix + tag
		if o.expand(fld.Name) {
			fields = append(fields,
				selectFields(
					reflect.New(fld.Type).Elem().Interface(),
					o.trim(fld.Name),
					name+".",
				)...,
			)
			continue
		}

		fields = append(fields, name)
	}

	return fields
}

func where(p *schema.Resource) string {
	parent := p.Item.(interface {
		GetResourceName() string
	})
	return FieldName(parent) + " = " + quoted(parent.GetResourceName())
}

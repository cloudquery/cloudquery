package util

import (
	"strings"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

// TODO: remove after https://github.com/cloudquery/plugin-sdk/pull/817 is released
func FieldChangePrettified(fc schema.FieldChange) string {
	switch fc.Type {
	case schema.TableColumnChangeTypeAdd:
		return "+ " + fieldPrettify(fc.Current)
	case schema.TableColumnChangeTypeRemove:
		return "- " + fieldPrettify(fc.Previous)
	case schema.TableColumnChangeTypeUpdate:
		return "~ " + fieldPrettify(fc.Previous) + " -> " + fieldPrettify(fc.Current)
	default:
		return "? " + fieldPrettify(fc.Previous) + " -> " + fieldPrettify(fc.Current)
	}
}

func fieldPrettify(field arrow.Field) string {
	builder := new(strings.Builder)
	builder.WriteString(field.Name)
	builder.WriteString(": ")

	if field.Nullable {
		builder.WriteString("nullable(")
	}
	builder.WriteString(field.Type.String())
	if field.Nullable {
		builder.WriteString(")")
	}

	if field.HasMetadata() {
		builder.WriteString(", metadata: ")
		builder.WriteString(field.Metadata.String())
	}
	return builder.String()
}

func SchemaChangesPrettified(tableName string, changes []schema.FieldChange) string {
	builder := new(strings.Builder)
	builder.WriteString(tableName + ":")
	for _, c := range changes {
		builder.WriteString("\n  ")
		builder.WriteString(FieldChangePrettified(c))
	}
	return builder.String()
}

func SchemasChangesPrettified(changes map[string][]schema.FieldChange) string {
	builder := new(strings.Builder)
	left := len(changes)
	for name, change := range changes {
		left--
		builder.WriteString(SchemaChangesPrettified(name, change))
		if left > 0 {
			builder.WriteString("\n")
		}
	}
	return builder.String()
}

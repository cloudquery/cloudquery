package util

import (
	"strings"

	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func prettifyChanges(tableName string, changes []schema.TableColumnChange, builder *strings.Builder) {
	builder.WriteString(tableName + ":")
	for _, c := range changes {
		builder.WriteString("\n  ")
		builder.WriteString(c.String())
	}
}

func ChangesPrettified(tableName string, changes []schema.TableColumnChange) string {
	builder := new(strings.Builder)
	prettifyChanges(tableName, changes, builder)
	return builder.String()
}

func SchemasChangesPrettified(changes map[string][]schema.TableColumnChange) string {
	builder := new(strings.Builder)
	left := len(changes)
	for name, change := range changes {
		left--
		prettifyChanges(name, change, builder)
		if left > 0 {
			builder.WriteString("\n")
		}
	}
	return builder.String()
}

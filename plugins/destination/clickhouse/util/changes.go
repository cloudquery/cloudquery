package util

import (
	"strings"

	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func ChangesPrettified(tableName string, changes []schema.TableColumnChange) string {
	builder := new(strings.Builder)
	builder.WriteString(tableName + ":")
	for _, c := range changes {
		builder.WriteString("\n  ")
		builder.WriteString(c.String())
	}
	return builder.String()
}

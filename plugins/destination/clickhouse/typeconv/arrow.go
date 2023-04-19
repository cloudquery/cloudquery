package typeconv

import (
	"strings"

	"github.com/apache/arrow/go/v12/arrow"
)

func ArrowField(name, typ string) arrow.Field {
	field := arrow.Field{Name: name}
	if strings.HasPrefix(typ, "Nullable(") {
		field.Nullable = true
		typ = strings.TrimSuffix(strings.TrimPrefix(typ, "Nullable("), ")")
	}
	field.Type = arrowDataType(typ)
	return field
}
func arrowDataType(typ string) arrow.DataType {
	return arrow.ListOfField(arrow.Field{Name: "123"})
}

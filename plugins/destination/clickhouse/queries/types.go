package queries

import (
	"strings"

	"github.com/apache/arrow/go/v12/arrow"
)

func chType(dataType arrow.DataType) string {
	switch dataType.ID() {
	case arrow.BOOL:
	}
	return "String"
}

func chFieldType(field arrow.Field) string {
	typ := chType(field.Type)

	// We allow nullable values in arrays, but arrays cannot be nullable themselves
	if field.Type.ID() == arrow.LIST || !field.Nullable {
		return typ
	}

	return "Nullable(" + typ + ")"
}

func arrowField(name, typ string) arrow.Field {
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

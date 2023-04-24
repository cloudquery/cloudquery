package definitions

import (
	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/util"
)

func FieldType(field arrow.Field) string {
	typ := dataType(field.Type)

	// We allow nullable values in arrays, but arrays cannot be nullable themselves
	if field.Type.ID() == arrow.LIST || !field.Nullable {
		return typ
	}

	return "Nullable(" + typ + ")"
}

func FieldDefinition(field arrow.Field) string {
	return util.SanitizeID(field.Name) + " " + FieldType(field)
}

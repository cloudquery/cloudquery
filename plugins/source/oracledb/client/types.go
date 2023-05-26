package client

import (
	"strings"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/types"
)

func SQLType(t arrow.DataType) string {
	switch {
	case arrow.TypeEqual(arrow.FixedWidthTypes.Boolean, t):
		return "char(1)"
	case arrow.TypeEqual(arrow.PrimitiveTypes.Int64, t):
		return "NUMBER(19)"
	case arrow.TypeEqual(arrow.PrimitiveTypes.Float64, t):
		return "binary_double"
	case arrow.TypeEqual(types.ExtensionTypes.UUID, t):
		return "raw(16)"
	case arrow.IsBinaryLike(t.ID()):
		return "blob"
	case arrow.TypeEqual(arrow.FixedWidthTypes.Timestamp_us, t):
		return "timestamp"
	default:
		return "clob"
	}
}

func SchemaType(tableName string, columnName string, dataType string) arrow.DataType {
	if strings.HasPrefix(columnName, "timestamp") {
		return arrow.FixedWidthTypes.Timestamp_us
	}

	switch dataType {
	case "raw(16)":
		return types.ExtensionTypes.UUID
	case "char(1)":
		return arrow.FixedWidthTypes.Boolean
	case "float", "binary_float", "binary_double":
		return arrow.PrimitiveTypes.Float64
	case "binary":
		return arrow.BinaryTypes.Binary
	case "number":
		return arrow.PrimitiveTypes.Int64
	case "blob", "raw", "long raw":
		return arrow.BinaryTypes.Binary
	}

	return arrow.BinaryTypes.String
}

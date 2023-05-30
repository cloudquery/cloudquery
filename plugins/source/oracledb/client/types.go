package client

import (
	"strings"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/types"
)

func SQLType(t arrow.DataType) string {
	switch t.(type) {
	case *arrow.BooleanType:
		return "char(1)"
	case *arrow.Int8Type, *arrow.Int16Type, *arrow.Int32Type, *arrow.Int64Type,
		*arrow.Uint8Type, *arrow.Uint16Type, *arrow.Uint32Type, *arrow.Uint64Type:
		return "NUMBER(19)"
	case *arrow.Float16Type, *arrow.Float32Type, *arrow.Float64Type:
		return "binary_double"
	case *types.UUIDType:
		return "raw(16)"
	case *arrow.BinaryType, *arrow.LargeBinaryType:
		return "blob"
	case *arrow.TimestampType:
		return "timestamp"
	default:
		return "clob"
	}
}

func SchemaType(dataType string) arrow.DataType {
	dataTypeLower := strings.ToLower(dataType)
	if strings.HasPrefix(dataTypeLower, "timestamp") {
		return arrow.FixedWidthTypes.Timestamp_us
	}

	if strings.HasPrefix(dataTypeLower, "number") {
		return arrow.PrimitiveTypes.Int64
	}

	switch dataTypeLower {
	case "raw(16)":
		return types.ExtensionTypes.UUID
	case "char(1)":
		return arrow.FixedWidthTypes.Boolean
	case "float", "binary_float", "binary_double":
		return arrow.PrimitiveTypes.Float64
	case "binary", "blob", "raw", "long raw":
		return arrow.BinaryTypes.Binary
	}

	return arrow.BinaryTypes.String
}

package client

import (
	"strconv"
	"strings"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

const defaultPrecision = 38
const defaultScale = 0

func getValue(str string, defaultValue int32) int32 {
	val, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return defaultValue
	}
	return int32(val)
}

func getPrecisionAndScale(dataType string) (precision, scale int32) {
	str := strings.TrimPrefix(dataType, "number")
	if str == "" {
		return defaultPrecision, defaultScale
	}
	str = strings.TrimPrefix(str, "(")
	str = strings.TrimSuffix(str, ")")
	parts := strings.Split(str, ",")

	switch len(parts) {
	case 1:
		precision = getValue(parts[0], defaultPrecision)
		scale = defaultScale
	case 2:
		precision = getValue(parts[0], defaultPrecision)
		scale = getValue(parts[1], defaultScale)
	default:
		precision = defaultPrecision
		scale = defaultScale
	}
	return precision, scale
}

func SQLType(t arrow.DataType) string {
	switch t.(type) {
	case *arrow.BooleanType:
		return "char(1)"
	case *arrow.Int8Type, *arrow.Int16Type, *arrow.Int32Type, *arrow.Int64Type,
		*arrow.Uint8Type, *arrow.Uint16Type, *arrow.Uint32Type, *arrow.Uint64Type:
		return "NUMBER(38)"
	case *arrow.Float16Type, *arrow.Float32Type:
		return "binary_float"
	case *arrow.Float64Type:
		return "binary_double"
	case *types.UUIDType:
		return "raw(16)"
	case *arrow.BinaryType, *arrow.LargeBinaryType, *arrow.FixedSizeBinaryType:
		return "blob"
	case *arrow.TimestampType:
		return "timestamp(9)"
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
		precision, scale := getPrecisionAndScale(dataTypeLower)
		return &arrow.Decimal128Type{Precision: precision, Scale: scale}
	}

	switch dataTypeLower {
	case "raw(16)":
		return types.ExtensionTypes.UUID
	case "char(1)":
		return arrow.FixedWidthTypes.Boolean
	case "binary_float":
		return arrow.PrimitiveTypes.Float32
	case "float", "binary_double":
		return arrow.PrimitiveTypes.Float64
	case "binary", "blob", "raw", "long raw":
		return arrow.BinaryTypes.Binary
	}

	return arrow.BinaryTypes.String
}

package client

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

const defaultPrecision = 10
const defaultScale = 0

func getValue(str string, defaultValue int32) int32 {
	val, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return defaultValue
	}
	return int32(val)
}

func getPrecisionAndScale(dataType string) (precision, scale int32) {
	str := strings.TrimPrefix(dataType, "decimal")
	str = strings.TrimPrefix(str, "numeric")
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
	switch v := t.(type) {
	case *arrow.BooleanType:
		return "bool"
	case *arrow.Int8Type:
		return "tinyint"
	case *arrow.Int16Type:
		return "smallint"
	case *arrow.Int32Type:
		return "int"
	case *arrow.Int64Type:
		return "bigint"
	case *arrow.Uint8Type:
		return "tinyint unsigned"
	case *arrow.Uint16Type:
		return "smallint unsigned"
	case *arrow.Uint32Type:
		return "int unsigned"
	case *arrow.Uint64Type:
		return "bigint unsigned"
	case *arrow.Float16Type, *arrow.Float32Type:
		return "float"
	case *arrow.Float64Type:
		return "double"
	case arrow.DecimalType:
		return fmt.Sprintf("decimal(%d,%d)", v.GetPrecision(), v.GetScale())
	case *types.UUIDType:
		return "binary(16)"
	case *arrow.BinaryType, *arrow.LargeBinaryType, *arrow.FixedSizeBinaryType:
		return "blob"
	case *arrow.TimestampType:
		return "datetime(6)"
	case *types.JSONType:
		return "json"
	default:
		return "text"
	}
}

func isUnsigned(columnType string) bool {
	return strings.HasSuffix(columnType, " unsigned")
}

func SchemaType(dataType string, columnType string) arrow.DataType {
	if columnType == "binary(16)" {
		return types.ExtensionTypes.UUID
	}
	if columnType == "tinyint(1)" {
		return arrow.FixedWidthTypes.Boolean
	}
	if strings.HasPrefix(columnType, "datetime") {
		return arrow.FixedWidthTypes.Timestamp_us
	}
	if strings.HasPrefix(columnType, "decimal") || strings.HasPrefix(columnType, "numeric") {
		precision, scale := getPrecisionAndScale(columnType)
		return &arrow.Decimal128Type{Precision: precision, Scale: scale}
	}
	if strings.HasPrefix(columnType, "tinyint") {
		if isUnsigned(columnType) {
			return arrow.PrimitiveTypes.Uint8
		}
		return arrow.PrimitiveTypes.Int8
	}
	if strings.HasPrefix(columnType, "smallint") {
		if isUnsigned(columnType) {
			return arrow.PrimitiveTypes.Uint16
		}
		return arrow.PrimitiveTypes.Int16
	}
	if strings.HasPrefix(columnType, "int") {
		if isUnsigned(columnType) {
			return arrow.PrimitiveTypes.Uint32
		}
		return arrow.PrimitiveTypes.Int32
	}
	if strings.HasPrefix(columnType, "bigint") {
		if isUnsigned(columnType) {
			return arrow.PrimitiveTypes.Uint64
		}
		return arrow.PrimitiveTypes.Int64
	}
	switch dataType {
	case "bool", "boolean":
		return arrow.FixedWidthTypes.Boolean
	case "float":
		return arrow.PrimitiveTypes.Float32
	case "double":
		return arrow.PrimitiveTypes.Float64
	case "timestamp":
		return arrow.FixedWidthTypes.Timestamp_us
	case "json":
		return types.ExtensionTypes.JSON
	case "binary", "blob":
		return arrow.BinaryTypes.Binary
	}

	return arrow.BinaryTypes.String
}

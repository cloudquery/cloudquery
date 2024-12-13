package client

import (
	"strconv"
	"strings"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

const defaultPrecision = 10
const defaultScale = 0

func isUnsigned(sqlType string) bool {
	return strings.Contains(sqlType, "unsigned")
}

func parseStringValue(str string, defaultValue int32) int32 {
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
		precision = parseStringValue(parts[0], defaultPrecision)
		scale = defaultScale
	case 2:
		precision = parseStringValue(parts[0], defaultPrecision)
		scale = parseStringValue(parts[1], defaultScale)
	default:
		precision = defaultPrecision
		scale = defaultScale
	}
	return precision, scale
}

func mySQLTypeToArrowType(sqlType string) arrow.DataType {
	if sqlType == "binary(16)" {
		return types.ExtensionTypes.UUID
	}
	if sqlType == "tinyint(1)" {
		return arrow.FixedWidthTypes.Boolean
	}
	if strings.HasPrefix(sqlType, "datetime") {
		return arrow.FixedWidthTypes.Timestamp_us
	}
	if strings.HasPrefix(sqlType, "decimal") || strings.HasPrefix(sqlType, "numeric") {
		precision, scale := getPrecisionAndScale(sqlType)
		return &arrow.Decimal128Type{Precision: precision, Scale: scale}
	}
	if strings.HasPrefix(sqlType, "tinyint") {
		if isUnsigned(sqlType) {
			return arrow.PrimitiveTypes.Uint8
		}
		return arrow.PrimitiveTypes.Int8
	}
	if strings.HasPrefix(sqlType, "smallint") {
		if isUnsigned(sqlType) {
			return arrow.PrimitiveTypes.Uint16
		}
		return arrow.PrimitiveTypes.Int16
	}
	if strings.HasPrefix(sqlType, "int") {
		if isUnsigned(sqlType) {
			return arrow.PrimitiveTypes.Uint32
		}
		return arrow.PrimitiveTypes.Int32
	}
	if strings.HasPrefix(sqlType, "bigint") {
		if isUnsigned(sqlType) {
			return arrow.PrimitiveTypes.Uint64
		}
		return arrow.PrimitiveTypes.Int64
	}
	switch sqlType {
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

func arrowTypeToMySqlStr(t arrow.DataType) string {
	switch t.(type) {
	case *arrow.BooleanType:
		// we can use `bool` which is an alias for `tinyint(1)` but since MySQL information schema returns `tinyint(1)` we use it here as well
		// to be aligned with `mySQLTypeToArrowType`
		return "tinyint(1)"
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
	case *arrow.Float32Type:
		return "float"
	case *arrow.Float64Type:
		return "double"
	case *arrow.StringType, *arrow.LargeStringType:
		return "text"
	case *arrow.BinaryType, *arrow.LargeBinaryType:
		return "blob"
	// nolint:typecheck
	case *types.UUIDType:
		return "binary(16)"
	case *arrow.TimestampType:
		return "datetime(6)"
	case *arrow.StructType, *arrow.ListType, *types.JSONType:
		return "json"
	default:
		return "text"
	}
}

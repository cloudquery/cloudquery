package client

import (
	"strconv"
	"strings"

	"github.com/cloudquery/plugin-sdk/v2/schema"
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

func SQLType(t schema.ValueType) string {
	switch t {
	case schema.TypeBool:
		return "char(1)"
	case schema.TypeInt:
		return "NUMBER(19)"
	case schema.TypeFloat:
		return "binary_double"
	case schema.TypeUUID:
		return "raw(16)"
	case schema.TypeByteArray:
		return "blob"
	case schema.TypeTimestamp:
		return "timestamp"
	case schema.TypeJSON,
		schema.TypeString,
		schema.TypeStringArray,
		schema.TypeUUIDArray,
		schema.TypeCIDRArray,
		schema.TypeMacAddrArray,
		schema.TypeInetArray,
		schema.TypeIntArray,
		schema.TypeCIDR,
		schema.TypeMacAddr,
		schema.TypeInet:
		return "clob"
	default:
		panic("unknown type " + t.String())
	}
}

func SchemaType(tableName string, columnName string, dataType string) schema.ValueType {
	if strings.HasPrefix(columnName, "timestamp") {
		return schema.TypeTimestamp
	}

	if strings.HasPrefix(dataType, "number") {
		_, scale := getPrecisionAndScale(dataType)
		if scale == 0 {
			return schema.TypeInt
		}
		return schema.TypeFloat
	}

	switch dataType {
	case "raw(16)":
		return schema.TypeUUID
	case "char(1)":
		return schema.TypeBool
	case "float", "binary_float", "binary_double":
		return schema.TypeFloat
	case "binary":
		return schema.TypeByteArray
	case "blob", "raw", "long raw":
		return schema.TypeByteArray
	}

	return schema.TypeString
}

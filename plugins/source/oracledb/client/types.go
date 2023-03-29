package client

import (
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
)

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

	switch dataType {
	case "raw(16)":
		return schema.TypeUUID
	case "char(1)":
		return schema.TypeBool
	case "float", "binary_float", "binary_double":
		return schema.TypeFloat
	case "binary":
		return schema.TypeByteArray
	case "number":
		return schema.TypeInt
	case "blob", "raw", "long raw":
		return schema.TypeByteArray
	}

	return schema.TypeString
}

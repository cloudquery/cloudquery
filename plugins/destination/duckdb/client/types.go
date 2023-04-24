package client

import (
	"strings"

	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func (*Client) SchemaTypeToDuckDB(t schema.ValueType) string {
	switch t {
	case schema.TypeBool:
		return "boolean"
	case schema.TypeInt:
		return "bigint"
	case schema.TypeFloat:
		return "real"
	case schema.TypeUUID:
		return "uuid"
	case schema.TypeString:
		return "text"
	case schema.TypeByteArray:
		return "blob"
	case schema.TypeStringArray:
		return "text[]"
	case schema.TypeTimestamp:
		return "timestamp"
	case schema.TypeJSON:
		return "text"
	case schema.TypeUUIDArray:
		return "uuid[]"
	case schema.TypeCIDR:
		return "text"
	case schema.TypeCIDRArray:
		return "text[]"
	case schema.TypeMacAddr:
		return "text"
	case schema.TypeMacAddrArray:
		return "text[]"
	case schema.TypeInet:
		return "text"
	case schema.TypeInetArray:
		return "text[]"
	case schema.TypeIntArray:
		return "int[]"
	default:
		panic("unknown type")
	}
}

func (*Client) duckdbTypeToSchema(t string) schema.ValueType {
	isArray := strings.HasSuffix(t, "[]")
	t = strings.TrimSuffix(t, "[]")
	switch t {
	case "int", "integer", "bigint", "int8", "long":
		if isArray {
			return schema.TypeIntArray
		}
		return schema.TypeInt
	case "ubigint", "uinteger", "usmallint", "utinyint":
		if isArray {
			return schema.TypeIntArray
		}
		return schema.TypeInt
	case "bit", "bitstring":
		if isArray {
			return schema.TypeStringArray
		}
		return schema.TypeString
	case "boolean":
		if isArray {
			panic("unsupported type " + t + "[]")
		}
		return schema.TypeBool
	case "double", "float8", "numeric", "decimal", "real", "float":
		if isArray {
			panic("unsupported type " + t + "[]")
		}
		return schema.TypeFloat
	case "text", "varchar", "char", "bpchar", "string":
		if isArray {
			return schema.TypeStringArray
		}
		return schema.TypeString
	case "blob":
		if isArray {
			panic("unsupported type " + t + "[]")
		}
		return schema.TypeByteArray
	case "uuid":
		if isArray {
			return schema.TypeUUIDArray
		}
		return schema.TypeUUID
	case "date", "timestamp":
		if isArray {
			panic("unsupported type " + t + "[]")
		}
		return schema.TypeTimestamp
	default:
		panic("unknown type: " + t)
	}
}

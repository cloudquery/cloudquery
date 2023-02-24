package client

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func SQLType(t schema.ValueType) string {
	switch t {
	case schema.TypeBool:
		return "bool"
	case schema.TypeInt:
		return "bigint"
	case schema.TypeFloat:
		return "float"
	case schema.TypeUUID:
		return "binary(16)"
	case schema.TypeByteArray:
		return "blob"
	case schema.TypeTimestamp:
		return "timestamp"
	case schema.TypeCIDR,
		schema.TypeMacAddr,
		schema.TypeInet:
		return "nvarchar(255)"
	case schema.TypeString,
		schema.TypeStringArray,
		schema.TypeUUIDArray,
		schema.TypeCIDRArray,
		schema.TypeMacAddrArray,
		schema.TypeInetArray,
		schema.TypeIntArray:
		return "text"
	case schema.TypeJSON:
		return "json"
	default:
		panic("unknown type " + t.String())
	}
}

func SchemaType(t string) schema.ValueType {
	switch t {
	case "bool", "tinyint(1)":
		return schema.TypeBool
	case "bigint":
		return schema.TypeInt
	case "float":
		return schema.TypeFloat
	case "binary(16)":
		return schema.TypeUUID
	case "blob":
		return schema.TypeByteArray
	case "timestamp":
		return schema.TypeTimestamp
	case "nvarchar(255)", "varchar(255)":
		return schema.TypeInet
	case "text":
		return schema.TypeString
	case "json":
		return schema.TypeJSON
	default:
		panic("unknown type " + t)
	}
}

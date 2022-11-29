package client

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func (*Client) SchemaTypeToSnowflake(t schema.ValueType) string {
	switch t {
	case schema.TypeBool:
		return "boolean"
	case schema.TypeInt:
		return "bigint"
	case schema.TypeFloat:
		return "float8"
	case schema.TypeUUID:
		return "text"
	case schema.TypeString:
		return "text"
	case schema.TypeByteArray:
		return "binary"
	case schema.TypeStringArray:
		return "array"
	case schema.TypeTimestamp:
		return "timestamp"
	case schema.TypeJSON:
		return "variant"
	case schema.TypeUUIDArray:
		return "array"
	case schema.TypeCIDR:
		return "text"
	case schema.TypeCIDRArray:
		return "array"
	case schema.TypeMacAddr:
		return "text"
	case schema.TypeMacAddrArray:
		return "array"
	case schema.TypeInet:
		return "text"
	case schema.TypeInetArray:
		return "array"
	case schema.TypeIntArray:
		return "array"
	default:
		panic("unknown type")
	}
}

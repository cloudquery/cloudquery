package client

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func (*Client) SchemaTypeToBigQuery(t schema.ValueType) string {
	switch t {
	case schema.TypeBool:
		return "bool"
	case schema.TypeInt:
		return "int64"
	case schema.TypeFloat:
		return "float64"
	case schema.TypeUUID:
		return "string"
	case schema.TypeString:
		return "string"
	case schema.TypeByteArray:
		return "bytes"
	case schema.TypeStringArray:
		return "array" // what do to here? TODO
	case schema.TypeTimestamp:
		return "timestamp"
	case schema.TypeJSON:
		return "json"
	case schema.TypeUUIDArray:
		return "array"
	case schema.TypeCIDR:
		return "string"
	case schema.TypeCIDRArray:
		return "array"
	case schema.TypeMacAddr:
		return "text"
	case schema.TypeMacAddrArray:
		return "array"
	case schema.TypeInet:
		return "string"
	case schema.TypeInetArray:
		return "array"
	case schema.TypeIntArray:
		return "array"
	default:
		panic("unknown type")
	}
}

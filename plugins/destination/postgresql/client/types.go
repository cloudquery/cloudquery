package client

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func SchemaTypeToPg(t schema.ValueType) string {
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
	case schema.TypeStringArray:
		return "text[]"
	case schema.TypeTimestamp:
		return "timestamp without time zone"
	case schema.TypeJSON:
		return "jsonb"
	case schema.TypeUUIDArray:
		return "uuid[]"
	case schema.TypeInetArray:
		return "inet[]"
	case schema.TypeCIDR:
		return "cidr"
	case schema.TypeCIDRArray:
		return "cidr[]"
	case schema.TypeMacAddr:
		return "macaddr"
	case schema.TypeMacAddrArray:
		return "macaddr[]"
	case schema.TypeInet:
		return "inet"
	case schema.TypeIntArray:
		return "bigint[]"
	default:
		return "text"
	}
}

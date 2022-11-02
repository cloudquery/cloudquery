package client

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) SchemaTypeToPg(t schema.ValueType) string {
	switch t {
	case schema.TypeBool:
		return "integer"
	case schema.TypeInt:
		return "integer"
	case schema.TypeFloat:
		return "real"
	case schema.TypeUUID:
		return "text"
	case schema.TypeString:
		return "text"
	case schema.TypeByteArray:
		return "blob"
	case schema.TypeStringArray:
		return "text"
	case schema.TypeTimestamp:
		return "text"
	case schema.TypeTimeInterval:
		return "text"
	case schema.TypeJSON:
		return "text"
	case schema.TypeUUIDArray:
		return "text"
	case schema.TypeCIDR:
		return "text"
	case schema.TypeCIDRArray:
		return "text"
	case schema.TypeMacAddr:
		return "text"
	case schema.TypeMacAddrArray:
		return "text"
	case schema.TypeInet:
		return "text"
	case schema.TypeInetArray:
		return "text"
	case schema.TypeIntArray:
		return "text"
	default:
		return ""
	}
}

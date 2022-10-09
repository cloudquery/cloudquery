package client

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) SchemaTypeToPg(t schema.ValueType) string {
	switch c.pgType {
	case pgTypeCockroachDB:
		return c.SchemaTypeToCockroach(t)
	default:
		return c.SchemaTypeToPg10(t)
	}
}

func (*Client) SchemaTypeToPg10(t schema.ValueType) string {
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

func (*Client) SchemaTypeToCockroach(t schema.ValueType) string {
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
	case schema.TypeInet:
		return "inet"
	case schema.TypeIntArray:
		return "bigint[]"
	case schema.TypeCIDR:
		return "inet"
	case schema.TypeCIDRArray:
		return "inet[]"
	case schema.TypeMacAddrArray:
		return "text[]"
	default:
		return "text"
	}
}

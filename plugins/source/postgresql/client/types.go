package client

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) PgToSchemaType(t string) schema.ValueType {
	switch c.pgType {
	case pgTypeCockroachDB:
		return c.CockroachToSchemaType(t)
	default:
		return c.Pg10ToSchemaType(t)
	}
}

func (*Client) Pg10ToSchemaType(t string) schema.ValueType {
	switch t {
	case "boolean":
		return schema.TypeBool
	case "bigint", "integer":
		return schema.TypeInt
	case "double precision":
		return schema.TypeFloat
	case "uuid":
		return schema.TypeUUID
	case "text":
		return schema.TypeString
	case "bytea":
		return schema.TypeByteArray
	case "text[]":
		return schema.TypeStringArray
	case "timestamp without time zone":
		return schema.TypeTimestamp
	case "jsonb":
		return schema.TypeJSON
	case "uuid[]":
		return schema.TypeUUIDArray
	case "cidr":
		return schema.TypeCIDR
	case "cidr[]":
		return schema.TypeCIDRArray
	case "macaddr":
		return schema.TypeMacAddr
	case "macaddr[]":
		return schema.TypeMacAddrArray
	case "inet":
		return schema.TypeInet
	case "inet[]":
		return schema.TypeInetArray
	case "bigint[]":
		return schema.TypeIntArray
	default:
		panic("unknown type " + t)
	}
}

func (*Client) CockroachToSchemaType(t string) schema.ValueType {
	switch t {
	case "boolean":
		return schema.TypeBool
	case "bigint":
		return schema.TypeInt
	case "double precision":
		return schema.TypeFloat
	case "uuid":
		return schema.TypeUUID
	case "text":
		return schema.TypeString
	case "bytea":
		return schema.TypeByteArray
	case "text[]":
		return schema.TypeStringArray
	case "timestamp without time zone":
		return schema.TypeTimestamp
	case "jsonb":
		return schema.TypeJSON
	case "uuid[]":
		return schema.TypeUUIDArray
	case "inet":
		return schema.TypeInet
	case "inet[]":
		return schema.TypeInetArray
	case "bigint[]":
		return schema.TypeIntArray
	default:
		panic("unknown type " + t)
	}
}

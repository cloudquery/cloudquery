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
		return "double precision"
	case schema.TypeUUID:
		return "uuid"
	case schema.TypeString:
		return "text"
	case schema.TypeByteArray:
		return "bytea"
	case schema.TypeStringArray:
		return "text[]"
	case schema.TypeTimestamp:
		return "timestamp without time zone"
	case schema.TypeJSON:
		return "jsonb"
	case schema.TypeUUIDArray:
		return "uuid[]"
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
	case schema.TypeInetArray:
		return "inet[]"
	case schema.TypeIntArray:
		return "bigint[]"
	default:
		panic("unknown type " + t.String())
	}
}

func (*Client) SchemaTypeToCockroach(t schema.ValueType) string {
	switch t {
	case schema.TypeBool:
		return "boolean"
	case schema.TypeInt:
		return "bigint"
	case schema.TypeFloat:
		return "double precision"
	case schema.TypeUUID:
		return "uuid"
	case schema.TypeString:
		return "text"
	case schema.TypeByteArray:
		return "bytea"
	case schema.TypeStringArray:
		return "text[]"
	case schema.TypeTimestamp:
		return "timestamp without time zone"
	case schema.TypeJSON:
		return "jsonb"
	case schema.TypeUUIDArray:
		return "uuid[]"
	case schema.TypeCIDR:
		return "inet"
	case schema.TypeCIDRArray:
		return "inet[]"
	case schema.TypeMacAddr:
		return "text"
	case schema.TypeMacAddrArray:
		return "text[]"
	case schema.TypeInet:
		return "inet"
	case schema.TypeInetArray:
		return "inet[]"
	case schema.TypeIntArray:
		return "bigint[]"
	default:
		panic("unknown type " + t.String())
	}
}

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

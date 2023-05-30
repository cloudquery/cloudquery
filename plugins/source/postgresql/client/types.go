package client

import (
	"strings"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v2/schema"
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

func (c *Client) PgToSchemaType(t string) arrow.DataType {
	switch c.pgType {
	case pgTypeCockroachDB:
		return c.CockroachToSchemaType(t)
	default:
		return c.Pg10ToSchemaType(t)
	}
}

func (*Client) Pg10ToSchemaType(t string) arrow.DataType {
	if strings.HasPrefix(t, "timestamp") {
		return schema.TypeTimestamp
	}

	switch t {
	case "boolean":
		return schema.TypeBool
	case "bigint", "integer", "bigserial", "smallint", "smallserial", "serial":
		return schema.TypeInt
	case "double precision", "float", "real", "numeric":
		return schema.TypeFloat
	case "uuid":
		return schema.TypeUUID
	case "bytea":
		return schema.TypeByteArray
	case "text[]":
		return schema.TypeStringArray
	case "json", "jsonb":
		return schema.TypeJSON
	case "uuid[]":
		return schema.TypeUUIDArray
	case "cidr":
		return schema.TypeCIDR
	case "cidr[]":
		return schema.TypeCIDRArray
	case "macaddr", "macaddr8":
		return schema.TypeMacAddr
	case "macaddr[]", "macaddr8[]":
		return schema.TypeMacAddrArray
	case "inet":
		return schema.TypeInet
	case "inet[]":
		return schema.TypeInetArray
	case "bigint[]", "integer[]", "smallint[]", "bigserial[]", "smallserial[]", "serial[]":
		return schema.TypeIntArray
	default:
		return schema.TypeString
	}
}

func (*Client) CockroachToSchemaType(t string) arrow.DataType {
	if strings.HasPrefix(t, "timestamp") {
		return schema.TypeTimestamp
	}

	switch t {
	case "boolean":
		return schema.TypeBool
	case "bigint", "int", "oid", "serial":
		return schema.TypeInt
	case "decimal", "float":
		return schema.TypeFloat
	case "uuid":
		return schema.TypeUUID
	case "bytea":
		return schema.TypeByteArray
	case "text[]":
		return schema.TypeStringArray
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
		return schema.TypeString
	}
}

package client

import (
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
	"golang.org/x/exp/maps"
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

func (c *Client) PgToSchemaType(tableName string, columnName string, t string) (schema.ValueType, error) {
	switch c.pgType {
	case pgTypeCockroachDB:
		return c.CockroachToSchemaType(tableName, columnName, t)
	default:
		return c.Pg10ToSchemaType(tableName, columnName, t)
	}
}

func (*Client) Pg10ToSchemaType(tableName string, columnName string, postgresType string) (schema.ValueType, error) {
	if strings.HasPrefix(postgresType, "timestamp") {
		return schema.TypeTimestamp, nil
	}

	pgToSchemaType := map[string]schema.ValueType{
		"boolean":          schema.TypeBool,
		"bigint":           schema.TypeInt,
		"integer":          schema.TypeInt,
		"float":            schema.TypeFloat,
		"real":             schema.TypeFloat,
		"double precision": schema.TypeFloat,
		"uuid":             schema.TypeUUID,
		"text":             schema.TypeString,
		"bytea":            schema.TypeByteArray,
		"text[]":           schema.TypeStringArray,
		"jsonb":            schema.TypeJSON,
		"uuid[]":           schema.TypeUUIDArray,
		"cidr":             schema.TypeCIDR,
		"cidr[]":           schema.TypeCIDRArray,
		"macaddr":          schema.TypeMacAddr,
		"macaddr[]":        schema.TypeMacAddrArray,
		"inet":             schema.TypeInet,
		"inet[]":           schema.TypeInetArray,
		"bigint[]":         schema.TypeIntArray,
	}

	if v, ok := pgToSchemaType[postgresType]; ok {
		return v, nil
	}

	return schema.TypeInvalid, fmt.Errorf("got unknown PostgreSQL type %q for column %q of table %q while trying to convert it to CloudQuery internal schema type. Supported PostgreSQL types are %q", postgresType, columnName, tableName, append(maps.Keys(pgToSchemaType), "timestamp"))
}

func (*Client) CockroachToSchemaType(tableName string, columnName string, cockroachType string) (schema.ValueType, error) {
	if strings.HasPrefix(cockroachType, "timestamp") {
		return schema.TypeTimestamp, nil
	}

	cockroachToSchemaType := map[string]schema.ValueType{
		"boolean":          schema.TypeBool,
		"bigint":           schema.TypeInt,
		"float":            schema.TypeFloat,
		"real":             schema.TypeFloat,
		"double precision": schema.TypeFloat,
		"uuid":             schema.TypeUUID,
		"text":             schema.TypeString,
		"bytea":            schema.TypeByteArray,
		"text[]":           schema.TypeStringArray,
		"jsonb":            schema.TypeJSON,
		"uuid[]":           schema.TypeUUIDArray,
		"inet":             schema.TypeInet,
		"inet[]":           schema.TypeInetArray,
		"bigint[]":         schema.TypeIntArray,
	}

	if v, ok := cockroachToSchemaType[cockroachType]; ok {
		return v, nil
	}

	return schema.TypeInvalid, fmt.Errorf("got unknown CockroachDB type %q for column %q of table %q while trying to convert it to CloudQuery internal schema type. Supported CockroachDB types are %q", cockroachType, columnName, tableName, append(maps.Keys(cockroachToSchemaType), "timestamp"))
}

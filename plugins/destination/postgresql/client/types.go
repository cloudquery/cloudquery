package client

import (
	"fmt"
	"strings"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/types"
)

func (c *Client) SchemaTypeToPg(t arrow.DataType) string {
	switch c.pgType {
	case pgTypeCockroachDB:
		return c.SchemaTypeToCockroach(t)
	default:
		return c.SchemaTypeToPg10(t)
	}
}

func (c *Client) SchemaTypeToPg10(t arrow.DataType) string {
	switch v := t.(type) {
	case *arrow.ListType:
		return c.SchemaTypeToPg10(v.Elem()) + "[]"
	case *arrow.FixedSizeListType:
		return c.SchemaTypeToPg10(v.Elem()) + fmt.Sprintf("[%d]", v.Len())
	case *arrow.BooleanType:
		return "boolean"
	case *arrow.Int8Type, *arrow.Uint8Type:
		return "smallint"
	case *arrow.Int16Type, *arrow.Uint16Type:
		return "smallint"
	case *arrow.Int32Type, *arrow.Uint32Type:
		return "integer"
	case *arrow.Int64Type, *arrow.Uint64Type:
		return "bigint"
	case *arrow.Float32Type:
		return "real"
	case *arrow.Float64Type:
		return "double precision"
	case *arrow.StringType, *arrow.LargeStringType:
		return "text"
	case *arrow.BinaryType, *arrow.LargeBinaryType:
		return "bytea"
	case *types.UUIDType:
		return "uuid"
	case *arrow.TimestampType:
		return "timestamp without time zone"
	case *types.JSONType:
		return "jsonb"
	case *arrow.StructType:
		return "jsonb"
	case *types.InetType:
		return "inet"
	case *types.MacType:
		return "macaddr"
	default:
		return "text"
	}
}

func (c *Client) SchemaTypeToCockroach(t arrow.DataType) string {
	switch v := t.(type) {
	case *arrow.ListType:
		return c.SchemaTypeToCockroach(v.Elem()) + "[]"
	case *arrow.FixedSizeListType:
		return c.SchemaTypeToCockroach(v.Elem()) + fmt.Sprintf("[%d]", v.Len())
	case *arrow.BooleanType:
		return "boolean"
	case *arrow.Int8Type, *arrow.Uint8Type:
		return "smallint"
	case *arrow.Int16Type, *arrow.Uint16Type:
		return "smallint"
	case *arrow.Int32Type, *arrow.Uint32Type:
		return "integer"
	case *arrow.Int64Type, *arrow.Uint64Type:
		return "bigint"
	case *arrow.Float32Type:
		return "real"
	case *arrow.Float64Type:
		return "double precision"
	case *arrow.StringType, *arrow.LargeStringType:
		return "text"
	case *arrow.BinaryType, *arrow.LargeBinaryType:
		return "bytea"
	case *types.UUIDType:
		return "uuid"
	case *arrow.TimestampType:
		return "timestamp without time zone"
	case *types.JSONType:
		return "jsonb"
	case *arrow.StructType:
		return "jsonb"
	case *types.InetType:
		return "inet"
	default:
		return "text"
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

func (c *Client) Pg10ToSchemaType(t string) arrow.DataType {
	if strings.HasPrefix(t, "timestamp") {
		return arrow.FixedWidthTypes.Timestamp_us
	}
	if strings.HasSuffix(t, "[]") {
		return arrow.ListOf(c.Pg10ToSchemaType(strings.TrimSuffix(t, "[]")))
	}

	switch t {
	case "boolean":
		return arrow.FixedWidthTypes.Boolean
	case "smallint":
		return arrow.PrimitiveTypes.Int16
	case "integer":
		return arrow.PrimitiveTypes.Int32
	case "bigint":
		return arrow.PrimitiveTypes.Int64
	case "real":
		return arrow.PrimitiveTypes.Float32
	case "double precision":
		return arrow.PrimitiveTypes.Float64
	case "uuid":
		return types.ExtensionTypes.UUID
	case "bytea":
		return arrow.BinaryTypes.Binary
	case "json", "jsonb":
		return types.ExtensionTypes.JSON
	case "cidr":
		return types.ExtensionTypes.Inet
	case "macaddr", "macaddr8":
		return types.ExtensionTypes.Mac
	case "inet":
		return types.ExtensionTypes.Inet
	default:
		return arrow.BinaryTypes.String
	}
}

func (*Client) CockroachToSchemaType(t string) arrow.DataType {
	if strings.HasPrefix(t, "timestamp") {
		return arrow.FixedWidthTypes.Timestamp_us
	}

	switch t {
	case "boolean":
		return arrow.FixedWidthTypes.Boolean
	case "bigint", "int", "oid", "serial":
		return arrow.PrimitiveTypes.Int64
	case "decimal", "float":
		return arrow.PrimitiveTypes.Float64
	case "uuid":
		return types.ExtensionTypes.UUID
	case "bytea":
		return arrow.BinaryTypes.Binary
	case "text[]":
		return arrow.ListOf(arrow.BinaryTypes.String)
	case "jsonb":
		return types.ExtensionTypes.UUID
	case "uuid[]":
		return arrow.ListOf(types.ExtensionTypes.UUID)
	case "inet":
		return types.ExtensionTypes.Inet
	case "inet[]":
		return arrow.ListOf(types.ExtensionTypes.Inet)
	case "bigint[]":
		return arrow.ListOf(arrow.PrimitiveTypes.Int64)
	default:
		return arrow.BinaryTypes.String
	}
}

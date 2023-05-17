package client

import (
	"fmt"
	"strings"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/types"
)

func (c *Client) SchemaTypeToPg(t arrow.DataType) string {
	switch c.pgType {
	case pgTypeCockroachDB:
		return c.SchemaTypeToCockroach(t)
	default:
		return c.SchemaTypeToPg10(t)
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
	case *arrow.Int16Type:
		return "smallint"
	case *arrow.Uint16Type, *arrow.Int32Type:
		return "integer"
	case *arrow.Uint32Type, *arrow.Int64Type:
		return "bigint"
	case *arrow.Uint64Type:
		return "numeric"
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
	case *types.MACType:
		return "macaddr"
	default:
		return "text"
	}
}

func (c *Client) Pg10ToSchemaType(t string) arrow.DataType {
	if strings.HasPrefix(t, "timestamp") {
		// we use microsecond per postgres documentation:
		// https://www.postgresql.org/docs/current/datatype-datetime.html
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
	case "numeric":
		return arrow.PrimitiveTypes.Uint64
	case "real":
		return arrow.PrimitiveTypes.Float32
	case "double precision":
		return arrow.PrimitiveTypes.Float64
	case "uuid":
		return types.ExtensionTypes.UUID
	case "bytea":
		return arrow.BinaryTypes.LargeBinary
	case "json", "jsonb":
		return types.ExtensionTypes.JSON
	case "cidr":
		return types.ExtensionTypes.Inet
	case "macaddr", "macaddr8":
		return types.ExtensionTypes.MAC
	case "inet":
		return types.ExtensionTypes.Inet
	default:
		return arrow.BinaryTypes.LargeString
	}
}

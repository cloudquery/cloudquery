package client

import (
	"fmt"
	"strings"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/types"
)

func (c *Client) SchemaTypeToCockroach(t arrow.DataType) string {
	switch v := t.(type) {
	case *arrow.ListType:
		return c.SchemaTypeToCockroach(v.Elem()) + "[]"
	case *arrow.FixedSizeListType:
		return c.SchemaTypeToCockroach(v.Elem()) + fmt.Sprintf("[%d]", v.Len())
	case *arrow.BooleanType:
		return "boolean"
	case *arrow.Int8Type, *arrow.Uint8Type, *arrow.Int16Type, *arrow.Uint16Type:
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

func (c *Client) CockroachToSchemaType(t string) arrow.DataType {
	if strings.HasPrefix(t, "timestamp") {
		return arrow.FixedWidthTypes.Timestamp_us
	}
	if strings.HasSuffix(t, "[]") {
		return arrow.ListOf(c.CockroachToSchemaType(strings.TrimSuffix(t, "[]")))
	}

	switch t {
	case "boolean":
		return arrow.FixedWidthTypes.Boolean
	case "bigint", "int", "oid", "serial", "integer":
		return arrow.PrimitiveTypes.Int64
	case "decimal", "float", "real", "double precision":
		return arrow.PrimitiveTypes.Float64
	case "uuid":
		return types.ExtensionTypes.UUID
	case "bytea":
		return arrow.BinaryTypes.LargeBinary
	case "jsonb", "json":
		return types.ExtensionTypes.JSON
	case "inet":
		return types.ExtensionTypes.Inet
	default:
		return arrow.BinaryTypes.LargeString
	}
}

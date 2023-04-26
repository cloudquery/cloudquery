package client

import (
	"strings"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/types"
)

func (c *Client) SchemaTypeToDuckDB(t arrow.DataType) string {
	switch v := t.(type) {
	case *arrow.ListType:
		return c.SchemaTypeToDuckDB(v.Elem()) + "[]"
	case *arrow.BooleanType:
		return "boolean"
	case *arrow.Int8Type:
		return "tinyint"
	case *arrow.Int16Type:
		return "smallint"
	case *arrow.Int32Type:
		return "int"
	case *arrow.Int64Type:
		return "bigint"
	case *arrow.Uint8Type:
		return "utinyint"
	case *arrow.Uint16Type:
		return "usmallint"
	case *arrow.Uint32Type:
		return "uint"
	case *arrow.Uint64Type:
		return "ubigint"
	case *arrow.Float32Type:
		return "float"
	case *arrow.Float64Type:
		return "double"
	case *arrow.BinaryType:
		return "blob"
	case *arrow.LargeBinaryType:
		return "blob"
	case *types.UUIDType:
		return "uuid"
	case *types.JSONType:
		return "json"
	case *arrow.TimestampType:
		return "timestamp"
	case *arrow.Date32Type:
		return "date"
	case *arrow.DayTimeIntervalType:
		return "interval"
	default:
		return "varchar"
	}
}

func (c *Client) duckdbTypeToSchema(t string) arrow.DataType {
	if strings.HasSuffix(t, "[]") {
		return arrow.ListOf(c.duckdbTypeToSchema(strings.TrimSuffix(t, "[]")))
	}
	switch t {
	case "tinyint", "int1":
		return arrow.PrimitiveTypes.Int8
	case "smallint":
		return arrow.PrimitiveTypes.Int16
	case "int", "int4", "integer", "signed":
		return arrow.PrimitiveTypes.Int32
	case "bigint", "int8", "long":
		return arrow.PrimitiveTypes.Int64
	case "utinyint":
		return arrow.PrimitiveTypes.Uint8
	case "usmallint":
		return arrow.PrimitiveTypes.Uint16
	case "uint", "uint4", "uinteger":
		return arrow.PrimitiveTypes.Uint32
	case "ubigint":
		return arrow.PrimitiveTypes.Uint64
	case "boolean", "bool", "logical":
		return arrow.FixedWidthTypes.Boolean
	case "double", "float8", "numeric", "decimal":
		return arrow.PrimitiveTypes.Float64
	case "float", "float4", "real":
		return arrow.PrimitiveTypes.Float32
	case "blob", "bytea", "binary", "varbinary":
		return arrow.BinaryTypes.Binary
	case "date":
		return arrow.FixedWidthTypes.Date32
	case "timestamp", "datetime", "timestamp with time zone", "timestamptz":
		return arrow.FixedWidthTypes.Timestamp_us
	case "interval":
		return arrow.FixedWidthTypes.DayTimeInterval
	case "json":
		return types.ExtensionTypes.JSON
	case "uuid":
		return types.ExtensionTypes.UUID
	default:
		return arrow.BinaryTypes.String
	}
}

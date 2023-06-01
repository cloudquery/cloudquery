package client

import (
	"strings"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/types"
)

type listLike interface {
	arrow.DataType
	Elem() arrow.DataType
}

func plainWrite(dt arrow.DataType) bool {
	return arrow.TypeEqual(dt, transformTypeForWriting(dt))
}

func transformTypeForWriting(dt arrow.DataType) arrow.DataType {
	switch dt := dt.(type) {
	case listLike:
		return arrow.ListOf(transformTypeForWriting(dt.Elem()))
	case *arrow.MapType:
		return arrow.ListOf(transformTypeForWriting(dt.ValueType()))
	}

	switch dt := transformType(dt).(type) {
	case *types.UUIDType, *types.JSONType:
		return arrow.BinaryTypes.String
	default:
		return dt
	}
}

func transformSchemaForWriting(sc *arrow.Schema) *arrow.Schema {
	fields := sc.Fields()
	for i := range fields {
		fields[i].Type = transformTypeForWriting(fields[i].Type)
	}
	md := sc.Metadata()
	return arrow.NewSchema(fields, &md)
}

func transformType(dt arrow.DataType) arrow.DataType {
	return duckDBToArrow(arrowToDuckDB(dt))
}

func arrowToDuckDB(t arrow.DataType) string {
	switch v := t.(type) {
	case *arrow.ListType:
		return arrowToDuckDB(v.Elem()) + "[]"
	case *arrow.BooleanType:
		return "boolean"
	case *arrow.Int8Type:
		return "tinyint"
	case *arrow.Int16Type:
		return "smallint"
	case *arrow.Int32Type:
		return "integer"
	case *arrow.Int64Type:
		return "bigint"
	case *arrow.Uint8Type:
		return "uinteger"
	case *arrow.Uint16Type:
		return "uinteger"
	case *arrow.Uint32Type:
		return "uinteger"
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
	case *arrow.Date32Type, *arrow.Date64Type:
		return "date"
	case *arrow.DayTimeIntervalType:
		return "interval"
	case *arrow.StructType:
		return "json"
	default:
		return "varchar"
	}
}

func duckDBToArrow(t string) arrow.DataType {
	if strings.HasSuffix(t, "[]") {
		return arrow.ListOf(duckDBToArrow(strings.TrimSuffix(t, "[]")))
	}
	if strings.HasPrefix(t, "struct") {
		return types.ExtensionTypes.JSON
	}
	switch t {
	case "tinyint", "int1":
		return arrow.PrimitiveTypes.Int8
	case "smallint", "int2", "short":
		return arrow.PrimitiveTypes.Int16
	case "integer", "int4", "signed", "int":
		return arrow.PrimitiveTypes.Int32
	case "bigint", "int8", "long":
		return arrow.PrimitiveTypes.Int64
	case "utinyint":
		return arrow.PrimitiveTypes.Uint8
	case "usmallint":
		return arrow.PrimitiveTypes.Uint16
	case "uinteger", "uint4":
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
		return arrow.FixedWidthTypes.Date64
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

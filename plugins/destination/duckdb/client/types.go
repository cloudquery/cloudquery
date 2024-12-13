package client

import (
	"slices"
	"strings"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

func transformSchemaForWriting(sc *arrow.Schema) *arrow.Schema {
	md := arrow.MetadataFrom(sc.Metadata().ToMap())
	return arrow.NewSchema(transformFieldsForWriting(sc.Fields()), &md)
}

func transformFieldsForWriting(fields []arrow.Field) []arrow.Field {
	for i := range fields {
		fields[i].Type = transformTypeForWriting(fields[i].Type)
	}
	return fields
}

func transformTypeForWriting(dt arrow.DataType) arrow.DataType {
	switch dt := dt.(type) {
	case *arrow.StructType, *arrow.MapType:
		return arrow.BinaryTypes.String
	case arrow.ListLikeType:
		return arrow.ListOf(transformTypeForWriting(dt.Elem()))
	case *types.UUIDType, *types.JSONType:
		return arrow.BinaryTypes.String
	default:
		return duckDBToArrow(arrowToDuckDB(dt))
	}
}

func arrowToDuckDB(dt arrow.DataType) string {
	switch dt := dt.(type) {
	case *arrow.StructType, *arrow.MapType:
		return "json"
	case arrow.ListLikeType:
		return arrowToDuckDB(dt.Elem()) + "[]"
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
	case *arrow.Date32Type, *arrow.Date64Type, *arrow.TimestampType:
		return "timestamp"
	default:
		return "varchar"
	}
}

func duckDBToArrow(t string) arrow.DataType {
	switch {
	case strings.HasSuffix(t, "[]"):
		return arrow.ListOf(duckDBToArrow(strings.TrimSuffix(t, "[]")))
	case strings.HasPrefix(t, "struct"):
		return duckDBStructToArrow(t)
	case strings.HasPrefix(t, "map"):
		return duckDBMapToArrow(t)
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
	case "timestamp", "datetime", "timestamp with time zone", "timestamptz":
		return arrow.FixedWidthTypes.Timestamp_us
	case "json":
		return types.ExtensionTypes.JSON
	case "uuid":
		return types.ExtensionTypes.UUID
	default:
		return arrow.BinaryTypes.String
	}
}

func duckDBStructToArrow(spec string) *arrow.StructType {
	params := strings.TrimPrefix(spec, "struct")
	params = strings.TrimSpace(params)
	params = strings.TrimSuffix(strings.TrimPrefix(params, "("), ")")

	fieldsSpec := splitParams(params)
	if len(fieldsSpec) == 0 {
		panic("unsupported struct spec: " + spec)
	}

	fields := make([]arrow.Field, len(fieldsSpec))
	for i, fieldSpec := range fieldsSpec {
		parts := strings.SplitN(fieldSpec, " ", 2)
		if len(parts) != 2 {
			panic("unsupported field spec: " + fieldSpec)
		}

		fields[i] = arrow.Field{
			Name:     strings.Trim(parts[0], `"`),
			Type:     duckDBToArrow(strings.TrimSpace(parts[1])),
			Nullable: true, // all duckdb columns are nullable
		}
	}

	return arrow.StructOf(fields...)
}

func duckDBMapToArrow(spec string) *arrow.MapType {
	params := strings.TrimPrefix(spec, "map")
	params = strings.TrimSpace(params)
	params = strings.TrimSuffix(strings.TrimPrefix(params, "("), ")")

	kv := splitParams(params)
	if len(kv) != 2 {
		panic("unsupported map spec: " + spec)
	}

	// these should only be types
	return arrow.MapOf(duckDBToArrow(kv[0]), duckDBToArrow(kv[1]))
}

func splitParams(params string) []string {
	params = strings.TrimSpace(params)

	var brackets int
	var parts []string
	elem := make([]rune, 0, len(params))

	for _, r := range params {
		switch r {
		case '(':
			brackets++
		case ')':
			brackets--
		case ',':
			if brackets == 0 {
				parts = append(parts, strings.TrimSpace(string(elem)))
				elem = elem[:0] // cleanup
				continue
			}
		}
		elem = append(elem, r)
	}
	parts = append(parts, strings.TrimSpace(string(elem)))

	return slices.Clip(parts)
}

func sanitizeID(id string) string {
	return `"` + id + `"`
}

func sanitized(ids []string) []string {
	res := make([]string, len(ids))
	for i, id := range ids {
		res[i] = sanitizeID(id)
	}
	return res
}

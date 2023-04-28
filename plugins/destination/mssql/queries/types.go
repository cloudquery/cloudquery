package queries

import (
	"fmt"
	"reflect"
	"time"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/types"
	"golang.org/x/exp/maps"
)

func SQLType(_type arrow.DataType) string {
	switch _type.(type) {
	case *arrow.BooleanType:
		return "bit"

	case *arrow.Uint8Type:
		return "tinyint" // uint8
	case *arrow.Int8Type, *arrow.Uint16Type, *arrow.Int16Type:
		return "smallint" // int16
	case *arrow.Uint32Type, *arrow.Int32Type:
		return "int" // int32
	case *arrow.Uint64Type, *arrow.Int64Type:
		return "bigint" // int64

	case *arrow.Float32Type:
		return "real"
	case *arrow.Float64Type:
		return "float" // == float(53)

	case *arrow.LargeStringType:
		return "nvarchar(max)" // we will also use it as the default type

	case *arrow.StringType, *types.InetType, *types.MacType:
		return "nvarchar(4000)" // feasible to see these as PK, so need to limit the value

	case arrow.BinaryDataType, *arrow.FixedSizeBinaryType:
		return "varbinary(max)"

	case *types.UUIDType:
		return "uniqueidentifier"

	case *arrow.TimestampType:
		return "datetime2"

	default:
		return "nvarchar(max)"
	}
}

func SchemaType(tableName, columnName, sqlType string) (arrow.DataType, error) {
	sqlToSchema := map[string]arrow.DataType{
		"bit":              new(arrow.BooleanType),
		"tinyint":          new(arrow.Uint8Type),
		"smallint":         new(arrow.Int16Type),
		"int":              new(arrow.Int32Type),
		"bigint":           new(arrow.Int64Type),
		"real":             new(arrow.Float32Type),
		"float":            new(arrow.Float64Type),
		"uniqueidentifier": types.NewUUIDType(),
		"varbinary(max)":   new(arrow.LargeBinaryType),
		"datetime2":        &arrow.TimestampType{Unit: arrow.Nanosecond}, // the precision is 100ns in MSSQL
		"nvarchar(4000)":   new(arrow.StringType),
		"nvarchar(max)":    new(arrow.LargeStringType),
	}

	if v, ok := sqlToSchema[sqlType]; ok {
		return v, nil
	}
	return nil, fmt.Errorf("got unknown MSSQL type %q of column %q for table %q while trying to convert it to Apache Arrow type. Supported MSSQL types are %q", sqlType, columnName, tableName, maps.Keys(sqlToSchema))
}

// columnGoType has to be in sync with SQLType
func columnGoType(_type arrow.DataType) reflect.Type {
	switch _type.(type) {
	case *arrow.BooleanType:
		return reflect.TypeOf(true)

	case *arrow.Uint8Type:
		return reflect.TypeOf(uint8(0))
	case *arrow.Int8Type, *arrow.Uint16Type, *arrow.Int16Type:
		return reflect.TypeOf(int16(0))
	case *arrow.Uint32Type, *arrow.Int32Type:
		return reflect.TypeOf(int32(0))
	case *arrow.Uint64Type, *arrow.Int64Type:
		return reflect.TypeOf(int64(0))

	case *arrow.Float32Type:
		return reflect.TypeOf(float32(0))
	case *arrow.Float64Type:
		return reflect.TypeOf(float64(0))

	case *arrow.LargeStringType:
		return reflect.TypeOf("")

	case *arrow.StringType, *types.InetType, *types.MacType:
		return reflect.TypeOf("")

	case arrow.BinaryDataType, *arrow.FixedSizeBinaryType:
		return reflect.TypeOf([]byte{})

	case *types.UUIDType:
		return reflect.TypeOf([]byte{})

	case *arrow.TimestampType:
		return reflect.TypeOf(time.Time{})

	default:
		return reflect.TypeOf("")
	}
}

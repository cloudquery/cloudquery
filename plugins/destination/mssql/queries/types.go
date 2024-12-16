package queries

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

const DefaultMaxLengthNvarchar = "4000"

func SQLType(dataType arrow.DataType, pk bool) string {
	switch dataType := dataType.(type) {
	case *arrow.BooleanType:
		return "bit"

	case *arrow.Uint8Type:
		return "tinyint" // uint8
	case *arrow.Int8Type, *arrow.Int16Type: // no special int8 type, upscale
		return "smallint" // int16
	case *arrow.Uint16Type, *arrow.Int32Type: // no special uint16 type, upscale
		return "int" // int32
	case *arrow.Uint32Type, *arrow.Int64Type: // no special uint32 type, upscale
		return "bigint" // int64
	case *arrow.Uint64Type: // we store this as int64, although it may produce overflow and negative numbers
		return "bigint" // int64

	case *arrow.Float32Type:
		return "real"
	case *arrow.Float64Type:
		return "float" // == float(53)

	case *arrow.StringType, *arrow.LargeStringType, *types.InetType, *types.MACType:
		if pk {
			return "nvarchar(" + DefaultMaxLengthNvarchar + ")"
		}
		return "nvarchar(max)" // we will also use it as the default type

	case *arrow.FixedSizeBinaryType:
		return "varbinary(" + strconv.Itoa(dataType.ByteWidth) + ")"

	case arrow.BinaryDataType:
		return "varbinary(max)"

	case *types.UUIDType:
		return "uniqueidentifier"

	case *arrow.TimestampType:
		return "datetime2"

	default:
		return "nvarchar(max)"
	}
}

func SchemaType(sqlType string) arrow.DataType {
	// this is for the types without precision
	simpleSQLToSchema := map[string]arrow.DataType{
		"bit":              arrow.FixedWidthTypes.Boolean,
		"tinyint":          arrow.PrimitiveTypes.Uint8,
		"smallint":         arrow.PrimitiveTypes.Int16,
		"int":              arrow.PrimitiveTypes.Int32,
		"bigint":           arrow.PrimitiveTypes.Int64,
		"real":             arrow.PrimitiveTypes.Float32,
		"float":            arrow.PrimitiveTypes.Float64,
		"uniqueidentifier": types.NewUUIDType(),
		"datetime2":        arrow.FixedWidthTypes.Timestamp_ns, // the precision is 100ns in MSSQL
		"datetimeoffset":   arrow.FixedWidthTypes.Timestamp_ns, // the precision is 100ns in MSSQL
	}

	if dt, ok := simpleSQLToSchema[sqlType]; ok {
		return dt
	}

	// 2 types left to check: nvarchar & varbinary
	colType, precision := sqlType, ""
	if parts := strings.SplitN(sqlType, "(", 2); len(parts) == 2 {
		colType, precision = parts[0], strings.TrimSuffix(parts[1], ")")
	}
	switch colType {
	case "nvarchar":
		if precision == "max" {
			return new(arrow.LargeStringType)
		}

		// we just return the arrow.String here
		return new(arrow.StringType)
	case "varbinary":
		if precision == "max" {
			return new(arrow.LargeBinaryType)
		}

		width, err := strconv.Atoi(precision)
		if err != nil {
			// should never happen
			panic(fmt.Errorf("failed to parse %q into int: %w", precision, err))
		}

		return &arrow.FixedSizeBinaryType{ByteWidth: width}
	}

	// default to LargeString (nvarchar(max))
	return new(arrow.LargeStringType)
}

// columnGoType has to be in sync with SQLType
func columnGoType(dataType arrow.DataType) reflect.Type {
	switch dataType.(type) {
	case *arrow.BooleanType:
		return reflect.TypeOf(true)

	case *arrow.Uint8Type:
		return reflect.TypeOf(uint8(0))
	case *arrow.Int8Type, *arrow.Int16Type: // no special int8 type, upscale
		return reflect.TypeOf(int16(0))
	case *arrow.Uint16Type, *arrow.Int32Type: // no special uint16 type, upscale
		return reflect.TypeOf(int32(0))
	case *arrow.Uint32Type, *arrow.Int64Type: // no special uint32 type, upscale
		return reflect.TypeOf(int64(0))
	case *arrow.Uint64Type: // we store this as int64, although it may produce overflow and negative numbers
		return reflect.TypeOf(int64(0))

	case *arrow.Float32Type:
		return reflect.TypeOf(float32(0))
	case *arrow.Float64Type:
		return reflect.TypeOf(float64(0))

	case *arrow.LargeStringType, *arrow.StringType, *types.InetType, *types.MACType:
		return reflect.TypeOf("")

	case arrow.BinaryDataType, *arrow.FixedSizeBinaryType, *types.UUIDType:
		return reflect.TypeOf([]byte{})

	case *arrow.TimestampType:
		return reflect.TypeOf(time.Time{})

	default:
		return reflect.TypeOf("")
	}
}

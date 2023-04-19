package client

import (
	"fmt"
	"sort"
	"strings"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/types"
	"golang.org/x/exp/maps"
)

func mySQLTypeToArrowType(tableName string, columnName string, sqlType string) (arrow.DataType, error) {
	if strings.HasPrefix(sqlType, "datetime") {
		// MySQL permits up to microseconds (6 digits) precision
		return arrow.FixedWidthTypes.Timestamp_us, nil
	}
	sqlTypeToSchemaType := map[string]arrow.DataType{
		"bool":              arrow.FixedWidthTypes.Boolean,
		"tinyint(1)":        arrow.FixedWidthTypes.Boolean,
		"tinyint":           arrow.PrimitiveTypes.Int8,
		"smallint":          arrow.PrimitiveTypes.Int16,
		"int":               arrow.PrimitiveTypes.Int32,
		"bigint":            arrow.PrimitiveTypes.Int64,
		"bigint(20)":        arrow.PrimitiveTypes.Int64,
		"tinyint unsigned":  arrow.PrimitiveTypes.Uint8,
		"smallint unsigned": arrow.PrimitiveTypes.Uint16,
		"int unsigned":      arrow.PrimitiveTypes.Uint32,
		"bigint unsigned":   arrow.PrimitiveTypes.Uint64,
		"float":             arrow.PrimitiveTypes.Float32,
		"double":            arrow.PrimitiveTypes.Float64,
		"binary(16)":        types.ExtensionTypes.UUID,
		"blob":              arrow.BinaryTypes.LargeBinary,
		"nvarchar(255)":     types.ExtensionTypes.Inet,
		"varchar(255)":      types.ExtensionTypes.Inet,
		"text":              arrow.BinaryTypes.LargeString,
		"json":              types.ExtensionTypes.JSON,
	}

	if v, ok := sqlTypeToSchemaType[sqlType]; ok {
		return v, nil
	}

	supportedTypes := maps.Keys(sqlTypeToSchemaType)
	supportedTypes = append(supportedTypes, "datetime")
	sort.Strings(supportedTypes)
	return nil, fmt.Errorf("got unknown MySQL type %q for column %q of table %q while trying to convert it to CloudQuery internal schema type. Supported MySQL types are %q", sqlType, columnName, tableName, supportedTypes)
}

func arrowTypeToMySqlStr(t arrow.DataType) string {
	switch t.(type) {
	case *arrow.BooleanType:
		return "bool"
	case *arrow.Int8Type:
		return "tinyint"
	case *arrow.Int16Type:
		return "smallint"
	case *arrow.Int32Type:
		return "int"
	case *arrow.Int64Type:
		return "bigint"
	case *arrow.Uint8Type:
		return "tinyint unsigned"
	case *arrow.Uint16Type:
		return "smallint unsigned"
	case *arrow.Uint32Type:
		return "int unsigned"
	case *arrow.Uint64Type:
		return "bigint unsigned"
	case *arrow.Float32Type:
		return "float"
	case *arrow.Float64Type:
		return "double"
	case *arrow.StringType, *arrow.LargeStringType:
		return "text"
	case *arrow.BinaryType, *arrow.LargeBinaryType:
		return "blob"
	case *types.UUIDType:
		return "binary(16)"
	case *arrow.TimestampType:
		return "datetime(6)"
	case *types.JSONType:
		return "json"
	case *arrow.StructType:
		return "json"
	case *types.InetType, *types.MacType:
		return "nvarchar(255)"
	default:
		return "text"
	}
}

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
		"bool":          arrow.FixedWidthTypes.Boolean,
		"tinyint(1)":    arrow.FixedWidthTypes.Boolean,
		"bigint":        arrow.PrimitiveTypes.Int64,
		"bigint(20)":    arrow.PrimitiveTypes.Int64,
		"float":         arrow.PrimitiveTypes.Float64,
		"binary(16)":    types.ExtensionTypes.UUID,
		"blob":          arrow.BinaryTypes.LargeBinary,
		"nvarchar(255)": types.ExtensionTypes.Inet,
		"varchar(255)":  types.ExtensionTypes.Inet,
		"text":          arrow.BinaryTypes.LargeString,
		"json":          types.ExtensionTypes.JSON,
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
	case *arrow.Int8Type, *arrow.Uint8Type, *arrow.Int16Type, *arrow.Uint16Type, *arrow.Int32Type, *arrow.Uint32Type, *arrow.Int64Type, *arrow.Uint64Type:
		return "bigint"
	case *arrow.Float16Type, *arrow.Float32Type, *arrow.Float64Type:
		return "float"
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

package ch

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/util"
	"github.com/cloudquery/plugin-sdk/v2/types"
)

func extensionType(extension arrow.ExtensionType) string {
	switch extension.(type) {
	// https://clickhouse.com/docs/en/sql-reference/data-types/uuid
	case *types.UUIDType:
		return "UUID"

	// https://clickhouse.com/docs/en/sql-reference/data-types/string
	case *types.InetType, *types.MacType:
		return "String"
	case *types.JSONType:
		// ClickHouse can't handle values like [{"x":{"y":"z"}}] at the moment.
		// https://github.com/ClickHouse/ClickHouse/issues/46190
		return "String"

	default:
		return "String"
	}
}

func arrayType(list arrow.DataType) (string, error) {
	var field arrow.Field
	switch list := list.(type) {
	case *arrow.ListType:
		field = list.ElemField()
	case *arrow.LargeListType:
		field = list.ElemField()
	case *arrow.FixedSizeListType:
		field = list.ElemField()
	default:
		return "", fmt.Errorf("unsupported Apache Arow list type: %s", list)
	}

	elem, err := fieldType(field) // adds Nullable
	if err != nil {
		return "", err
	}

	return "Array(" + elem + ")", nil
}

func structType(_struct *arrow.StructType) (string, error) {
	// https://clickhouse.com/docs/en/sql-reference/data-types/tuple
	fields, err := Definitions(_struct.Fields()...)
	if err != nil {
		return "", err
	}

	return "Tuple(" + strings.Join(fields, ", ") + ")", nil
}

func mapType(_map *arrow.MapType) (string, error) {
	// https://clickhouse.com/docs/en/sql-reference/data-types/map
	// Keys can only be: String, Integer, LowCardinality, FixedString, UUID, Date, DateTime, Date32, Enum.
	keyType, err := fieldType(_map.KeyField())
	if err != nil {
		return "", err
	}
	switch {
	case keyType == "String",
		keyType == "UUID",
		keyType == "Date",
		keyType == "Date32",
		keyType == "DateTime":
	case strings.HasPrefix(keyType, "Int"):
	case strings.HasPrefix(keyType, "UInt"):
	case strings.HasPrefix(keyType, "LowCardinality"):
	case strings.HasPrefix(keyType, "FixedString"):
	case strings.HasPrefix(keyType, "Enum"):
	default:
		return "", fmt.Errorf("unsupported Apache Arraw type for ClickHouse map key: %s", keyType)
	}

	itemType, err := fieldType(_map.ItemField()) // adds Nullable, too
	if err != nil {
		return "", err
	}

	return "Map(" + keyType + ", " + itemType + ")", nil
}

func columnType(dataType arrow.DataType) (string, error) {
	switch dataType.ID() {
	// https://clickhouse.com/docs/en/sql-reference/data-types/boolean
	case arrow.BOOL:
		return "Bool", nil

	// https://clickhouse.com/docs/en/sql-reference/data-types/int-uint
	case arrow.UINT8:
		return "UInt8", nil
	case arrow.UINT16:
		return "UInt16", nil
	case arrow.UINT32:
		return "UInt32", nil
	case arrow.UINT64:
		return "UInt64", nil
	case arrow.INT8:
		return "Int8", nil
	case arrow.INT16:
		return "Int16", nil
	case arrow.INT32:
		return "Int32", nil
	case arrow.INT64:
		return "Int64", nil

	// https://clickhouse.com/docs/en/sql-reference/data-types/float
	case arrow.FLOAT16, arrow.FLOAT32:
		return "Float32", nil
	case arrow.FLOAT64:
		return "Float64", nil

	// https://clickhouse.com/docs/en/sql-reference/data-types/string
	case arrow.STRING, arrow.BINARY, arrow.LARGE_STRING, arrow.LARGE_BINARY:
		return "String", nil

	// https://clickhouse.com/docs/en/sql-reference/data-types/fixedstring
	case arrow.FIXED_SIZE_BINARY:
		return "FixedString(" + strconv.Itoa(dataType.(*arrow.FixedSizeBinaryType).ByteWidth) + ")", nil

	// https://clickhouse.com/docs/en/sql-reference/data-types/date32
	case arrow.DATE32:
		return "Date32", nil

	// https://clickhouse.com/docs/en/sql-reference/data-types/datetime64
	case arrow.DATE64:
		return "DateTime64(3)", nil // 3 = milliseconds
	case arrow.TIMESTAMP:
		switch unit := dataType.(*arrow.TimestampType).Unit; unit {
		case arrow.Second:
			return "DateTime64(0)", nil // 0 = seconds
		case arrow.Millisecond:
			return "DateTime64(3)", nil // 3 = milliseconds
		case arrow.Microsecond:
			return "DateTime64(6)", nil // 3 = milliseconds
		case arrow.Nanosecond:
			return "DateTime64(9)", nil // 3 = milliseconds
		default:
			return "", fmt.Errorf("unsupported Apache Arrow Timestamp resolution: %s", unit)
		}

	// https://clickhouse.com/docs/en/sql-reference/data-types/decimal
	case arrow.DECIMAL128:
		return "Decimal128(" + strconv.FormatInt(int64(dataType.(*arrow.Decimal128Type).Scale), 10) + ")", nil
	case arrow.DECIMAL256:
		return "Decimal256(" + strconv.FormatInt(int64(dataType.(*arrow.Decimal256Type).Scale), 10) + ")", nil

	// https://clickhouse.com/docs/en/sql-reference/data-types/array
	case arrow.LIST, arrow.LARGE_LIST, arrow.FIXED_SIZE_LIST:
		return arrayType(dataType)

	// https://clickhouse.com/docs/en/sql-reference/data-types/tuple
	case arrow.STRUCT:
		return structType(dataType.(*arrow.StructType))

	// https://clickhouse.com/docs/en/sql-reference/data-types/map
	case arrow.MAP:
		return mapType(dataType.(*arrow.MapType))

	// Only support CQ extensions for now
	case arrow.EXTENSION:
		return extensionType(dataType.(arrow.ExtensionType)), nil

	// everything else that's not supported ATM
	default:
		return "String", nil
	}
}

func fieldType(field arrow.Field) (string, error) {
	typ, err := columnType(field.Type)
	if err != nil {
		return "", err
	}

	// We allow nullable values in arrays, but arrays cannot be nullable themselves
	if field.Type.ID() == arrow.LIST || !field.Nullable {
		return typ, nil
	}

	return "Nullable(" + typ + ")", nil
}

func fieldDefinition(field arrow.Field) (string, error) {
	fieldType, err := fieldType(field)
	if err != nil {
		return "", err
	}
	return util.SanitizeID(field.Name) + " " + fieldType, err
}

func Definitions(fields ...arrow.Field) ([]string, error) {
	res := make([]string, len(fields))
	for i, field := range fields {
		fieldDef, err := fieldDefinition(field)
		if err != nil {
			return nil, err
		}
		res[i] = fieldDef
	}
	return res, nil
}

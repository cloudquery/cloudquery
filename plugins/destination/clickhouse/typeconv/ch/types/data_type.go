package types

import (
	"strconv"

	"github.com/apache/arrow/go/v12/arrow"
)

func dataType(_type arrow.DataType) (string, error) {
	switch _type.ID() {
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
		return "FixedString(" + strconv.Itoa(_type.(*arrow.FixedSizeBinaryType).ByteWidth) + ")", nil

	// https://clickhouse.com/docs/en/sql-reference/data-types/date32
	case arrow.DATE32:
		return "Date32", nil

	// https://clickhouse.com/docs/en/sql-reference/data-types/datetime64
	case arrow.DATE64:
		return "DateTime64(3)", nil // 3 = milliseconds
	case arrow.TIMESTAMP:
		return timestampType(_type.(*arrow.TimestampType))

	// https://clickhouse.com/docs/en/sql-reference/data-types/decimal
	case arrow.DECIMAL128, arrow.DECIMAL256:
		return decimalType(_type.(arrow.DecimalType))

	// https://clickhouse.com/docs/en/sql-reference/data-types/array
	case arrow.LIST, arrow.LARGE_LIST, arrow.FIXED_SIZE_LIST:
		return listType(_type.(listDataType))

	// https://clickhouse.com/docs/en/sql-reference/data-types/tuple
	case arrow.STRUCT:
		return structType(_type.(*arrow.StructType))

	// Only support CQ extensions for now
	case arrow.EXTENSION:
		return extensionType(_type.(arrow.ExtensionType)), nil

	case arrow.MAP:
		// TODO: support https://clickhouse.com/docs/en/sql-reference/data-types/map
		return "String", nil

	// everything else that's not supported ATM
	default:
		return "String", nil
	}
}

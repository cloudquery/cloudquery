package definitions

import (
	"strconv"

	"github.com/apache/arrow/go/v12/arrow"
)

func dataType(dataType arrow.DataType) string {
	switch dataType.ID() {
	// https://clickhouse.com/docs/en/sql-reference/data-types/boolean
	case arrow.BOOL:
		return "Bool"

	// https://clickhouse.com/docs/en/sql-reference/data-types/int-uint
	case arrow.UINT8:
		return "UInt8"
	case arrow.UINT16:
		return "UInt16"
	case arrow.UINT32:
		return "UInt32"
	case arrow.UINT64:
		return "UInt64"
	case arrow.INT8:
		return "Int8"
	case arrow.INT16:
		return "Int16"
	case arrow.INT32:
		return "Int32"
	case arrow.INT64:
		return "Int64"

	// https://clickhouse.com/docs/en/sql-reference/data-types/float
	case arrow.FLOAT16, arrow.FLOAT32:
		return "Float32"
	case arrow.FLOAT64:
		return "Float64"

	// https://clickhouse.com/docs/en/sql-reference/data-types/string
	case arrow.STRING, arrow.BINARY, arrow.LARGE_STRING, arrow.LARGE_BINARY:
		return "String"

	// https://clickhouse.com/docs/en/sql-reference/data-types/fixedstring
	case arrow.FIXED_SIZE_BINARY:
		return "FixedString(" + strconv.Itoa(dataType.(*arrow.FixedSizeBinaryType).ByteWidth) + ")"

	// https://clickhouse.com/docs/en/sql-reference/data-types/date32
	case arrow.DATE32:
		return "Date32"

	// https://clickhouse.com/docs/en/sql-reference/data-types/datetime64
	case arrow.DATE64:
		return "DateTime64(3)" // 3 = milliseconds
	case arrow.TIMESTAMP:
		switch unit := dataType.(*arrow.TimestampType).Unit; unit {
		case arrow.Second:
			return "DateTime64(0)" // 0 = seconds
		case arrow.Millisecond:
			return "DateTime64(3)" // 3 = milliseconds
		case arrow.Microsecond:
			return "DateTime64(6)" // 3 = milliseconds
		default:
			return "DateTime64(9)" // 3 = milliseconds
		}

	// https://clickhouse.com/docs/en/sql-reference/data-types/decimal
	case arrow.DECIMAL128:
		const (
			minPrecision = 19
			maxPrecision = 38
		)
		decimal := dataType.(*arrow.Decimal128Type)
		return decimalType(decimal.Precision, decimal.Scale, minPrecision, maxPrecision)
	case arrow.DECIMAL256:
		const (
			minPrecision = 39
			maxPrecision = 76
		)
		decimal := dataType.(*arrow.Decimal256Type)
		return decimalType(decimal.Precision, decimal.Scale, minPrecision, maxPrecision)

	// https://clickhouse.com/docs/en/sql-reference/data-types/array
	case arrow.LIST, arrow.LARGE_LIST, arrow.FIXED_SIZE_LIST:
		return listType(dataType)

	// https://clickhouse.com/docs/en/sql-reference/data-types/tuple
	case arrow.STRUCT:
		return structType(dataType.(*arrow.StructType))

	// Only support CQ extensions for now
	case arrow.EXTENSION:
		return extensionType(dataType.(arrow.ExtensionType))

	case arrow.MAP:
		// TODO: support https://clickhouse.com/docs/en/sql-reference/data-types/map
		return "String"

	// everything else that's not supported ATM
	default:
		return "String"
	}
}

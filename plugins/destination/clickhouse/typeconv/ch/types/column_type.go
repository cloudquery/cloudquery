package types

import (
	"strconv"

	"github.com/apache/arrow-go/v18/arrow"
)

func ColumnType(dataType arrow.DataType) (string, error) {
	switch dataType := dataType.(type) {
	// https://clickhouse.com/docs/en/sql-reference/data-types/boolean
	case *arrow.BooleanType:
		return "Bool", nil

	// https://clickhouse.com/docs/en/sql-reference/data-types/int-uint
	case *arrow.Uint8Type:
		return "UInt8", nil
	case *arrow.Uint16Type:
		return "UInt16", nil
	case *arrow.Uint32Type:
		return "UInt32", nil
	case *arrow.Uint64Type:
		return "UInt64", nil
	case *arrow.Int8Type:
		return "Int8", nil
	case *arrow.Int16Type:
		return "Int16", nil
	case *arrow.Int32Type:
		return "Int32", nil
	case *arrow.Int64Type:
		return "Int64", nil

	// https://clickhouse.com/docs/en/sql-reference/data-types/float
	case *arrow.Float16Type, *arrow.Float32Type:
		return "Float32", nil
	case *arrow.Float64Type:
		return "Float64", nil

	// https://clickhouse.com/docs/en/sql-reference/data-types/string
	case arrow.BinaryDataType:
		return "String", nil

	// https://clickhouse.com/docs/en/sql-reference/data-types/fixedstring
	case *arrow.FixedSizeBinaryType:
		return "FixedString(" + strconv.Itoa(dataType.ByteWidth) + ")", nil

	// https://clickhouse.com/docs/en/sql-reference/data-types/date32
	case *arrow.Date32Type:
		return "Date32", nil

	// https://clickhouse.com/docs/en/sql-reference/data-types/datetime
	// Although `date64` claims millisecond precision, having it not on the date border is UB.
	case *arrow.Date64Type:
		return "DateTime", nil

	// https://clickhouse.com/docs/en/sql-reference/data-types/datetime64
	case *arrow.Time32Type:
		return timeType(dataType.Unit, nil)
	case *arrow.Time64Type:
		return timeType(dataType.Unit, nil)
	case *arrow.TimestampType:
		return timestampType(dataType)

	// https://clickhouse.com/docs/en/sql-reference/data-types/decimal
	case arrow.DecimalType:
		return decimalType(dataType)

	// https://clickhouse.com/docs/en/sql-reference/data-types/map
	case *arrow.MapType:
		return mapType(dataType)

	// https://clickhouse.com/docs/en/sql-reference/data-types/array
	case arrow.ListLikeType:
		return listType(dataType)

	// https://clickhouse.com/docs/en/sql-reference/data-types/tuple
	case *arrow.StructType:
		return structType(dataType)

	// Only support CQ extensions for now
	case arrow.ExtensionType:
		return extensionType(dataType), nil

	// everything else that's not supported ATM
	default:
		return "String", nil
	}
}

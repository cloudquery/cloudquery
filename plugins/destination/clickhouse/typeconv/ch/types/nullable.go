package types

import (
	"github.com/apache/arrow/go/v12/arrow"
)

// https://clickhouse.com/docs/en/sql-reference/data-types/nullable
// Array & Tuple can't be Nullable
func canBeNullable(dataType arrow.DataType) bool {
	switch dataType.(type) {
	case *arrow.MapType:
		return false
	case listDataType:
		return false
	case *arrow.StructType:
		return false
	default:
		return true
	}
}

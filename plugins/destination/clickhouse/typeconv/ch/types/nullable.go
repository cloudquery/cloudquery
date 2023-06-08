package types

import (
	"github.com/apache/arrow/go/v13/arrow"
)

// https://clickhouse.com/docs/en/sql-reference/data-types/nullable
// Array, Map & Tuple can't be Nullable
func canBeNullable(dataType arrow.DataType) bool {
	switch dataType.(type) {
	case arrow.ListLikeType, // matches all lists (large, fixed size, normal) + maps
		*arrow.StructType:
		return false
	default:
		return true
	}
}

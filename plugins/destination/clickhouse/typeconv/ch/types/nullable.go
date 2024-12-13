package types

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// CanBeNullable is used to check if data type can be wrapped in `Nullable`
// https://clickhouse.com/docs/en/sql-reference/data-types/nullable
// Array, Map & Tuple can't be Nullable
func CanBeNullable(dataType arrow.DataType) bool {
	switch dataType.(type) {
	case arrow.ListLikeType, // matches all lists (large, fixed size, normal) + maps
		*arrow.StructType:
		return false
	default:
		return true
	}
}

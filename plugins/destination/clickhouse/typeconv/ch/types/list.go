package types

import (
	"github.com/apache/arrow-go/v18/arrow"
)

// https://clickhouse.com/docs/en/sql-reference/data-types/array
func listType(listType arrow.ListLikeType) (string, error) {
	item, err := FieldType(listType.ElemField())
	if err != nil {
		return "", err
	}

	return "Array(" + item + ")", nil // adds Nullable to the item, if required
}

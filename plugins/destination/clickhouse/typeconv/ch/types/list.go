package types

import (
	"github.com/apache/arrow/go/v13/arrow"
)

type listDataType interface {
	arrow.DataType
	ElemField() arrow.Field
}

// https://clickhouse.com/docs/en/sql-reference/data-types/array
func listType(listType listDataType) (string, error) {
	item, err := FieldType(listType.ElemField())
	if err != nil {
		return "", err
	}

	return "Array(" + item + ")", nil // adds Nullable to the item, if required
}

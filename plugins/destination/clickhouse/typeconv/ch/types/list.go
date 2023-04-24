package types

import (
	"github.com/apache/arrow/go/v12/arrow"
)

type listDataType interface {
	arrow.DataType
	ElemField() arrow.Field
}

// https://clickhouse.com/docs/en/sql-reference/data-types/array
func listType(_type listDataType) string {
	return "Array(" + FieldType(_type.ElemField()) + ")" // adds Nullable to the item, if required
}

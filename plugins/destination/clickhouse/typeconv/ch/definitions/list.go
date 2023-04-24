package definitions

import (
	"github.com/apache/arrow/go/v12/arrow"
)

type listDataType interface {
	ElemField() arrow.Field
}

func listType(list arrow.DataType) string {
	return "Array(" + FieldType(list.(listDataType).ElemField()) + ")" // adds Nullable
}

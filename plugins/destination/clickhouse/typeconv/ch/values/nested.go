package values

import (
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
)

// sanitizeNested will replace all null entries with empty ones as in CH nested types aren't nullable themselves
// https://clickhouse.com/docs/en/sql-reference/data-types/nullable
// array passed should be of arrow.NestedType
func sanitizeNested(arr arrow.Array) arrow.Array {
	if arr.NullN() == 0 {
		return arr
	}

	if _, ok := arr.DataType().(arrow.NestedType); !ok {
		// This can happen only if the parent is nested & we need to construct the empty value
		return array.MakeFromData(array.NewData(
			arr.DataType(), arr.Len(),
			append([]*memory.Buffer{nil}, arr.Data().Buffers()[1:]...), // the first elem is validity
			arr.Data().Children(),
			0, arr.Data().Offset(),
		))
	}

	children := make([]arrow.ArrayData, len(arr.Data().Children()))
	for i, child := range arr.Data().Children() {
		children[i] = sanitizeNested(array.MakeFromData(child)).Data()
	}

	return array.MakeFromData(array.NewData(
		arr.DataType(), arr.Len(),
		append([]*memory.Buffer{nil}, arr.Data().Buffers()[1:]...), // the first elem is validity
		children,
		0, arr.Data().Offset(),
	))
}

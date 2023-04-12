package client

import (
	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
)

func timestampValues(arr *array.Timestamp) []any {
	unit := arr.DataType().(*arrow.TimestampType).Unit
	res := make([]any, arr.Len())
	for i := range res {
		if arr.IsValid(i) && !arr.IsNull(i) {
			res[i] = arr.Value(i).ToTime(unit).UTC()
		}
	}
	return res
}

func getValues(arr arrow.Array) []any {
	if tsArr, ok := arr.(*array.Timestamp); ok {
		return timestampValues(tsArr)
	}

	res := make([]any, arr.Len())
	for i := range res {
		if arr.IsValid(i) && !arr.IsNull(i) {
			res[i] = arr.GetOneForMarshal(i)
		}
	}
	return res
}

func transpose(m map[string][]any, l int) []map[string]any {
	rows := make([]map[string]any, l)
	for i := range rows {
		rows[i] = make(map[string]any)
	}

	for i, row := range rows {
		for k, v := range m {
			row[k] = v[i]
		}
	}

	return rows
}

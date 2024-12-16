package values

import (
	"fmt"
	"reflect"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
)

func structValue(arr *array.Struct) (any, error) {
	arr = sanitizeNested(arr).(*array.Struct)

	fields := arr.DataType().(*arrow.StructType).Fields()
	columns := make(map[string][]any, len(fields))
	for i, field := range fields {
		data, err := FromArray(arr.Field(i))
		if err != nil {
			return nil, err
		}

		columns[field.Name], err = toSlice(data)
		if err != nil {
			return nil, err
		}
	}

	rows := make([]map[string]any, arr.Len())
	for i := 0; i < arr.Len(); i++ {
		row := make(map[string]any, len(fields))
		for _, field := range fields {
			row[field.Name] = columns[field.Name][i]
		}
		rows[i] = row
	}

	return rows, nil
}

func toSlice(data any) ([]any, error) {
	val := reflect.ValueOf(data)
	if val.Kind() != reflect.Slice {
		return nil, fmt.Errorf("expected slice, got %T", data)
	}

	res := make([]any, val.Len())
	for i := 0; i < len(res); i++ {
		elem := val.Index(i)
		if elem.CanInterface() {
			res[i] = elem.Interface()
		}
	}
	return res, nil
}

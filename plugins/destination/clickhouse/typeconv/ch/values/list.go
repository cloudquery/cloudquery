package values

import (
	"reflect"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2/lib/column"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/apache/arrow/go/v13/arrow/memory"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/typeconv/ch/types"
)

func listValue(arr array.ListLike) (any, error) {
	colType, err := types.ColumnType(arr.DataType())
	if err != nil {
		return nil, err
	}
	// Need to create slice of the proper type.
	// We could infer in from elements, but sometimes array is empty
	col, err := column.Type(colType).Column("tmp", time.UTC)
	if err != nil {
		return nil, err
	}
	valueType := col.ScanType()

	sanitized, err := sanitizeNested(arr)
	if err != nil {
		return nil, err
	}
	arr = sanitized.(array.ListLike)
	elems := make([]any, arr.Len())
	for i := 0; i < arr.Len(); i++ {
		from, to := arr.ValueOffsets(i)
		elems[i], err = FromArray(array.NewSlice(arr.ListValues(), from, to))
		if err != nil {
			return nil, err
		}
	}

	res := reflect.MakeSlice(reflect.SliceOf(reflect.PointerTo(valueType)), len(elems), len(elems)) // we do []*(type) for nullable assignment
	for i, elem := range elems {
		// lists aren't nullable themselves
		// https://clickhouse.com/docs/en/sql-reference/data-types/nullable
		val := reflect.New(valueType)
		val.Elem().Set(reflect.ValueOf(elem))
		res.Index(i).Set(val)
	}

	return res.Interface(), nil
}

// sanitizeNested will replace all null entries with empty ones as in CH nested types aren't nullable themselves
// https://clickhouse.com/docs/en/sql-reference/data-types/nullable
func sanitizeNested(arr arrow.Array) (arrow.Array, error) {
	if arr.NullN() == 0 {
		return arr, nil
	}

	builder := array.NewBuilder(memory.DefaultAllocator, arr.DataType())
	for i := 0; i < arr.Len(); i++ {
		if arr.IsNull(i) {
			builder.AppendEmptyValue()
			continue
		}
		if err := builder.AppendValueFromString(arr.ValueStr(i)); err != nil {
			return nil, err
		}
	}
	return builder.NewArray(), nil
}

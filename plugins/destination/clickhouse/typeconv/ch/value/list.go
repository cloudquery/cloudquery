package value

import (
	"reflect"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2/lib/column"
	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/typeconv/ch/definitions"
)

func listValue(arr array.ListLike) (any, error) {
	var err error
	elems := make([]any, arr.Len())
	for i := 0; i < arr.Len(); i++ {
		if arr.IsNull(i) || !arr.IsValid(i) {
			continue
		}

		from, to := arr.ValueOffsets(i)
		elems[i], err = FromArray(array.NewSlice(arr.ListValues(), from, to))
		if err != nil {
			return nil, err
		}
	}

	// Need to create slice of the proper type.
	// We could infer in from elements, but sometimes array is empty
	colType, err := column.Type(definitions.FieldType(arrow.Field{Type: arr.DataType()})).Column("tmp", time.UTC)
	if err != nil {
		return nil, err
	}
	_type := colType.ScanType()

	res := reflect.MakeSlice(reflect.SliceOf(reflect.PointerTo(_type)), len(elems), len(elems)) // we do []*(type) for nullable assignment
	for i, elem := range elems {
		// we need to fill in for the in-depth recursive parsing by ClickHouse SDK
		val := reflect.New(_type)
		if elem != nil {
			val.Elem().Set(reflect.ValueOf(elem))
		}
		res.Index(i).Set(val)
	}

	return res.Interface(), nil
}

type listWrapper struct {
	*array.FixedSizeList
}

var _ array.ListLike = listWrapper{}

func (l listWrapper) ValueOffsets(i int) (start, end int64) {
	n := int64(l.DataType().(*arrow.FixedSizeListType).Len())
	off := int64(l.Offset())
	return (off + int64(i)) * n, (off + int64(i+1)) * n
}

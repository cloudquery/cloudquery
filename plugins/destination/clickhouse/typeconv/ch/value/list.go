package value

import (
	"reflect"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
)

func listValue(arr array.ListLike) (any, error) {
	if arr.Len() == 0 {
		return nil, nil
	}

	elems := make([]any, arr.Len())
	var _type reflect.Type
	for i := 0; i < arr.Len(); i++ {
		if arr.IsNull(i) || !arr.IsValid(i) {
			continue
		}
		from, to := arr.ValueOffsets(i)
		elem, err := FromArray(array.NewSlice(arr.ListValues(), from, to))
		if err != nil {
			return nil, err
		}
		elems[i] = elem
		if _type == nil && elem != nil {
			_type = reflect.TypeOf(elem)
		}
	}

	if _type == nil {
		return elems, nil
	}

	res := reflect.MakeSlice(reflect.SliceOf(reflect.PointerTo(_type)), len(elems), len(elems)) // we do []*(type) for nullable assignment
	for i, elem := range elems {
		if elem == nil {
			continue
		}
		val := reflect.New(_type)
		val.Elem().Set(reflect.ValueOf(elem))
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

package value

import (
	"reflect"

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
		if _type == nil {
			_type = reflect.PointerTo(reflect.TypeOf(elem)) // we do []*(type) for nullable assignment
		}
	}

	if _type == nil {
		// all elements are nil, so just return nil (NB: not fully equivalent, but highly unlikely)
		return nil, nil
	}

	res := reflect.MakeSlice(reflect.SliceOf(_type), len(elems), len(elems))
	for i, elem := range elems {
		res.Index(i).Set(reflect.Indirect(reflect.ValueOf(elem))) // we do []*(type) for nullable assignment
	}

	return res.Interface(), nil
}

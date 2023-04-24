package values

import (
	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
)

func marshalValue[A any](arr arrow.Array) []*A {
	res := make([]*A, arr.Len())
	for i := 0; i < arr.Len(); i++ {
		if arr.IsValid(i) && !arr.IsNull(i) {
			val := arr.GetOneForMarshal(i).(A)
			res[i] = &val
		}
	}
	return res
}

func valueStrData(arr arrow.Array) []*string {
	res := make([]*string, arr.Len())
	for i := 0; i < arr.Len(); i++ {
		if arr.IsValid(i) && !arr.IsNull(i) {
			val := arr.ValueStr(i)
			res[i] = &val
		}
	}
	return res
}

func float16Value(arr *array.Float16) []*float32 {
	res := make([]*float32, arr.Len())
	for i := 0; i < arr.Len(); i++ {
		if arr.IsValid(i) && !arr.IsNull(i) {
			val := arr.Value(i).Float32()
			res[i] = &val
		}
	}
	return res
}

func byteArrValue(arr primitive[[]byte]) [][]byte {
	res := make([][]byte, arr.Len())
	for i := 0; i < arr.Len(); i++ {
		if arr.IsValid(i) && !arr.IsNull(i) {
			res[i] = arr.Value(i)
		}
	}
	return res
}

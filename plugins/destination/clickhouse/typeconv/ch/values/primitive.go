package values

import (
	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
)

type primitive[A any] interface {
	arrow.Array
	Value(int) A
}

func primitiveValue[A any](arr primitive[A]) []*A {
	res := make([]*A, arr.Len())
	for i := 0; i < arr.Len(); i++ {
		if arr.IsValid(i) && !arr.IsNull(i) {
			val := arr.Value(i)
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

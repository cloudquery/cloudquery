package values

import (
	"github.com/apache/arrow/go/v15/arrow"
)

type primitive[A any] interface {
	arrow.Array
	Value(int) A
}

func primitiveValue[A any, ARR primitive[A]](arr ARR) []*A {
	res := make([]*A, arr.Len())
	for i := 0; i < arr.Len(); i++ {
		if arr.IsValid(i) {
			val := arr.Value(i)
			res[i] = &val
		}
	}
	return res
}

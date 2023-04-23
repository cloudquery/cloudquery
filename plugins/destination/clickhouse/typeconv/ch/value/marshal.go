package value

import (
	"github.com/apache/arrow/go/v12/arrow"
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

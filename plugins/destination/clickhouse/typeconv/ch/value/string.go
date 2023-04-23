package value

import (
	"github.com/apache/arrow/go/v12/arrow"
)

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

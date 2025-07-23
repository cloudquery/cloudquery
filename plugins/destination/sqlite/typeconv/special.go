package typeconv

import (
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
)

func valueStrData(arr arrow.Array) []string {
	res := make([]string, arr.Len())
	for i := 0; i < arr.Len(); i++ {
		if arr.IsValid(i) {
			res[i] = arr.ValueStr(i)
		}
	}
	return res
}

func float16Value(arr *array.Float16) []float32 {
	res := make([]float32, arr.Len())
	for i := 0; i < arr.Len(); i++ {
		if arr.IsValid(i) {
			res[i] = arr.Value(i).Float32()
		}
	}
	return res
}

func byteArrValue(arr primitive[[]byte]) []string {
	res := make([]string, arr.Len())
	for i := 0; i < arr.Len(); i++ {
		if arr.IsValid(i) {
			res[i] = string(arr.Value(i))
		}
	}
	return res
}

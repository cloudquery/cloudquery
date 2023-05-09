package values

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/goccy/go-json"
)

func ptr[A any](a A) *A { return &a }

func marshalValuesToStrings(arr arrow.Array) ([]*string, error) {
	res := make([]*string, arr.Len())
	for i := 0; i < arr.Len(); i++ {
		if arr.IsNull(i) {
			continue
		}
		data, err := json.MarshalWithOption(arr.GetOneForMarshal(i), json.DisableHTMLEscape())
		if err != nil {
			return nil, err
		}
		res[i] = ptr(string(data))
	}
	return res, nil
}

func valueStrData(arr arrow.Array) []*string {
	res := make([]*string, arr.Len())
	for i := 0; i < arr.Len(); i++ {
		if arr.IsValid(i) {
			val := arr.ValueStr(i)
			res[i] = &val
		}
	}
	return res
}

func float16Value(arr *array.Float16) []*float32 {
	res := make([]*float32, arr.Len())
	for i := 0; i < arr.Len(); i++ {
		if arr.IsValid(i) {
			val := arr.Value(i).Float32()
			res[i] = &val
		}
	}
	return res
}

func byteArrValue(arr primitive[[]byte]) [][]byte {
	res := make([][]byte, arr.Len())
	for i := 0; i < arr.Len(); i++ {
		if arr.IsValid(i) {
			res[i] = arr.Value(i)
		}
	}
	return res
}

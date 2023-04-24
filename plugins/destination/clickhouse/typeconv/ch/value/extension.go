package value

import (
	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
	"github.com/cloudquery/plugin-sdk/v2/types"
	"github.com/google/uuid"
)

func extensionValue(arr array.ExtensionArray) any {
	switch arr := arr.(type) {
	case *types.UUIDArray:
		return marshalValue[uuid.UUID](arr)
	case *types.InetArray, *types.MacArray, *types.JSONArray:
		return valueStrData(arr)
	default:
		// we fallback here to string representation
		return valueStrData(arr)
	}
}

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

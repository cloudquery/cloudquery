package typeconv

import (
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
)

func FromArray(arr arrow.Array) (any, error) {
	switch arr := arr.(type) {
	case *array.Boolean:
		return primitiveValue(arr), nil

	case *array.Uint8:
		return primitiveValue(arr), nil
	case *array.Uint16:
		return primitiveValue(arr), nil
	case *array.Uint32:
		return primitiveValue(arr), nil
	case *array.Uint64:
		return primitiveValue(arr), nil

	case *array.Int8:
		return primitiveValue(arr), nil
	case *array.Int16:
		return primitiveValue(arr), nil
	case *array.Int32:
		return primitiveValue(arr), nil
	case *array.Int64:
		return primitiveValue(arr), nil

	case *array.Float16:
		return float16Value(arr), nil
	case *array.Float32:
		return primitiveValue(arr), nil
	case *array.Float64:
		return primitiveValue(arr), nil

	case *array.String:
		return primitiveValue(arr), nil

	case *array.Binary:
		return byteArrValue(arr), nil
	case *array.FixedSizeBinary:
		return byteArrValue(arr), nil
	case *array.LargeBinary:
		return byteArrValue(arr), nil

	default:
		return valueStrData(arr), nil
	}
}

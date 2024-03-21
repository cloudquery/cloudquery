package values

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
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

	case *array.Date32:
		return dateValue(arr), nil
	case *array.Date64:
		return dateValue(arr), nil

	case *array.Time32:
		return timeValue(arr, arr.DataType().(*arrow.Time32Type).Unit), nil
	case *array.Time64:
		return timeValue(arr, arr.DataType().(*arrow.Time64Type).Unit), nil
	case *array.Timestamp:
		return timestampValue(arr)

	case *array.Decimal128:
		return decimalValue(arr), nil
	case *array.Decimal256:
		return decimalValue(arr), nil

	case array.ExtensionArray:
		return extensionValue(arr), nil

	case *array.Struct:
		return structValue(arr)

	case *array.Map:
		// it also matches array.ListLike, so we check it before the array.ListLike case
		return mapValue(arr)

	case array.ListLike:
		return listValue(arr)

	default:
		return valueStrData(arr), nil
	}
}

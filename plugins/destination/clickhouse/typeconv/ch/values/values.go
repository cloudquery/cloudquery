package values

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/apache/arrow/go/v13/arrow/decimal128"
	"github.com/apache/arrow/go/v13/arrow/decimal256"
)

func FromArray(arr arrow.Array) (any, error) {
	switch arr := arr.(type) {
	case *array.Boolean:
		return primitiveValue[bool](arr), nil

	case *array.Uint8:
		return primitiveValue[uint8](arr), nil
	case *array.Uint16:
		return primitiveValue[uint16](arr), nil
	case *array.Uint32:
		return primitiveValue[uint32](arr), nil
	case *array.Uint64:
		return primitiveValue[uint64](arr), nil

	case *array.Int8:
		return primitiveValue[int8](arr), nil
	case *array.Int16:
		return primitiveValue[int16](arr), nil
	case *array.Int32:
		return primitiveValue[int32](arr), nil
	case *array.Int64:
		return primitiveValue[int64](arr), nil

	case *array.Float16:
		return float16Value(arr), nil
	case *array.Float32:
		return primitiveValue[float32](arr), nil
	case *array.Float64:
		return primitiveValue[float64](arr), nil

	case *array.String:
		return primitiveValue[string](arr), nil

	case *array.Binary:
		return byteArrValue(arr), nil
	case *array.FixedSizeBinary:
		return byteArrValue(arr), nil
	case *array.LargeBinary:
		return byteArrValue(arr), nil

	case *array.Date32:
		return dateValue[arrow.Date32](arr), nil
	case *array.Date64:
		return dateValue[arrow.Date64](arr), nil

	case *array.Timestamp:
		return timestampValue(arr)

	case *array.Decimal128:
		return decimalValue[decimal128.Num](arr), nil
	case *array.Decimal256:
		return decimalValue[decimal256.Num](arr), nil

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

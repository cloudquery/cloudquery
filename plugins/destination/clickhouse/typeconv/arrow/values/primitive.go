package values

import (
	"github.com/apache/arrow/go/v12/arrow/array"
	"github.com/apache/arrow/go/v12/arrow/float16"
)

type primitiveBuilder[A any] interface {
	array.Builder
	Append(A)
}

func buildPrimitiveValues[A any](builder primitiveBuilder[A], value *A) {
	switch {
	case value == nil, value == (*A)(nil):
		builder.AppendNull()
	default:
		builder.Append(*value)
	}
}

func buildFloat16Values(builder primitiveBuilder[float16.Num], value *float32) {
	switch {
	case value == nil, value == (*float32)(nil):
		builder.AppendNull()
	default:
		builder.Append(float16.New(*value))
	}
}

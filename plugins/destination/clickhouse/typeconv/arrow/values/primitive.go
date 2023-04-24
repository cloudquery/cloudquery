package values

import (
	"github.com/apache/arrow/go/v12/arrow/array"
)

type primitiveBuilder[A any] interface {
	array.Builder
	Append(A)
}

func buildPrimitive[A any](builder primitiveBuilder[A], value *A) {
	switch {
	case value == nil, value == (*A)(nil):
		builder.AppendNull()
	default:
		builder.Append(*value)
	}
}

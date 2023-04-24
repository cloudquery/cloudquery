package values

import (
	"github.com/apache/arrow/go/v12/arrow/array"
)

type primitiveBuilder[A any] interface {
	array.Builder
	Append(A)
}

func buildPrimitive[A any](builder primitiveBuilder[A], value any) {
	v, ok := unwrap[A](value)
	if !ok {
		builder.AppendNull()
		return
	}
	builder.Append(v)
}

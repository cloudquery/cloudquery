package values

import (
	"github.com/apache/arrow-go/v18/arrow/array"
)

type primitiveBuilder[A any] interface {
	array.Builder
	Append(A)
}

func buildPrimitive[A any, B primitiveBuilder[A]](builder B, value any) {
	v, ok := unwrap[A](value)
	if !ok {
		builder.AppendNull()
		return
	}
	builder.Append(v)
}

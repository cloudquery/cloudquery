package values

import (
	"github.com/apache/arrow-go/v18/arrow/float16"
)

func buildFloat16(builder primitiveBuilder[float16.Num], value any) {
	v, ok := unwrap[float32](value)
	if !ok {
		builder.AppendNull()
		return
	}
	builder.Append(float16.New(v))
}

func buildBinary(builder primitiveBuilder[[]byte], value any) {
	v, ok := unwrap[string](value)
	if !ok {
		builder.AppendNull()
		return
	}
	builder.Append([]byte(v))
}

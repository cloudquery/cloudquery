package values

import (
	"github.com/apache/arrow/go/v12/arrow/float16"
	"github.com/cloudquery/plugin-sdk/v2/types"
	"github.com/google/uuid"
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

func buildUUID(builder *types.UUIDBuilder, value any) {
	v, ok := unwrap[uuid.UUID](value)
	if !ok {
		builder.AppendNull()
		return
	}
	builder.Append(v)
}

func unwrap[A any](value any) (A, bool) {
	var unwrapped A
	switch value := value.(type) {
	case **A:
		if value != (**A)(nil) {
			return unwrap[A](*value)
		}
	case *A:
		switch {
		case value == nil, value == (*A)(nil):
		default:
			return *value, true
		}
	}
	return unwrapped, false
}

package values

import (
	"strings"

	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/apache/arrow/go/v13/arrow/float16"
	"github.com/goccy/go-json"
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

func buildUnmarshalOne(builder array.Builder, value any) error {
	v, ok := unwrap[string](value)
	if !ok {
		builder.AppendNull()
		return nil
	}

	return builder.UnmarshalOne(json.NewDecoder(strings.NewReader(v)))
}

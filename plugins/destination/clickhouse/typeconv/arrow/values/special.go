package values

import (
	"github.com/apache/arrow/go/v12/arrow/float16"
)

func buildFloat16(builder primitiveBuilder[float16.Num], value *float32) {
	switch {
	case value == nil, value == (*float32)(nil):
		builder.AppendNull()
	default:
		builder.Append(float16.New(*value))
	}
}

func buildBinary(builder primitiveBuilder[[]byte], value *string) {
	switch {
	case value == nil, value == (*string)(nil):
		builder.AppendNull()
	default:
		builder.Append([]byte(*value))
	}
}

package values

import (
	"github.com/apache/arrow/go/v12/arrow/array"
)

func buildFromString(builder array.Builder, value any) error {
	v, ok := unwrap[string](value)
	if !ok {
		builder.AppendNull()
		return nil
	}
	return builder.AppendValueFromString(v)
}

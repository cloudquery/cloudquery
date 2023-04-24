package values

import (
	"github.com/apache/arrow/go/v12/arrow/array"
)

func buildFromString(builder array.Builder, value *string) error {
	switch {
	case value == nil, value == (*string)(nil):
		builder.AppendNull()
		return nil
	default:
		return builder.AppendValueFromString(*value)
	}
}

package values

import (
	"github.com/apache/arrow/go/v12/arrow/array"
)

func buildFromString(builder array.Builder, value *string) error {
	if value == nil {
		builder.AppendNull()
		return nil
	}
	return builder.AppendValueFromString(*value)
}

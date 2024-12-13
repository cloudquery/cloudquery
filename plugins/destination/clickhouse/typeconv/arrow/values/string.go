package values

import (
	"github.com/apache/arrow-go/v18/arrow/array"
)

func buildFromString(builder array.Builder, value any) error {
	return buildFromStringWithZero(builder, value, "")
}

// buildFromStringWithZero will use builder.AppendEmptyValue if the v is "" or matches the passed empty value
func buildFromStringWithZero(builder array.Builder, value any, zero string) error {
	v, ok := unwrap[string](value)
	if !ok {
		builder.AppendNull()
		return nil
	}

	if len(v) == 0 || v == zero {
		builder.AppendEmptyValue()
		return nil
	}

	return builder.AppendValueFromString(v)
}

package values

import (
	"github.com/apache/arrow/go/v13/arrow/array"
)

func buildFromString(builder array.Builder, value any) error {
	v, ok := unwrap[string](value)
	if !ok {
		builder.AppendNull()
		return nil
	}

	if len(v) > 0 {
		return builder.AppendValueFromString(v)
	}

	// binary types are handled separately, so here we have a builder that most likely can't handle empty string.
	// having empty string in CH means that this was a null value
	builder.AppendNull()
	return nil
}

package client

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thoas/go-funk"
)

func TestStringToNullableArrayPathResolver(t *testing.T) {
	// this test will show that we actually get a typed nil
	s := struct {
		A []string
	}{
		A: []string{"a", "", "None", "b"},
	}

	value := funk.Get(s, "A", funk.WithAllowZero())

	stringArrayValue, ok := value.([]string)
	require.True(t, ok)

	sanitized := make([]*string, len(stringArrayValue))

	for i := range stringArrayValue {
		switch stringArrayValue[i] {
		case "", "None":
		// nop, as already nil
		default:
			sanitized[i] = &stringArrayValue[i]
		}
	}

	for i := range sanitized {
		require.False(t, isNilTrivial(sanitized[i]))
	}
}

func isNilTrivial(a any) bool {
	return a == nil
}

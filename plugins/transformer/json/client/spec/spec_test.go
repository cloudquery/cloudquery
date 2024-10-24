package spec

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetDefaults(t *testing.T) {
	tests := []struct {
		name     string
		input    Spec
		expected Spec
	}{
		{
			name: "Adds * as default pattern",
			input: Spec{
				Tables: nil,
			},
			expected: Spec{
				Tables: []string{"*"},
			},
		},
		{
			name: "Does nothing",
			input: Spec{
				Tables: []string{"*"},
			},
			expected: Spec{
				Tables: []string{"*"},
			},
		},
		{
			name: "Does nothing when a specific pattern is set",
			input: Spec{
				Tables: []string{"table1"},
			},
			expected: Spec{
				Tables: []string{"table1"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.input.SetDefaults()
			require.Equal(t, tt.expected, tt.input)
		})
	}
}

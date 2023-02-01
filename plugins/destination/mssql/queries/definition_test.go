package queries

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDefinition_Constraint(t *testing.T) {
	type testCase struct {
		def      *Definition
		expected string
	}

	for _, c := range []testCase{
		{
			def: &Definition{},
		},
		{
			def:      &Definition{unique: true},
			expected: "UNIQUE",
		},
		{
			def:      &Definition{notNull: true},
			expected: "NOT NULL",
		},
		{
			def:      &Definition{notNull: true, unique: true},
			expected: "UNIQUE NOT NULL",
		},
	} {
		require.Equal(t, c.expected, c.def.Constraint())
	}
}

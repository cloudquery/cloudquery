package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnique(t *testing.T) {
	for _, tc := range []struct {
		Input    []string
		Expected []string
	}{
		{
			Input:    nil,
			Expected: nil,
		},
		{
			Input:    []string{},
			Expected: []string{},
		},
		{
			Input:    []string{"a"},
			Expected: []string{"a"},
		},
		{
			Input:    []string{"b", "a"},
			Expected: []string{"a", "b"},
		},
		{
			Input:    []string{"a", "b"},
			Expected: []string{"a", "b"},
		},
		{
			Input:    []string{"a", "c", "b", "d"},
			Expected: []string{"a", "b", "c", "d"},
		},
		{
			Input:    []string{"a", "c", "b", "d", "b", "c"},
			Expected: []string{"a", "b", "c", "d"},
		},
		{
			Input:    []string{"a", "c", "b", "d", "b", "c", "e"},
			Expected: []string{"a", "b", "c", "d", "e"},
		},
	} {
		ret := Unique(tc.Input)
		assert.EqualValues(t, tc.Expected, ret)
	}
}

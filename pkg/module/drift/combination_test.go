package drift

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombination(t *testing.T) {
	cases := []struct {
		InputBase      [][]string
		InputMul       []string
		ExpectedResult [][]string
	}{
		{ // 0 x 2
			nil,
			[]string{"a", "b"},
			[][]string{{"a"}, {"b"}},
		},
		{ // 2 x 0
			[][]string{{"a"}, {"b"}},
			nil,
			[][]string{{"a"}, {"b"}},
		},
		{ // 2 x 2
			[][]string{{"a"}, {"b"}},
			[]string{"A", "B"},
			[][]string{{"a", "A"}, {"b", "A"}, {"a", "B"}, {"b", "B"}},
		},
		{ // 3 x 2
			[][]string{{"a"}, {"b"}, {"c"}},
			[]string{"A", "B"},
			[][]string{{"a", "A"}, {"b", "A"}, {"c", "A"}, {"a", "B"}, {"b", "B"}, {"c", "B"}},
		},
		{ // 2 x 3
			[][]string{{"a"}, {"b"}},
			[]string{"A", "B", "C"},
			[][]string{{"a", "A"}, {"b", "A"}, {"a", "B"}, {"b", "B"}, {"a", "C"}, {"b", "C"}},
		},
		{ // 3 x 3
			[][]string{{"a"}, {"b"}, {"c"}},
			[]string{"A", "B", "C"},
			[][]string{{"a", "A"}, {"b", "A"}, {"c", "A"}, {"a", "B"}, {"b", "B"}, {"c", "B"}, {"a", "C"}, {"b", "C"}, {"c", "C"}},
		},
	}
	for _, tc := range cases {
		res := MatrixProduct(tc.InputBase, tc.InputMul)
		assert.Equal(t, tc.ExpectedResult, res)
	}
}

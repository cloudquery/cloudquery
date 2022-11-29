package recipes

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppendNoRepeat(t *testing.T) {
	tc := []struct {
		Input    string
		Add      string
		Expected string
	}{
		{
			Input:    "foo",
			Add:      "bar",
			Expected: "foo_bar",
		},
		{
			Input:    "foo",
			Add:      "foo_bar",
			Expected: "foo_bar",
		},
		{
			Input:    "foo_bar",
			Add:      "bar_baz",
			Expected: "foo_bar_baz",
		},
		{
			Input:    "foo_bar",
			Add:      "baz_bax",
			Expected: "foo_bar_baz_bax",
		},
		{
			Input:    "foo_bar_baz",
			Add:      "bar_baz_bax",
			Expected: "foo_bar_baz_bax",
		},
	}
	const sep = "_"
	for _, tt := range tc {
		val := appendNoRepeat(strings.Split(tt.Input, sep), strings.Split(tt.Add, sep)...)
		if !assert.Equal(t, tt.Expected, strings.Join(val, sep)) {
			t.Logf("Failed inputs: %v + %v != %v ???\n", tt.Input, tt.Add, tt.Expected)
		}

	}
}

func TestSliceEndsWith(t *testing.T) {
	tc := []struct {
		Haystack []string
		Needle   []string
		Expected bool
	}{
		{
			Haystack: []string{"a"},
			Needle:   []string{"a"},
			Expected: true,
		},
		{
			Haystack: []string{"a"},
			Needle:   []string{"b"},
			Expected: false,
		},
		{
			Haystack: []string{"a", "b"},
			Needle:   []string{"b"},
			Expected: true,
		},
		{
			Haystack: []string{"a", "b", "c"},
			Needle:   []string{"b"},
			Expected: false,
		},
		{
			Haystack: []string{"a", "b", "c"},
			Needle:   []string{"a", "b", "c"},
			Expected: true,
		},
		{
			Haystack: []string{"a", "b", "c"},
			Needle:   []string{"b", "c"},
			Expected: true,
		},
		{
			Haystack: []string{"a", "b", "c"},
			Needle:   []string{"d", "b", "c"},
			Expected: false,
		},
	}
	for _, tt := range tc {
		val := sliceEndsWith(tt.Haystack, tt.Needle)
		if !assert.Equal(t, tt.Expected, val) {
			t.Logf("Failed inputs: %v + %v != %v ???\n", tt.Haystack, tt.Needle, tt.Expected)
		}
	}
}

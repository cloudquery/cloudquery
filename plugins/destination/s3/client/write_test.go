package client

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSanitizeJSONKeys(t *testing.T) {
	m := map[string]any{
		"foo": "bar",
		"bar": map[string]any{
			"foo-bar": "baz",
		},
		"foo:bar":     "baz",
		"foo-bar-baz": []any{"baz", map[string]any{"foo:bar": "baz"}},
		"string": map[string]string{
			"foo-bar": "baz",
		},
		"int": map[string]int{
			"foo-bar": 123,
		},
		"pointer": map[string]*string{
			"foo-bar": &[]string{"baz"}[0],
		},
	}
	sanitizeJSONKeys(m)
	want := map[string]any{
		"foo": "bar",
		"bar": map[string]any{
			"foo_bar": "baz",
		},
		"foo_bar":     "baz",
		"foo_bar_baz": []any{"baz", map[string]any{"foo_bar": "baz"}},
		"string": map[string]string{
			"foo_bar": "baz",
		},
		"int": map[string]int{
			"foo_bar": 123,
		},
		"pointer": map[string]*string{
			"foo_bar": &[]string{"baz"}[0],
		},
	}
	if diff := cmp.Diff(want, m); diff != "" {
		t.Errorf("sanitizeJSONKeys() mismatch (-want +got):\n%s", diff)
	}
}

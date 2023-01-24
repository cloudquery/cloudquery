package client

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSanitizeJSONKeys(t *testing.T) {
	m := map[string]any{
		"foo": "bar",
		"baz": map[string]any{
			"foo-bar": "baz",
		},
		"foo:bar":     "baz",
		"foo-bar-baz": []any{"baz", map[string]any{"foo:bar": "baz"}},
	}
	sanitizeJSONKeys(m)
	want := map[string]any{
		"foo": "bar",
		"baz": map[string]any{
			"foo_bar": "baz",
		},
		"foo_bar":     "baz",
		"foo_bar_baz": []any{"baz", map[string]any{"foo_bar": "baz"}},
	}
	if diff := cmp.Diff(want, m); diff != "" {
		t.Errorf("sanitizeJSONKeys() mismatch (-want +got):\n%s", diff)
	}
}

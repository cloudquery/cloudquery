package specs

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReplaceSourcePlaceholders(t *testing.T) {
	cases := []struct {
		name        string
		give        map[string]any
		values      []ReplacementValue
		expectError bool
		expected    map[string]any
	}{
		{
			name: "complex",
			give: map[string]any{
				"a":              "b",
				"number":         json.Number("123"),
				"not_connection": "name@email.con",
				"connection":     []string{"@my_destination.connection"},
				"inner": map[string]any{
					"c":          123,
					"connection": "@my_destination.connection",
				},
			},
			expected: map[string]any{
				"a":              "b",
				"number":         json.Number("123"),
				"not_connection": "name@email.con",
				"connection":     []string{"postgres://localhost:5432"},
				"inner": map[string]any{
					"c":          123,
					"connection": "postgres://localhost:5432",
				},
			},
			values: []ReplacementValue{
				{
					PluginName: "my_destination",
					Connection: "postgres://localhost:5432",
				},
			},
		},
		{
			name: "error_plugin_not_found",
			give: map[string]any{
				"a":          "b",
				"connection": []string{"@invalid_destination.connection"},
			},
			expected: map[string]any{
				"a":          "b",
				"connection": []string{"@invalid_destination.connection"},
			},
			expectError: true,
			values: []ReplacementValue{
				{
					PluginName: "my_destination",
					Connection: "postgres://localhost:5432",
				},
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := ReplacePlaceholders(tc.give, tc.values)
			if err != nil && !tc.expectError {
				t.Fatalf("unexpected error: %v", err)
			} else if tc.expectError && err == nil {
				t.Fatalf("expected error but got none")
			}
			// replacement happens in-place
			got := tc.give
			if diff := cmp.Diff(tc.expected, got); diff != "" {
				t.Errorf("diff (+got, -want): %v", diff)
			}
		})
	}
}

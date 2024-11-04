package util

import "testing"

func TestSortedKeys(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		expected []string
	}{
		{
			name:     "Empty map",
			input:    map[string]int{},
			expected: []string{},
		},
		{
			name:     "Single key",
			input:    map[string]int{"a": 1},
			expected: []string{"a"},
		},
		{
			name:     "Multiple keys",
			input:    map[string]int{"c": 3, "a": 1, "b": 2},
			expected: []string{"a", "b", "c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SortedKeys(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("Expected %d keys, got %d", len(tt.expected), len(result))
			}
			for i, key := range result {
				if key != tt.expected[i] {
					t.Errorf("Expected key %s at position %d, got %s", tt.expected[i], i, key)
				}
			}
		})
	}
}

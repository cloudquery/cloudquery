package keyvals

import (
	"errors"
	"testing"
)

func TestToMap(t *testing.T) {
	err := errors.New("error")
	actual := ToMap([]interface{}{"key", "value", "error", err})

	expected := map[string]interface{}{
		"key":   "value",
		"error": err,
	}

	if len(expected) != len(actual) {
		t.Fatalf("expected logging fields to be equal\ngot:  %v\nwant: %v", actual, expected)
	}

	for key, value := range expected {
		if actual[key] != value {
			t.Errorf("expected logging fields to be equal\ngot:  %v\nwant: %v", actual, expected)
		}
	}
}

package tableoptions

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestProcessRelativeTimeProcessing(t *testing.T) {
	valCases := []struct {
		name     string
		input    string
		now      time.Time
		expected time.Time
	}{
		{
			name:     "now",
			input:    "now",
			now:      time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC),
			expected: time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC),
		},
		{
			name:     "yesterday",
			input:    "now-1d",
			now:      time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC),
			expected: time.Date(2020, 1, 1, 3, 4, 5, 0, time.UTC),
		},
		{
			name:     "now truncated by 1 hour",
			input:    "now%1h",
			now:      time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC),
			expected: time.Date(2020, 1, 2, 3, 0, 0, 0, time.UTC),
		},
		{
			name:     "yesterday truncated by day",
			input:    "now-1d%1d",
			now:      time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC),
			expected: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, tc := range valCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			obj := map[string]any{
				"test_time": tc.input,
			}
			err := processRelativeTimes(obj, tc.now, []string{"test_time"})
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tc.expected, obj["test_time"])
		})
	}
}

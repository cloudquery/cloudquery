package types

import (
	"testing"
	"time"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/stretchr/testify/require"
)

func Test_timestampType(t *testing.T) {
	type testCase struct {
		_type    *arrow.TimestampType
		expected string
	}

	loc, err := time.LoadLocation("America/New_York")
	require.NoError(t, err)
	timeZone := loc.String()

	for _, tc := range []testCase{
		{_type: &arrow.TimestampType{Unit: 128}, expected: "String"},
		{_type: &arrow.TimestampType{TimeZone: "NeverWhen"}, expected: "String"},
		{_type: &arrow.TimestampType{Unit: arrow.Second}, expected: "DateTime64(0)"},
		{
			_type:    &arrow.TimestampType{Unit: arrow.Second, TimeZone: timeZone},
			expected: "DateTime64(0, 'America/New_York')",
		},
		{_type: &arrow.TimestampType{Unit: arrow.Millisecond}, expected: "DateTime64(3)"},
		{
			_type:    &arrow.TimestampType{Unit: arrow.Millisecond, TimeZone: timeZone},
			expected: "DateTime64(3, 'America/New_York')",
		},
		{_type: &arrow.TimestampType{Unit: arrow.Microsecond}, expected: "DateTime64(6)"},
		{
			_type:    &arrow.TimestampType{Unit: arrow.Microsecond, TimeZone: timeZone},
			expected: "DateTime64(6, 'America/New_York')",
		},
		{_type: &arrow.TimestampType{Unit: arrow.Nanosecond}, expected: "DateTime64(9)"},
		{
			_type:    &arrow.TimestampType{Unit: arrow.Nanosecond, TimeZone: timeZone},
			expected: "DateTime64(9, 'America/New_York')",
		},
	} {
		ensureDefinition(t, tc._type, tc.expected)
	}
}

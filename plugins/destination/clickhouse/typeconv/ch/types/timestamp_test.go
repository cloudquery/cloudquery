package types

import (
	"testing"
	"time"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/stretchr/testify/require"
)

func Test_timestampType(t *testing.T) {
	loc, err := time.LoadLocation("America/New_York")
	require.NoError(t, err)
	timeZone := loc.String()

	for _, tc := range []testCase{
		{dataType: &arrow.Time32Type{Unit: 128}, expected: "String"},
		{dataType: &arrow.Time32Type{Unit: arrow.Second}, expected: "String"},
		{dataType: &arrow.Time32Type{Unit: arrow.Millisecond}, expected: "String"},
		{dataType: &arrow.Time64Type{Unit: 128}, expected: "String"},
		{dataType: &arrow.Time64Type{Unit: arrow.Microsecond}, expected: "String"},
		{dataType: &arrow.Time64Type{Unit: arrow.Nanosecond}, expected: "String"},
		{dataType: &arrow.TimestampType{Unit: arrow.Second}, expected: "DateTime64(0)"},
		{
			dataType: &arrow.TimestampType{Unit: arrow.Second, TimeZone: timeZone},
			expected: "DateTime64(0, 'America/New_York')",
		},
		{dataType: &arrow.TimestampType{Unit: arrow.Millisecond}, expected: "DateTime64(3)"},
		{
			dataType: &arrow.TimestampType{Unit: arrow.Millisecond, TimeZone: timeZone},
			expected: "DateTime64(3, 'America/New_York')",
		},
		{dataType: &arrow.TimestampType{Unit: arrow.Microsecond}, expected: "DateTime64(6)"},
		{
			dataType: &arrow.TimestampType{Unit: arrow.Microsecond, TimeZone: timeZone},
			expected: "DateTime64(6, 'America/New_York')",
		},
		{dataType: &arrow.TimestampType{Unit: arrow.Nanosecond}, expected: "DateTime64(9)"},
		{
			dataType: &arrow.TimestampType{Unit: arrow.Nanosecond, TimeZone: timeZone},
			expected: "DateTime64(9, 'America/New_York')",
		},
	} {
		ensureDefinition(t, tc)
	}
}

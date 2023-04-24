package arrow

import (
	"testing"
	"time"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/stretchr/testify/require"
)

func Test_dateTimeType(t *testing.T) {
	type testCase struct {
		_type    string
		expected arrow.DataType
	}

	tz, err := time.LoadLocation("America/New_York")
	require.NoError(t, err)
	for _, tc := range []testCase{
		{_type: "DateTime", expected: &arrow.TimestampType{Unit: arrow.Second}},
		{_type: "DateTime('America/New_York')", expected: &arrow.TimestampType{Unit: arrow.Second, TimeZone: tz.String()}},
	} {
		ensureField(t, tc._type, tc.expected)
	}
}

func Test_dateTime64Type(t *testing.T) {
	type testCase struct {
		_type    string
		expected arrow.DataType
	}

	tz, err := time.LoadLocation("America/New_York")
	require.NoError(t, err)
	for _, tc := range []testCase{
		{_type: "DateTime64(0)", expected: &arrow.TimestampType{Unit: arrow.Second}},
		{_type: "DateTime64(1)", expected: &arrow.TimestampType{Unit: arrow.Millisecond}},
		{_type: "DateTime64(2)", expected: &arrow.TimestampType{Unit: arrow.Millisecond}},
		{_type: "DateTime64(3)", expected: &arrow.TimestampType{Unit: arrow.Millisecond}},
		{_type: "DateTime64(4)", expected: &arrow.TimestampType{Unit: arrow.Microsecond}},
		{_type: "DateTime64(5)", expected: &arrow.TimestampType{Unit: arrow.Microsecond}},
		{_type: "DateTime64(6)", expected: &arrow.TimestampType{Unit: arrow.Microsecond}},
		{_type: "DateTime64(7)", expected: &arrow.TimestampType{Unit: arrow.Nanosecond}},
		{_type: "DateTime64(8)", expected: &arrow.TimestampType{Unit: arrow.Nanosecond}},
		{_type: "DateTime64(9)", expected: &arrow.TimestampType{Unit: arrow.Nanosecond}},
		{_type: "DateTime64(7, 'America/New_York')", expected: &arrow.TimestampType{Unit: arrow.Nanosecond, TimeZone: tz.String()}},
	} {
		ensureField(t, tc._type, tc.expected)
	}
}

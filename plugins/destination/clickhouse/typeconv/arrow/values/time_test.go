package values

import (
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/stretchr/testify/require"
)

func Test_timestamp(t *testing.T) {
	const (
		amount = 100
		unit   = arrow.Millisecond
	)
	for i := 0; i < amount; i++ {
		ts := time.Now().UTC().AddDate(0, i, i).Add(time.Duration(i * 123))
		t.Run(ts.String(), func(t *testing.T) {
			aTs, err := arrow.TimestampFromTime(ts, unit)
			require.NoError(t, err)
			ensureRecord(t, testCase{dataType: &arrow.TimestampType{Unit: arrow.Millisecond}, value: &ts, expected: aTs})
		})
	}
}

func Test_time32(t *testing.T) {
	const (
		amount = 100
		unit   = arrow.Millisecond
	)
	for i := 0; i < amount; i++ {
		ts := time.Unix(int64(i), int64(i-(i%int(time.Millisecond)))).UTC()
		aTs, err := arrow.Time32FromString(ts.Format("15:04:05.999"), unit)
		require.NoError(t, err)
		tt := aTs.ToTime(unit)
		require.Exactly(t, ts, tt)
		ensureRecord(t, testCase{dataType: &arrow.Time32Type{Unit: unit}, value: &ts, expected: aTs})
	}
}

func Test_time64(t *testing.T) {
	const (
		amount = 100
		unit   = arrow.Nanosecond
	)
	for i := 0; i < amount; i++ {
		ts := time.Unix(int64(i), int64(i)).UTC()
		aTs, err := arrow.Time64FromString(ts.Format("15:04:05.999999999"), unit)
		require.NoError(t, err)
		tt := aTs.ToTime(unit)
		require.Exactly(t, ts, tt)
		ensureRecord(t, testCase{dataType: &arrow.Time64Type{Unit: unit}, value: &ts, expected: aTs})
	}
}

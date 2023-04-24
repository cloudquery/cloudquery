package value

import (
	"testing"
	"time"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
	"github.com/apache/arrow/go/v12/arrow/memory"
	"github.com/stretchr/testify/require"
)

func Test_timeValue(t *testing.T) {
	const N = 100
	values := make([]time.Time, N)

	loc, err := time.LoadLocation("America/New_York")
	require.NoError(t, err)

	bld := array.NewTimestampBuilder(memory.DefaultAllocator, &arrow.TimestampType{Unit: arrow.Nanosecond, TimeZone: loc.String()})
	for i := range values {
		now := time.Now().In(loc)
		values[i] = now

		arrowTimestamp, ok, err := arrow.TimestampFromStringInLocation(now.Format(time.RFC3339Nano), arrow.Nanosecond, loc)
		require.NoError(t, err)
		require.True(t, ok) // check that zone formatted correctly
		bld.Append(arrowTimestamp)
	}

	// we can use timestampValue here, but we want to check the returned data type from the topmost func, too
	data, err := FromArray(bld.NewTimestampArray())
	require.NoError(t, err)

	elems := data.([]*time.Time)
	require.Equal(t, N, len(elems))
	for i, elem := range elems {
		require.NotNil(t, elem)
		require.Exactly(t, values[i], *elem)
	}
}

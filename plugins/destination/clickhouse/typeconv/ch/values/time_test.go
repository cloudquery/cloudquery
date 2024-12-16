package values

import (
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/stretchr/testify/require"
)

func Test_timeValue(t *testing.T) {
	const amount = 100
	values := make([]time.Time, amount)

	loc, err := time.LoadLocation("America/New_York")
	require.NoError(t, err)

	builder := array.NewTimestampBuilder(memory.DefaultAllocator, &arrow.TimestampType{Unit: arrow.Nanosecond, TimeZone: loc.String()})
	for i := range values {
		now := time.Now().In(loc)
		values[i] = now

		arrowTimestamp, ok, err := arrow.TimestampFromStringInLocation(now.Format(time.RFC3339Nano), arrow.Nanosecond, loc)
		require.NoError(t, err)
		require.True(t, ok) // check that zone formatted correctly
		builder.Append(arrowTimestamp)
	}

	// we can use timestampValue here, but we want to check the returned data type from the topmost func, too
	data, err := FromArray(builder.NewArray())
	require.NoError(t, err)

	elems := data.([]*time.Time)
	require.Equal(t, amount, len(elems))
	for i, elem := range elems {
		require.NotNil(t, elem)
		require.Exactly(t, values[i], *elem)
	}
}

func Test_time32Value(t *testing.T) {
	const amount = 100
	values := make([]time.Time, amount)
	builder := array.NewTime32Builder(memory.DefaultAllocator, &arrow.Time32Type{Unit: arrow.Millisecond})
	for i := range values {
		values[i] = time.Unix(int64(i), int64(i-(i%int(time.Millisecond)))).UTC()
		builder.Append(arrow.Time32(values[i].UnixMilli()))
	}

	// we can use timestampValue here, but we want to check the returned data type from the topmost func, too
	data, err := FromArray(builder.NewArray())
	require.NoError(t, err)

	elems := data.([]*time.Time)
	require.Equal(t, amount, len(elems))
	for i, elem := range elems {
		require.NotNil(t, elem)
		require.Exactly(t, values[i], *elem)
	}
}

func Test_time64Value(t *testing.T) {
	const amount = 100
	values := make([]time.Time, amount)
	builder := array.NewTime64Builder(memory.DefaultAllocator, &arrow.Time64Type{Unit: arrow.Nanosecond})
	for i := range values {
		values[i] = time.Unix(int64(i), int64(i)).UTC()
		builder.Append(arrow.Time64(values[i].UnixNano()))
	}

	// we can use timestampValue here, but we want to check the returned data type from the topmost func, too
	data, err := FromArray(builder.NewArray())
	require.NoError(t, err)

	elems := data.([]*time.Time)
	require.Equal(t, amount, len(elems))
	for i, elem := range elems {
		require.NotNil(t, elem)
		require.Exactly(t, values[i], *elem)
	}
}

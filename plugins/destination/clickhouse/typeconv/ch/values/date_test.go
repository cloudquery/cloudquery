package values

import (
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/stretchr/testify/require"
)

func Test_date32Value(t *testing.T) {
	const amount = 100
	values := make([]arrow.Date32, amount)

	builder := array.NewDate32Builder(memory.DefaultAllocator)
	for i := range values {
		values[i] = arrow.Date32FromTime(time.Now().UTC().AddDate(0, i, i))
		builder.Append(values[i])
	}

	// we can use timestampValue here, but we want to check the returned data type from the topmost func, too
	data, err := FromArray(builder.NewArray())
	require.NoError(t, err)

	elems := data.([]*time.Time)
	require.Equal(t, amount, len(elems))
	for i, elem := range elems {
		require.NotNil(t, elem)
		require.Exactly(t, values[i].ToTime(), *elem)
		require.Exactly(t, values[i], arrow.Date32FromTime(*elem))
	}
}

func Test_date64Value(t *testing.T) {
	const amount = 100
	values := make([]arrow.Date64, amount)

	builder := array.NewDate64Builder(memory.DefaultAllocator)
	for i := range values {
		values[i] = arrow.Date64FromTime(time.Now().UTC().AddDate(0, i, i))
		builder.Append(values[i])
	}

	// we can use timestampValue here, but we want to check the returned data type from the topmost func, too
	data, err := FromArray(builder.NewArray())
	require.NoError(t, err)

	elems := data.([]*time.Time)
	require.Equal(t, amount, len(elems))
	for i, elem := range elems {
		require.NotNil(t, elem)
		require.Exactly(t, values[i].ToTime(), *elem)
		require.Exactly(t, values[i], arrow.Date64FromTime(*elem))
	}
}

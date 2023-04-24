package values

import (
	"math"
	"math/rand"
	"testing"

	"github.com/apache/arrow/go/v12/arrow/array"
	"github.com/apache/arrow/go/v12/arrow/memory"
	"github.com/stretchr/testify/require"
)

func TestFromArray(t *testing.T) {
	const amount = 100

	values := make([]float64, amount)
	for i := range values {
		values[i] = rand.Float64()*(math.MaxFloat64-1) + rand.Float64()
	}

	builder := array.NewFloat64Builder(memory.DefaultAllocator)
	for _, f := range values {
		builder.Append(f)
	}

	data, err := FromArray(builder.NewArray())
	require.NoError(t, err)

	elems := data.([]*float64)
	require.Equal(t, amount, len(elems))
	for i, elem := range elems {
		require.NotNil(t, elem)
		require.Exactly(t, values[i], *elem)
	}
}

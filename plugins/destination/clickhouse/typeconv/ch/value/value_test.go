package value

import (
	"math"
	"math/rand"
	"testing"

	"github.com/apache/arrow/go/v12/arrow/array"
	"github.com/apache/arrow/go/v12/arrow/memory"
	"github.com/stretchr/testify/require"
)

func TestFromArray(t *testing.T) {
	const N = 100

	floats := make([]float64, N)
	for i := range floats {
		floats[i] = rand.Float64()*(math.MaxFloat64-1) + rand.Float64()
	}

	bld := array.NewFloat64Builder(memory.DefaultAllocator)
	for _, f := range floats {
		bld.Append(f)
	}

	data, err := FromArray(bld.NewArray())
	require.NoError(t, err)

	elems := data.([]*float64)
	require.Equal(t, N, len(elems))
	for i, elem := range elems {
		require.NotNil(t, elem)
		require.Exactly(t, floats[i], *elem)
	}
}

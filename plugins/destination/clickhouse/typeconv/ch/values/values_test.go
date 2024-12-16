package values

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/schema"
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

func BenchmarkFromArray(b *testing.B) {
	table := schema.TestTable("test", schema.TestSourceOptions{})
	sourceName := "test-source"
	syncTime := time.Now().UTC().Round(time.Second)
	opts := schema.GenTestDataOptions{
		SourceName: sourceName,
		SyncTime:   syncTime,
		MaxRows:    b.N,
	}
	record := schema.NewTestDataGenerator(0).Generate(table, opts)
	b.ResetTimer()
	for _, col := range record.Columns() {
		_, _ = FromArray(col)
	}
}

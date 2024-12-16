package values

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func Test_listValue(t *testing.T) {
	const (
		amount = 200
		width  = 5
	)

	values := make([][]uuid.UUID, amount)
	for i := range values {
		if i%2 == 0 {
			continue
		}
		row := make([]uuid.UUID, width)
		for j := range row {
			row[j] = uuid.New()
		}
		values[i] = row
	}

	builder := array.NewListBuilder(memory.DefaultAllocator, types.NewUUIDType())
	uidBuilder := builder.ValueBuilder().(*types.UUIDBuilder)
	for i, row := range values {
		if i%2 == 0 {
			builder.AppendNull()
			continue
		}
		builder.Append(true)
		for _, uid := range row {
			uidBuilder.Append(uid)
		}
	}

	data, err := listValue(builder.NewListArray())
	require.NoError(t, err)

	uidSlices := data.([][]*uuid.UUID)

	require.Equal(t, amount, len(uidSlices))
	for i, row := range uidSlices {
		require.NotNil(t, row)
		if i%2 == 0 {
			require.Empty(t, row)
			continue
		}
		require.Equal(t, width, len(row))
		for j, uid := range row {
			require.NotNil(t, uid)
			require.Exactly(t, values[i][j], *uid)
		}
	}
}

func Test_largeListValue(t *testing.T) {
	const (
		amount = 200
		width  = 5
	)

	values := make([][]uuid.UUID, amount)
	for i := range values {
		if i%2 == 0 {
			continue
		}
		row := make([]uuid.UUID, width)
		for j := range row {
			row[j] = uuid.New()
		}
		values[i] = row
	}

	builder := array.NewLargeListBuilder(memory.DefaultAllocator, types.NewUUIDType())
	uidBuilder := builder.ValueBuilder().(*types.UUIDBuilder)
	for i, row := range values {
		if i%2 == 0 {
			builder.AppendNull()
			continue
		}
		builder.Append(true)
		for _, uid := range row {
			uidBuilder.Append(uid)
		}
	}

	data, err := listValue(builder.NewLargeListArray())
	require.NoError(t, err)

	uidSlices := data.([][]*uuid.UUID)

	require.Equal(t, amount, len(uidSlices))
	for i, row := range uidSlices {
		require.NotNil(t, row)
		if i%2 == 0 {
			require.Empty(t, row)
			continue
		}
		require.Equal(t, width, len(row))
		for j, uid := range row {
			require.NotNil(t, uid)
			require.Exactly(t, values[i][j], *uid)
		}
	}
}

func Test_fixedSizeListValue(t *testing.T) {
	const (
		amount = 200
		width  = 5
	)

	values := make([][]uuid.UUID, amount)
	for i := range values {
		row := make([]uuid.UUID, width)
		for j := range row {
			if i%2 == 0 {
				continue
			}
			row[j] = uuid.New()
		}
		values[i] = row
	}

	builder := array.NewFixedSizeListBuilder(memory.DefaultAllocator, width, types.NewUUIDType())
	uidBuilder := builder.ValueBuilder().(*types.UUIDBuilder)
	for i, row := range values {
		if i%2 == 0 {
			builder.AppendNull()
			continue
		}
		builder.Append(true)
		for _, uid := range row {
			uidBuilder.Append(uid)
		}
	}

	data, err := listValue(builder.NewListArray())
	require.NoError(t, err)

	uidSlices := data.([][]*uuid.UUID)

	require.Equal(t, amount, len(uidSlices))
	for i, row := range uidSlices {
		require.NotNil(t, row)
		require.Equal(t, width, len(row))
		for j, uid := range row {
			require.NotNil(t, uid)
			require.Exactly(t, values[i][j], *uid)
		}
	}
}

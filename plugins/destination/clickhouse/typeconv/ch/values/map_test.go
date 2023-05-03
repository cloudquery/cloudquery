package values

import (
	"testing"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
	"github.com/apache/arrow/go/v12/arrow/memory"
	"github.com/stretchr/testify/require"
)

func Test_mapValue(t *testing.T) {
	mapType := arrow.MapOf(
		new(arrow.StringType),
		arrow.StructOf(
			arrow.Field{Name: "nullable_bool", Type: new(arrow.BooleanType), Nullable: true},
			arrow.Field{Name: "non_nullable_bool", Type: new(arrow.BooleanType)},
		),
	)

	builder := array.NewMapBuilderWithType(memory.DefaultAllocator, mapType)
	keyBuilder, itemBuilder := builder.KeyBuilder().(*array.StringBuilder), builder.ValueBuilder().(*array.StructBuilder)
	nullableBld := itemBuilder.FieldBuilder(0).(*array.BooleanBuilder)
	nonNullableBld := itemBuilder.FieldBuilder(1).(*array.BooleanBuilder)

	// single proper value
	builder.Append(true)
	keyBuilder.Append("proper")
	itemBuilder.Append(true)
	nullableBld.Append(true)
	nonNullableBld.Append(true)

	// single empty value
	builder.Append(true)
	keyBuilder.Append("empty")
	itemBuilder.AppendNull()

	// 2 values: proper & null
	builder.Append(true)
	keyBuilder.Append("proper")
	itemBuilder.Append(true)
	nullableBld.Append(true)
	nonNullableBld.Append(true)
	keyBuilder.Append("empty")
	itemBuilder.AppendNull()

	// null
	builder.AppendNull()

	data, err := mapValue(builder.NewMapArray())
	require.NoError(t, err)

	elems := data.([]*map[string]*map[string]any)
	require.Equal(t, 4, len(elems))

	// single proper value
	elem := elems[0]
	require.NotNil(t, elem)
	val, ok := (*elem)["proper"]
	require.True(t, ok)
	require.NotNil(t, val)
	require.Equal(t, map[string]any{"non_nullable_bool": true, "nullable_bool": true}, *val)

	// single empty value
	elem = elems[1]
	require.NotNil(t, elem)
	val, ok = (*elem)["empty"]
	require.True(t, ok)
	require.Nil(t, val)

	// 2 values: proper & null
	elem = elems[2]
	require.NotNil(t, elem)
	val, ok = (*elem)["proper"]
	require.True(t, ok)
	require.NotNil(t, val)
	require.Equal(t, map[string]any{"non_nullable_bool": true, "nullable_bool": true}, *val)
	val, ok = (*elem)["empty"]
	require.True(t, ok)
	require.Nil(t, val)

	// null
	elem = elems[3]
	require.Nil(t, elem)
}

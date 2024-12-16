package values

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/stretchr/testify/require"
)

func Test_structValue(t *testing.T) {
	structType := arrow.StructOf(
		arrow.Field{Name: "nullable_bool", Type: new(arrow.BooleanType), Nullable: true},
		arrow.Field{Name: "non_nullable_bool", Type: new(arrow.BooleanType)},
	)

	builder := array.NewStructBuilder(memory.DefaultAllocator, structType)
	nullableBld := builder.FieldBuilder(0).(*array.BooleanBuilder)
	nonNullableBld := builder.FieldBuilder(1).(*array.BooleanBuilder)

	builder.Append(true)
	nullableBld.Append(true)
	nonNullableBld.Append(true)

	builder.Append(true)
	nullableBld.AppendNull()
	nonNullableBld.Append(true)

	builder.AppendNull()

	data, err := structValue(builder.NewStructArray())
	require.NoError(t, err)

	elems := data.([]map[string]any)
	require.Equal(t, 3, len(elems))

	// both filled in
	elem := elems[0]
	require.NotNil(t, elem)
	require.Equal(t, map[string]any{
		"nullable_bool":     ptr(true),
		"non_nullable_bool": ptr(true),
	}, elem)

	// 1 filled in, the other is nil
	elem = elems[1]
	require.NotNil(t, elem)
	require.Equal(t, map[string]any{
		"nullable_bool":     ptr(false),
		"non_nullable_bool": ptr(true),
	}, elem)

	// nil, but for CH we have a map for every key with nil value
	elem = elems[2]
	require.Equal(t, map[string]any{
		"nullable_bool":     ptr(false),
		"non_nullable_bool": ptr(false),
	}, elem)
}

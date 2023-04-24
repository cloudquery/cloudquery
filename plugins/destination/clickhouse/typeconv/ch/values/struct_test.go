package values

import (
	"testing"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
	"github.com/apache/arrow/go/v12/arrow/memory"
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

	data, err := structValue(builder.NewStructArray())
	require.NoError(t, err)

	elems := data.([]*map[string]any)
	require.Equal(t, 2, len(elems))

	// both filled in
	require.NotNil(t, elems[0])
	nullable, ok := (*elems[0])["nullable_bool"]
	require.True(t, ok)
	require.True(t, *nullable.(*bool))

	nonNullable, ok := (*elems[0])["non_nullable_bool"]
	require.True(t, ok)
	require.True(t, *nonNullable.(*bool))

	// 1 filled in, the other is nil
	require.NotNil(t, elems[1])
	nullable, ok = (*elems[1])["nullable_bool"]
	require.True(t, ok)
	require.Nil(t, nullable)

	nonNullable, ok = (*elems[1])["non_nullable_bool"]
	require.True(t, ok)
	require.True(t, *nonNullable.(*bool))
}

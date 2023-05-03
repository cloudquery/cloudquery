package values

import (
	"testing"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
	"github.com/apache/arrow/go/v12/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v2/types"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func Test_mapValue(t *testing.T) {
	mapType := arrow.MapOf(
		new(arrow.StringType),
		arrow.StructOf(
			arrow.Field{Name: "bool_n", Type: new(arrow.BooleanType), Nullable: true},
			arrow.Field{Name: "bool", Type: new(arrow.BooleanType)},
			arrow.Field{Name: "list", Type: arrow.ListOf(types.NewUUIDType())},
			arrow.Field{
				Name: "map",
				Type: arrow.MapOf(new(arrow.Int32Type), new(arrow.Float64Type)),
			},
			arrow.Field{
				Name:     "map_n",
				Type:     arrow.MapOf(new(arrow.Int32Type), new(arrow.Float64Type)),
				Nullable: true,
			},
		),
	)

	builder := array.NewMapBuilderWithType(memory.DefaultAllocator, mapType)
	keyBuilder, itemBuilder := builder.KeyBuilder().(*array.StringBuilder), builder.ItemBuilder().(*array.StructBuilder)

	boolBld := itemBuilder.FieldBuilder(0).(*array.BooleanBuilder)
	boolNBld := itemBuilder.FieldBuilder(1).(*array.BooleanBuilder)
	listBld := itemBuilder.FieldBuilder(2).(*array.ListBuilder)
	listValBld := listBld.ValueBuilder().(*types.UUIDBuilder)

	mapBld := itemBuilder.FieldBuilder(3).(*array.MapBuilder)
	mapKeyBld, mapItemBld := mapBld.KeyBuilder().(*array.Int32Builder), mapBld.ItemBuilder().(*array.Float64Builder)

	mapNBld := itemBuilder.FieldBuilder(4).(*array.MapBuilder)
	mapNKeyBld, mapNItemBld := mapNBld.KeyBuilder().(*array.Int32Builder), mapNBld.ItemBuilder().(*array.Float64Builder)

	// single proper value
	builder.Append(true)
	keyBuilder.Append("proper")
	itemBuilder.Append(true)
	boolBld.Append(false)
	boolNBld.Append(true)
	listBld.Append(true)
	listValBld.Append(uuid.NameSpaceDNS)
	mapBld.Append(true)
	mapKeyBld.Append(123)
	mapItemBld.Append(123.456)
	mapNBld.Append(true)
	mapNKeyBld.Append(321)
	mapNItemBld.Append(654.321)

	// single empty value
	builder.Append(true)
	keyBuilder.Append("empty")
	itemBuilder.AppendNull()

	// 2 values: proper & null
	builder.Append(true)
	keyBuilder.Append("proper")
	itemBuilder.Append(true)
	boolBld.Append(false)
	boolNBld.Append(true)
	listBld.Append(true)
	listValBld.Append(uuid.NameSpaceDNS)
	mapBld.Append(true)
	mapKeyBld.Append(123)
	mapItemBld.Append(123.456)
	mapNBld.AppendNull() // nil map

	keyBuilder.Append("empty")
	itemBuilder.AppendNull()

	// null
	builder.AppendNull()

	data, err := mapValue(builder.NewMapArray())
	require.NoError(t, err)

	elems := data.([]*map[string]map[string]any)
	require.Equal(t, 4, len(elems))

	// single proper value
	elem := elems[0]
	require.NotNil(t, elem)
	val, ok := (*elem)["proper"]
	require.True(t, ok)
	require.NotNil(t, val)
	require.EqualValues(t, map[string]any{
		"bool":   ptr(true),
		"bool_n": ptr(false),
		"list":   ptr([]*uuid.UUID{&uuid.NameSpaceDNS}),
		"map":    ptr(map[int32]*float64{123: ptr(float64(123.456))}),
		"map_n":  ptr(map[int32]*float64{321: ptr(float64(654.321))}),
	}, val)

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
	require.EqualValues(t, map[string]any{
		"bool":   ptr(true),
		"bool_n": ptr(false),
		"list":   ptr([]*uuid.UUID{&uuid.NameSpaceDNS}),
		"map":    ptr(map[int32]*float64{123: ptr(float64(123.456))}),
		"map_n":  ptr(map[int32]*float64(nil)),
	}, val)
	val, ok = (*elem)["empty"]
	require.True(t, ok)
	require.Nil(t, val)

	// null
	elem = elems[3]
	require.NotNil(t, elem)
	require.Nil(t, *elem)
}

func ptr[A any](a A) *A { return &a }

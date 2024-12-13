package values

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/types"
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
			arrow.Field{Name: "map", Type: arrow.MapOf(new(arrow.Int32Type), new(arrow.Float64Type))},
			arrow.Field{Name: "map_n", Type: arrow.MapOf(new(arrow.Int32Type), new(arrow.Float64Type)), Nullable: true},
			arrow.Field{Name: "map_uuid", Type: arrow.MapOf(types.NewUUIDType(), types.NewUUIDType())},
			arrow.Field{Name: "mapped_to_string", Type: arrow.MapOf(new(arrow.Float64Type), new(arrow.StringType))},
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

	mapUUIDBld := itemBuilder.FieldBuilder(5).(*array.MapBuilder)
	mapUUIDKeyBld, mapUUIDItemBld := mapUUIDBld.KeyBuilder().(*types.UUIDBuilder), mapUUIDBld.ItemBuilder().(*types.UUIDBuilder)

	mapToStrBld := itemBuilder.FieldBuilder(6).(*array.MapBuilder)
	mapToStrKeyBld, mapToStrItemBld := mapToStrBld.KeyBuilder().(*array.Float64Builder), mapToStrBld.ItemBuilder().(*array.StringBuilder)

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
	mapUUIDBld.Append(true)
	mapUUIDKeyBld.Append(uuid.NameSpaceURL)
	mapUUIDItemBld.Append(uuid.NameSpaceURL)
	mapToStrBld.Append(true)
	mapToStrKeyBld.Append(1010.543)
	mapToStrItemBld.Append("some string")

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
	mapUUIDBld.Append(true)
	mapUUIDKeyBld.Append(uuid.NameSpaceURL)
	mapUUIDItemBld.AppendNull()
	mapToStrBld.Append(true)
	mapToStrKeyBld.Append(1010.543)
	mapToStrItemBld.AppendNull()

	keyBuilder.Append("empty")
	itemBuilder.AppendNull()

	// null
	builder.AppendNull()

	data, err := mapValue(builder.NewMapArray())
	require.NoError(t, err)

	elems := data.([]map[string]map[string]any)
	require.Equal(t, 4, len(elems))

	// single proper value
	elem := elems[0]
	require.NotNil(t, elem)
	val, ok := elem["proper"]
	require.True(t, ok)
	require.NotNil(t, val)
	require.EqualValues(t, map[string]any{
		"bool":             ptr(true),
		"bool_n":           ptr(false),
		"list":             []*uuid.UUID{&uuid.NameSpaceDNS},
		"map":              map[int32]*float64{123: ptr(float64(123.456))},
		"map_n":            map[int32]*float64{321: ptr(float64(654.321))},
		"map_uuid":         map[uuid.UUID]*uuid.UUID{uuid.NameSpaceURL: ptr(uuid.NameSpaceURL)},
		"mapped_to_string": ptr(`[{"key":1010.543,"value":"some string"}]`),
	}, val)

	// single empty value
	elem = elems[1]
	require.NotNil(t, elem)
	val, ok = elem["empty"]
	require.True(t, ok)
	// tuples are non-nullable in CH
	require.NotNil(t, val)
	require.EqualValues(t, map[string]any{
		"bool":             ptr(false),
		"bool_n":           ptr(false),
		"list":             []*uuid.UUID{},
		"map":              map[int32]*float64{},
		"map_n":            map[int32]*float64{},
		"map_uuid":         map[uuid.UUID]*uuid.UUID{},
		"mapped_to_string": ptr("[]"),
	}, val)

	// 2 values: proper & null
	elem = elems[2]
	require.NotNil(t, elem)
	val, ok = elem["proper"]
	require.True(t, ok)
	require.NotNil(t, val)
	require.EqualValues(t, map[string]any{
		"bool":             ptr(true),
		"bool_n":           ptr(false),
		"list":             []*uuid.UUID{&uuid.NameSpaceDNS},
		"map":              map[int32]*float64{123: ptr(float64(123.456))},
		"map_n":            map[int32]*float64{},
		"map_uuid":         map[uuid.UUID]*uuid.UUID{uuid.NameSpaceURL: (*uuid.UUID)(nil)},
		"mapped_to_string": ptr(`[{"key":1010.543,"value":null}]`),
	}, val)
	val, ok = elem["empty"]
	require.True(t, ok)
	// tuples are non-nullable in CH
	require.NotNil(t, val)
	require.EqualValues(t, map[string]any{
		"bool":             ptr(false),
		"bool_n":           ptr(false),
		"list":             []*uuid.UUID{},
		"map":              map[int32]*float64{},
		"map_n":            map[int32]*float64{},
		"map_uuid":         map[uuid.UUID]*uuid.UUID{},
		"mapped_to_string": ptr("[]"),
	}, val)

	// null
	elem = elems[3]
	// maps are non-nullable in CH
	require.NotNil(t, elem)
	require.Empty(t, elem)
}

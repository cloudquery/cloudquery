package client

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestBsonDocToMap_BsonM(t *testing.T) {
	input := bson.M{"key": "value", "num": int32(42)}
	result := bsonDocToMap(input)
	require.Equal(t, map[string]any{"key": "value", "num": int32(42)}, result)
}

func TestBsonDocToMap_BsonD(t *testing.T) {
	input := bson.D{{Key: "key", Value: "value"}, {Key: "num", Value: int32(42)}}
	result := bsonDocToMap(input)
	require.Equal(t, map[string]any{"key": "value", "num": int32(42)}, result)
}

func TestBsonDocToMap_NestedBsonD(t *testing.T) {
	input := bson.D{
		{Key: "outer", Value: bson.D{
			{Key: "inner_key", Value: "inner_value"},
			{Key: "deep", Value: bson.D{
				{Key: "level3", Value: int32(3)},
			}},
		}},
	}
	result := bsonDocToMap(input)
	require.Equal(t, map[string]any{
		"outer": map[string]any{
			"inner_key": "inner_value",
			"deep": map[string]any{
				"level3": int32(3),
			},
		},
	}, result)
}

func TestBsonDocToMap_NestedBsonA(t *testing.T) {
	input := bson.D{
		{Key: "list", Value: bson.A{
			bson.D{{Key: "name", Value: "first"}},
			bson.D{{Key: "name", Value: "second"}},
		}},
	}
	result := bsonDocToMap(input)
	require.Equal(t, map[string]any{
		"list": []any{
			map[string]any{"name": "first"},
			map[string]any{"name": "second"},
		},
	}, result)
}

func TestBsonDocToMap_BsonMWithNestedBsonD(t *testing.T) {
	input := bson.M{
		"nested": bson.D{{Key: "a", Value: "b"}},
	}
	result := bsonDocToMap(input)
	require.Equal(t, map[string]any{
		"nested": map[string]any{"a": "b"},
	}, result)
}

func TestBsonDocToMap_NilInput(t *testing.T) {
	result := bsonDocToMap(nil)
	require.Nil(t, result)
}

func TestBsonDocToMap_UnsupportedType(t *testing.T) {
	result := bsonDocToMap("not a bson doc")
	require.Nil(t, result)
}

func TestConvertBSONValue_Primitives(t *testing.T) {
	require.Equal(t, "hello", convertBSONValue("hello"))
	require.Equal(t, int32(42), convertBSONValue(int32(42)))
	require.Equal(t, true, convertBSONValue(true))
	require.Nil(t, convertBSONValue(nil))
}

func TestConvertBSONValue_BsonD(t *testing.T) {
	input := bson.D{{Key: "k", Value: "v"}}
	result := convertBSONValue(input)
	require.Equal(t, map[string]any{"k": "v"}, result)
}

func TestConvertBSONValue_BsonA(t *testing.T) {
	input := bson.A{"a", int32(1), bson.D{{Key: "nested", Value: true}}}
	result := convertBSONValue(input)
	expected := []any{"a", int32(1), map[string]any{"nested": true}}
	require.Equal(t, expected, result)
}

func TestReverseTransform_BoolField(t *testing.T) {
	bldr := array.NewBooleanBuilder(memory.DefaultAllocator)
	defer bldr.Release()
	err := (&Client{}).reverseTransform(arrow.Field{Type: arrow.FixedWidthTypes.Boolean}, bldr, true)
	require.NoError(t, err)
	arr := bldr.NewArray()
	defer arr.Release()
	require.Equal(t, 1, arr.Len())
	require.True(t, arr.(*array.Boolean).Value(0))
}

func TestReverseTransform_NullValue(t *testing.T) {
	bldr := array.NewInt32Builder(memory.DefaultAllocator)
	defer bldr.Release()
	err := (&Client{}).reverseTransform(arrow.Field{Type: arrow.PrimitiveTypes.Int32}, bldr, nil)
	require.NoError(t, err)
	arr := bldr.NewArray()
	defer arr.Release()
	require.Equal(t, 1, arr.Len())
	require.True(t, arr.IsNull(0))
}

func TestReverseTransform_BinaryField(t *testing.T) {
	bldr := array.NewBinaryBuilder(memory.DefaultAllocator, arrow.BinaryTypes.Binary)
	defer bldr.Release()
	data := []byte{0x01, 0x02, 0x03}
	err := (&Client{}).reverseTransform(arrow.Field{Type: arrow.BinaryTypes.Binary}, bldr, bson.Binary{Data: data})
	require.NoError(t, err)
	arr := bldr.NewArray()
	defer arr.Release()
	require.Equal(t, 1, arr.Len())
	require.Equal(t, data, arr.(*array.Binary).Value(0))
}

func TestReverseTransform_TimestampMillisecond(t *testing.T) {
	bldr := array.NewTimestampBuilder(memory.DefaultAllocator, &arrow.TimestampType{Unit: arrow.Millisecond})
	defer bldr.Release()
	dt := bson.DateTime(1704067200000) // 2024-01-01T00:00:00Z in milliseconds
	err := (&Client{}).reverseTransform(
		arrow.Field{Type: &arrow.TimestampType{Unit: arrow.Millisecond}},
		bldr, dt,
	)
	require.NoError(t, err)
	arr := bldr.NewArray()
	defer arr.Release()
	require.Equal(t, 1, arr.Len())
	require.Equal(t, arrow.Timestamp(1704067200000), arr.(*array.Timestamp).Value(0))
}

func TestReverseTransform_ListField(t *testing.T) {
	bldr := array.NewListBuilder(memory.DefaultAllocator, arrow.PrimitiveTypes.Int32)
	defer bldr.Release()
	err := (&Client{}).reverseTransform(
		arrow.Field{Type: arrow.ListOf(arrow.PrimitiveTypes.Int32)},
		bldr,
		bson.A{int32(1), int32(2), int32(3)},
	)
	require.NoError(t, err)
	arr := bldr.NewArray()
	defer arr.Release()
	require.Equal(t, 1, arr.Len())
}

func TestReverseTransformer(t *testing.T) {
	table := &schema.Table{
		Name: "test_table",
		Columns: schema.ColumnList{
			{Name: "name", Type: arrow.BinaryTypes.String},
			{Name: "age", Type: arrow.PrimitiveTypes.Int32},
		},
	}

	values := bson.M{
		"name": "Alice",
		"age":  int32(30),
	}

	c := &Client{}
	rec, err := c.reverseTransformer(table, values)
	require.NoError(t, err)
	defer rec.Release()
	require.Equal(t, int64(1), rec.NumRows())
	require.Equal(t, int64(2), rec.NumCols())
	require.Equal(t, "Alice", rec.Column(0).(*array.String).Value(0))
	require.Equal(t, int32(30), rec.Column(1).(*array.Int32).Value(0))
}

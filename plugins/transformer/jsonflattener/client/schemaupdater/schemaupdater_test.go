package schemaupdater

import (
	"bytes"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/require"
)

func TestAddJSONFlattenedFields(t *testing.T) {
	sc := createTestSchema(t)
	updater := New(sc)

	fieldTypeSchemas := map[string]map[string]string{
		"col1": {
			"key_a": "utf8",
			"key_b": "int64",
			"key_c": "bool",
		},
	}

	updatedSchema, err := updater.AddJSONFlattenedFields(fieldTypeSchemas)
	require.NoError(t, err)

	require.Equal(t, 4, updatedSchema.NumFields(), "Expected 3 fields")

	require.Equal(t, "col1", updatedSchema.Field(0).Name, "Expected field name 'col1'")
	require.Equal(t, "json", updatedSchema.Field(0).Type.String(), "Expected type 'json'")

	require.Equal(t, "col1__key_a", updatedSchema.Field(1).Name, "Expected field name 'col1__key_a'")
	require.Equal(t, "utf8", updatedSchema.Field(1).Type.String(), "Expected type 'utf8'")
	require.Equal(t, "col1__key_b", updatedSchema.Field(2).Name, "Expected field name 'col1__key_b'")
	require.Equal(t, "int64", updatedSchema.Field(2).Type.String(), "Expected type 'int64'")
	require.Equal(t, "col1__key_c", updatedSchema.Field(3).Name, "Expected field name 'col1__key_c'")
	require.Equal(t, "bool", updatedSchema.Field(3).Type.String(), "Expected type 'bool'")
}

func createTestSchema(t *testing.T) *arrow.Schema {
	md := arrow.NewMetadata([]string{schema.MetadataTableName}, []string{"testTable"})
	bs := []byte(`{"key_a": "value", "key_b": 2, "key_c": true}`)
	arr := createArray(t, bs)
	fieldMD := map[string]string{
		schema.MetadataTypeSchema: `{"key_a": "utf8", "key_b": "int64", "key_c": "bool"}`,
	}
	return arrow.NewSchema(
		[]arrow.Field{
			{Name: "col1", Type: arr.DataType(), Nullable: true, Metadata: arrow.MetadataFrom(fieldMD)},
		},
		&md,
	)
}

func createArray(t *testing.T, bs []byte) arrow.Array {
	b := types.NewJSONBuilder(memory.NewGoAllocator())
	defer b.Release()
	dec := json.NewDecoder(bytes.NewReader(bs))
	err := b.UnmarshalOne(dec)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	return b.NewArray()
}

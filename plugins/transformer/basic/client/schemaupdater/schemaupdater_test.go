package schemaupdater

import (
	"testing"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/stretchr/testify/require"
)

func TestRemoveColumnIndices(t *testing.T) {
	sc := createTestSchema()
	updater := New(sc)

	colIndices := map[int]struct{}{0: {}}
	updatedSchema := updater.RemoveColumnIndices(colIndices)

	require.Equal(t, 1, updatedSchema.NumFields(), "Expected 1 field")
	require.Equal(t, "col2", updatedSchema.Field(0).Name, "Expected field name 'col2'")
}

func TestAddStringColumnAtPos(t *testing.T) {
	sc := createTestSchema()
	updater := New(sc)

	updatedSchema, err := updater.AddStringColumnAtPos("col3", 1, false)
	require.NoError(t, err)

	require.Equal(t, 3, updatedSchema.NumFields(), "Expected 3 fields")
	require.Equal(t, "col3", updatedSchema.Field(1).Name, "Expected field name 'col3'")
	require.False(t, updatedSchema.Field(1).Nullable, "Expected field 'col3' to be non-nullable")
}

func TestAddStringColumnAtEnd(t *testing.T) {
	sc := createTestSchema()
	updater := New(sc)

	updatedSchema, err := updater.AddStringColumnAtPos("col3", -1, true)
	require.NoError(t, err)

	require.Equal(t, 3, updatedSchema.NumFields(), "Expected 3 fields")
	require.Equal(t, "col3", updatedSchema.Field(2).Name, "Expected field name 'col3'")
	require.True(t, updatedSchema.Field(2).Nullable, "Expected field 'col3' to be nullable")
}

func TestTransformMaintainsMetadata(t *testing.T) {
	md1 := arrow.NewMetadata([]string{"key1", "key2"}, []string{"value1", "value2"})
	md2 := arrow.NewMetadata([]string{"key3", "key4"}, []string{"value3", "value4"})
	md3 := arrow.NewMetadata([]string{"key5", "key6"}, []string{"value5", "value6"})
	sc := arrow.NewSchema(
		[]arrow.Field{
			{Name: "col1", Type: arrow.BinaryTypes.String, Nullable: true, Metadata: md1},
			{Name: "col2", Type: arrow.BinaryTypes.String, Nullable: true, Metadata: md2},
		},
		&md3,
	)
	updater := New(sc)
	updatedSchema, err := updater.AddStringColumnAtPos("col3", -1, true)
	require.NoError(t, err)

	require.Equal(t, sc.Metadata(), updatedSchema.Metadata(), "Expected metadata to be retained")
	require.Equal(t, sc.Field(0).Metadata, updatedSchema.Field(0).Metadata, "Expected metadata to be retained")
	require.Equal(t, sc.Field(1).Metadata, updatedSchema.Field(1).Metadata, "Expected metadata to be retained")
}

func TestChangeTableName(t *testing.T) {
	testSchema := createTestSchema()
	updater := New(testSchema)

	updatedSchema, err := updater.ChangeTableName("cq_sync_{{.OldName}}")
	require.NoError(t, err)

	newTableName, ok := updatedSchema.Metadata().GetValue(schema.MetadataTableName)
	require.True(t, ok, "Expected table name to be present in metadata")
	require.Equal(t, "cq_sync_testTable", newTableName, "Expected table name to be 'cq_sync_testTable'")
	require.Equal(t, testSchema.NumFields(), updatedSchema.NumFields(), "Expected number of fields to remain the same")
	require.Equal(t, testSchema.Field(0).Name, updatedSchema.Field(0).Name, "Expected field name to remain the same")
	require.Equal(t, testSchema.Field(1).Name, updatedSchema.Field(1).Name, "Expected field name to remain the same")
}

func createTestSchema() *arrow.Schema {
	md := arrow.NewMetadata([]string{schema.MetadataTableName}, []string{"testTable"})
	return arrow.NewSchema(
		[]arrow.Field{
			{Name: "col1", Type: arrow.BinaryTypes.String, Nullable: true},
			{Name: "col2", Type: arrow.BinaryTypes.String, Nullable: true},
		},
		&md,
	)
}

package recordupdater

import (
	"bytes"
	"testing"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/apache/arrow/go/v17/arrow/array"
	"github.com/apache/arrow/go/v17/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/require"
)

const sampleRowContent = `{"key_a": "value", "key_b": 2, "key_c": true}`

func TestFlattenJSONFields(t *testing.T) {
	record := createTestRecord(t)
	updater := New(record)

	updatedRecord, err := updater.FlattenJSONFields()
	require.NoError(t, err)

	require.Equal(t, int64(4), updatedRecord.NumCols())
	require.Equal(t, int64(1), updatedRecord.NumRows())
	requireAllColsLenMatchRecordsLen(t, updatedRecord)
	require.Equal(t, "col1__key_a", updatedRecord.ColumnName(1))
	require.Equal(t, "value", updatedRecord.Column(1).(*array.String).Value(0))
	require.Equal(t, "col1__key_b", updatedRecord.ColumnName(2))
	require.Equal(t, int64(2), updatedRecord.Column(2).(*array.Int64).Value(0))
	require.Equal(t, "col1__key_c", updatedRecord.ColumnName(3))
	require.Equal(t, true, updatedRecord.Column(3).(*array.Boolean).Value(0))
}

func createTestRecord(t *testing.T) arrow.Record {
	return array.NewRecord(createTestSchema(t), []arrow.Array{createArray(t, []byte(sampleRowContent))}, 1)
}

func requireAllColsLenMatchRecordsLen(t *testing.T, record arrow.Record) {
	for i := 0; i < int(record.NumCols()); i++ {
		require.Equal(t, int(record.NumRows()), record.Column(i).Len(), "Expected length of %d for column %d", record.NumRows(), i)
	}
}

func createTestSchema(t *testing.T) *arrow.Schema {
	md := arrow.NewMetadata([]string{schema.MetadataTableName}, []string{"testTable"})
	arr := createArray(t, []byte(sampleRowContent))
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
	b := types.NewJSONBuilder(array.NewExtensionBuilder(memory.NewGoAllocator(), types.NewJSONType()))
	defer b.Release()
	dec := json.NewDecoder(bytes.NewReader(bs))
	err := b.UnmarshalOne(dec)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	return b.NewArray()
}

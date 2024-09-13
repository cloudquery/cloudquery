package recordupdater

import (
	"testing"
	"time"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/apache/arrow/go/v17/arrow/array"
	"github.com/apache/arrow/go/v17/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/stretchr/testify/require"
)

func TestRemoveColumns(t *testing.T) {
	record := createTestRecord()
	updater := New(record)

	updatedRecord, err := updater.RemoveColumns([]string{"col1"})
	require.NoError(t, err)

	require.Equal(t, int64(1), updatedRecord.NumCols())
	require.Equal(t, int64(2), updatedRecord.NumRows())
	requireAllColsLenMatchRecordsLen(t, updatedRecord)
	require.Equal(t, "col2", updatedRecord.ColumnName(0))
}

func TestAddLiteralStringColumn(t *testing.T) {
	record := createTestRecord()
	updater := New(record)

	updatedRecord, err := updater.AddLiteralStringColumn("col3", "literal", -1)
	require.NoError(t, err)

	require.Equal(t, int64(3), updatedRecord.NumCols())
	require.Equal(t, int64(2), updatedRecord.NumRows())
	requireAllColsLenMatchRecordsLen(t, updatedRecord)
	require.Equal(t, "col3", updatedRecord.ColumnName(2))
	require.Equal(t, "literal", updatedRecord.Column(2).(*array.String).Value(0))
	require.Equal(t, "literal", updatedRecord.Column(2).(*array.String).Value(1))
}

func TestAddTimestampColumn(t *testing.T) {
	record := createTestRecord()
	updater := New(record)
	initial := time.Now()
	// Sleep to ensure that the timestamp is different, otherwise it fails on GitHub Actions, but succeeds locally
	time.Sleep(10 * time.Millisecond)
	updatedRecord, err := updater.AddTimestampColumn("col3", -1)
	time.Sleep(10 * time.Millisecond)
	after := time.Now()
	require.NoError(t, err)

	require.Equal(t, int64(3), updatedRecord.NumCols())
	require.Equal(t, int64(2), updatedRecord.NumRows())
	requireAllColsLenMatchRecordsLen(t, updatedRecord)
	require.Equal(t, "col3", updatedRecord.ColumnName(2))
	unit := updatedRecord.Column(2).DataType().(*arrow.TimestampType).Unit

	colVal := updatedRecord.Column(2).(*array.Timestamp).Value(0).ToTime(unit).UTC()
	// Check if the timestamp is within the expected range
	require.True(t, colVal.Before(after))
	require.True(t, colVal.After(initial))
}

func TestObfuscateColumns(t *testing.T) {
	record := createTestRecord()
	updater := New(record)

	updatedRecord, err := updater.ObfuscateColumns([]string{"col1"})
	require.NoError(t, err)

	require.Equal(t, int64(2), updatedRecord.NumCols())
	require.Equal(t, int64(2), updatedRecord.NumRows())
	requireAllColsLenMatchRecordsLen(t, updatedRecord)
	require.Equal(t, "col1", updatedRecord.ColumnName(0))
	require.Equal(t, "col2", updatedRecord.ColumnName(1))
	require.Equal(t, "cc1d9c865e8380c2d566dc724c66369051acfaa3e9e8f36ad6c67d7d9b8461a5", updatedRecord.Column(0).(*array.String).Value(0))
	require.Equal(t, "528e5290f8ff0eb0325f0472b9c1a9ef4fac0b02ff6094b64d9382af4a10444b", updatedRecord.Column(0).(*array.String).Value(1))
}

func TestChangeTableName(t *testing.T) {
	record := createTestRecord()
	updater := New(record)

	updatedRecord, err := updater.ChangeTableName("cq_sync_{{.OldName}}")
	require.NoError(t, err)

	require.Equal(t, int64(2), updatedRecord.NumCols())
	require.Equal(t, int64(2), updatedRecord.NumRows())
	requireAllColsLenMatchRecordsLen(t, updatedRecord)

	newTableName, ok := updatedRecord.Schema().Metadata().GetValue(schema.MetadataTableName)
	require.True(t, ok, "Expected table name to be present in metadata")
	require.Equal(t, "cq_sync_testTable", newTableName)
}

func createTestRecord() arrow.Record {
	md := arrow.NewMetadata([]string{schema.MetadataTableName}, []string{"testTable"})
	bld := array.NewRecordBuilder(memory.DefaultAllocator, arrow.NewSchema(
		[]arrow.Field{
			{Name: "col1", Type: arrow.BinaryTypes.String},
			{Name: "col2", Type: arrow.BinaryTypes.String},
		},
		&md,
	))
	defer bld.Release()

	bld.Field(0).(*array.StringBuilder).AppendValues([]string{"val1", "val2"}, nil)
	bld.Field(1).(*array.StringBuilder).AppendValues([]string{"val3", "val4"}, nil)

	return bld.NewRecord()
}

func requireAllColsLenMatchRecordsLen(t *testing.T, record arrow.Record) {
	for i := 0; i < int(record.NumCols()); i++ {
		require.Equal(t, int(record.NumRows()), record.Column(i).Len(), "Expected length of %d for column %d", record.NumRows(), i)
	}
}

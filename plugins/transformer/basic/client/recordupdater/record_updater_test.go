package recordupdater

import (
	"testing"
	"time"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/apache/arrow/go/v17/arrow/array"
	"github.com/apache/arrow/go/v17/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRemoveColumns(t *testing.T) {
	record := createTestRecord()
	updater := New(record)

	updatedRecord, err := updater.RemoveColumns([]string{"col1", "col3.foo.bar.0", "col3.hello"})
	require.NoError(t, err)

	require.Equal(t, int64(2), updatedRecord.NumCols())
	require.Equal(t, int64(2), updatedRecord.NumRows())
	requireAllColsLenMatchRecordsLen(t, updatedRecord)
	require.Equal(t, "col2", updatedRecord.ColumnName(0))
	assert.Equal(t, `{"foo":{"bar":["b","c"]}}`, updatedRecord.Column(1).ValueStr(0))
	assert.Equal(t, `{"foo":{"bar":["e","f"]}}`, updatedRecord.Column(1).ValueStr(1))
}

func TestAddLiteralStringColumn(t *testing.T) {
	record := createTestRecord()
	updater := New(record)

	updatedRecord, err := updater.AddLiteralStringColumn("col4", "literal", -1)
	require.NoError(t, err)

	require.Equal(t, int64(4), updatedRecord.NumCols())
	require.Equal(t, int64(2), updatedRecord.NumRows())
	requireAllColsLenMatchRecordsLen(t, updatedRecord)
	require.Equal(t, "col4", updatedRecord.ColumnName(3))
	require.Equal(t, "literal", updatedRecord.Column(3).(*array.String).Value(0))
	require.Equal(t, "literal", updatedRecord.Column(3).(*array.String).Value(1))
}

func TestAddTimestampColumn(t *testing.T) {
	record := createTestRecord()
	updater := New(record)
	initial := time.Now()
	// Sleep to ensure that the timestamp is different, otherwise it fails on GitHub Actions, but succeeds locally
	time.Sleep(10 * time.Millisecond)
	updatedRecord, err := updater.AddTimestampColumn("col4", -1)
	time.Sleep(10 * time.Millisecond)
	after := time.Now()
	require.NoError(t, err)

	require.Equal(t, int64(4), updatedRecord.NumCols())
	require.Equal(t, int64(2), updatedRecord.NumRows())
	requireAllColsLenMatchRecordsLen(t, updatedRecord)
	require.Equal(t, "col4", updatedRecord.ColumnName(3))
	unit := updatedRecord.Column(3).DataType().(*arrow.TimestampType).Unit

	colVal := updatedRecord.Column(3).(*array.Timestamp).Value(0).ToTime(unit).UTC()
	// Check if the timestamp is within the expected range
	require.True(t, colVal.Before(after))
	require.True(t, colVal.After(initial))
}

func TestObfuscateColumns(t *testing.T) {
	record := createTestRecord()
	updater := New(record)

	updatedRecord, err := updater.ObfuscateColumns([]string{"col1", "col3.foo.bar.0", "col3.foo.bar.1"})
	require.NoError(t, err)

	require.Equal(t, int64(3), updatedRecord.NumCols())
	require.Equal(t, int64(2), updatedRecord.NumRows())
	requireAllColsLenMatchRecordsLen(t, updatedRecord)
	require.Equal(t, "col1", updatedRecord.ColumnName(0))
	require.Equal(t, "col2", updatedRecord.ColumnName(1))
	require.Equal(t, "cc1d9c865e8380c2d566dc724c66369051acfaa3e9e8f36ad6c67d7d9b8461a5", updatedRecord.Column(0).(*array.String).Value(0))
	require.Equal(t, "528e5290f8ff0eb0325f0472b9c1a9ef4fac0b02ff6094b64d9382af4a10444b", updatedRecord.Column(0).(*array.String).Value(1))
	assert.Equal(t, `{"foo":{"bar":["ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb","3e23e8160039594a33894f6564e1b1348bbd7a0088d42c4acb73eeaed59c009d","c"]},"hello":"world"}`, updatedRecord.Column(2).ValueStr(0))
	assert.Equal(t, `{"foo":{"bar":["18ac3e7343f016890c510e93f935261169d9e3f565436429830faf0934f4f8e4","3f79bb7b435b05321651daefd374cdc681dc06faa65e374e38337b88ca046dea","f"]}}`, updatedRecord.Column(2).ValueStr(1))
}

func TestChangeTableName(t *testing.T) {
	record := createTestRecord()
	updater := New(record)

	updatedRecord, err := updater.ChangeTableName("cq_sync_{{.OldName}}")
	require.NoError(t, err)

	require.Equal(t, int64(3), updatedRecord.NumCols())
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
			{Name: "col3", Type: types.NewJSONType()},
		},
		&md,
	))
	defer bld.Release()

	bld.Field(0).(*array.StringBuilder).AppendValues([]string{"val1", "val2"}, nil)
	bld.Field(1).(*array.StringBuilder).AppendValues([]string{"val3", "val4"}, nil)
	bld.Field(2).(*types.JSONBuilder).AppendBytes([]byte(`{"foo":{"bar":["a","b","c"]},"hello":"world"}`))
	bld.Field(2).(*types.JSONBuilder).AppendBytes([]byte(`{"foo":{"bar":["d","e","f"]}}`))

	return bld.NewRecord()
}

func requireAllColsLenMatchRecordsLen(t *testing.T, record arrow.Record) {
	for i := 0; i < int(record.NumCols()); i++ {
		require.Equal(t, int(record.NumRows()), record.Column(i).Len(), "Expected length of %d for column %d", record.NumRows(), i)
	}
}

package recordupdater

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
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
	require.False(t, updatedRecord.Schema().Field(3).Nullable, "Expected column to be non-nullable")
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
	require.Equal(t,
		fmt.Sprintf("%s cc1d9c865e8380c2d566dc724c66369051acfaa3e9e8f36ad6c67d7d9b8461a5", redactedByCQMessage),
		updatedRecord.Column(0).(*array.String).Value(0))
	require.Equal(t,
		fmt.Sprintf("%s 528e5290f8ff0eb0325f0472b9c1a9ef4fac0b02ff6094b64d9382af4a10444b", redactedByCQMessage),
		updatedRecord.Column(0).(*array.String).Value(1))
	assert.Equal(t,
		fmt.Sprintf(`{"foo":{"bar":["%s ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb","%s 3e23e8160039594a33894f6564e1b1348bbd7a0088d42c4acb73eeaed59c009d","c"]},"hello":"world"}`, redactedByCQMessage, redactedByCQMessage),
		updatedRecord.Column(2).ValueStr(0))
	assert.Equal(t,
		fmt.Sprintf(`{"foo":{"bar":["%s 18ac3e7343f016890c510e93f935261169d9e3f565436429830faf0934f4f8e4","%s 3f79bb7b435b05321651daefd374cdc681dc06faa65e374e38337b88ca046dea","f"]}}`, redactedByCQMessage, redactedByCQMessage),
		updatedRecord.Column(2).ValueStr(1))
}

func TestAutoObfuscateColumns(t *testing.T) {
	sc := []string{"col1", "col3.foo.bar.0", "col3.foo.bar.1", "col4"}
	scJSON, err := json.Marshal(sc)
	require.NoError(t, err)
	md := arrow.NewMetadata(
		[]string{schema.MetadataTableName, schema.MetadataTableSensitiveColumns},
		[]string{"testTable", string(scJSON)})
	record := createTestRecordWithMetadata(&md)
	updater := New(record)

	updatedRecord, err := updater.ObfuscateSensitiveColumns()
	require.NoError(t, err)

	require.Equal(t, int64(4), updatedRecord.NumCols())
	require.Equal(t, int64(2), updatedRecord.NumRows())
	requireAllColsLenMatchRecordsLen(t, updatedRecord)
	require.Equal(t, "col1", updatedRecord.ColumnName(0))
	require.Equal(t, "col2", updatedRecord.ColumnName(1))
	require.Equal(t,
		fmt.Sprintf("%s cc1d9c865e8380c2d566dc724c66369051acfaa3e9e8f36ad6c67d7d9b8461a5", redactedByCQMessage),
		updatedRecord.Column(0).(*array.String).Value(0))
	require.Equal(t,
		fmt.Sprintf("%s 528e5290f8ff0eb0325f0472b9c1a9ef4fac0b02ff6094b64d9382af4a10444b", redactedByCQMessage),
		updatedRecord.Column(0).(*array.String).Value(1))
	assert.Equal(t,
		fmt.Sprintf(`{"foo":{"bar":["%s ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb","%s 3e23e8160039594a33894f6564e1b1348bbd7a0088d42c4acb73eeaed59c009d","c"]},"hello":"world"}`, redactedByCQMessage, redactedByCQMessage),
		updatedRecord.Column(2).ValueStr(0))
	assert.Equal(t,
		fmt.Sprintf(`{"foo":{"bar":["%s 18ac3e7343f016890c510e93f935261169d9e3f565436429830faf0934f4f8e4","%s 3f79bb7b435b05321651daefd374cdc681dc06faa65e374e38337b88ca046dea","f"]}}`, redactedByCQMessage, redactedByCQMessage),
		updatedRecord.Column(2).ValueStr(1))
	assert.Equal(t,
		fmt.Sprintf("%s e017a4a3db0b278a196f65c94c2af6c86820a0dd870af9cbb67f4af639085d76", redactedByCQMessage),
		string(updatedRecord.Column(3).(*array.Binary).Value(0)))
	assert.Equal(t,
		fmt.Sprintf("%s aeccf98bf8926a3c2580cb2f4161e0c64379294d30dc2ea7756121b28353b863", redactedByCQMessage),
		string(updatedRecord.Column(3).(*array.Binary).Value(1)))
}

func TestAutoObfuscateEntireJSONColumn(t *testing.T) {
	sc := []string{"col3"}
	scJSON, err := json.Marshal(sc)
	require.NoError(t, err)
	md := arrow.NewMetadata(
		[]string{schema.MetadataTableName, schema.MetadataTableSensitiveColumns},
		[]string{"testTable", string(scJSON)})
	record := createTestRecordWithMetadata(&md)
	updater := New(record)

	updatedRecord, err := updater.ObfuscateSensitiveColumns()
	require.NoError(t, err)

	require.Equal(t, int64(4), updatedRecord.NumCols())
	require.Equal(t, int64(2), updatedRecord.NumRows())
	requireAllColsLenMatchRecordsLen(t, updatedRecord)
	require.Equal(t, "col1", updatedRecord.ColumnName(0))
	require.Equal(t, "col2", updatedRecord.ColumnName(1))
	assert.Equal(t,
		fmt.Sprintf(`{"%s":"81f2a9ddc7ae49a6b585358c6ff54bbd26613c4a46a988b614e42bc5729eda36"}`, redactedByCQJSONName),
		updatedRecord.Column(2).ValueStr(0))
	assert.Equal(t,
		fmt.Sprintf(`{"%s":"b56ea9a87c46567fc64564f68461e8f1068ffa515eee20c3387b97bc17f24cda"}`, redactedByCQJSONName),
		updatedRecord.Column(2).ValueStr(1))
}

func TestAutoObfuscateEntireJSONColumnSkipsJsonPath(t *testing.T) {
	sc := []string{"col3.foo", "col3"}
	scJSON, err := json.Marshal(sc)
	require.NoError(t, err)
	md := arrow.NewMetadata(
		[]string{schema.MetadataTableName, schema.MetadataTableSensitiveColumns},
		[]string{"testTable", string(scJSON)})
	record := createTestRecordWithMetadata(&md)
	updater := New(record)

	updatedRecord, err := updater.ObfuscateSensitiveColumns()
	require.NoError(t, err)

	require.Equal(t, int64(4), updatedRecord.NumCols())
	require.Equal(t, int64(2), updatedRecord.NumRows())
	requireAllColsLenMatchRecordsLen(t, updatedRecord)
	require.Equal(t, "col1", updatedRecord.ColumnName(0))
	require.Equal(t, "col2", updatedRecord.ColumnName(1))
	assert.Equal(t,
		fmt.Sprintf(`{"%s":"81f2a9ddc7ae49a6b585358c6ff54bbd26613c4a46a988b614e42bc5729eda36"}`, redactedByCQJSONName),
		updatedRecord.Column(2).ValueStr(0))
	assert.Equal(t,
		fmt.Sprintf(`{"%s":"b56ea9a87c46567fc64564f68461e8f1068ffa515eee20c3387b97bc17f24cda"}`, redactedByCQJSONName),
		updatedRecord.Column(2).ValueStr(1))
}

func TestRenameColumn(t *testing.T) {
	record := createTestRecord()
	updater := New(record)

	updatedRecord, err := updater.RenameColumn("col1", "newCol1")
	require.NoError(t, err)

	require.Equal(t, int64(3), updatedRecord.NumCols())
	require.Equal(t, int64(2), updatedRecord.NumRows())
	requireAllColsLenMatchRecordsLen(t, updatedRecord)
	require.Equal(t, "newCol1", updatedRecord.ColumnName(0))
	require.Equal(t, "col2", updatedRecord.ColumnName(1))
	require.Equal(t, "col3", updatedRecord.ColumnName(2))
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

func createTestRecordWithMetadata(metadata *arrow.Metadata) arrow.Record {
	bld := array.NewRecordBuilder(memory.DefaultAllocator, arrow.NewSchema(
		[]arrow.Field{
			{Name: "col1", Type: arrow.BinaryTypes.String},
			{Name: "col2", Type: arrow.BinaryTypes.String},
			{Name: "col3", Type: types.NewJSONType()},
			{Name: "col4", Type: &arrow.BinaryType{}},
		},
		metadata,
	))
	defer bld.Release()

	bld.Field(0).(*array.StringBuilder).AppendValues([]string{"val1", "val2"}, nil)
	bld.Field(1).(*array.StringBuilder).AppendValues([]string{"val3", "val4"}, nil)
	bld.Field(2).(*types.JSONBuilder).AppendBytes([]byte(`{"foo":{"bar":["a","b","c"]},"hello":"world"}`))
	bld.Field(2).(*types.JSONBuilder).AppendBytes([]byte(`{"foo":{"bar":["d","e","f"]}}`))
	bld.Field(3).(*array.BinaryBuilder).AppendValues([][]byte{[]byte("val5"), []byte("val6")}, nil)

	return bld.NewRecord()
}

func requireAllColsLenMatchRecordsLen(t *testing.T, record arrow.Record) {
	for i := 0; i < int(record.NumCols()); i++ {
		require.Equal(t, int(record.NumRows()), record.Column(i).Len(), "Expected length of %d for column %d", record.NumRows(), i)
	}
}

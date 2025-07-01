package recordupdater

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/cloudquery/plugins/transformer/basic/client/spec"
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
		fmt.Sprintf("%s cc1d9c865e8380c2d566dc724c66369051acfaa3e9e8f36ad6c67d7d9b8461a5", redactedByCQMessage),
		string(updatedRecord.Column(3).(*array.Binary).Value(0)))
	assert.Equal(t,
		fmt.Sprintf("%s 44a036a895f1f40e3bf8cf930f287edc1cf0a0d0c75b36d1d25b777577f37e7e", redactedByCQMessage),
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

func TestDropRow(t *testing.T) {
	record := createTestRecord()
	updater := New(record)

	updatedRecord, err := updater.DropRows([]string{"col1"}, &[]string{"val1"}[0])
	require.NoError(t, err)

	require.Equal(t, int64(3), updatedRecord.NumCols())
	require.Equal(t, int64(1), updatedRecord.NumRows())
	requireAllColsLenMatchRecordsLen(t, updatedRecord)
	require.Equal(t, "col1", updatedRecord.ColumnName(0))
	require.Equal(t, "col2", updatedRecord.ColumnName(1))
	require.Equal(t, "val2", updatedRecord.Column(0).(*array.String).Value(0))
	assert.Equal(t, `{"foo":{"bar":["d","e","f"]}}`, updatedRecord.Column(2).ValueStr(0))
}

func TestDropRowTimestamp(t *testing.T) {
	record := createTestRecordWithTS()
	updater := New(record)
	updatedRecord, err := updater.DropRows([]string{"col4"}, &[]string{"2025-06-27 10:40:35Z"}[0])
	require.NoError(t, err)

	require.Equal(t, int64(4), updatedRecord.NumCols())
	require.Equal(t, int64(1), updatedRecord.NumRows())
	requireAllColsLenMatchRecordsLen(t, updatedRecord)
	require.Equal(t, "2026-01-01 00:00:00Z", updatedRecord.Column(3).(*array.Timestamp).ValueStr(0))
}

func TestComprehensiveDropRow(t *testing.T) {
	table := schema.TestTable("test_drop_row", schema.TestSourceOptions{})
	tg := schema.NewTestDataGenerator(5)
	record := tg.Generate(table, schema.GenTestDataOptions{
		MaxRows:    10,
		StableTime: time.Date(2025, 6, 27, 10, 40, 35, 914319, time.UTC),
	})
	updater := New(record)
	updatedRecord, err := updater.DropRows([]string{"uuid"}, &[]string{"3831f26b-7a87-577a-ba61-77c84f262922"}[0])
	require.NoError(t, err)
	require.Equal(t, "dae677ed-5012-5bc8-8067-a8374a14edfa", updatedRecord.Column(14).(*types.UUIDArray).ValueStr(0))
	require.Equal(t, int64(9), updatedRecord.NumRows())

	updatedRecord, err = updater.DropRows([]string{"mac"}, &[]string{"a6:ae:92:fb:b5:2c"}[0])
	require.NoError(t, err)
	require.Equal(t, int64(8), updatedRecord.NumRows())
	require.Equal(t, "aa:f1:cb:2e:55:8f", updatedRecord.Column(16).(*types.MACArray).ValueStr(0))

	updatedRecord, err = updater.DropRows([]string{"inet"}, &[]string{"139.0.16.60/10"}[0])
	require.NoError(t, err)
	require.Equal(t, int64(7), updatedRecord.NumRows())
	require.Equal(t, "30.233.221.51/25", updatedRecord.Column(15).(*types.InetArray).ValueStr(0))

	updatedRecord, err = updater.DropRows([]string{"json"}, &[]string{`{"test":["a","b",52011]}`}[0])
	require.NoError(t, err)
	require.Equal(t, `{"test":["a","b",16309]}`, updatedRecord.Column(17).(*types.JSONArray).ValueStr(0))
	require.Equal(t, int64(6), updatedRecord.NumRows())

	updatedRecord, err = updater.DropRows([]string{"uint64"}, &[]string{"1492571184685610752"}[0])
	require.NoError(t, err)
	require.Equal(t, `4019863684675753984`, updatedRecord.Column(8).(*array.Uint64).ValueStr(0))
	require.Equal(t, int64(5), updatedRecord.NumRows())

	updatedRecord, err = updater.DropRows([]string{"date64"}, &[]string{"2023-06-12"}[0])
	require.NoError(t, err)
	require.Equal(t, `2023-04-25`, updatedRecord.Column(19).(*array.Date64).ValueStr(0))
	require.Equal(t, int64(4), updatedRecord.NumRows())

	updatedRecord, err = updater.DropRows([]string{"timestamp_ns"}, &[]string{"2025-06-27 10:40:35.000914Z"}[0])
	require.NoError(t, err)
	require.Equal(t, int64(0), updatedRecord.NumRows())
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

func createTestRecordWithTS() arrow.Record {
	md := arrow.NewMetadata([]string{schema.MetadataTableName}, []string{"testTable"})
	bld := array.NewRecordBuilder(memory.DefaultAllocator, arrow.NewSchema(
		[]arrow.Field{
			{Name: "col1", Type: arrow.BinaryTypes.String},
			{Name: "col2", Type: arrow.BinaryTypes.String},
			{Name: "col3", Type: types.NewJSONType()},
			{Name: "col4", Type: &arrow.TimestampType{}},
		},
		&md,
	))
	defer bld.Release()

	bld.Field(0).(*array.StringBuilder).AppendValues([]string{"val1", "val2"}, nil)
	bld.Field(1).(*array.StringBuilder).AppendValues([]string{"val3", "val4"}, nil)
	bld.Field(2).(*types.JSONBuilder).AppendBytes([]byte(`{"foo":{"bar":["a","b","c"]},"hello":"world"}`))
	bld.Field(2).(*types.JSONBuilder).AppendBytes([]byte(`{"foo":{"bar":["d","e","f"]}}`))
	bld.Field(3).(*array.TimestampBuilder).AppendTime(time.Date(2025, 6, 27, 10, 40, 35, 914319000, time.UTC))
	bld.Field(3).(*array.TimestampBuilder).AppendTime(time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC))

	return bld.NewRecord()
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
	bld.Field(3).(*array.BinaryBuilder).AppendValues([][]byte{[]byte("val1"), []byte("val5")}, nil)

	return bld.NewRecord()
}

func requireAllColsLenMatchRecordsLen(t *testing.T, record arrow.Record) {
	for i := 0; i < int(record.NumCols()); i++ {
		require.Equal(t, int(record.NumRows()), record.Column(i).Len(), "Expected length of %d for column %d", record.NumRows(), i)
	}
}

func TestChangeCaseStringTransformations(t *testing.T) {
	record := createTestRecord()
	updater := New(record)

	updatedRecord, err := updater.ChangeCase(spec.KindUppercase, []string{"col1", "col2"})
	require.NoError(t, err)

	require.Equal(t, int64(3), updatedRecord.NumCols())
	require.Equal(t, int64(2), updatedRecord.NumRows())
	requireAllColsLenMatchRecordsLen(t, updatedRecord)
	require.Equal(t, "col1", updatedRecord.ColumnName(0))
	require.Equal(t, "col2", updatedRecord.ColumnName(1))
	require.Equal(t, "VAL1", updatedRecord.Column(0).(*array.String).Value(0))
	require.Equal(t, "VAL2", updatedRecord.Column(0).(*array.String).Value(1))
	require.Equal(t, "VAL3", updatedRecord.Column(1).(*array.String).Value(0))
	require.Equal(t, "VAL4", updatedRecord.Column(1).(*array.String).Value(1))

	updatedRecord, err = updater.ChangeCase(spec.KindLowercase, []string{"col1", "col2"})
	require.NoError(t, err)
	require.Equal(t, int64(3), updatedRecord.NumCols())
	require.Equal(t, int64(2), updatedRecord.NumRows())
	requireAllColsLenMatchRecordsLen(t, updatedRecord)
	require.Equal(t, "col1", updatedRecord.ColumnName(0))
	require.Equal(t, "col2", updatedRecord.ColumnName(1))
	require.Equal(t, "val1", updatedRecord.Column(0).(*array.String).Value(0))
	require.Equal(t, "val2", updatedRecord.Column(0).(*array.String).Value(1))
	require.Equal(t, "val3", updatedRecord.Column(1).(*array.String).Value(0))
	require.Equal(t, "val4", updatedRecord.Column(1).(*array.String).Value(1))
}

func TestChangeCaseJsonPath(t *testing.T) {
	record := createTestRecord()
	updater := New(record)

	updatedRecord, err := updater.ChangeCase(spec.KindUppercase, []string{"col3.foo.bar.0", "col3.foo.bar.1", "col3.hello"})
	require.NoError(t, err)

	require.Equal(t, int64(3), updatedRecord.NumCols())
	require.Equal(t, int64(2), updatedRecord.NumRows())
	requireAllColsLenMatchRecordsLen(t, updatedRecord)
	require.Equal(t, "col1", updatedRecord.ColumnName(0))
	require.Equal(t, "col2", updatedRecord.ColumnName(1))
	assert.Equal(t, `{"foo":{"bar":["A","B","c"]},"hello":"WORLD"}`, updatedRecord.Column(2).ValueStr(0))
	assert.Equal(t, `{"foo":{"bar":["D","E","f"]}}`, updatedRecord.Column(2).ValueStr(1))
	require.Equal(t, "val1", updatedRecord.Column(0).(*array.String).Value(0))
	require.Equal(t, "val2", updatedRecord.Column(0).(*array.String).Value(1))
	require.Equal(t, "val3", updatedRecord.Column(1).(*array.String).Value(0))
	require.Equal(t, "val4", updatedRecord.Column(1).(*array.String).Value(1))

	updatedRecord, err = updater.ChangeCase(spec.KindLowercase, []string{"col3.foo.bar.0", "col3.foo.bar.1", "col3.hello"})
	require.NoError(t, err)
	require.Equal(t, int64(3), updatedRecord.NumCols())
	require.Equal(t, int64(2), updatedRecord.NumRows())
	requireAllColsLenMatchRecordsLen(t, updatedRecord)
	require.Equal(t, "col1", updatedRecord.ColumnName(0))
	require.Equal(t, "col2", updatedRecord.ColumnName(1))
	assert.Equal(t, `{"foo":{"bar":["a","b","c"]},"hello":"world"}`, updatedRecord.Column(2).ValueStr(0))
	assert.Equal(t, `{"foo":{"bar":["d","e","f"]}}`, updatedRecord.Column(2).ValueStr(1))
	require.Equal(t, "val1", updatedRecord.Column(0).(*array.String).Value(0))
	require.Equal(t, "val2", updatedRecord.Column(0).(*array.String).Value(1))
	require.Equal(t, "val3", updatedRecord.Column(1).(*array.String).Value(0))
	require.Equal(t, "val4", updatedRecord.Column(1).(*array.String).Value(1))
}

func TestChangeCaseEntireJson(t *testing.T) {
	record := createTestRecord()
	updater := New(record)
	updatedRecord, err := updater.ChangeCase(spec.KindUppercase, []string{"col3"})
	require.NoError(t, err)
	require.Equal(t, int64(3), updatedRecord.NumCols())
	require.Equal(t, int64(2), updatedRecord.NumRows())
	requireAllColsLenMatchRecordsLen(t, updatedRecord)
	require.Equal(t, "col1", updatedRecord.ColumnName(0))
	require.Equal(t, "col2", updatedRecord.ColumnName(1))
	assert.Equal(t, `{"FOO":{"BAR":["A","B","C"]},"HELLO":"WORLD"}`, updatedRecord.Column(2).ValueStr(0))
	assert.Equal(t, `{"FOO":{"BAR":["D","E","F"]}}`, updatedRecord.Column(2).ValueStr(1))
	require.Equal(t, "val1", updatedRecord.Column(0).(*array.String).Value(0))
	require.Equal(t, "val2", updatedRecord.Column(0).(*array.String).Value(1))
	require.Equal(t, "val3", updatedRecord.Column(1).(*array.String).Value(0))
	require.Equal(t, "val4", updatedRecord.Column(1).(*array.String).Value(1))

	updatedRecord, err = updater.ChangeCase(spec.KindLowercase, []string{"col3"})
	require.NoError(t, err)
	require.Equal(t, int64(3), updatedRecord.NumCols())
	require.Equal(t, int64(2), updatedRecord.NumRows())
	requireAllColsLenMatchRecordsLen(t, updatedRecord)
	require.Equal(t, "col1", updatedRecord.ColumnName(0))
	require.Equal(t, "col2", updatedRecord.ColumnName(1))
	assert.Equal(t, `{"foo":{"bar":["a","b","c"]},"hello":"world"}`, updatedRecord.Column(2).ValueStr(0))
	assert.Equal(t, `{"foo":{"bar":["d","e","f"]}}`, updatedRecord.Column(2).ValueStr(1))
	require.Equal(t, "val1", updatedRecord.Column(0).(*array.String).Value(0))
	require.Equal(t, "val2", updatedRecord.Column(0).(*array.String).Value(1))
	require.Equal(t, "val3", updatedRecord.Column(1).(*array.String).Value(0))
	require.Equal(t, "val4", updatedRecord.Column(1).(*array.String).Value(1))
}

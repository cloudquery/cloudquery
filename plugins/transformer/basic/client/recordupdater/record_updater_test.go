package recordupdater

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/cloudquery/plugins/transformer/basic/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/scalar"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
	"github.com/tidwall/pretty"
)

func sortJSON(jsonStr string) string {
	opts := pretty.Options{SortKeys: true}
	return string(pretty.Ugly(pretty.PrettyOptions([]byte(jsonStr), &opts)))
}

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

	updatedRecord, err := updater.ObfuscateColumns([]string{"col1", "col3.foo.bar.0", "col3.foo.bar.1"}, true)
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
		fmt.Sprintf(`{"foo":{"bar":["%s ac8d8342bbb2362d13f0a559a3621bb407011368895164b628a54f7fc33fc43c","%s c100f95c1913f9c72fc1f4ef0847e1e723ffe0bde0b36e5f36c13f81fe8c26ed","c"]},"hello":"world"}`, redactedByCQMessage, redactedByCQMessage),
		updatedRecord.Column(2).ValueStr(0))
	assert.Equal(t,
		fmt.Sprintf(`{"foo":{"bar":["%s 3fa5834dc920d385ca9b099c9fe55dcca163a6b256a261f8f147291b0e7cf633","%s 8c8656c5d114d7f8b2a412d2d5fd03accce3ed050624a0493734591a9666b110","f"]}}`, redactedByCQMessage, redactedByCQMessage),
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

	updatedRecord, err := updater.ObfuscateSensitiveColumns(true)
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
		fmt.Sprintf(`{"foo":{"bar":["%s ac8d8342bbb2362d13f0a559a3621bb407011368895164b628a54f7fc33fc43c","%s c100f95c1913f9c72fc1f4ef0847e1e723ffe0bde0b36e5f36c13f81fe8c26ed","c"]},"hello":"world"}`, redactedByCQMessage, redactedByCQMessage),
		updatedRecord.Column(2).ValueStr(0))
	assert.Equal(t,
		fmt.Sprintf(`{"foo":{"bar":["%s 3fa5834dc920d385ca9b099c9fe55dcca163a6b256a261f8f147291b0e7cf633","%s 8c8656c5d114d7f8b2a412d2d5fd03accce3ed050624a0493734591a9666b110","f"]}}`, redactedByCQMessage, redactedByCQMessage),
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

	updatedRecord, err := updater.ObfuscateSensitiveColumns(true)
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

	expectedTimestamp := record.Column(3).(*array.Timestamp).Value(1)
	require.Equal(t, expectedTimestamp, updatedRecord.Column(3).(*array.Timestamp).Value(0))
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

	updatedRecord, err = updater.DropRows([]string{"timestamp_ns"}, &[]string{"2025-06-27T10:40:35.000914Z"}[0])
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

	updatedRecord, err := updater.ObfuscateSensitiveColumns(true)
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

func createTestRecordWithTS() arrow.RecordBatch {
	md := arrow.NewMetadata([]string{schema.MetadataTableName}, []string{"testTable"})
	s := arrow.NewSchema(
		[]arrow.Field{
			{Name: "col1", Type: arrow.BinaryTypes.String},
			{Name: "col2", Type: arrow.BinaryTypes.String},
			{Name: "col3", Type: types.NewJSONType()},
			{Name: "col4", Type: &arrow.TimestampType{}},
		},
		&md,
	)

	col1Builder := array.NewStringBuilder(memory.DefaultAllocator)
	col1Builder.AppendValues([]string{"val1", "val2"}, nil)

	col2Builder := array.NewStringBuilder(memory.DefaultAllocator)
	col2Builder.AppendValues([]string{"val3", "val4"}, nil)

	col3Builder := types.NewJSONBuilder(memory.DefaultAllocator)
	col3Builder.AppendBytes([]byte(`{"foo":{"bar":["a","b","c"]},"hello":"world"}`))
	col3Builder.AppendBytes([]byte(`{"foo":{"bar":["d","e","f"]}}`))

	col4Builder := array.NewTimestampBuilderWithValueStrLayout(memory.DefaultAllocator, &arrow.TimestampType{}, scalar.TimestampStringLayout)
	col4Builder.AppendTime(time.Date(2025, 6, 27, 10, 40, 35, 914319000, time.UTC))
	col4Builder.AppendTime(time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC))

	values := []arrow.Array{
		col1Builder.NewArray(),
		col2Builder.NewArray(),
		col3Builder.NewArray(),
		col4Builder.NewArray(),
	}

	return array.NewRecordBatch(s, values, int64(2))
}

func createTestRecord() arrow.RecordBatch {
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

	return bld.NewRecordBatch()
}

func createTestRecordWithMetadata(metadata *arrow.Metadata) arrow.RecordBatch {
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

	return bld.NewRecordBatch()
}

func requireAllColsLenMatchRecordsLen(t *testing.T, record arrow.RecordBatch) {
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

func TestObfuscateNestedColumnsWithGjsonSyntax(t *testing.T) {
	// Create test record with nested JSON structure
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
	bld.Field(2).(*types.JSONBuilder).AppendBytes([]byte(`{"top_foo":[{"foo":"baz0"},{"foo":"baz1"},{"foo":"baz2"}]}`))
	bld.Field(2).(*types.JSONBuilder).AppendBytes([]byte(`{"top_foo":[{"foo":"baz3"},{"foo":"baz4"},{"foo":"baz5"}]}`))

	record := bld.NewRecordBatch()
	updater := New(record)

	// Test obfuscation using gjson syntax with # for array elements
	updatedRecord, err := updater.ObfuscateColumns([]string{"col3.top_foo.#.foo"}, true)
	require.NoError(t, err)

	require.Equal(t, int64(3), updatedRecord.NumCols())
	require.Equal(t, int64(2), updatedRecord.NumRows())
	requireAllColsLenMatchRecordsLen(t, updatedRecord)

	// Check that the nested foo values are obfuscated
	col3Val := updatedRecord.Column(2).ValueStr(0)
	require.Contains(t, col3Val, redactedByCQMessage, "Expected obfuscated values to contain redacted message")
	require.Contains(t, col3Val, "top_foo", "Expected top_foo structure to be maintained")
	// Verify that all three "foo" values in the array are obfuscated
	require.Equal(t, 3, strings.Count(col3Val, redactedByCQMessage), "Expected 3 obfuscated values for the 3 foo items")

	// Check second row as well
	col3Val2 := updatedRecord.Column(2).ValueStr(1)
	require.Contains(t, col3Val2, redactedByCQMessage, "Expected obfuscated values to contain redacted message")
	require.Equal(t, 3, strings.Count(col3Val2, redactedByCQMessage), "Expected 3 obfuscated values for the 3 foo items")
}

func TestObfuscateDeeplyNestedColumnsWithGjsonSyntax(t *testing.T) {
	// Create test record with deeply nested JSON structure
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
	// First row: has 2 objects in object2 array, each with 2 nested2_object1 values = 4 total
	bld.Field(2).(*types.JSONBuilder).AppendBytes([]byte(`{"object1":{"object2":[{"nested_object1":{"nested_object2":[{"nested2_object1":1},{"nested2_object1":2}]}},{"nested_object1":{"nested_object2":[{"nested2_object1":3},{"nested2_object1":4}]}}]}}`))
	// Second row: has 1 object in object2 array, with 2 nested2_object1 values = 2 total
	bld.Field(2).(*types.JSONBuilder).AppendBytes([]byte(`{"object1":{"object2":[{"nested_object1":{"nested_object2":[{"nested2_object1":5},{"nested2_object1":6}]}}]}}`))

	record := bld.NewRecordBatch()
	updater := New(record)

	// Test obfuscation using gjson syntax with multiple # for nested arrays
	updatedRecord, err := updater.ObfuscateColumns([]string{"col3.object1.object2.#.nested_object1.nested_object2.#.nested2_object1"}, true)
	require.NoError(t, err)

	require.Equal(t, int64(3), updatedRecord.NumCols())
	require.Equal(t, int64(2), updatedRecord.NumRows())
	requireAllColsLenMatchRecordsLen(t, updatedRecord)

	// Check first row: should have 4 obfuscated values
	col3Val := updatedRecord.Column(2).ValueStr(0)
	require.Contains(t, col3Val, redactedByCQMessage, "Expected obfuscated values to contain redacted message")
	require.Contains(t, col3Val, "object1", "Expected object1 structure to be maintained")
	require.Contains(t, col3Val, "nested_object1", "Expected nested_object1 structure to be maintained")
	require.Equal(t, 4, strings.Count(col3Val, redactedByCQMessage), "Expected 4 obfuscated values for the 4 nested2_object1 items in first row")

	// Check second row: should have 2 obfuscated values
	col3Val2 := updatedRecord.Column(2).ValueStr(1)
	require.Contains(t, col3Val2, redactedByCQMessage, "Expected obfuscated values to contain redacted message")
	require.Equal(t, 2, strings.Count(col3Val2, redactedByCQMessage), "Expected 2 obfuscated values for the 2 nested2_object1 items in second row")
}

func TestRemoveNestedColumnsWithGjsonSyntax(t *testing.T) {
	// Create test record with nested JSON structure
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
	bld.Field(2).(*types.JSONBuilder).AppendBytes([]byte(`{"top_foo":[{"foo":"baz0","keep":"value0"},{"foo":"baz1","keep":"value1"},{"foo":"baz2","keep":"value2"}],"other":"data"}`))
	bld.Field(2).(*types.JSONBuilder).AppendBytes([]byte(`{"top_foo":[{"foo":"baz3","keep":"value3"},{"foo":"baz4","keep":"value4"},{"foo":"baz5","keep":"value5"}],"other":"data"}`))

	record := bld.NewRecordBatch()
	updater := New(record)

	// Test removal using gjson syntax with # for array elements
	updatedRecord, err := updater.RemoveColumns([]string{"col3.top_foo.#.foo"})
	require.NoError(t, err)

	require.Equal(t, int64(3), updatedRecord.NumCols())
	require.Equal(t, int64(2), updatedRecord.NumRows())
	requireAllColsLenMatchRecordsLen(t, updatedRecord)

	// Check that the nested foo values are removed but keep values remain
	expectedJSON1 := `{"top_foo":[{"keep":"value0"},{"keep":"value1"},{"keep":"value2"}],"other":"data"}`
	actualJSON1 := updatedRecord.Column(2).ValueStr(0)
	require.Equal(t, sortJSON(expectedJSON1), sortJSON(actualJSON1), "Expected foo fields to be removed from first row")

	expectedJSON2 := `{"top_foo":[{"keep":"value3"},{"keep":"value4"},{"keep":"value5"}],"other":"data"}`
	actualJSON2 := updatedRecord.Column(2).ValueStr(1)
	require.Equal(t, sortJSON(expectedJSON2), sortJSON(actualJSON2), "Expected foo fields to be removed from second row")
}

func TestRemoveDeeplyNestedColumnsWithGjsonSyntax(t *testing.T) {
	// Create test record with deeply nested JSON structure
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
	// First row: has 2 objects in object2 array, each with 2 nested2_object1 values = 4 total
	bld.Field(2).(*types.JSONBuilder).AppendBytes([]byte(`{"object1":{"object2":[{"nested_object1":{"nested_object2":[{"nested2_object1":1,"keep":"a"},{"nested2_object1":2,"keep":"b"}]}},{"nested_object1":{"nested_object2":[{"nested2_object1":3,"keep":"c"},{"nested2_object1":4,"keep":"d"}]}}]}}`))
	// Second row: has 1 object in object2 array, with 2 nested2_object1 values = 2 total
	bld.Field(2).(*types.JSONBuilder).AppendBytes([]byte(`{"object1":{"object2":[{"nested_object1":{"nested_object2":[{"nested2_object1":5,"keep":"e"},{"nested2_object1":6,"keep":"f"}]}}]}}`))

	record := bld.NewRecordBatch()
	updater := New(record)

	// Test removal using gjson syntax with multiple # for nested arrays
	updatedRecord, err := updater.RemoveColumns([]string{"col3.object1.object2.#.nested_object1.nested_object2.#.nested2_object1"})
	require.NoError(t, err)

	require.Equal(t, int64(3), updatedRecord.NumCols())
	require.Equal(t, int64(2), updatedRecord.NumRows())
	requireAllColsLenMatchRecordsLen(t, updatedRecord)

	// Check first row: nested2_object1 values should be removed but keep values should remain
	expectedJSON1 := `{"object1":{"object2":[{"nested_object1":{"nested_object2":[{"keep":"a"},{"keep":"b"}]}},{"nested_object1":{"nested_object2":[{"keep":"c"},{"keep":"d"}]}}]}}`
	actualJSON1 := updatedRecord.Column(2).ValueStr(0)
	require.Equal(t, sortJSON(expectedJSON1), sortJSON(actualJSON1), "Expected nested2_object1 fields to be removed from first row")

	// Check second row: nested2_object1 values should be removed but keep values should remain
	expectedJSON2 := `{"object1":{"object2":[{"nested_object1":{"nested_object2":[{"keep":"e"},{"keep":"f"}]}}]}}`
	actualJSON2 := updatedRecord.Column(2).ValueStr(1)
	require.Equal(t, sortJSON(expectedJSON2), sortJSON(actualJSON2), "Expected nested2_object1 fields to be removed from second row")
}

func TestRemoveNestedArrayWithGjsonSyntax(t *testing.T) {
	// Create test record with nested array structure like user described
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
	// Test structure: [{"env": [{"name": "AWS_ACCESS_KEY_ID", "value": "test"}]}]
	bld.Field(2).(*types.JSONBuilder).AppendBytes([]byte(`[{"env": [{"name": "AWS_ACCESS_KEY_ID", "value": "test"}, {"name": "AWS_SECRET_KEY", "value": "secret"}]}, {"env": [{"name": "DB_PASSWORD", "value": "password"}]}]`))
	bld.Field(2).(*types.JSONBuilder).AppendBytes([]byte(`[{"env": [{"name": "API_KEY", "value": "api-key-value"}]}]`))

	record := bld.NewRecordBatch()
	updater := New(record)

	// Test removal using gjson syntax: #.env.#.value (remove all "value" fields from nested env arrays)
	updatedRecord, err := updater.RemoveColumns([]string{"col3.#.env.#.value"})
	require.NoError(t, err)

	require.Equal(t, int64(3), updatedRecord.NumCols())
	require.Equal(t, int64(2), updatedRecord.NumRows())
	requireAllColsLenMatchRecordsLen(t, updatedRecord)

	// Check first row: "value" fields should be removed but "name" fields should remain
	expectedJSON1 := `[{"env": [{"name": "AWS_ACCESS_KEY_ID"}, {"name": "AWS_SECRET_KEY"}]}, {"env": [{"name": "DB_PASSWORD"}]}]`
	actualJSON1 := updatedRecord.Column(2).ValueStr(0)
	require.Equal(t, sortJSON(expectedJSON1), sortJSON(actualJSON1), "Expected value fields to be removed from first row")

	// Check second row: "value" field should be removed but "name" field should remain
	expectedJSON2 := `[{"env": [{"name": "API_KEY"}]}]`
	actualJSON2 := updatedRecord.Column(2).ValueStr(1)
	require.Equal(t, sortJSON(expectedJSON2), sortJSON(actualJSON2), "Expected value fields to be removed from second row")
}

func createTestRecordWithCQ() arrow.RecordBatch {
	md := arrow.NewMetadata([]string{schema.MetadataTableName}, []string{"testTable"})
	bld := array.NewRecordBuilder(memory.DefaultAllocator, arrow.NewSchema(
		[]arrow.Field{
			{Name: "_cq_id", Type: arrow.BinaryTypes.String},
			{Name: "_cq_sync_time", Type: arrow.BinaryTypes.String},
			{Name: "name", Type: arrow.BinaryTypes.String},
			{Name: "secret", Type: arrow.BinaryTypes.String},
		},
		&md,
	))
	defer bld.Release()

	bld.Field(0).(*array.StringBuilder).AppendValues([]string{"id1", "id2"}, nil)
	bld.Field(1).(*array.StringBuilder).AppendValues([]string{"t1", "t2"}, nil)
	bld.Field(2).(*array.StringBuilder).AppendValues([]string{"n1", "n2"}, nil)
	bld.Field(3).(*array.StringBuilder).AppendValues([]string{"s1", "s2"}, nil)

	return bld.NewRecordBatch()
}

func createTestRecordWithPK() arrow.RecordBatch {
	md := arrow.NewMetadata([]string{schema.MetadataTableName}, []string{"testTable"})
	pkMeta := arrow.NewMetadata([]string{schema.MetadataPrimaryKey}, []string{schema.MetadataTrue})
	bld := array.NewRecordBuilder(memory.DefaultAllocator, arrow.NewSchema(
		[]arrow.Field{
			{Name: "_cq_id", Type: arrow.BinaryTypes.String},
			{Name: "uid", Type: arrow.BinaryTypes.String, Metadata: pkMeta},
			{Name: "name", Type: arrow.BinaryTypes.String},
			{Name: "secret", Type: arrow.BinaryTypes.String},
		},
		&md,
	))
	defer bld.Release()

	bld.Field(0).(*array.StringBuilder).AppendValues([]string{"cq1", "cq2"}, nil)
	bld.Field(1).(*array.StringBuilder).AppendValues([]string{"uid-1", "uid-2"}, nil)
	bld.Field(2).(*array.StringBuilder).AppendValues([]string{"pod-a", "pod-b"}, nil)
	bld.Field(3).(*array.StringBuilder).AppendValues([]string{"s1", "s2"}, nil)

	return bld.NewRecordBatch()
}

func TestObfuscateColumnsExcept_UnknownKeepColumnErrors(t *testing.T) {
	record := createTestRecord()
	_, err := New(record).ObfuscateColumnsExcept([]string{"col1", "does_not_exist"}, true, "redact")
	require.Error(t, err)
	require.Contains(t, err.Error(), "does_not_exist")
}

func TestObfuscateColumnsExcept_NestedPathOnNonJSONErrors(t *testing.T) {
	record := createTestRecord()
	_, err := New(record).ObfuscateColumnsExcept([]string{"col1.foo"}, true, "redact")
	require.Error(t, err)
	require.Contains(t, err.Error(), "col1")
}

func TestObfuscateColumnsExcept_PKNotAllowlistedIncludeSHAFalseErrors(t *testing.T) {
	record := createTestRecordWithPK()
	_, err := New(record).ObfuscateColumnsExcept([]string{"name"}, false, "redact")
	require.Error(t, err)
	require.Contains(t, err.Error(), "uid")
	require.Contains(t, err.Error(), "include_sha")
}

func TestObfuscateColumnsExcept_PKNotAllowlistedIncludeSHATrueKeepsDistinct(t *testing.T) {
	record := createTestRecordWithPK()
	updated, err := New(record).ObfuscateColumnsExcept([]string{"name"}, true, "redact")
	require.NoError(t, err)
	uid0 := updated.Column(1).(*array.String).Value(0)
	uid1 := updated.Column(1).(*array.String).Value(1)
	require.True(t, strings.HasPrefix(uid0, redactedByCQMessage))
	require.NotEqual(t, uid0, uid1, "redacted PK must remain distinct with include_sha=true")
}

func TestObfuscateColumnsExcept_PKAllowlistedIncludeSHAFalseOK(t *testing.T) {
	record := createTestRecordWithPK()
	updated, err := New(record).ObfuscateColumnsExcept([]string{"uid"}, false, "redact")
	require.NoError(t, err)
	require.Equal(t, "uid-1", updated.Column(1).(*array.String).Value(0))
	require.Equal(t, redactedByCQMessageNoSHA, updated.Column(2).(*array.String).Value(0))
}

func TestObfuscateColumnsExcept_DropEveryColumnErrors(t *testing.T) {
	md := arrow.NewMetadata([]string{schema.MetadataTableName}, []string{"testTable"})
	bld := array.NewRecordBuilder(memory.DefaultAllocator, arrow.NewSchema(
		[]arrow.Field{{Name: "count", Type: arrow.PrimitiveTypes.Int64}}, &md,
	))
	defer bld.Release()
	bld.Field(0).(*array.Int64Builder).AppendValues([]int64{1, 2}, nil)
	record := bld.NewRecordBatch()

	_, err := New(record).ObfuscateColumnsExcept([]string{}, true, "redact")
	require.Error(t, err)
	require.Contains(t, err.Error(), "drop every column")
}

func TestObfuscateColumnsExcept_TopLevelKeepRedact(t *testing.T) {
	record := createTestRecord()
	updater := New(record)

	updatedRecord, err := updater.ObfuscateColumnsExcept([]string{"col1"}, true, "redact")
	require.NoError(t, err)

	require.Equal(t, int64(3), updatedRecord.NumCols())
	require.Equal(t, int64(2), updatedRecord.NumRows())
	requireAllColsLenMatchRecordsLen(t, updatedRecord)

	require.Equal(t, "val1", updatedRecord.Column(0).(*array.String).Value(0))
	require.Equal(t, "val2", updatedRecord.Column(0).(*array.String).Value(1))
	require.Equal(t, redactValue([]byte("val3"), true), updatedRecord.Column(1).(*array.String).Value(0))
	require.Equal(t, redactValue([]byte("val4"), true), updatedRecord.Column(1).(*array.String).Value(1))
	require.Equal(t,
		fmt.Sprintf(`{"%s":"81f2a9ddc7ae49a6b585358c6ff54bbd26613c4a46a988b614e42bc5729eda36"}`, redactedByCQJSONName),
		updatedRecord.Column(2).ValueStr(0))
}

func TestObfuscateColumnsExcept_KeepJSONSubpath(t *testing.T) {
	record := createTestRecord()
	updater := New(record)

	updatedRecord, err := updater.ObfuscateColumnsExcept([]string{"col1", "col2", "col3.foo.bar.0"}, true, "redact")
	require.NoError(t, err)
	require.Equal(t, int64(3), updatedRecord.NumCols())

	got := updatedRecord.Column(2).ValueStr(0)
	require.Equal(t, "a", gjson.Get(got, "foo.bar.0").String(), "kept leaf must be preserved")
	require.True(t, strings.HasPrefix(gjson.Get(got, "foo.bar.1").String(), redactedByCQMessage), "sibling array leaf must be redacted")
	require.True(t, strings.HasPrefix(gjson.Get(got, "foo.bar.2").String(), redactedByCQMessage), "sibling array leaf must be redacted")
	require.True(t, strings.HasPrefix(gjson.Get(got, "hello").String(), redactedByCQMessage), "sibling object leaf must be redacted")
}

func TestObfuscateColumnsExcept_PerElementDistinctHashes(t *testing.T) {
	record := createTestRecord()
	updater := New(record)

	updatedRecord, err := updater.ObfuscateColumnsExcept([]string{"col3.hello"}, true, "redact")
	require.NoError(t, err)

	got := updatedRecord.Column(2).ValueStr(0)
	b0 := gjson.Get(got, "foo.bar.0").String()
	b1 := gjson.Get(got, "foo.bar.1").String()
	b2 := gjson.Get(got, "foo.bar.2").String()
	require.True(t, strings.HasPrefix(b0, redactedByCQMessage))
	require.NotEqual(t, b0, b1, "each array element must hash to a distinct value")
	require.NotEqual(t, b1, b2, "each array element must hash to a distinct value")
	require.Equal(t, "world", gjson.Get(got, "hello").String(), "kept leaf must be preserved")
}

func TestObfuscateColumnsExcept_DropUnhashable(t *testing.T) {
	record := createTestRecordWithTS()
	updater := New(record)

	updatedRecord, err := updater.ObfuscateColumnsExcept([]string{"col1"}, true, "redact")
	require.NoError(t, err)

	require.Equal(t, int64(3), updatedRecord.NumCols())
	names := make([]string, 0, updatedRecord.NumCols())
	for i := 0; i < int(updatedRecord.NumCols()); i++ {
		names = append(names, updatedRecord.ColumnName(i))
	}
	require.Equal(t, []string{"col1", "col2", "col3"}, names)
	require.Equal(t, "val1", updatedRecord.Column(0).(*array.String).Value(0))
}

func TestObfuscateColumnsExcept_CQColumnsPassthrough(t *testing.T) {
	record := createTestRecordWithCQ()
	updater := New(record)

	updatedRecord, err := updater.ObfuscateColumnsExcept([]string{"name"}, true, "redact")
	require.NoError(t, err)

	require.Equal(t, int64(4), updatedRecord.NumCols())
	require.Equal(t, "id1", updatedRecord.Column(0).(*array.String).Value(0), "_cq_id must pass through")
	require.Equal(t, "t1", updatedRecord.Column(1).(*array.String).Value(0), "_cq_sync_time must pass through")
	require.Equal(t, "n1", updatedRecord.Column(2).(*array.String).Value(0), "allowlisted column must be kept")
	require.Equal(t, redactValue([]byte("s1"), true), updatedRecord.Column(3).(*array.String).Value(0), "non-allowlisted column must be redacted")
}

func TestObfuscateColumnsExcept_IncludeSHAFalse(t *testing.T) {
	record := createTestRecord()
	updater := New(record)

	updatedRecord, err := updater.ObfuscateColumnsExcept([]string{"col1"}, false, "redact")
	require.NoError(t, err)

	require.Equal(t, redactedByCQMessageNoSHA, updatedRecord.Column(1).(*array.String).Value(0), "no SHA appended when include_sha=false")
	require.Equal(t,
		fmt.Sprintf(`{"%s":"%s"}`, redactedByCQJSONName, redactedByCQMessageNoSHA),
		updatedRecord.Column(2).ValueStr(0))
}

func TestObfuscateColumnsExcept_TopLevelArrayKeepImage(t *testing.T) {
	md := arrow.NewMetadata([]string{schema.MetadataTableName}, []string{"testTable"})
	bld := array.NewRecordBuilder(memory.DefaultAllocator, arrow.NewSchema(
		[]arrow.Field{
			{Name: "_cq_id", Type: arrow.BinaryTypes.String},
			{Name: "spec_containers", Type: types.NewJSONType()},
		},
		&md,
	))
	defer bld.Release()
	bld.Field(0).(*array.StringBuilder).AppendValues([]string{"id1"}, nil)
	bld.Field(1).(*types.JSONBuilder).AppendBytes([]byte(`[{"image":"nginx","env":[{"name":"A","value":"1"}]},{"image":"redis","env":[{"name":"B","value":"2"}]}]`))
	record := bld.NewRecordBatch()
	updater := New(record)

	updatedRecord, err := updater.ObfuscateColumnsExcept([]string{"spec_containers.#.image"}, true, "redact")
	require.NoError(t, err)

	got := updatedRecord.Column(1).ValueStr(0)
	require.Equal(t, "nginx", gjson.Get(got, "0.image").String(), "image kept for element 0")
	require.Equal(t, "redis", gjson.Get(got, "1.image").String(), "image kept for element 1 via # wildcard")
	require.True(t, strings.HasPrefix(gjson.Get(got, "0.env.0.value").String(), redactedByCQMessage), "env value must be redacted")
	require.True(t, strings.HasPrefix(gjson.Get(got, "0.env.0.name").String(), redactedByCQMessage), "non-allowlisted env name must be redacted")
}

func TestObfuscateColumnsExcept_DropMode(t *testing.T) {
	record := createTestRecord()
	updated, err := New(record).ObfuscateColumnsExcept([]string{"col1", "col3.hello"}, true, spec.UnmatchedDrop)
	require.NoError(t, err)

	names := make([]string, 0, updated.NumCols())
	for i := 0; i < int(updated.NumCols()); i++ {
		names = append(names, updated.ColumnName(i))
	}
	require.Equal(t, []string{"col1", "col3"}, names, "non-allowlisted col2 dropped entirely")
	require.Equal(t, "val1", updated.Column(0).(*array.String).Value(0))
	require.Equal(t, `{"hello":"world"}`, updated.Column(1).ValueStr(0), "non-kept foo subtree dropped")
	require.Zero(t, strings.Count(updated.Column(1).ValueStr(0), redactedByCQMessage), "no markers in drop mode")
}

func TestObfuscateColumnsExcept_DropModeArray(t *testing.T) {
	md := arrow.NewMetadata([]string{schema.MetadataTableName}, []string{"testTable"})
	bld := array.NewRecordBuilder(memory.DefaultAllocator, arrow.NewSchema(
		[]arrow.Field{
			{Name: "_cq_id", Type: arrow.BinaryTypes.String},
			{Name: "spec_containers", Type: types.NewJSONType()},
		},
		&md,
	))
	defer bld.Release()
	bld.Field(0).(*array.StringBuilder).AppendValues([]string{"id1"}, nil)
	bld.Field(1).(*types.JSONBuilder).AppendBytes([]byte(`[{"image":"nginx","env":[{"name":"A","value":"1"}]},{"image":"redis"}]`))
	record := bld.NewRecordBatch()

	updated, err := New(record).ObfuscateColumnsExcept([]string{"spec_containers.#.image"}, true, spec.UnmatchedDrop)
	require.NoError(t, err)

	got := updated.Column(1).ValueStr(0)
	require.Equal(t, "nginx", gjson.Get(got, "0.image").String())
	require.Equal(t, "redis", gjson.Get(got, "1.image").String())
	require.False(t, gjson.Get(got, "0.env").Exists(), "non-allowlisted env dropped in drop mode")
	require.Zero(t, strings.Count(got, redactedByCQMessage), "no markers in drop mode")
}

func TestObfuscateColumnsExcept_CollapseMode(t *testing.T) {
	record := createTestRecord()
	updated, err := New(record).ObfuscateColumnsExcept([]string{"col3.hello"}, true, spec.UnmatchedCollapse)
	require.NoError(t, err)

	got := updated.Column(2).ValueStr(0)
	require.Equal(t, "world", gjson.Get(got, "hello").String(), "kept leaf preserved")
	require.Equal(t, gjson.String, gjson.Get(got, "foo").Type, "non-kept subtree collapsed to a single marker")
	require.True(t, strings.HasPrefix(gjson.Get(got, "foo").String(), redactedByCQMessage))
	require.Equal(t, 1, strings.Count(got, redactedByCQMessage), "one marker for the whole collapsed subtree")
}

func TestObfuscateColumnsExcept_TypedColumnsNeverStringified(t *testing.T) {
	md := arrow.NewMetadata([]string{schema.MetadataTableName}, []string{"testTable"})
	bld := array.NewRecordBuilder(memory.DefaultAllocator, arrow.NewSchema(
		[]arrow.Field{
			{Name: "keep_me", Type: arrow.BinaryTypes.String},
			{Name: "a_bool", Type: arrow.FixedWidthTypes.Boolean},
			{Name: "a_ts", Type: arrow.FixedWidthTypes.Timestamp_us},
			{Name: "an_int", Type: arrow.PrimitiveTypes.Int64},
		},
		&md,
	))
	defer bld.Release()
	bld.Field(0).(*array.StringBuilder).AppendValues([]string{"v1"}, nil)
	bld.Field(1).(*array.BooleanBuilder).AppendValues([]bool{true}, nil)
	bld.Field(2).(*array.TimestampBuilder).AppendValues([]arrow.Timestamp{1}, nil)
	bld.Field(3).(*array.Int64Builder).AppendValues([]int64{7}, nil)
	record := bld.NewRecordBatch()

	for _, mode := range []string{spec.UnmatchedRedact, spec.UnmatchedCollapse, spec.UnmatchedDrop} {
		t.Run(mode, func(t *testing.T) {
			updated, err := New(record).ObfuscateColumnsExcept([]string{"keep_me"}, true, mode)
			require.NoError(t, err)

			names := make([]string, 0, updated.NumCols())
			for i := 0; i < int(updated.NumCols()); i++ {
				names = append(names, updated.ColumnName(i))
				require.Equal(t, record.Schema().Field(record.Schema().FieldIndices(updated.ColumnName(i))[0]).Type,
					updated.Schema().Field(i).Type, "surviving column must keep its original Arrow type")
			}
			require.Equal(t, []string{"keep_me"}, names,
				"non-allowlisted bool/timestamp/int columns must be dropped, never redacted into a string")
		})
	}
}

func TestObfuscateColumnsExcept_DropModePKErrors(t *testing.T) {
	record := createTestRecordWithPK()
	_, err := New(record).ObfuscateColumnsExcept([]string{"name"}, true, spec.UnmatchedDrop)
	require.Error(t, err)
	require.Contains(t, err.Error(), "uid")
	require.Contains(t, err.Error(), "drop")
}

func TestObfuscateColumns_IncludeSHAFalse(t *testing.T) {
	record := createTestRecord()
	updater := New(record)

	updatedRecord, err := updater.ObfuscateColumns([]string{"col1"}, false)
	require.NoError(t, err)

	require.Equal(t, redactedByCQMessageNoSHA, updatedRecord.Column(0).(*array.String).Value(0))
	require.Equal(t, redactedByCQMessageNoSHA, updatedRecord.Column(0).(*array.String).Value(1))
}

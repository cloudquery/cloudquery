package recordupdater

import (
	"encoding/json"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestFlattenJSONFields(t *testing.T) {
	record := testRecord(
		[]string{"col1"},
		map[string]string{"col1": `{"key_a": "utf8", "key_b": "int64", "key_c": "bool"}`},
		[]arrow.Array{buildJSONColumn([]*any{toP(`{"key_a": "value", "key_b": 2, "key_c": true}`)})},
	)
	updater := New(zerolog.Nop(), record)

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

func TestFlattenJSONFieldsWithTimestamp(t *testing.T) {
	record := testRecord(
		[]string{"col1"},
		map[string]string{"col1": `{"key_a": "timestamp[us, tz=UTC]"}`},
		[]arrow.Array{buildJSONColumn([]*any{toP(`{"key_a": "2024-01-02T03:04:05.006Z"}`)})},
	)
	updater := New(zerolog.Nop(), record)

	updatedRecord, err := updater.FlattenJSONFields()
	require.NoError(t, err)

	require.Equal(t, int64(2), updatedRecord.NumCols())
	require.Equal(t, int64(1), updatedRecord.NumRows())
	requireAllColsLenMatchRecordsLen(t, updatedRecord)
	require.Equal(t, "col1__key_a", updatedRecord.ColumnName(1))
	require.Equal(t, "2024-01-02T03:04:05.006Z", updatedRecord.Column(1).(*array.Timestamp).Value(0).ToTime(arrow.Microsecond).Format("2006-01-02T15:04:05.000Z"))
}

func TestFlattenJSONFieldsDoesntFlattenFieldsKeyedUTF8(t *testing.T) {
	record := testRecord(
		[]string{"col1"},
		map[string]string{"col1": `{"key_a": "utf8", "key_b": "int64", "utf8": "any"}`},
		[]arrow.Array{buildJSONColumn([]*any{toP(`{"key_a": "value", "key_b": 2, "utf8": "any"}`)})},
	)
	updater := New(zerolog.Nop(), record)

	updatedRecord, err := updater.FlattenJSONFields()
	require.NoError(t, err)

	require.Equal(t, int64(3), updatedRecord.NumCols())
	require.Equal(t, int64(1), updatedRecord.NumRows())
	requireAllColsLenMatchRecordsLen(t, updatedRecord)
	require.Equal(t, "col1__key_a", updatedRecord.ColumnName(1))
	require.Equal(t, "value", updatedRecord.Column(1).(*array.String).Value(0))
	require.Equal(t, "col1__key_b", updatedRecord.ColumnName(2))
	require.Equal(t, int64(2), updatedRecord.Column(2).(*array.Int64).Value(0))
}

func TestNestedJSONFlattenedToFirstLevel(t *testing.T) {
	record := testRecord(
		[]string{"col1"},
		map[string]string{"col1": `{"nested": {"key_a": "utf8", "key_b": "int64", "key_c": "bool"}}`},
		[]arrow.Array{buildJSONColumn([]*any{toP(`{"nested": {"key_a": "value", "key_b": 2, "key_c": true}}`)})},
	)
	updater := New(zerolog.Nop(), record)

	updatedRecord, err := updater.FlattenJSONFields()
	require.NoError(t, err)

	require.Equal(t, int64(2), updatedRecord.NumCols())
	require.Equal(t, int64(1), updatedRecord.NumRows())
	requireAllColsLenMatchRecordsLen(t, updatedRecord)
	require.Equal(t, "col1", updatedRecord.ColumnName(0))
	require.Equal(t, "col1__nested", updatedRecord.ColumnName(1))
	require.Equal(t, "json", updatedRecord.Schema().Field(0).Type.String())
	require.Equal(t, json.RawMessage(`{"key_a":"value","key_b":2,"key_c":true}`), updatedRecord.Column(1).(*types.JSONArray).GetOneForMarshal(0))
}

func TestDifferentCasingWorks(t *testing.T) {
	record := testRecord(
		[]string{"col"},
		map[string]string{"col": `{"subcolumn_one": "utf8"}`}, // Note the different casing
		[]arrow.Array{buildJSONColumn([]*any{toP(`{"subcolumnOne": "value", "unknownColumn": 2}`)})},
	)
	updater := New(zerolog.Nop(), record)

	updatedRecord, err := updater.FlattenJSONFields()
	require.NoError(t, err)

	require.Equal(t, int64(2), updatedRecord.NumCols())
	require.Equal(t, int64(1), updatedRecord.NumRows())
	requireAllColsLenMatchRecordsLen(t, updatedRecord)
	require.Equal(t, "col", updatedRecord.ColumnName(0))
	require.Equal(t, "col__subcolumn_one", updatedRecord.ColumnName(1))
	require.Equal(t, "utf8", updatedRecord.Schema().Field(1).Type.String())
	require.Equal(t, "value", updatedRecord.Column(1).(*array.String).Value(0))
}

func TestDifferentCasingWorksEvenWhenFirstRowIsNull(t *testing.T) {
	record := testRecord(
		[]string{"col"},
		map[string]string{"col": `{"subcolumn_one": "utf8"}`}, // Note the different casing
		// Note first 3 rows are nil
		[]arrow.Array{buildJSONColumn([]*any{nil, nil, nil, toP(`{"subcolumnOne": "value", "unknownColumn": 2}`)})},
	)
	updater := New(zerolog.Nop(), record)

	updatedRecord, err := updater.FlattenJSONFields()
	require.NoError(t, err)

	require.Equal(t, int64(2), updatedRecord.NumCols())
	require.Equal(t, int64(4), updatedRecord.NumRows())
	requireAllColsLenMatchRecordsLen(t, updatedRecord)
	require.Equal(t, "col", updatedRecord.ColumnName(0))
	require.Equal(t, "col__subcolumn_one", updatedRecord.ColumnName(1))
	require.Equal(t, "utf8", updatedRecord.Schema().Field(1).Type.String())
	require.Equal(t, "value", updatedRecord.Column(1).(*array.String).Value(3))
}

func requireAllColsLenMatchRecordsLen(t *testing.T, record arrow.Record) {
	for i := 0; i < int(record.NumCols()); i++ {
		require.Equal(t, int(record.NumRows()), record.Column(i).Len(), "Expected length of %d for column %d", record.NumRows(), i)
	}
}

func testRecord(fieldNames []string, metadataTypeSchema map[string]string, rows []arrow.Array) arrow.Record {
	tableMD := arrow.NewMetadata([]string{schema.MetadataTableName}, []string{"testTable"})
	fields := make([]arrow.Field, len(fieldNames))
	for i, name := range fieldNames {
		fieldMD := map[string]string{schema.MetadataTypeSchema: metadataTypeSchema[name]}
		fields[i] = arrow.Field{Name: name, Type: rows[i].DataType(), Nullable: true, Metadata: arrow.MetadataFrom(fieldMD)}
	}
	return array.NewRecord(arrow.NewSchema(fields, &tableMD), rows, int64(rows[0].Len()))
}

func toP(s string) *any {
	a := any(json.RawMessage(s))
	return &a
}

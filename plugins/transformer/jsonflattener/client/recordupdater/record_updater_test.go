package recordupdater

import (
	"encoding/json"
	"testing"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/apache/arrow/go/v17/arrow/array"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/stretchr/testify/require"
)

func TestFlattenJSONFields(t *testing.T) {
	record := testRecord(
		[]string{"col1"},
		map[string]string{"col1": `{"key_a": "utf8", "key_b": "int64", "key_c": "bool"}`},
		[]arrow.Array{buildJSONColumn([]*any{toP(`{"key_a": "value", "key_b": 2, "key_c": true}`)})},
	)
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

func TestNestedJSONFlattenedToFirstLevel(t *testing.T) {
	record := testRecord(
		[]string{"col1"},
		map[string]string{"col1": `{"nested": {"key_a": "utf8", "key_b": "int64", "key_c": "bool"}}`},
		[]arrow.Array{buildJSONColumn([]*any{toP(`{"nested": {"key_a": "value", "key_b": 2, "key_c": true}}`)})},
	)
	updater := New(record)

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
	updater := New(record)

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
	updater := New(record)

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

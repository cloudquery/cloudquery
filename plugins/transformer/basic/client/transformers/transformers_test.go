package transformers

import (
	"strings"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/cloudquery/plugins/transformer/basic/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/stretchr/testify/require"
)

func TestNewFromSpec(t *testing.T) {
	tests := []struct {
		name    string
		spec    spec.TransformationSpec
		wantErr bool
	}{
		{
			name: "AddLiteralStringColumn",
			spec: spec.TransformationSpec{
				Kind:  spec.KindAddColumn,
				Name:  "new_col",
				Value: &[]string{"default"}[0],
			},
			wantErr: false,
		},
		{
			name: "RemoveColumns",
			spec: spec.TransformationSpec{
				Kind:    spec.KindRemoveColumns,
				Columns: []string{"col1"},
			},
			wantErr: false,
		},
		{
			name: "ObfuscateColumns",
			spec: spec.TransformationSpec{
				Kind:    spec.KindObfuscateColumns,
				Columns: []string{"col2"},
			},
			wantErr: false,
		},
		{
			name: "ChangeTableNames",
			spec: spec.TransformationSpec{
				Kind:                 spec.KindChangeTableNames,
				NewTableNameTemplate: "cq_sync_{{.OldName}}",
			},
			wantErr: false,
		},

		{
			name: "AddTimestampColumn",
			spec: spec.TransformationSpec{
				Kind:                 spec.KindAddTimestampColumn,
				NewTableNameTemplate: "_last_updated",
			},
			wantErr: false,
		},
		{
			name: "InvalidKind",
			spec: spec.TransformationSpec{
				Kind: "invalid_kind",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewFromSpec(tt.spec)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFromSpec() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTransform(t *testing.T) {
	tests := []struct {
		name     string
		spec     spec.TransformationSpec
		record   arrow.Record
		validate func(t *testing.T, record arrow.Record)
	}{
		{
			name: "AddLiteralStringColumn",
			spec: spec.TransformationSpec{
				Kind:   spec.KindAddColumn,
				Name:   "new_col",
				Value:  &[]string{"default"}[0],
				Tables: []string{"*"},
			},
			record: createTestRecord(),
			validate: func(t *testing.T, record arrow.Record) {
				require.Equal(t, int64(3), record.NumCols(), "Expected 3 columns")
				require.Equal(t, int64(2), record.NumRows(), "Expected 2 rows")
				require.Equal(t, "default", record.Column(2).(*array.String).Value(0), "Expected 'default' value in new_col column")
				require.Equal(t, "default", record.Column(2).(*array.String).Value(1), "Expected 'default' value in new_col column")
			},
		},
		{
			name: "AddTimestampColumn",
			spec: spec.TransformationSpec{
				Kind:   spec.KindAddTimestampColumn,
				Name:   "new_col",
				Tables: []string{"*"},
			},
			record: createTestRecord(),
			validate: func(t *testing.T, record arrow.Record) {
				require.Equal(t, int64(3), record.NumCols(), "Expected 3 columns")
				require.Equal(t, int64(2), record.NumRows(), "Expected 2 rows")
			},
		},
		{
			name: "RemoveColumns",
			spec: spec.TransformationSpec{
				Kind:    spec.KindRemoveColumns,
				Columns: []string{"col1"},
				Tables:  []string{"*"},
			},
			record: createTestRecord(),
			validate: func(t *testing.T, record arrow.Record) {
				require.Equal(t, int64(1), record.NumCols(), "Expected 1 column")
				require.Equal(t, int64(2), record.NumRows(), "Expected 2 rows")
			},
		},
		{
			name: "ObfuscateColumns",
			spec: spec.TransformationSpec{
				Kind:    spec.KindObfuscateColumns,
				Columns: []string{"col2"},
				Tables:  []string{"*"},
			},
			record: createTestRecord(),
			validate: func(t *testing.T, record arrow.Record) {
				require.Equal(t, "Redacted by CloudQuery | bac8d4414984861d5199b7a97699c728bee36c4084299b2ca905434cf65d8944", record.Column(1).(*array.String).Value(0), "Expected sha256 value in new_col column")
				require.Equal(t, "Redacted by CloudQuery | dd0fff6ac351dd46cd26e2d5c61e479ce7c68ef12489e04284c0fd66648723cb", record.Column(1).(*array.String).Value(1), "Expected sha256 value in new_col column")
				require.Equal(t, int64(2), record.NumCols(), "Expected 2 columns")
				require.Equal(t, int64(2), record.NumRows(), "Expected 2 rows")
			},
		},
		{
			name: "ObfuscateNestedColumnsWithGjsonSyntax",
			spec: spec.TransformationSpec{
				Kind:    spec.KindObfuscateColumns,
				Columns: []string{"col3.top_foo.#.foo"},
				Tables:  []string{"*"},
			},
			record: createTestRecordWithNestedArray(),
			validate: func(t *testing.T, record arrow.Record) {
				require.Equal(t, int64(3), record.NumCols(), "Expected 3 columns")
				require.Equal(t, int64(2), record.NumRows(), "Expected 2 rows")
				// Check that the nested foo values are obfuscated
				col3Val := record.Column(2).ValueStr(0)
				require.Contains(t, col3Val, "Redacted by CloudQuery |", "Expected obfuscated values to contain redacted message")
				require.Contains(t, col3Val, "top_foo", "Expected top_foo structure to be maintained")
				// Verify that all three "foo" values in the array are obfuscated
				require.Equal(t, 3, strings.Count(col3Val, "Redacted by CloudQuery |"), "Expected 3 obfuscated values for the 3 foo items")
			},
		},
		{
			name: "ObfuscateDeeplyNestedColumnsWithGjsonSyntax",
			spec: spec.TransformationSpec{
				Kind:    spec.KindObfuscateColumns,
				Columns: []string{"col3.object1.object2.#.nested_object1.nested_object2.#.nested2_object1"},
				Tables:  []string{"*"},
			},
			record: createTestRecordWithDeeplyNestedArray(),
			validate: func(t *testing.T, record arrow.Record) {
				require.Equal(t, int64(3), record.NumCols(), "Expected 3 columns")
				require.Equal(t, int64(2), record.NumRows(), "Expected 2 rows")
				// Check first row: should have 4 obfuscated values
				col3Val := record.Column(2).ValueStr(0)
				require.Contains(t, col3Val, "Redacted by CloudQuery |", "Expected obfuscated values to contain redacted message")
				require.Contains(t, col3Val, "object1", "Expected object1 structure to be maintained")
				require.Contains(t, col3Val, "nested_object1", "Expected nested_object1 structure to be maintained")
				require.Equal(t, 4, strings.Count(col3Val, "Redacted by CloudQuery |"), "Expected 4 obfuscated values for the 4 nested2_object1 items in first row")

				// Check second row: should have 2 obfuscated values
				col3Val2 := record.Column(2).ValueStr(1)
				require.Contains(t, col3Val2, "Redacted by CloudQuery |", "Expected obfuscated values to contain redacted message")
				require.Equal(t, 2, strings.Count(col3Val2, "Redacted by CloudQuery |"), "Expected 2 obfuscated values for the 2 nested2_object1 items in second row")
			},
		},
		{
			name: "UppercaseColumns",
			spec: spec.TransformationSpec{
				Kind:    spec.KindUppercase,
				Columns: []string{"col1"},
				Tables:  []string{"*"},
			},
			record: createTestRecord(),
			validate: func(t *testing.T, record arrow.Record) {
				require.Equal(t, "VAL1", record.Column(0).(*array.String).Value(0), "Expected uppercase value in col1 column")
				require.Equal(t, "VAL2", record.Column(0).(*array.String).Value(1), "Expected uppercase value in col1 column")
				require.Equal(t, int64(2), record.NumCols(), "Expected 2 columns")
				require.Equal(t, int64(2), record.NumRows(), "Expected 2 rows")
			},
		},
		{
			name: "LowercaseColumns",
			spec: spec.TransformationSpec{
				Kind:    spec.KindLowercase,
				Columns: []string{"col1"},
				Tables:  []string{"*"},
			},
			record: createUppercaseTestRecord(),
			validate: func(t *testing.T, record arrow.Record) {
				require.Equal(t, "val1", record.Column(0).(*array.String).Value(0), "Expected lowercase value in col1 column")
				require.Equal(t, "val2", record.Column(0).(*array.String).Value(1), "Expected lowercase value in col1 column")
				require.Equal(t, int64(2), record.NumCols(), "Expected 2 columns")
				require.Equal(t, int64(2), record.NumRows(), "Expected 2 rows")
			},
		},
		{
			name: "ChangeTableName",
			spec: spec.TransformationSpec{
				Kind:                 spec.KindChangeTableNames,
				NewTableNameTemplate: "cq_sync_{{.OldName}}",
				Tables:               []string{"*"},
			},
			record: createTestRecord(),
			validate: func(t *testing.T, record arrow.Record) {
				newTableName, ok := record.Schema().Metadata().GetValue(schema.MetadataTableName)
				require.True(t, ok, "Expected table name to be present in metadata")
				require.Equal(t, "cq_sync_table1", newTableName)
			},
		},
		{
			name: "DropRow-DropFirstRow",
			spec: spec.TransformationSpec{
				Kind:    spec.KindDropRows,
				Tables:  []string{"*"},
				Columns: []string{"col1"},
				Value:   &[]string{"val1"}[0],
			},
			record: createTestRecord(),
			validate: func(t *testing.T, record arrow.Record) {
				require.Equal(t, int64(2), record.NumCols(), "Expected 2 columns")
				require.Equal(t, int64(1), record.NumRows(), "Expected 1 rows")
				require.Equal(t, "val2", record.Column(0).(*array.String).Value(0), "Expected `val2` in col1 column row 0")
				require.Equal(t, "val4", record.Column(1).(*array.String).Value(0), "Expected `val4` in col2 column row 0")
			},
		},
		{
			name: "DropRow-DropSecondRow",
			spec: spec.TransformationSpec{
				Kind:    spec.KindDropRows,
				Tables:  []string{"*"},
				Columns: []string{"col2"},
				Value:   &[]string{"val4"}[0],
			},
			record: createTestRecord(),
			validate: func(t *testing.T, record arrow.Record) {
				require.Equal(t, int64(2), record.NumCols(), "Expected 2 columns")
				require.Equal(t, int64(1), record.NumRows(), "Expected 1 rows")
				require.Equal(t, "val1", record.Column(0).(*array.String).Value(0), "Expected `val1` in col1 column row 0")
				require.Equal(t, "val3", record.Column(1).(*array.String).Value(0), "Expected `val4` in col2 column row 0")
			},
		},
		{
			name: "DropRow-DoNotDropAny",
			spec: spec.TransformationSpec{
				Kind:    spec.KindDropRows,
				Tables:  []string{"*"},
				Columns: []string{"col6"},
				Value:   &[]string{"val1"}[0],
			},
			record: createTestRecord(),
			validate: func(t *testing.T, record arrow.Record) {
				require.Equal(t, int64(2), record.NumCols(), "Expected 2 columns")
				require.Equal(t, int64(2), record.NumRows(), "Expected 2 rows")
				require.Equal(t, "val1", record.Column(0).(*array.String).Value(0), "Expected `val1` in col1 column row 0")
				require.Equal(t, "val2", record.Column(0).(*array.String).Value(1), "Expected `val2` in col1 column row 1")
				require.Equal(t, "val3", record.Column(1).(*array.String).Value(0), "Expected `val3` in col2 column row 0")
				require.Equal(t, "val4", record.Column(1).(*array.String).Value(1), "Expected `val4` in col2 column row 0")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transformer, err := NewFromSpec(tt.spec)
			require.NoError(t, err, "NewFromSpec() should not return an error")

			transformedRecord, err := transformer.Transform(tt.record)
			require.NoError(t, err, "Transform() should not return an error")

			requireAllColsLenMatchRecordsLen(t, transformedRecord)
			tt.validate(t, transformedRecord)
		})
	}
}

func createTestRecord() arrow.Record {
	md := arrow.NewMetadata([]string{schema.MetadataTableName}, []string{"table1"})
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

func createUppercaseTestRecord() arrow.Record {
	md := arrow.NewMetadata([]string{schema.MetadataTableName}, []string{"table1"})
	bld := array.NewRecordBuilder(memory.DefaultAllocator, arrow.NewSchema(
		[]arrow.Field{
			{Name: "col1", Type: arrow.BinaryTypes.String},
			{Name: "col2", Type: arrow.BinaryTypes.String},
		},
		&md,
	))
	defer bld.Release()

	bld.Field(0).(*array.StringBuilder).AppendValues([]string{"VAL1", "VAL2"}, nil)
	bld.Field(1).(*array.StringBuilder).AppendValues([]string{"val3", "val4"}, nil)

	return bld.NewRecord()
}

func createTestRecordWithNestedArray() arrow.Record {
	md := arrow.NewMetadata([]string{schema.MetadataTableName}, []string{"table1"})
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

	return bld.NewRecord()
}

func createTestRecordWithDeeplyNestedArray() arrow.Record {
	md := arrow.NewMetadata([]string{schema.MetadataTableName}, []string{"table1"})
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
	bld.Field(2).(*types.JSONBuilder).AppendBytes([]byte(`{"object1":{"object2":[{"nested_object1":{"nested_object2":[{"nested2_object1":1},{"nested2_object1":2}]}},{"nested_object1":{"nested_object2":[{"nested2_object1":3},{"nested2_object1":4}]}}]}}`))
	bld.Field(2).(*types.JSONBuilder).AppendBytes([]byte(`{"object1":{"object2":[{"nested_object1":{"nested_object2":[{"nested2_object1":5},{"nested2_object1":6}]}}]}}`))

	return bld.NewRecord()
}

func requireAllColsLenMatchRecordsLen(t *testing.T, record arrow.Record) {
	for i := 0; i < int(record.NumCols()); i++ {
		require.Equal(t, int(record.NumRows()), record.Column(i).Len(), "Expected length of %d for column %d", record.NumRows(), i)
	}
}

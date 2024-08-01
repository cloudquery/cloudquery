package tablematcher

import (
	"testing"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/apache/arrow/go/v17/arrow/array"
	"github.com/apache/arrow/go/v17/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/stretchr/testify/require"
)

func TestIsTableMatch(t *testing.T) {
	patterns := []string{"table_*", "data_*"}
	matcher := New(patterns)

	tests := []struct {
		tableName string
		expected  bool
	}{
		{"table_1", true},
		{"data_1", true},
		{"other_table", false},
	}

	for _, test := range tests {
		result := matcher.isTableMatch(test.tableName)
		require.Equal(t, test.expected, result, "Expected %v for table name %s, got %v", test.expected, test.tableName, result)
	}
}

func TestIsSchemasTableMatch(t *testing.T) {
	tests := []struct {
		patterns  []string
		tableName string
		expected  bool
	}{
		{[]string{"table_*", "data_*"}, "table_1", true},
		{[]string{"data_*"}, "table_1", false},
		{[]string{"table_*", "data_*"}, "data_1", true},
		{[]string{"table_*"}, "data_1", false},
		{[]string{"table_*", "data_*"}, "other_table", false},
		{[]string{"*"}, "other_table", true},
	}
	for _, tt := range tests {
		matcher := New(tt.patterns)

		record := createTestRecordWithMetadata(tt.tableName)
		result, err := matcher.IsSchemasTableMatch(record.Schema())
		require.NoError(t, err, "Unexpected error")
		require.Equal(t, tt.expected, result, "Expected %v for table name %s, got %v", tt.expected, tt.tableName, result)
	}
}

func TestIsSchemasTableMatch_NoMetadata(t *testing.T) {
	patterns := []string{"table_*", "data_*"}
	matcher := New(patterns)

	bld := array.NewRecordBuilder(memory.DefaultAllocator, arrow.NewSchema(
		[]arrow.Field{
			{Name: "col1", Type: arrow.BinaryTypes.String},
		},
		nil,
	))
	defer bld.Release()

	bld.Field(0).(*array.StringBuilder).AppendValues([]string{"val1"}, nil)
	record := bld.NewRecord()

	_, err := matcher.IsSchemasTableMatch(record.Schema())
	require.Error(t, err, "Expected error")
	require.EqualError(t, err, "table name not found in record's metadata", "Expected error message 'table name not found in record's metadata'")
}

func createTestRecordWithMetadata(tableName string) arrow.Record {
	md := arrow.NewMetadata([]string{schema.MetadataTableName}, []string{tableName})

	bld := array.NewRecordBuilder(memory.DefaultAllocator, arrow.NewSchema(
		[]arrow.Field{
			{Name: "col1", Type: arrow.BinaryTypes.String},
		},
		&md,
	))
	defer bld.Release()

	bld.Field(0).(*array.StringBuilder).AppendValues([]string{"val1"}, nil)

	return bld.NewRecord()
}

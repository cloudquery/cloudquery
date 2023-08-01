package queries

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/apache/arrow/go/v13/arrow/memory"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetRows(t *testing.T) {
	mem := memory.NewCheckedAllocator(memory.DefaultAllocator)
	defer mem.AssertSize(t, 0)

	schema := arrow.NewSchema([]arrow.Field{
		{Name: "str", Type: arrow.BinaryTypes.String},
		{Name: "int8", Type: arrow.PrimitiveTypes.Int8},
		{Name: "int16", Type: arrow.PrimitiveTypes.Int16},
		{Name: "int32", Type: arrow.PrimitiveTypes.Int32},
		{Name: "int64", Type: arrow.PrimitiveTypes.Int64},
	}, nil)

	builder := array.NewRecordBuilder(mem, schema)
	defer builder.Release()

	const rowsPerRecord = 100
	for _, fBuilder := range builder.Fields() {
		fBuilder.AppendEmptyValues(rowsPerRecord)
	}
	record := builder.NewRecord()
	defer record.Release()

	table := array.NewTableFromRecords(schema, []arrow.Record{record, record, record})
	defer table.Release()

	require.Equal(t, int64(schema.NumFields()), table.NumCols())
	require.Equal(t, int64(3*rowsPerRecord), table.NumRows())

	rows, err := GetRows(table)
	require.NoError(t, err)
	require.Equal(t, 3*rowsPerRecord, len(rows))
	for _, row := range rows {
		require.Equal(t, schema.NumFields(), len(row))
	}
}

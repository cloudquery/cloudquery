package queries

import (
	"math/rand"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/stretchr/testify/require"
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

	records := make([]arrow.Record, 7)
	defer func() {
		for _, rec := range records {
			rec.Release()
		}
	}()

	var total int64
	for i := range records {
		rec := genRecord(mem, schema, rand.Intn(500))
		records[i] = rec
		t.Logf("generated record with %d rows\n", rec.NumRows())
		total += rec.NumRows()
	}
	t.Logf("generated %d rows in total\n", total)

	table := array.NewTableFromRecords(schema, records)
	defer table.Release()

	require.Equal(t, int64(schema.NumFields()), table.NumCols())
	require.Equal(t, total, table.NumRows())

	rows, err := GetRows(table)
	require.NoError(t, err)
	require.Equal(t, int(total), len(rows))
	for _, row := range rows {
		require.Equal(t, schema.NumFields(), len(row))
	}
}

func genRecord(mem memory.Allocator, schema *arrow.Schema, rows int) arrow.Record {
	builder := array.NewRecordBuilder(mem, schema)
	defer builder.Release()
	for _, fBuilder := range builder.Fields() {
		fBuilder.AppendEmptyValues(rows)
	}
	return builder.NewRecord()
}

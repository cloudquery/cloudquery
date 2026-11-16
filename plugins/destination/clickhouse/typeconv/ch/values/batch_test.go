package values

import (
	"context"
	"errors"
	"testing"

	"github.com/ClickHouse/clickhouse-go/v2/lib/column"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/stretchr/testify/require"
)

type headerBatch struct {
	driver.Batch
	columns       []column.Interface
	rejectedNulls error
}

func (b *headerBatch) Columns() []column.Interface {
	return b.columns
}

func (b *headerBatch) Column(n int) driver.BatchColumn {
	return headerColumn{column: b.columns[n], rejectedNulls: b.rejectedNulls}
}

type headerColumn struct {
	column        column.Interface
	rejectedNulls error
}

func (c headerColumn) Append(v any) error {
	_, err := c.column.Append(v)
	return err
}

func (c headerColumn) AppendRow(v any) error {
	if v == nil && c.rejectedNulls != nil {
		return c.rejectedNulls
	}
	return c.column.AppendRow(v)
}

func newHeaderBatch(t *testing.T, names ...string) *headerBatch {
	t.Helper()
	batch := &headerBatch{}
	for _, name := range names {
		col, err := column.Type("Nullable(String)").Column(name, nil)
		require.NoError(t, err)
		batch.columns = append(batch.columns, col)
	}
	return batch
}

func stringSchema(names ...string) *arrow.Schema {
	fields := make([]arrow.Field, len(names))
	for i, name := range names {
		fields[i] = arrow.Field{Name: name, Type: arrow.BinaryTypes.String, Nullable: true}
	}
	return arrow.NewSchema(fields, nil)
}

func stringRecord(sc *arrow.Schema, values ...string) arrow.RecordBatch {
	builder := array.NewRecordBuilder(memory.DefaultAllocator, sc)
	for i, value := range values {
		builder.Field(i).(*array.StringBuilder).Append(value)
	}
	return builder.NewRecordBatch()
}

func TestSchemaFieldIndicesMatchesByNameAndMarksExtraColumns(t *testing.T) {
	batch := newHeaderBatch(t, "b", "old", "a")
	indices, err := schemaFieldIndices(batch, stringSchema("a", "b"))
	require.NoError(t, err)
	require.Equal(t, []int{1, columnNotInSchema, 0}, indices)
}

func TestSchemaFieldIndicesMatchesNamesWithoutNUL(t *testing.T) {
	batch := newHeaderBatch(t, "ab")
	indices, err := schemaFieldIndices(batch, stringSchema("a\x00b"))
	require.NoError(t, err)
	require.Equal(t, []int{0}, indices)
}

func TestSchemaFieldIndicesFailsWhenTableLacksSchemaColumn(t *testing.T) {
	batch := newHeaderBatch(t, "a")
	_, err := schemaFieldIndices(batch, stringSchema("a", "b", "c"))
	require.EqualError(t, err, "table is missing 2 schema columns")
}

func TestBatchAddRecordsFillsExtraColumnsWithNulls(t *testing.T) {
	batch := newHeaderBatch(t, "old", "a")
	sc := stringSchema("a")
	records := []arrow.RecordBatch{stringRecord(sc, "x"), stringRecord(sc, "y")}

	require.NoError(t, BatchAddRecords(context.Background(), batch, sc, records))

	old, a := batch.columns[0], batch.columns[1]
	require.Equal(t, 2, old.Rows())
	require.Nil(t, old.Row(0, true))
	require.Nil(t, old.Row(1, true))
	require.Equal(t, 2, a.Rows())
	require.Equal(t, "x", *a.Row(0, true).(*string))
	require.Equal(t, "y", *a.Row(1, true).(*string))
}

func TestBatchAddRecordsFailsWhenExtraColumnRejectsNull(t *testing.T) {
	batch := newHeaderBatch(t, "old", "a")
	batch.rejectedNulls = errors.New("not nullable")
	sc := stringSchema("a")

	err := BatchAddRecords(context.Background(), batch, sc, []arrow.RecordBatch{stringRecord(sc, "x")})

	require.EqualError(t, err, `column "old" is not in the table schema: not nullable`)
}

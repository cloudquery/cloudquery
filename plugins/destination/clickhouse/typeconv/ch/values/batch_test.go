package values

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2/lib/column"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/stretchr/testify/require"
)

type headerBatch struct {
	driver.Batch
	columns  []column.Interface
	rejected map[string]error
}

func (b *headerBatch) Columns() []column.Interface {
	return b.columns
}

func (b *headerBatch) Column(n int) driver.BatchColumn {
	return &headerColumn{column: b.columns[n], err: b.rejected[b.columns[n].Name()]}
}

type headerColumn struct {
	column column.Interface
	err    error
}

func (c *headerColumn) Append(v any) error {
	if c.err != nil {
		return c.err
	}
	_, c.err = c.column.Append(v)
	return c.err
}

func (c *headerColumn) AppendRow(v any) error {
	if c.err != nil {
		return c.err
	}
	c.err = c.column.AppendRow(v)
	return c.err
}

func newHeaderBatch(t *testing.T, names ...string) *headerBatch {
	t.Helper()
	batch := &headerBatch{}
	for _, name := range names {
		batch.addColumn(t, name, "Nullable(String)")
	}
	return batch
}

func (b *headerBatch) addColumn(t *testing.T, name string, chType column.Type) {
	t.Helper()
	col, err := chType.Column(name, nil)
	require.NoError(t, err)
	b.columns = append(b.columns, col)
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

func TestBatchAddRecordsFillsExtraColumnsThatRejectNullWithEmptyValues(t *testing.T) {
	batch := newHeaderBatch(t)
	batch.addColumn(t, "extra", "Array(String)")
	batch.addColumn(t, "pair", "Tuple(s String, n Int64)")
	batch.addColumn(t, "a", "Nullable(String)")
	sc := stringSchema("a")
	records := []arrow.RecordBatch{stringRecord(sc, "x"), stringRecord(sc, "y")}

	require.NoError(t, BatchAddRecords(context.Background(), batch, sc, records))

	extra, pair := batch.columns[0], batch.columns[1]
	for row := range 2 {
		require.Equal(t, []string{}, extra.Row(row, false))
		require.Equal(t, map[string]any{"s": "", "n": int64(0)}, pair.Row(row, false))
	}
}

func TestAppendEmptyRowsFillsSupportedColumnTypes(t *testing.T) {
	for _, chType := range []column.Type{
		"String",
		"Int64",
		"DateTime",
		"DateTime64(3)",
		"UUID",
		"FixedString(4)",
		"Nullable(String)",
		"LowCardinality(String)",
		"Array(String)",
		"Array(Array(Int64))",
		"Array(Dynamic)",
		"Array(JSON)",
		"Map(String, String)",
		"Tuple(s String, n Int64)",
		"Tuple(s String, tags Array(String))",
		"Tuple(at DateTime64(3), n Int64)",
	} {
		t.Run(string(chType), func(t *testing.T) {
			header, err := chType.Column("extra", &column.ServerContext{Timezone: time.UTC})
			require.NoError(t, err)

			require.NoError(t, appendEmptyRows(&headerColumn{column: header}, header, 2))

			require.Equal(t, 2, header.Rows())
		})
	}
}

func TestAppendEmptyRowsFailsForUnsupportedNestedTupleShapes(t *testing.T) {
	for _, chType := range []column.Type{
		"Tuple(s LowCardinality(String), n Int64)",
		"Tuple(e Enum8('a' = 1))",
		"Tuple(Tuple(Int64, String), Int64)",
		"Tuple(d Dynamic)",
	} {
		t.Run(string(chType), func(t *testing.T) {
			header, err := chType.Column("extra", &column.ServerContext{Timezone: time.UTC})
			require.NoError(t, err)

			require.Error(t, appendEmptyRows(&headerColumn{column: header}, header, 2))
		})
	}
}

func TestBatchAddRecordsFailsWhenExtraColumnRejectsEmptyValue(t *testing.T) {
	batch := newHeaderBatch(t, "old", "a")
	batch.rejected = map[string]error{"old": errors.New("rejected")}
	sc := stringSchema("a")

	err := BatchAddRecords(context.Background(), batch, sc, []arrow.RecordBatch{stringRecord(sc, "x")})

	require.EqualError(t, err, `column "old" is not in the table schema: rejected`)
}

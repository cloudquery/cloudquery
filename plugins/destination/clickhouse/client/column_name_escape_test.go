package client

import (
	"context"
	"fmt"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v8/client/spec"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v8/util"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/stretchr/testify/require"
)

func TestWriteColumnNamesWithBackslash(t *testing.T) {
	ctx := context.Background()
	p := initPlugin(t)
	table := &schema.Table{
		Name: fmt.Sprintf("cq_test_column_name_backslash_%d", time.Now().UnixNano()),
		Columns: schema.ColumnList{
			{Name: "_cq_id", Type: arrow.BinaryTypes.String, NotNull: true, PrimaryKey: true},
			{Name: `a\q`, Type: arrow.BinaryTypes.String},
			{Name: `C:\Users`, Type: arrow.BinaryTypes.String},
		},
	}

	for i, id := range []string{"first", "second"} {
		require.NoError(t, p.WriteAll(ctx, []message.WriteMessage{
			&message.WriteMigrateTable{Table: table},
			&message.WriteInsert{Record: stringRecord(table, id, fmt.Sprintf("q%d", i), fmt.Sprintf("users%d", i))},
		}))
	}

	require.ElementsMatch(t, []string{"first|q0|users0", "second|q1|users1"}, readRows(t, p, table))
}

func TestWriteToTableThatKeepsDroppedColumn(t *testing.T) {
	for _, name := range []string{`a\q`, "a\x00b"} {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			p := initPlugin(t)
			tableName := fmt.Sprintf("cq_test_kept_column_%d", time.Now().UnixNano())
			columns := schema.ColumnList{
				{Name: "_cq_id", Type: arrow.BinaryTypes.String, NotNull: true, PrimaryKey: true},
				{Name: name, Type: arrow.BinaryTypes.String},
			}
			withOld := &schema.Table{Name: tableName, Columns: append(slices.Clone(columns), schema.Column{Name: "old", Type: arrow.BinaryTypes.String})}
			table := &schema.Table{Name: tableName, Columns: columns}

			require.NoError(t, p.WriteAll(ctx, []message.WriteMessage{&message.WriteMigrateTable{Table: withOld}}))
			require.NoError(t, p.WriteAll(ctx, []message.WriteMessage{
				&message.WriteMigrateTable{Table: table},
				&message.WriteInsert{Record: stringRecord(table, "first", "value")},
			}))

			require.Equal(t, []string{"first|value"}, readRows(t, p, table))
		})
	}
}

func TestWriteToTableWithHandAddedNestedColumns(t *testing.T) {
	ctx := context.Background()
	p := initPlugin(t)
	c := newClient(t, spec.Spec{ConnectionString: getTestConnection()}, plugin.NewClientOptions{})
	table := &schema.Table{
		Name: fmt.Sprintf("cq_test_hand_added_columns_%d", time.Now().UnixNano()),
		Columns: schema.ColumnList{
			{Name: "_cq_id", Type: arrow.BinaryTypes.String, NotNull: true, PrimaryKey: true},
			{Name: `a\q`, Type: arrow.BinaryTypes.String},
		},
	}
	require.NoError(t, migrateTable(c, table, false))
	t.Cleanup(func() { require.NoError(t, c.dropTable(ctx, table)) })
	for _, column := range []string{"extra Array(String)", "pair Tuple(s String, n Int64)"} {
		require.NoError(t, c.conn.Exec(ctx, fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s", util.SanitizeID(table.Name), column)))
	}

	require.NoError(t, c.WriteTableBatch(ctx, table.Name, message.WriteInserts{
		&message.WriteInsert{Record: stringRecord(table, "first", "value")},
	}))

	require.Equal(t, []string{"first|value"}, readRows(t, p, table))
}

func TestWriteColumnNamesWithParentheses(t *testing.T) {
	ctx := context.Background()
	p := initPlugin(t)
	table := &schema.Table{
		Name: fmt.Sprintf("cq_test_column_name_parentheses_%d", time.Now().UnixNano()),
		Columns: schema.ColumnList{
			{Name: "_cq_id", Type: arrow.BinaryTypes.String, NotNull: true, PrimaryKey: true},
			{Name: "a)b", Type: arrow.BinaryTypes.String},
			{Name: "x (y)", Type: arrow.BinaryTypes.String},
		},
	}

	require.NoError(t, p.WriteAll(ctx, []message.WriteMessage{
		&message.WriteMigrateTable{Table: table},
		&message.WriteInsert{Record: stringRecord(table, "first", "ab", "xy")},
	}))

	require.Equal(t, []string{"first|ab|xy"}, readRows(t, p, table))
}

func readRows(t *testing.T, p *plugin.Plugin, table *schema.Table) []string {
	t.Helper()
	read := make(chan arrow.RecordBatch)
	readErr := make(chan error, 1)
	go func() {
		readErr <- p.Read(context.Background(), table, read)
		close(read)
	}()
	var rows []string
	for record := range read {
		for row := 0; row < int(record.NumRows()); row++ {
			values := make([]string, record.NumCols())
			for col := range values {
				values[col] = record.Column(col).ValueStr(row)
			}
			rows = append(rows, strings.Join(values, "|"))
		}
	}
	require.NoError(t, <-readErr)
	return rows
}

func stringRecord(table *schema.Table, values ...string) arrow.RecordBatch {
	builder := array.NewRecordBuilder(memory.DefaultAllocator, table.ToArrowSchema())
	for i, value := range values {
		builder.Field(i).(*array.StringBuilder).Append(value)
	}
	return builder.NewRecordBatch()
}

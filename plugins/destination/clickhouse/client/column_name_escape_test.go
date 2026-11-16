package client

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/message"
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

	read := make(chan arrow.RecordBatch)
	readErr := make(chan error, 1)
	go func() {
		readErr <- p.Read(ctx, table, read)
		close(read)
	}()
	var rows []string
	for record := range read {
		for row := 0; row < int(record.NumRows()); row++ {
			rows = append(rows, fmt.Sprintf("%s|%s|%s",
				record.Column(0).ValueStr(row), record.Column(1).ValueStr(row), record.Column(2).ValueStr(row)))
		}
	}
	require.NoError(t, <-readErr)
	require.ElementsMatch(t, []string{"first|q0|users0", "second|q1|users1"}, rows)
}

func stringRecord(table *schema.Table, values ...string) arrow.RecordBatch {
	builder := array.NewRecordBuilder(memory.DefaultAllocator, table.ToArrowSchema())
	for i, value := range values {
		builder.Field(i).(*array.StringBuilder).Append(value)
	}
	return builder.NewRecordBatch()
}

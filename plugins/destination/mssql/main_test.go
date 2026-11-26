//go:debug x509negativeserial=1
package main

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/cloudquery/plugins/destination/mssql/v5/client"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/require"
)

func testConnectionString() string {
	if conn := os.Getenv("CQ_DEST_MS_SQL_TEST_CONN"); conn != "" {
		return conn
	}
	return "server=localhost;user id=SA;password=yourStrongP@ssword;port=1433;database=cloudquery"
}

func flushTestTable(name string) *schema.Table {
	return &schema.Table{
		Name: name,
		Columns: schema.ColumnList{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64, PrimaryKey: true, NotNull: true},
			schema.CqSourceNameColumn,
			schema.CqSyncTimeColumn,
		},
	}
}

func flushTestInsert(table *schema.Table) *message.WriteInsert {
	record := schema.NewTestDataGenerator(0).Generate(table, schema.GenTestDataOptions{
		SourceName: "test",
		SyncTime:   time.Now().UTC(),
		MaxRows:    1,
	})
	return &message.WriteInsert{Record: record}
}

func TestFlushReportsTablesWithFailedWrites(t *testing.T) {
	ctx := context.Background()
	p := newPlugin()
	require.True(t, p.SupportsFlush())

	spec, err := json.Marshal(client.Spec{ConnectionString: testConnectionString()})
	require.NoError(t, err)
	require.NoError(t, p.Init(ctx, spec, plugin.NewClientOptions{}))
	t.Cleanup(func() { require.NoError(t, p.Close(ctx)) })

	suffix := time.Now().UnixNano()
	written := flushTestTable(fmt.Sprintf("flush_written_%d", suffix))
	notMigrated := flushTestTable(fmt.Sprintf("flush_not_migrated_%d", suffix))
	require.NoError(t, p.WriteAll(ctx, []message.WriteMessage{&message.WriteMigrateTable{Table: written}}))

	flush := &message.WriteFlush{}
	msgs := make(chan message.WriteMessage, 3)
	msgs <- flushTestInsert(written)
	msgs <- flushTestInsert(notMigrated)
	msgs <- flush
	close(msgs)
	require.NoError(t, p.Write(ctx, msgs))

	flushErr := <-flush.Result()
	require.Equal(t, []string{notMigrated.Name}, message.FailedTables(flushErr), flushErr)
}

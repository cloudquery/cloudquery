package client

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/stretchr/testify/require"
)

func getConnectionString() string {
	if testConn := os.Getenv("CQ_DEST_MYSQL_TEST_CONNECTION_STRING"); len(testConn) > 0 {
		return testConn
	}

	return `root:test@/cloudquery`
}

func TestPlugin(t *testing.T) {
	ctx := context.Background()
	p := plugin.NewPlugin("mysql", "development", New)
	s := &Spec{
		ConnectionString: getConnectionString(),
	}
	specBytes, err := json.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	if err := p.Init(ctx, specBytes, plugin.NewClientOptions{}); err != nil {
		t.Fatal(err)
	}
	plugin.TestWriterSuiteRunner(t,
		p,
		plugin.WriterTestSuiteTests{
			SkipDeleteRecord: true,
			SafeMigrations: plugin.SafeMigrations{
				AddColumn:              true,
				AddColumnNotNull:       false,
				RemoveColumn:           true,
				RemoveColumnNotNull:    false,
				ChangeColumn:           false,
				RemoveUniqueConstraint: true,
			},
		},
	)
}

func writeManyRows() error {
	ctx := context.Background()
	p := plugin.NewPlugin("mysql", "development", New)
	s := &Spec{
		ConnectionString: getConnectionString(),
	}
	specBytes, err := json.Marshal(s)
	if err != nil {
		return err
	}
	if err := p.Init(ctx, specBytes, plugin.NewClientOptions{}); err != nil {
		return err
	}

	const numberOfTables = 5
	const recordsPerTable = 20
	msgs := make([]message.WriteMessage, 0)
	for i := 0; i < numberOfTables; i++ {
		tableName := fmt.Sprintf("table_%d", i)
		table := schema.TestTable(tableName, schema.TestSourceOptions{})
		table.Columns = append(table.Columns, schema.Column{Name: "name", Type: arrow.BinaryTypes.String, PrimaryKey: true})
		msgs = append(msgs, &message.WriteMigrateTable{Table: table})

		tg := schema.NewTestDataGenerator(0)
		for i := 0; i < recordsPerTable; i++ {
			record := tg.Generate(table, schema.GenTestDataOptions{
				MaxRows: 50,
			})
			msgs = append(msgs, &message.WriteInsert{Record: record})
		}
	}

	err = p.WriteAll(ctx, msgs)
	return err
}

func TestWriteManyRows(t *testing.T) {
	err := writeManyRows()
	require.NoError(t, err)
}

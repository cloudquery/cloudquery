package client

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
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
				AddColumn:           true,
				AddColumnNotNull:    false,
				RemoveColumn:        true,
				RemoveColumnNotNull: false,
				ChangeColumn:        false,
			},
			SkipSpecificMigrations: plugin.Migrations{
				RemoveUniqueConstraint: true,
			},
		},
	)
}

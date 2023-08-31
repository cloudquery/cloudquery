package client

import (
	"context"
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
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
	// We have to skip some data types each time because a single MySQL table cannot hold all the data types.
	for _, skipOpts := range []schema.TestSourceOptions{
		{
			SkipMaps: true,
		},
		{
			SkipLists: true,
		},
	} {
		plugin.TestWriterSuiteRunner(t,
			p,
			plugin.WriterTestSuiteTests{
				SafeMigrations: plugin.SafeMigrations{
					AddColumn:    true,
					RemoveColumn: true,
				},
			},
			plugin.WithTestDataOptions(skipOpts),
		)
		// This is necessary because tables are named based on the current time
		// As we iterate through the tests, if we don't sleep here then tables can be created with the same name
		time.Sleep(1 * time.Second)
	}
}

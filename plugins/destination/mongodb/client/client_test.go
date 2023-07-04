package client

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

func getTestConnection() string {
	testConn := os.Getenv("CQ_DEST_MONGODB_TEST_CONN")
	if testConn == "" {
		return "mongodb://localhost:27017"
	}
	return testConn
}

func TestPlugin(t *testing.T) {
	ctx := context.Background()
	p := plugin.NewPlugin("mongodb", "development", New)
	s := &Spec{
		ConnectionString: getTestConnection(),
		Database:         "destination_mongodb_test",
	}
	b, err := json.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	if err := p.Init(ctx, b, plugin.NewClientOptions{}); err != nil {
		t.Fatal(err)
	}
	plugin.TestWriterSuiteRunner(t,
		p,
		plugin.WriterTestSuiteTests{
			SkipMigrate: true,
		},
		// plugin.WithTestSourceTimePrecision(time.Millisecond),
	)
}

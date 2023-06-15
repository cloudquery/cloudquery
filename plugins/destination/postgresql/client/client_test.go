package client

import (
	"context"
	"os"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func getTestConnection() string {
	testConn := os.Getenv("CQ_DEST_PG_TEST_CONN")
	if testConn == "" {
		return "postgresql://postgres:pass@localhost:5432/postgres?sslmode=disable"
	}
	return testConn
}

var safeMigrations = plugin.SafeMigrations{
	AddColumn:           true,
	AddColumnNotNull:    false,
	RemoveColumn:        true,
	RemoveColumnNotNull: false,
	ChangeColumn:        false,
}

func TestPgPlugin(t *testing.T) {
	ctx := context.Background()
	p := plugin.NewPlugin("postgresql", "development", New)
	p.Init(ctx, &Spec{
		ConnectionString: getTestConnection(),
		PgxLogLevel:      LogLevelTrace,
	})
	testOpts := schema.TestSourceOptions{
		SkipMaps: true,
	}
	plugin.TestWriterSuiteRunner(t,
		p,
		plugin.PluginTestSuiteTests{
			SafeMigrations: safeMigrations,
		},
		plugin.WithTestDataOptions(testOpts),
	)
}

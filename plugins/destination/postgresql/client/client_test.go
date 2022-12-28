package client

import (
	"os"
	"testing"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
)

func getTestConnection() string {
	testConn := os.Getenv("CQ_DEST_PG_TEST_CONN")
	if testConn == "" {
		return "postgresql://postgres:pass@localhost:5432/postgres?sslmode=disable"
	}
	return testConn
}

func TestPgPlugin(t *testing.T) {
	p := destination.NewPlugin("postgresql", "development", New)
	destination.PluginTestSuiteRunner(t, p,
		Spec{
			ConnectionString: getTestConnection(),
			PgxLogLevel:      LogLevelTrace,
		},
		destination.PluginTestSuiteTests{})
}

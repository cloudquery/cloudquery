package client

import (
	"os"
	"testing"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/specs"
)

func getTestConnection() string {
	testConn := os.Getenv("CQ_DEST_PG_TEST_CONN")
	if testConn == "" {
		return "postgresql://postgres:pass@localhost:5432/postgres?sslmode=disable"
	}
	return testConn
}

var strategy = destination.MigrateStrategy{
	AddColumn:           specs.MigrateModeSafe,
	AddColumnNotNull:    specs.MigrateModeForced,
	RemoveColumn:        specs.MigrateModeSafe,
	RemoveColumnNotNull: specs.MigrateModeForced,
	ChangeColumn:        specs.MigrateModeForced,
}

func TestPgPlugin(t *testing.T) {
	destination.PluginTestSuiteRunner(t,
		func() *destination.Plugin {
			return destination.NewPlugin("postgresql", "development", New)
		},
		Spec{
			ConnectionString: getTestConnection(),
			PgxLogLevel:      LogLevelTrace,
		},
		destination.PluginTestSuiteTests{
			MigrateStrategyOverwrite: strategy,
			MigrateStrategyAppend:    strategy,
		})
}

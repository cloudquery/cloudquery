package client

import (
	"os"
	"testing"

	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/plugins/destination"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

func getTestConnection() string {
	testConn := os.Getenv("CQ_DEST_PG_TEST_CONN")
	if testConn == "" {
		return "postgresql://postgres:pass@localhost:5432/postgres?sslmode=disable"
	}
	return testConn
}

var strategy = plugin.MigrateStrategy{
	AddColumn:           plugin.MigrateModeSafe,
	AddColumnNotNull:    plugin.MigrateModeForce,
	RemoveColumn:        plugin.MigrateModeSafe,
	RemoveColumnNotNull: plugin.MigrateModeForce,
	ChangeColumn:        plugin.MigrateModeForce,
}

func TestPgPlugin(t *testing.T) {
	plugin.PluginTestSuiteRunner(t,
		func() *destination.Plugin {
			return destination.NewPlugin("postgresql", "development", New)
		},
		specs.Destination{
			Spec: &Spec{
				ConnectionString: getTestConnection(),
				PgxLogLevel:      LogLevelTrace,
			},
		},
		destination.PluginTestSuiteTests{
			MigrateStrategyOverwrite: strategy,
			MigrateStrategyAppend:    strategy,
		},
		destination.WithTestSourceSkipMaps(),
	)
}

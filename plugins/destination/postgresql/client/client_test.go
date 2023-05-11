package client

import (
	"os"
	"testing"

	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v2/plugins/destination"
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
		specs.Destination{
			Spec: &Spec{
				ConnectionString: getTestConnection(),
				PgxLogLevel:      LogLevelTrace,
			},
		},
		destination.PluginTestSuiteTests{
			MigrateStrategyOverwrite: strategy,
			MigrateStrategyAppend:    strategy,
		})
}

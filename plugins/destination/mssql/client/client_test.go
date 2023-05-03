package client

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/destination/mssql/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v2/plugins/destination"
	"github.com/cloudquery/plugin-sdk/v2/specs"
)

var migrateStrategy = destination.MigrateStrategy{
	AddColumn:           specs.MigrateModeSafe,
	AddColumnNotNull:    specs.MigrateModeForced,
	RemoveColumn:        specs.MigrateModeSafe,
	RemoveColumnNotNull: specs.MigrateModeForced,
	ChangeColumn:        specs.MigrateModeForced,
}

func getTestConnection() string {
	if testConn := os.Getenv("CQ_DEST_MS_SQL_TEST_CONN"); len(testConn) > 0 {
		return testConn
	}

	return fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		"localhost", "SA", "yourStrongP@ssword", 1433, "cloudquery",
	)
}

func TestPlugin(t *testing.T) {
	destination.PluginTestSuiteRunner(t,
		func() *destination.Plugin {
			return destination.NewPlugin("mssql", plugin.Version, New, destination.WithManagedWriter())
		},
		specs.Destination{
			Spec: &Spec{
				ConnectionString: getTestConnection(),
			},
		},
		destination.PluginTestSuiteTests{
			MigrateStrategyOverwrite: migrateStrategy,
			MigrateStrategyAppend:    migrateStrategy,
		},
	)
}

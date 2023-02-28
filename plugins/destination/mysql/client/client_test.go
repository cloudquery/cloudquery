package client

import (
	"os"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/destination/mysql/resources/plugin"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/specs"
)

var migrateStrategy = destination.MigrateStrategy{
	AddColumn:           specs.MigrateModeSafe,
	AddColumnNotNull:    specs.MigrateModeForced,
	RemoveColumn:        specs.MigrateModeSafe,
	RemoveColumnNotNull: specs.MigrateModeForced,
	ChangeColumn:        specs.MigrateModeForced,
}

func getConnectionString() string {
	if testConn := os.Getenv("CQ_DEST_MYSQL_TEST_CONNECTION_STRING"); len(testConn) > 0 {
		return testConn
	}

	return `root:test@/cloudquery`
}

func TestPlugin(t *testing.T) {
	destination.PluginTestSuiteRunner(t,
		func() *destination.Plugin {
			return destination.NewPlugin("mysql", plugin.Version, New, destination.WithManagedWriter())
		},
		Spec{ConnectionString: getConnectionString()},
		destination.PluginTestSuiteTests{
			MigrateStrategyOverwrite: migrateStrategy,
			MigrateStrategyAppend:    migrateStrategy,
		},
	)
}

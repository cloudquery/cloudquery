package client

import (
	"os"
	"testing"

	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v2/plugins/destination"
)

var migrateStrategy = destination.MigrateStrategy{
	AddColumn:           specs.MigrateModeSafe,
	AddColumnNotNull:    specs.MigrateModeForced,
	RemoveColumn:        specs.MigrateModeSafe,
	RemoveColumnNotNull: specs.MigrateModeForced,
	ChangeColumn:        specs.MigrateModeForced,
}

func getTestConnection() string {
	testConn := os.Getenv("CQ_DEST_MONGODB_TEST_CONN")
	if testConn == "" {
		return "mongodb://localhost:27017"
	}
	return testConn
}

func TestPlugin(t *testing.T) {
	destination.PluginTestSuiteRunner(t,
		func() *destination.Plugin {
			return destination.NewPlugin("mongodb", "development", New, destination.WithManagedWriter())
		},
		specs.Destination{
			Spec: &Spec{
				ConnectionString: getTestConnection(),
				Database:         "destination_mongodb_test",
			},
		},
		destination.PluginTestSuiteTests{
			SkipMigrateOverwriteForce: true,
			SkipMigrateAppendForce:    true,

			MigrateStrategyOverwrite: migrateStrategy,
			MigrateStrategyAppend:    migrateStrategy,
		})
}

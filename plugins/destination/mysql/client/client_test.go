package client

import (
	"os"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/destination/mysql/resources/plugin"
	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/plugins/destination"
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
		specs.Destination{
			Spec: &Spec{
				ConnectionString: getConnectionString(),
			},
		},
		destination.PluginTestSuiteTests{
			SkipOverwrite: true,

			SkipDeleteStale: true,
			SkipAppend:      false,

			SkipSecondAppend: true,

			// SkipMigrateAppend skips a test for the migrate function where a column is added,
			// data is appended, then the column is removed and more data appended, checking that the migrations handle
			// this correctly.
			SkipMigrateAppend: true,
			// SkipMigrateAppendForce skips a test for the migrate function where a column is changed in force mode
			SkipMigrateAppendForce: true,

			// SkipMigrateOverwrite skips a test for the migrate function where a column is added,
			// data is appended, then the column is removed and more data overwritten, checking that the migrations handle
			// this correctly.
			SkipMigrateOverwrite: true,
			// SkipMigrateOverwriteForce skips a test for the migrate function where a column is changed in force mode
			SkipMigrateOverwriteForce: true,
			MigrateStrategyOverwrite:  migrateStrategy,
			MigrateStrategyAppend:     migrateStrategy,
		},
	)
}

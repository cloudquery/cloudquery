package client

import (
	"net/url"
	"os"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/resources/plugin"
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
	if testConn := os.Getenv("CQ_DEST_CH_TEST_CONN"); len(testConn) > 0 {
		return testConn
	}

	return (&url.URL{
		User: url.UserPassword("cq", "test"),
		Host: "localhost:9000",
		Path: "cloudquery", // database
	}).String()
}

func TestPlugin(t *testing.T) {
	destination.PluginTestSuiteRunner(t,
		func() *destination.Plugin {
			return destination.NewPlugin(
				"clickhouse",
				plugin.Version,
				New,
				destination.WithManagedWriter(),
			)
		},
		specs.Destination{
			Spec: &Spec{
				ConnectionString: getTestConnection(),
			},
		},
		destination.PluginTestSuiteTests{
			SkipOverwrite:             true,
			SkipMigrateOverwrite:      true,
			SkipMigrateOverwriteForce: true,

			MigrateStrategyOverwrite: migrateStrategy,
			MigrateStrategyAppend:    migrateStrategy,
		},
	)
}

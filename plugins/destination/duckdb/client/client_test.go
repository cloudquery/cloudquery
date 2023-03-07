package client

import (
	"testing"

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

func TestPlugin(t *testing.T) {
	destination.PluginTestSuiteRunner(t,
		func() *destination.Plugin {
			return destination.NewPlugin("duckdb", "development", New)
		},
		Spec{
			ConnectionString: "", // in-memory database
		},
		destination.PluginTestSuiteTests{
			SkipOverwrite:             true,
			SkipDeleteStale:           true,
			SkipMigrateOverwrite:      true,
			SkipMigrateOverwriteForce: true,
			MigrateStrategyOverwrite:  migrateStrategy,
			MigrateStrategyAppend:     migrateStrategy,
		})
}

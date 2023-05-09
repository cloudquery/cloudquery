package client

import (
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

func TestPlugin(t *testing.T) {
	destination.PluginTestSuiteRunner(t,
		func() *destination.Plugin {
			return destination.NewPlugin("duckdb", "development", New, destination.WithManagedWriter())
		},
		specs.Destination{
			Spec: &Spec{
				ConnectionString: "",
			},
		},
		destination.PluginTestSuiteTests{
			MigrateStrategyOverwrite: migrateStrategy,
			MigrateStrategyAppend:    migrateStrategy,
		})
}

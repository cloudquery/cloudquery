package client

import (
	"testing"

	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/plugins/destination"
	"github.com/cloudquery/plugin-sdk/v3/types"
)

var migrateStrategy = destination.MigrateStrategy{
	AddColumn:           specs.MigrateModeSafe,
	AddColumnNotNull:    specs.MigrateModeForced,
	RemoveColumn:        specs.MigrateModeSafe,
	RemoveColumnNotNull: specs.MigrateModeForced,
	ChangeColumn:        specs.MigrateModeForced,
}

func TestPlugin(t *testing.T) {
	if err := types.RegisterAllExtensions(); err != nil {
		t.Fatal(err)
	}

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
			// SkipOverwrite:            true,
		},
		// not supported in Parquet Writer
		destination.WithTestSourceSkipIntervals(),
		destination.WithTestSourceSkipDurations(),

		// not supported in duckDB for now
		destination.WithTestSourceSkipTimes(),
		destination.WithTestSourceSkipDates(),
		destination.WithTestSourceSkipLargeTypes(),
	)
}

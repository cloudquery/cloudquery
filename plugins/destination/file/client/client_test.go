package client

import (
	"testing"

	"github.com/cloudquery/filetypes"
	"github.com/cloudquery/plugin-sdk/v2/plugins/destination"
	"github.com/cloudquery/plugin-sdk/v2/specs"
)

var migrateStrategy = destination.MigrateStrategy{
	AddColumn:           specs.MigrateModeForced,
	AddColumnNotNull:    specs.MigrateModeForced,
	RemoveColumn:        specs.MigrateModeForced,
	RemoveColumnNotNull: specs.MigrateModeForced,
	ChangeColumn:        specs.MigrateModeForced,
}

func TestPluginCSV(t *testing.T) {
	destination.PluginTestSuiteRunner(t,
		func() *destination.Plugin {
			return destination.NewPlugin("file", "development", New, destination.WithManagedWriter())
		},
		specs.Destination{
			Spec: &Spec{
				Directory: t.TempDir(),
				FileSpec: &filetypes.FileSpec{
					Format: filetypes.FormatTypeCSV,
				},
				NoRotate: true,
			},
		},
		destination.PluginTestSuiteTests{
			SkipOverwrite:             true,
			SkipDeleteStale:           true,
			SkipMigrateOverwrite:      true,
			SkipMigrateOverwriteForce: true,
			SkipMigrateAppendForce:    true,

			MigrateStrategyOverwrite: migrateStrategy,
			MigrateStrategyAppend:    migrateStrategy,
		},
	)
}

func TestPluginJSON(t *testing.T) {
	destination.PluginTestSuiteRunner(t,
		func() *destination.Plugin {
			return destination.NewPlugin("file", "development", New, destination.WithManagedWriter())
		},
		specs.Destination{
			Spec: &Spec{
				Directory: t.TempDir(),
				FileSpec: &filetypes.FileSpec{
					Format: filetypes.FormatTypeJSON,
				},
				NoRotate: true,
			},
		},
		destination.PluginTestSuiteTests{
			SkipOverwrite:             true,
			SkipDeleteStale:           true,
			SkipMigrateOverwrite:      true,
			SkipMigrateOverwriteForce: true,
			SkipMigrateAppendForce:    true,

			MigrateStrategyOverwrite: migrateStrategy,
			MigrateStrategyAppend:    migrateStrategy,
		},
	)
}

func TestPluginParquet(t *testing.T) {
	destination.PluginTestSuiteRunner(t,
		func() *destination.Plugin {
			return destination.NewPlugin("file", "development", New, destination.WithManagedWriter())
		},
		specs.Destination{
			Spec: &Spec{
				Directory: t.TempDir(),
				FileSpec: &filetypes.FileSpec{
					Format: filetypes.FormatTypeParquet,
				},
				NoRotate: true,
			},
		},
		destination.PluginTestSuiteTests{
			SkipOverwrite:             true,
			SkipDeleteStale:           true,
			SkipSecondAppend:          true,
			SkipMigrateOverwrite:      true,
			SkipMigrateOverwriteForce: true,
			SkipMigrateAppendForce:    true,

			MigrateStrategyOverwrite: migrateStrategy,
			MigrateStrategyAppend:    migrateStrategy,
		},
	)
}

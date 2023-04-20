package client

import (
	"path"
	"testing"

	"github.com/cloudquery/filetypes/v2"
	"github.com/cloudquery/filetypes/v2/csv"
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
	t.Run("Directory", func(t *testing.T) {
		spec := &Spec{
			Directory: t.TempDir(),
			FileSpec: &filetypes.FileSpec{
				Format: filetypes.FormatTypeCSV,
				FormatSpec: csv.Spec{
					SkipHeader: true,
					Delimiter:  ",",
				},
			},
			NoRotate: true,
		}
		testPlugin(t, spec)
	})

	t.Run("Path", func(t *testing.T) {
		spec := &Spec{
			Path: path.Join(t.TempDir(), "{{TABLE}}.{{FORMAT}}"),
			FileSpec: &filetypes.FileSpec{
				Format: filetypes.FormatTypeCSV,
				FormatSpec: csv.Spec{
					SkipHeader: true,
					Delimiter:  ",",
				},
			},
			NoRotate: true,
		}
		testPlugin(t, spec)
	})
}

func TestPluginJSON(t *testing.T) {
	t.Run("Directory", func(t *testing.T) {
		spec := &Spec{
			Directory: t.TempDir(),
			FileSpec: &filetypes.FileSpec{
				Format: filetypes.FormatTypeJSON,
			},
			NoRotate: true,
		}
		testPlugin(t, spec)
	})
	t.Run("Path", func(t *testing.T) {
		spec := &Spec{
			Path: path.Join(t.TempDir(), "{{TABLE}}.{{FORMAT}}"),
			FileSpec: &filetypes.FileSpec{
				Format: filetypes.FormatTypeJSON,
			},
			NoRotate: true,
		}
		testPlugin(t, spec)
	})
}

func TestPluginParquet(t *testing.T) {
	t.Run("Directory", func(t *testing.T) {
		spec := &Spec{
			Directory: t.TempDir(),
			FileSpec: &filetypes.FileSpec{
				Format: filetypes.FormatTypeParquet,
			},
			NoRotate: true,
		}
		testPlugin(t, spec)
	})
	t.Run("Path", func(t *testing.T) {
		spec := &Spec{
			Path: path.Join(t.TempDir(), "{{TABLE}}.{{FORMAT}}"),
			FileSpec: &filetypes.FileSpec{
				Format: filetypes.FormatTypeParquet,
			},
			NoRotate: true,
		}
		testPlugin(t, spec)
	})
}

func testPlugin(t *testing.T, spec *Spec) {
	destination.PluginTestSuiteRunner(t,
		func() *destination.Plugin {
			return destination.NewPlugin("file", "development", New, destination.WithManagedWriter())
		},
		specs.Destination{
			Spec: spec,
		},
		destination.PluginTestSuiteTests{
			SkipOverwrite:             true,
			SkipSecondAppend:          true,
			SkipDeleteStale:           true,
			SkipMigrateOverwrite:      true,
			SkipMigrateOverwriteForce: true,
			SkipMigrateAppendForce:    true,

			MigrateStrategyOverwrite: migrateStrategy,
			MigrateStrategyAppend:    migrateStrategy,
		},
	)
}

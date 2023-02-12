package client

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
)

func TestPluginCSV(t *testing.T) {
	destination.PluginTestSuiteRunner(t,
		func() *destination.Plugin {
			return destination.NewPlugin("file", "development", New, destination.WithManagedWriter())
		},
		Spec{
			Directory: t.TempDir(),
			Format:    FormatTypeCSV,
			NoRotate:  true,
		},
		destination.PluginTestSuiteTests{
			SkipOverwrite:             true,
			SkipDeleteStale:           true,
			SkipMigrateAppend:         true,
			SkipMigrateOverwrite:      true,
			SkipMigrateOverwriteForce: true,
		},
	)
}

func TestPluginJSON(t *testing.T) {
	destination.PluginTestSuiteRunner(t,
		func() *destination.Plugin {
			return destination.NewPlugin("file", "development", New, destination.WithManagedWriter())
		},
		Spec{
			Directory: t.TempDir(),
			Format:    FormatTypeJSON,
			NoRotate:  true,
		},
		destination.PluginTestSuiteTests{
			SkipOverwrite:             true,
			SkipDeleteStale:           true,
			SkipMigrateOverwrite:      true,
			SkipMigrateOverwriteForce: true,
		},
	)
}

func TestPluginParquet(t *testing.T) {
	destination.PluginTestSuiteRunner(t,
		func() *destination.Plugin {
			return destination.NewPlugin("file", "development", New, destination.WithManagedWriter())
		},
		Spec{
			Directory: t.TempDir(),
			Format:    FormatTypeParquet,
			NoRotate:  true,
		},
		destination.PluginTestSuiteTests{
			SkipOverwrite:        true,
			SkipDeleteStale:      true,
			SkipSecondAppend:     true,
			SkipMigrateAppend:    true,
			SkipMigrateOverwrite: true,
		},
	)
}

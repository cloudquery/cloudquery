package client

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
)

func TestPluginCSV(t *testing.T) {
	p := destination.NewPlugin("file", "development", New, destination.WithManagedWriter())

	destination.PluginTestSuiteRunner(t, p,
		Spec{
			Directory: t.TempDir(),
			Format:    FormatTypeCSV,
			NoRotate:  true,
		},
		destination.PluginTestSuiteTests{
			SkipOverwrite:        true,
			SkipDeleteStale:      true,
			SkipMigrateAppend:    true,
			SkipMigrateOverwrite: true,
		},
	)
}

func TestPluginJSON(t *testing.T) {
	p := destination.NewPlugin("file", "development", New, destination.WithManagedWriter())

	destination.PluginTestSuiteRunner(t, p,
		Spec{
			Directory: t.TempDir(),
			Format:    FormatTypeJSON,
			NoRotate:  true,
		},
		destination.PluginTestSuiteTests{
			SkipOverwrite:        true,
			SkipDeleteStale:      true,
			SkipMigrateOverwrite: true,
		},
	)
}

func TestPluginParquet(t *testing.T) {
	p := destination.NewPlugin("file", "development", New, destination.WithManagedWriter())

	destination.PluginTestSuiteRunner(t, p,
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

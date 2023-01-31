package client

import (
	"testing"

	"github.com/cloudquery/filetypes"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
)

func TestPluginCSV(t *testing.T) {
	p := destination.NewPlugin("file", "development", New, destination.WithManagedWriter())

	destination.PluginTestSuiteRunner(t, p,
		Spec{
			Directory: t.TempDir(),
			FileSpec:  &filetypes.FileSpec{Format: filetypes.FormatTypeCSV},
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
			FileSpec:  &filetypes.FileSpec{Format: filetypes.FormatTypeJSON},
			NoRotate:  true,
		},
		destination.PluginTestSuiteTests{
			SkipOverwrite:        true,
			SkipDeleteStale:      true,
			SkipMigrateOverwrite: true,
		},
	)
}

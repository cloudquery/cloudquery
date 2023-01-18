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
			SkipOverwrite:     true,
			SkipDeleteStale:   true,
			SkipMigrateAppend: true,
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
			SkipOverwrite:   true,
			SkipDeleteStale: true,
		},
	)
}

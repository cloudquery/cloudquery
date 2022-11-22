package client

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/plugins"
)

func TestPluginCSVLocal(t *testing.T) {
	p := plugins.NewDestinationPlugin("file", "development", New)
	plugins.DestinationPluginTestSuiteRunner(t, p,
		Spec{
			Directory: t.TempDir(),
			Backend:   BackendTypeLocal,
			Format:    FormatTypeCSV,
		},
		plugins.DestinationTestSuiteTests{
			SkipOverwrite:   true,
			SkipDeleteStale: true,
		},
	)
}

func TestPluginCSVGCS(t *testing.T) {
	p := plugins.NewDestinationPlugin("file", "development", New)
	plugins.DestinationPluginTestSuiteRunner(t, p,
		Spec{
			Directory: "cq-yev-test/dest-plugin-file",
			Backend:   BackendTypeGCS,
			Format:    FormatTypeCSV,
		},
		plugins.DestinationTestSuiteTests{
			SkipOverwrite:   true,
			SkipDeleteStale: true,
			SkipAppendTwice: true,
		},
	)
}

func TestPluginCSVS3(t *testing.T) {
	p := plugins.NewDestinationPlugin("file", "development", New)
	plugins.DestinationPluginTestSuiteRunner(t, p,
		Spec{
			Directory: "cq-yev-test/dest-plugin-file",
			Backend:   BackendTypeS3,
			Format:    FormatTypeCSV,
		},
		plugins.DestinationTestSuiteTests{
			SkipOverwrite:   true,
			SkipDeleteStale: true,
			SkipAppendTwice: true,
		},
	)
}
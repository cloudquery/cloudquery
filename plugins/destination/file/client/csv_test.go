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
			NoRotate: true,
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
			Directory: "cq-yev-test/test-csv-gcs",
			Backend:   BackendTypeGCS,
			Format:    FormatTypeCSV,
			NoRotate: true,
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
			Directory: "cq-playground-test/dest-plugin-file",
			Backend:   BackendTypeS3,
			Format:    FormatTypeCSV,
			NoRotate: true,
		},
		plugins.DestinationTestSuiteTests{
			SkipOverwrite:   true,
			SkipDeleteStale: true,
			SkipAppendTwice: true,
		},
	)
}
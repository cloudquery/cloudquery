package client

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/plugins"
)

func TestPluginJSONLocal(t *testing.T) {
	p := plugins.NewDestinationPlugin("file", "development", New)
	plugins.DestinationPluginTestSuiteRunner(t, p,
		Spec{
			Directory: t.TempDir(),
			Backend:   BackendTypeLocal,
			Format:    FormatTypeJSON,
			NoRotate:  true,
		},
		plugins.DestinationTestSuiteTests{
			SkipOverwrite:   true,
			SkipDeleteStale: true,
		},
	)
}

func TestPluginJSONGCS(t *testing.T) {
	p := plugins.NewDestinationPlugin("file", "development", New)
	plugins.DestinationPluginTestSuiteRunner(t, p,
		Spec{
			Directory: "cq-yev-test/dest-plugin-file",
			Backend:   BackendTypeGCS,
			Format:    FormatTypeJSON,
			NoRotate:  true,
		},
		plugins.DestinationTestSuiteTests{
			SkipOverwrite:   true,
			SkipDeleteStale: true,
			SkipAppendTwice: true,
		},
	)
}

func TestPluginJSONAWS(t *testing.T) {
	p := plugins.NewDestinationPlugin("file", "development", New)
	plugins.DestinationPluginTestSuiteRunner(t, p,
		Spec{
			Directory: "cq-playground-test/dest-plugin-file",
			Backend:   BackendTypeS3,
			Format:    FormatTypeJSON,
			NoRotate:  true,
		},
		plugins.DestinationTestSuiteTests{
			SkipOverwrite:   true,
			SkipDeleteStale: true,
			SkipAppendTwice: true,
		},
	)
}

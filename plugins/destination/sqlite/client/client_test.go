package client

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/specs"
)

func TestPlugin(t *testing.T) {
	p := plugins.NewDestinationPlugin("postgresql", "development", New)

	plugins.DestinationPluginTestSuiteRunner(t, p, specs.Destination{
		WriteMode: specs.WriteModeAppend,
		Spec: Spec{
			ConnectionString: ":memory:",
		},
	})
}

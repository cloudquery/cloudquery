package client

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/plugins"
)

func TestPlugin(t *testing.T) {
	p := plugins.NewDestinationPlugin("sqlite", "development", New)

	plugins.DestinationPluginTestSuiteRunner(t, p,
		Spec{
			ConnectionString: ":memory:",
		},
		plugins.DestinationTestSuiteTests{})
}

package client

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
)

func TestPlugin(t *testing.T) {
	p := destination.NewPlugin("sqlite", "development", New)

	destination.PluginTestSuiteRunner(t, p,
		Spec{
			ConnectionString: ":memory:",
		},
		destination.PluginTestSuiteTests{})
}

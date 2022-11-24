package client

import (
	"os"
	"testing"

	"github.com/cloudquery/plugin-sdk/plugins"
)

func TestPlugin(t *testing.T) {
	p := plugins.NewDestinationPlugin("postgresql", "development", New)
	plugins.DestinationPluginTestSuiteRunner(t, p,
		Spec{
			DSN: os.Getenv("SNOW_TEST_DSN"),
		},
		plugins.DestinationTestSuiteTests{
			SkipOverwrite:   true,
		})
}

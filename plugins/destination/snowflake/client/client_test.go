package client

import (
	"os"
	"testing"

	"github.com/cloudquery/plugin-sdk/plugins"
)

func TestPlugin(t *testing.T) {
	p := plugins.NewDestinationPlugin("snowflake", "development", New)
	plugins.DestinationPluginTestSuiteRunner(t, p,
		Spec{
			ConnectionString: os.Getenv("SNOW_TEST_DSN"),
		},
		plugins.DestinationTestSuiteTests{
			SkipOverwrite: true,
		})
}

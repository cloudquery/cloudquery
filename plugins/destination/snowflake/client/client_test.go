package client

import (
	"os"
	"testing"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
)

func TestPlugin(t *testing.T) {
	p := destination.NewPlugin("snowflake", "development", New, destination.WithManagedWriter())
	destination.PluginTestSuiteRunner(t, p,
		Spec{
			ConnectionString: os.Getenv("SNOW_TEST_DSN"),
		},
		destination.PluginTestSuiteTests{
			SkipOverwrite: true,
		})
}

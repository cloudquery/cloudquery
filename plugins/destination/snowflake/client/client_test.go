package client

import (
	"os"
	"testing"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
)

func TestPlugin(t *testing.T) {
	destination.PluginTestSuiteRunner(t,
		func() *destination.Plugin {
			return destination.NewPlugin("snowflake", "development", New, destination.WithManagedWriter())
		},
		Spec{
			ConnectionString: os.Getenv("SNOW_TEST_DSN"),
		},
		destination.PluginTestSuiteTests{
			SkipOverwrite:             true,
			SkipMigrateAppend:         true, // fails with `invalid identifier '"new_column"'`, maybe because delays in schema propagation?
			SkipMigrateOverwrite:      true,
			SkipMigrateOverwriteForce: true,
			SkipMigrateAppendForce:    true,
		})
}

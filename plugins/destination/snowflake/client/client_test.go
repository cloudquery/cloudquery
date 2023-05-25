package client

import (
	"os"
	"testing"

	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/plugins/destination"
)

func TestPlugin(t *testing.T) {
	destination.PluginTestSuiteRunner(t,
		func() *destination.Plugin {
			return destination.NewPlugin("snowflake", "development", New, destination.WithManagedWriter())
		},
		specs.Destination{
			Spec: &Spec{
				ConnectionString: os.Getenv("SNOW_TEST_DSN"),
			},
		},
		destination.PluginTestSuiteTests{
			SkipOverwrite:             true,
			SkipMigrateAppend:         true, // fails with `invalid identifier '"new_column"'`, maybe because delays in schema propagation?
			SkipMigrateOverwrite:      true,
			SkipMigrateOverwriteForce: true,
			SkipMigrateAppendForce:    true,
		},
		destination.WithTestSourceSkipIntervals(),
		destination.WithTestSourceSkipDates(),
		
	)
}

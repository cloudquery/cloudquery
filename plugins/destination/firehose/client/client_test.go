package client

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/v2/plugins/destination"
	"github.com/cloudquery/plugin-sdk/v2/specs"
)

const streamARN = "cq-playground-test"

func TestPluginJSON(t *testing.T) {
	destination.PluginTestSuiteRunner(t,
		func() *destination.Plugin {
			return destination.NewPlugin("firehose", "development", New, destination.WithDefaultBatchSize(500))
		},
		specs.Destination{
			Spec: &Spec{
				NoRotate:  true,
				StreamARN: streamARN,
			},
		},
		destination.PluginTestSuiteTests{
			SkipAppend:                true,
			SkipDeleteStale:           true,
			SkipMigrateAppend:         true,
			SkipOverwrite:             true,
			SkipSecondAppend:          true,
			SkipMigrateOverwrite:      true,
			SkipMigrateOverwriteForce: true,
		},
	)
}

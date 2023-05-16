package client

import (
	"testing"

	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v2/plugins/destination"
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
			SkipOverwrite:             true,
			SkipDeleteStale:           true,
			SkipAppend:                true,
			SkipSecondAppend:          true,
			SkipMigrateAppend:         true,
			SkipMigrateAppendForce:    true,
			SkipMigrateOverwrite:      true,
			SkipMigrateOverwriteForce: true,
		},
	)
}

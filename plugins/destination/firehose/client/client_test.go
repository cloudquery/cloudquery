package client

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
)

const streamARN = "cq-playground-test"

func TestPluginJSON(t *testing.T) {
	destination.PluginTestSuiteRunner(t,
		func() *destination.Plugin {
			return destination.NewPlugin("firehose", "development", New, destination.WithDefaultBatchSize(500))
		},
		Spec{
			NoRotate:  true,
			StreamARN: streamARN,
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

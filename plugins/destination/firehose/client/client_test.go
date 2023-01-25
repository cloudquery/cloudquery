package client

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
)

const streamARN = "cq-playground-test"

func TestPluginJSON(t *testing.T) {
	p := destination.NewPlugin("firehose", "development", New, destination.WithDefaultBatchSize(500))
	destination.PluginTestSuiteRunner(t, p,
		Spec{
			NoRotate:  true,
			StreamARN: streamARN,
		},
		destination.PluginTestSuiteTests{
			SkipOverwrite:     true,
			SkipDeleteStale:   true,
			SkipSecondAppend:  true,
			SkipMigrateAppend: true,
			SkipAppend:        true,
		},
	)
}

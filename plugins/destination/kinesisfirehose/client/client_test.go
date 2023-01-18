package client

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
)

const streamARN = "cq-playground-test"

func TestPluginJSON(t *testing.T) {
	p := destination.NewPlugin("kinesisfirehose", "development", New, destination.WithManagedWriter())

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

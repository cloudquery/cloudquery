package client

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
)

const bucket = "cqdestinationazblob"
const container = "test"

func TestPluginCSV(t *testing.T) {
	p := destination.NewPlugin("azblob", "development", New, destination.WithManagedWriter())

	destination.PluginTestSuiteRunner(t, p,
		Spec{
			Bucket:    bucket,
			Container: container,
			Path:      t.TempDir(),
			Format:    FormatTypeCSV,
			NoRotate:  true,
		},
		destination.PluginTestSuiteTests{
			SkipOverwrite:    true,
			SkipDeleteStale:  true,
			SkipSecondAppend: true,
		},
	)
}

func TestPluginJSON(t *testing.T) {
	p := destination.NewPlugin("azblob", "development", New, destination.WithManagedWriter())

	destination.PluginTestSuiteRunner(t, p,
		Spec{
			Bucket:    bucket,
			Container: container,
			Path:      t.TempDir(),
			Format:    FormatTypeJSON,
			NoRotate:  true,
		},
		destination.PluginTestSuiteTests{
			SkipOverwrite:    true,
			SkipDeleteStale:  true,
			SkipSecondAppend: true,
		},
	)
}

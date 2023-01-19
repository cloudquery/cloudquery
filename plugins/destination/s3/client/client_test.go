package client

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
)

const bucket = "cq-playground-test"

func TestPluginCSV(t *testing.T) {
	p := destination.NewPlugin("s3", "development", New, destination.WithManagedWriter())
	spec := Spec{
		Bucket:   bucket,
		Path:     t.TempDir(),
		Format:   FormatTypeCSV,
		NoRotate: true,
	}
	spec.SetDefaults()
	destination.PluginTestSuiteRunner(t, p,
		spec,
		destination.PluginTestSuiteTests{
			SkipOverwrite:     true,
			SkipDeleteStale:   true,
			SkipSecondAppend:  true,
			SkipMigrateAppend: true,
		},
	)
}

func TestPluginJSON(t *testing.T) {
	p := destination.NewPlugin("s3", "development", New, destination.WithManagedWriter())
	spec := Spec{
		Bucket:   bucket,
		Path:     t.TempDir(),
		Format:   FormatTypeJSON,
		NoRotate: true,
	}
	spec.SetDefaults()
	destination.PluginTestSuiteRunner(t, p,
		spec,
		destination.PluginTestSuiteTests{
			SkipOverwrite:     true,
			SkipDeleteStale:   true,
			SkipSecondAppend:  true,
			SkipMigrateAppend: true,
		},
	)
}

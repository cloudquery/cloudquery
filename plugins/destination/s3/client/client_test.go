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
		Path:     t.TempDir()[1:],
		Format:   FormatTypeCSV,
		NoRotate: true,
	}
	spec.SetDefaults()
	destination.PluginTestSuiteRunner(t, p,
		spec,
		destination.PluginTestSuiteTests{
			SkipOverwrite:        true,
			SkipDeleteStale:      true,
			SkipSecondAppend:     true,
			SkipMigrateAppend:    true,
			SkipMigrateOverwrite: true,
		},
	)
}

func TestPluginJSON(t *testing.T) {
	p := destination.NewPlugin("s3", "development", New, destination.WithManagedWriter())
	spec := Spec{
		Bucket:   bucket,
		Path:     t.TempDir()[1:],
		Format:   FormatTypeJSON,
		NoRotate: true,
	}
	spec.SetDefaults()
	destination.PluginTestSuiteRunner(t, p,
		spec,
		destination.PluginTestSuiteTests{
			SkipOverwrite:        true,
			SkipDeleteStale:      true,
			SkipSecondAppend:     true,
			SkipMigrateAppend:    true,
			SkipMigrateOverwrite: true,
		},
	)
}

package client

import (
	"testing"

	"github.com/cloudquery/filetypes"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
)

const bucket = "cq-dest-gcs"

func TestPluginCSV(t *testing.T) {
	spec := Spec{
		Bucket:   bucket,
		Path:     t.TempDir(),
		NoRotate: true,
		FileSpec: &filetypes.FileSpec{
			Format: filetypes.FormatTypeCSV,
		},
	}
	spec.SetDefaults()
	destination.PluginTestSuiteRunner(t,
		func() *destination.Plugin {
			return destination.NewPlugin("gcs", "development", New, destination.WithManagedWriter())
		},
		spec,
		destination.PluginTestSuiteTests{
			SkipOverwrite:             true,
			SkipDeleteStale:           true,
			SkipSecondAppend:          true,
			SkipMigrateAppend:         true,
			SkipMigrateOverwrite:      true,
			SkipMigrateOverwriteForce: true,
			SkipMigrateAppendForce:    true,
		},
	)
}

func TestPluginJSON(t *testing.T) {
	spec := Spec{
		Bucket:   bucket,
		Path:     t.TempDir(),
		NoRotate: true,
		FileSpec: &filetypes.FileSpec{
			Format: filetypes.FormatTypeJSON,
		},
	}
	spec.SetDefaults()
	destination.PluginTestSuiteRunner(t,
		func() *destination.Plugin {
			return destination.NewPlugin("gcs", "development", New, destination.WithManagedWriter())
		},
		spec,
		destination.PluginTestSuiteTests{
			SkipOverwrite:             true,
			SkipDeleteStale:           true,
			SkipSecondAppend:          true,
			SkipMigrateAppend:         true,
			SkipMigrateOverwrite:      true,
			SkipMigrateOverwriteForce: true,
			SkipMigrateAppendForce:    true,
		},
	)
}

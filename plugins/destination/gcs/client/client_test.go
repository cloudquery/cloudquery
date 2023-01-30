package client

import (
	"testing"

	"github.com/cloudquery/filetypes"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
)

const bucket = "cq-dest-gcs"

func TestPluginCSV(t *testing.T) {
	p := destination.NewPlugin("gcs", "development", New, destination.WithManagedWriter())

	spec := Spec{
		Bucket:   bucket,
		Path:     t.TempDir(),
		NoRotate: true,
		FileSpec: &filetypes.FileSpec{
			Format: filetypes.FormatTypeCSV,
		},
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
	p := destination.NewPlugin("gcs", "development", New, destination.WithManagedWriter())
	spec := Spec{
		Bucket:   bucket,
		Path:     t.TempDir(),
		NoRotate: true,
		FileSpec: &filetypes.FileSpec{
			Format: filetypes.FormatTypeJSON,
		},
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

package client

import (
	"testing"

	"github.com/cloudquery/filetypes/v3"
	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/plugins/destination"
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
		specs.Destination{
			Spec: &spec,
		},
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
		specs.Destination{
			Spec: &spec,
		},
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

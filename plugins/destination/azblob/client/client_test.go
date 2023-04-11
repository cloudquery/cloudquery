package client

import (
	"testing"

	"github.com/cloudquery/filetypes"
	"github.com/cloudquery/plugin-sdk/v2/plugins/destination"
	"github.com/cloudquery/plugin-sdk/v2/specs"
)

const storage_account = "cqdestinationazblob"
const container = "test"

func TestPluginCSV(t *testing.T) {
	spec := Spec{
		StorageAccount: storage_account,
		Container:      container,
		Path:           t.TempDir(),
		NoRotate:       true,
		FileSpec: &filetypes.FileSpec{
			Format: filetypes.FormatTypeCSV,
		},
	}
	spec.SetDefaults()
	destination.PluginTestSuiteRunner(t,
		func() *destination.Plugin {
			return destination.NewPlugin("azblob", "development", New, destination.WithManagedWriter())
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
		StorageAccount: storage_account,
		Container:      container,
		Path:           t.TempDir(),
		NoRotate:       true,
		FileSpec: &filetypes.FileSpec{
			Format: filetypes.FormatTypeJSON,
		},
	}
	spec.SetDefaults()
	destination.PluginTestSuiteRunner(t,
		func() *destination.Plugin {
			return destination.NewPlugin("azblob", "development", New, destination.WithManagedWriter())
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

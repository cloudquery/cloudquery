package client

import (
	"testing"

	"github.com/cloudquery/filetypes"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
)

const storage_account = "cqdestinationazblob"
const container = "test"

func TestPluginCSV(t *testing.T) {
	p := destination.NewPlugin("azblob", "development", New, destination.WithManagedWriter())
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
	p := destination.NewPlugin("azblob", "development", New, destination.WithManagedWriter())
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

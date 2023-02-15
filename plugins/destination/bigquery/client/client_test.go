package client

import (
	"os"
	"testing"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
)

func TestPlugin(t *testing.T) {
	destination.PluginTestSuiteRunner(t,
		func() *destination.Plugin {
			return destination.NewPlugin("bigquery", "development", New, destination.WithManagedWriter())
		},
		Spec{
			ProjectID:        os.Getenv("BIGQUERY_PROJECT_ID"),
			DatasetID:        os.Getenv("BIGQUERY_DATASET_ID"),
			TimePartitioning: "none",
		},
		destination.PluginTestSuiteTests{
			SkipOverwrite:             true,
			SkipMigrateOverwrite:      true,
			SkipMigrateOverwriteForce: true,
			SkipMigrateAppendForce:    true,
		})
}

package client

import (
	"os"
	"testing"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
)

func TestPlugin(t *testing.T) {
	p := destination.NewPlugin("bigquery", "development", New, destination.WithManagedWriter())
	destination.PluginTestSuiteRunner(t, p,
		Spec{
			ProjectID:        os.Getenv("BIGQUERY_PROJECT_ID"),
			DatasetID:        os.Getenv("BIGQUERY_DATASET_ID"),
			TimePartitioning: "none",
		},
		destination.PluginTestSuiteTests{
			SkipOverwrite: true,
		})
}

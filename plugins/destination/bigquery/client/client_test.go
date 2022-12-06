package client

import (
	"os"
	"testing"

	"github.com/cloudquery/plugin-sdk/plugins"
)

func TestPlugin(t *testing.T) {
	t.Skip("Skipping until fixed: https://github.com/cloudquery/cloudquery/issues/5248")
	p := plugins.NewDestinationPlugin("bigquery", "development", New)
	plugins.DestinationPluginTestSuiteRunner(t, p,
		Spec{
			ProjectID:        os.Getenv("BIGQUERY_PROJECT_ID"),
			DatasetID:        os.Getenv("BIGQUERY_DATASET_ID"),
			TimePartitioning: "none",
		},
		plugins.DestinationTestSuiteTests{
			SkipOverwrite: true,
		})
}

package client

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func TestPlugin(t *testing.T) {
	ctx := context.Background()

	p := plugin.NewPlugin("bigquery", "development", New)
	spec := &Spec{
		ProjectID:        os.Getenv("BIGQUERY_PROJECT_ID"),
		DatasetID:        os.Getenv("BIGQUERY_DATASET_ID"),
		DatasetLocation:  os.Getenv("BIGQUERY_DATASET_LOCATION"),
		TimePartitioning: "none",
	}
	specBytes, err := json.Marshal(spec)
	if err != nil {
		t.Fatal(err)
	}
	if err := p.Init(ctx, specBytes, plugin.NewClientOptions{}); err != nil {
		t.Fatal(err)
	}

	plugin.TestWriterSuiteRunner(t,
		p,
		plugin.WriterTestSuiteTests{
			SkipUpsert:      true,
			SkipMigrate:     true,
			SkipDeleteStale: true,
		},
		plugin.WithTestDataOptions(schema.TestSourceOptions{
			SkipMaps: true,
			// https://github.com/cloudquery/cloudquery/issues/12022
			SkipTimes: true,
		}),
		plugin.WithTestIgnoreNullsInLists(),
	)
}

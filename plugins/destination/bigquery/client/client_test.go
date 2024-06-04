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
		ProjectID:        getEnvDefault("BIGQUERY_PROJECT_ID", "test"),
		DatasetID:        getEnvDefault("BIGQUERY_DATASET_ID", "test"),
		DatasetLocation:  os.Getenv("BIGQUERY_DATASET_LOCATION"),
		Endpoint:         getEnvDefault("BIGQUERY_ENDPOINT", "http://localhost:9050"),
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
			SkipUpsert:       true,
			SkipMigrate:      true,
			SkipDeleteStale:  true,
			SkipDeleteRecord: true,
		},
		plugin.WithTestDataOptions(schema.TestSourceOptions{
			SkipMaps: true,
			// https://github.com/cloudquery/cloudquery/issues/12022
			SkipTimes: true,
		}),
		plugin.WithTestIgnoreNullsInLists(),
	)
}

func getEnvDefault(key string, value string) string {
	v, ok := os.LookupEnv(key)
	if ok {
		return v
	}
	return value
}

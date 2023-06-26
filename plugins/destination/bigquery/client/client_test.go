package client

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

// var migrateStrategy = destination.MigrateStrategy{
// 	AddColumn:           specs.MigrateModeSafe,
// 	AddColumnNotNull:    specs.MigrateModeForced,
// 	RemoveColumn:        specs.MigrateModeSafe,
// 	RemoveColumnNotNull: specs.MigrateModeForced,
// 	ChangeColumn:        specs.MigrateModeForced,
// }

func TestPlugin(t *testing.T) {
	ctx := context.Background()

	p := plugin.NewPlugin("bigquery", "development", New)
	spec := &Spec{
		ProjectID:        os.Getenv("BIGQUERY_PROJECT_ID"),
		DatasetID:        os.Getenv("BIGQUERY_DATASET_ID"),
		TimePartitioning: "none",
	}
	specBytes, err := json.Marshal(spec)
	if err != nil {
		t.Fatal(err)
	}
	if err := p.Init(ctx, specBytes); err != nil {
		t.Fatal(err)
	}

	plugin.TestWriterSuiteRunner(t,
		p,
		plugin.WriterTestSuiteTests{
			SkipUpsert:  true,
			SkipMigrate: true,
		},
		plugin.WithTestDataOptions(schema.TestSourceOptions{
			SkipMaps: true,
		}),
		plugin.WithTestIgnoreNullsInLists(),
	)
}

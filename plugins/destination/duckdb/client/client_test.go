package client

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func TestPlugin(t *testing.T) {
	ctx := context.Background()
	p := plugin.NewPlugin("duckdb", "development", New)
	spec := Spec{
		ConnectionString: "?threads=1",
		Debug:            true,
	}
	specBytes, err := json.Marshal(spec)
	if err != nil {
		t.Fatal(err)
	}

	delayAfterDeleteStale = true
	if err := p.Init(ctx, specBytes, plugin.NewClientOptions{}); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := p.Close(ctx); err != nil {
			t.Logf("failed to close plugin: %v", err)
		}
	})

	plugin.TestWriterSuiteRunner(t,
		p,
		plugin.WriterTestSuiteTests{
			SkipDeleteRecord: true,
			SafeMigrations: plugin.SafeMigrations{
				AddColumn:    true,
				RemoveColumn: true,
			},
		},
		plugin.WithTestDataOptions(schema.TestSourceOptions{
			// not supported in Parquet Writer
			SkipDurations: true,
			SkipIntervals: true,
			// not supported in duckDB for now
			SkipLargeTypes: true,
			// not supported in Appender
			SkipMaps:    true,
			SkipStructs: true,
			// test:
			//SkipDates: true,
			SkipLists: true,
			//SkipDecimals: true,
			//SkipTimes:    true,
		}),
	)
}

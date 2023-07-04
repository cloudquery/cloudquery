package client

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

func TestPlugin(t *testing.T) {
	ctx := context.Background()
	if err := types.RegisterAllExtensions(); err != nil {
		t.Fatal(err)
	}

	p := plugin.NewPlugin("duckdb", "development", New)
	spec := Spec{
		ConnectionString: "?threads=1",
		Debug:            true,
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
			SkipTimes:      true,
			SkipDates:      true,
			SkipLargeTypes: true,
		}),
	)
}

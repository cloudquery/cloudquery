package client

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"
)

func TestPlugin(t *testing.T) {
	ctx := context.Background()
	p := plugin.NewPlugin("motherduck", "development", New)
	spec := Spec{
		ConnectionString: "?threads=1",
		Debug:            true,
	}
	specBytes, err := json.Marshal(spec)
	if err != nil {
		t.Fatal(err)
	}

	p.SetLogger(zerolog.New(zerolog.NewTestWriter(t)).Level(zerolog.DebugLevel))

	delayAfterDeleteStale = true
	localDuckDB = true // use local DuckDB to test
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
			SkipSpecificMigrations: plugin.Migrations{
				RemoveUniqueConstraint: true,
			},
		},
		plugin.WithTestDataOptions(schema.TestSourceOptions{
			// not supported in Parquet Writer
			SkipDurations: true,
			SkipIntervals: true,
		}),
	)
}

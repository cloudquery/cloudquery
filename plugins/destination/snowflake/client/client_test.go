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
	p := plugin.NewPlugin("snowflake", "development", New)
	spec := &Spec{
		ConnectionString: os.Getenv("SNOW_TEST_DSN"),
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
			SkipDeleteRecord: true,
			SkipSpecificWriteTests: plugin.WriteTests{
				DuplicatePK: true,
			},
			SafeMigrations: plugin.SafeMigrations{
				AddColumn:              true,
				AddColumnNotNull:       false,
				RemoveColumn:           true,
				RemoveColumnNotNull:    false,
				RemoveUniqueConstraint: false,
				MovePKToCQOnly:         true,
			},
		},
		plugin.WithTestDataOptions(schema.TestSourceOptions{
			SkipIntervals:  true,
			SkipMaps:       true,
			SkipLargeTypes: true,
			SkipLists:      true,
		}),
	)
}

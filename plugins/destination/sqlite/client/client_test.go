package client

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

func TestPlugin(t *testing.T) {
	ctx := context.Background()
	p := plugin.NewPlugin("sqlite", "development", New)
	spec := Spec{
		ConnectionString: ":memory:",
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
			SkipSpecificMigrations: plugin.Migrations{
				RemoveUniqueConstraint: true,
			},
		})
}

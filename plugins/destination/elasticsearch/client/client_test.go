package client

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

func TestPlugin(t *testing.T) {
	address := os.Getenv("ELASTICSEARCH_ADDRESS")
	if address == "" {
		address = "http://localhost:9200"
	}
	p := plugin.NewPlugin("elasticsearch", "development", New)
	spec := &Spec{
		Addresses: []string{address},
	}
	specBytes, err := json.Marshal(spec)
	if err != nil {
		t.Fatal(err)
	}
	if err := p.Init(context.Background(), specBytes, plugin.NewClientOptions{}); err != nil {
		t.Fatal(err)
	}
	plugin.TestWriterSuiteRunner(t,
		p,
		plugin.WriterTestSuiteTests{
			SafeMigrations: plugin.SafeMigrations{
				AddColumn:           true,
				RemoveColumn:        true,
				RemoveColumnNotNull: true,
				AddColumnNotNull:    true,
			},
		},
	)
}

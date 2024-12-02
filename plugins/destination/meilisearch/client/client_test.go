package client

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/meilisearch/v2/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

func getTestSpec() *Spec {
	apiKey := os.Getenv("CQ_DEST_MEILI_TEST_API_KEY")
	if len(apiKey) == 0 {
		apiKey = "test"
	}
	host := os.Getenv("CQ_DEST_MEILI_TEST_HOST")
	if len(host) == 0 {
		host = "http://localhost:7700"
	}

	return &Spec{Host: host, APIKey: apiKey}
}

func TestPlugin(t *testing.T) {
	ctx := context.Background()
	p := plugin.NewPlugin("meilisearch", internalPlugin.Version, New)
	spec := getTestSpec()
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
			SkipDeleteStale:  true,
			SkipDeleteRecord: true,
			SkipMigrate:      true,
			SafeMigrations: plugin.SafeMigrations{
				AddColumn:           true,
				AddColumnNotNull:    true,
				RemoveColumn:        true,
				RemoveColumnNotNull: true,
			},
		},
	)
}

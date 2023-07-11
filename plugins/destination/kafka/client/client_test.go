package client

import (
	"context"
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/cloudquery/filetypes/v4"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

const (
	defaultConnectionString = "localhost:29092"
)

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func TestPlugin(t *testing.T) {
	ctx := context.Background()
	p := plugin.NewPlugin("kafka", "development", New)
	b, err := json.Marshal(&Spec{
		Brokers:            strings.Split(getenv("CQ_DEST_KAFKA_CONNECTION_STRING", defaultConnectionString), ","),
		SaslUsername:       getenv("CQ_DEST_KAFKA_SASL_USERNAME", ""),
		SaslPassword:       getenv("CQ_DEST_KAFKA_SASL_PASSWORD", ""),
		Verbose:            true,
		MaxMetadataRetries: 15,
		FileSpec: &filetypes.FileSpec{
			Format: filetypes.FormatTypeJSON,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := p.Init(ctx, b, plugin.NewClientOptions{}); err != nil {
		t.Fatal(err)
	}
	plugin.TestWriterSuiteRunner(t,
		p,
		plugin.WriterTestSuiteTests{
			SkipUpsert:      true,
			SkipMigrate:     true,
			SkipDeleteStale: true,
		},
	)
}

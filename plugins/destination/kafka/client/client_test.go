package client

import (
	"os"
	"strings"
	"testing"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
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

func TestPgPlugin(t *testing.T) {
	p := destination.NewPlugin("kafka", "development", New)
	destination.PluginTestSuiteRunner(t, p,
		Spec{
			Brokers:            strings.Split(getenv("CQ_DEST_KAFKA_CONNECTION_STRING", defaultConnectionString), ","),
			SaslUsername:       getenv("CQ_DEST_KAFKA_SASL_USERNAME", ""),
			SaslPassword:       getenv("CQ_DEST_KAFKA_SASL_PASSWORD", ""),
			Format:             FormatTypeJSON,
			Verbose:            true,
			MaxMetadataRetries: 15,
		},
		destination.PluginTestSuiteTests{
			SkipOverwrite:        true,
			SkipMigrateAppend:    true,
			SkipMigrateOverwrite: true,
		})
}

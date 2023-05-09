package client

import (
	"os"
	"strings"
	"testing"

	"github.com/cloudquery/filetypes/v2"
	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v2/plugins/destination"
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
	destination.PluginTestSuiteRunner(t,
		func() *destination.Plugin {
			return destination.NewPlugin("kafka", "development", New)
		},
		specs.Destination{
			Spec: &Spec{
				Brokers:            strings.Split(getenv("CQ_DEST_KAFKA_CONNECTION_STRING", defaultConnectionString), ","),
				SaslUsername:       getenv("CQ_DEST_KAFKA_SASL_USERNAME", ""),
				SaslPassword:       getenv("CQ_DEST_KAFKA_SASL_PASSWORD", ""),
				Verbose:            true,
				MaxMetadataRetries: 15,
				FileSpec: &filetypes.FileSpec{
					Format: filetypes.FormatTypeJSON,
				},
			},
		},
		destination.PluginTestSuiteTests{
			SkipOverwrite:             true,
			SkipMigrateAppend:         true,
			SkipMigrateOverwrite:      true,
			SkipMigrateOverwriteForce: true,
			SkipMigrateAppendForce:    true,
		})
}

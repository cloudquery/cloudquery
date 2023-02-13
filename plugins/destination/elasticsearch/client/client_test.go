package client

import (
	"os"
	"testing"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
)

func TestPlugin(t *testing.T) {
	address := os.Getenv("ELASTICSEARCH_ADDRESS")
	if address == "" {
		address = "http://localhost:9200"
	}
	destination.PluginTestSuiteRunner(t,
		func() *destination.Plugin {
			return destination.NewPlugin("elasticsearch", "development", New, destination.WithManagedWriter())
		},
		Spec{
			Addresses: []string{address},
		},
		destination.PluginTestSuiteTests{
			SkipMigrateOverwriteForce: true,
			SkipMigrateAppendForce:    true,
		})
}

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
	p := destination.NewPlugin("elasticsearch", "development", New, destination.WithManagedWriter())
	destination.PluginTestSuiteRunner(t, p,
		Spec{
			Addresses: []string{address},
		},
		destination.PluginTestSuiteTests{})
}

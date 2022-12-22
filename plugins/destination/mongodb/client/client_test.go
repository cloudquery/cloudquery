package client

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
)

func TestPlugin(t *testing.T) {
	p := destination.NewPlugin("mongodb", "development", New, destination.WithManagedWriter())
	destination.PluginTestSuiteRunner(t, p,
		Spec{
			ConnectionString: "mongodb://localhost:27017",
			Database:         "destination_mongodb_test",
		},
		destination.PluginTestSuiteTests{})
}

package client

import (
	"os"
	"testing"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
)

const (
	defaultConnectionString = "bolt://localhost:7687"
	defaultUsername         = "neo4j"
	defaultPassword         = "test1234"
)

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func TestPlugin(t *testing.T) {
	p := destination.NewPlugin("neo4j", "development", New, destination.WithManagedWriter())
	destination.PluginTestSuiteRunner(t, p,
		Spec{
			Username:         getenv("CQ_DEST_NEO4J_USERNAME", defaultUsername),
			Password:         getenv("CQ_DEST_NEO4J_PASSWORD", defaultPassword),
			ConnectionString: getenv("CQ_DEST_NEO4J_CONNECTION_STRING", defaultConnectionString),
		},
		destination.PluginTestSuiteTests{})
}

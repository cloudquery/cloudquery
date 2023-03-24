package client

import (
	"os"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/destination/meilisearch/resources/plugin"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/specs"
)

var migrateStrategy = destination.MigrateStrategy{
	AddColumn:           specs.MigrateModeSafe,
	AddColumnNotNull:    specs.MigrateModeSafe,
	RemoveColumn:        specs.MigrateModeSafe,
	RemoveColumnNotNull: specs.MigrateModeSafe,
	ChangeColumn:        specs.MigrateModeSafe,
}

func getTestSpec() *Spec {
	apiKey := os.Getenv("CQ_DEST_MEILI_TEST_API_KEY")
	if len(apiKey) == 0 {
		apiKey = "test"
	}
	host := os.Getenv("CQ_DEST_MEILI_TEST_HOST")
	if len(host) == 0 {
		host = "http://localhost:7700"
	}

	return &Spec{Host: host, APIKey: apiKey, Timeout: time.Minute}
}

func TestPlugin(t *testing.T) {
	destination.PluginTestSuiteRunner(t,
		func() *destination.Plugin {
			return destination.NewPlugin("meilisearch", plugin.Version, New, destination.WithManagedWriter())
		},
		specs.Destination{Spec: getTestSpec()},
		destination.PluginTestSuiteTests{
			SkipDeleteStale:          true,
			MigrateStrategyOverwrite: migrateStrategy,
			MigrateStrategyAppend:    migrateStrategy,
		},
	)
}

package client

import (
	"os"
	"testing"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/specs"
)

var migrateStrategy = destination.MigrateStrategy{
	AddColumn:           specs.MigrateModeSafe,
	AddColumnNotNull:    specs.MigrateModeForced,
	RemoveColumn:        specs.MigrateModeSafe,
	RemoveColumnNotNull: specs.MigrateModeForced,
	ChangeColumn:        specs.MigrateModeForced,
}

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
	destination.PluginTestSuiteRunner(t,
		func() *destination.Plugin {
			return destination.NewPlugin("neo4j", "development", New, destination.WithManagedWriter())
		},
		Spec{
			Username:         getenv("CQ_DEST_NEO4J_USERNAME", defaultUsername),
			Password:         getenv("CQ_DEST_NEO4J_PASSWORD", defaultPassword),
			ConnectionString: getenv("CQ_DEST_NEO4J_CONNECTION_STRING", defaultConnectionString),
		},
		destination.PluginTestSuiteTests{
			SkipMigrateOverwriteForce: true,
			SkipMigrateAppendForce:    true,

			MigrateStrategyOverwrite: migrateStrategy,
			MigrateStrategyAppend:    migrateStrategy,
		})
}

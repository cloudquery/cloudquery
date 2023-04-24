package client

import (
	"os"
	"testing"

	"github.com/cloudquery/plugin-sdk/v2/plugins/destination"
	"github.com/cloudquery/plugin-sdk/v2/specs"
)

var migrateStrategy = destination.MigrateStrategy{
	AddColumn:           specs.MigrateModeSafe,
	AddColumnNotNull:    specs.MigrateModeForced,
	RemoveColumn:        specs.MigrateModeSafe,
	RemoveColumnNotNull: specs.MigrateModeForced,
	ChangeColumn:        specs.MigrateModeForced,
}

func TestPlugin(t *testing.T) {
	address := os.Getenv("ELASTICSEARCH_ADDRESS")
	if address == "" {
		address = "http://localhost:9200"
	}
	destination.PluginTestSuiteRunner(t,
		func() *destination.Plugin {
			return destination.NewPlugin("elasticsearch", "development", New, destination.WithManagedWriter())
		},
		specs.Destination{
			Spec: &Spec{
				Addresses: []string{address},
			},
		},
		destination.PluginTestSuiteTests{
			SkipMigrateOverwriteForce: true,
			SkipMigrateAppendForce:    true,

			MigrateStrategyOverwrite: migrateStrategy,
			MigrateStrategyAppend:    migrateStrategy,
		})
}

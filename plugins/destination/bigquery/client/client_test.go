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
	destination.PluginTestSuiteRunner(t,
		func() *destination.Plugin {
			return destination.NewPlugin("bigquery", "development", New, destination.WithManagedWriter())
		},
		specs.Destination{
			Spec: &Spec{
				ProjectID:        os.Getenv("BIGQUERY_PROJECT_ID"),
				DatasetID:        os.Getenv("BIGQUERY_DATASET_ID"),
				TimePartitioning: "none",
			},
		},
		destination.PluginTestSuiteTests{
			SkipOverwrite:             true,
			SkipMigrateOverwrite:      true,
			SkipMigrateOverwriteForce: true,
			SkipMigrateAppendForce:    true,

			// This fails due to a delay in schema propagation. Another solution is to wait a few minutes, but that makes tests super slow.
			SkipMigrateAppend: true,

			MigrateStrategyOverwrite: migrateStrategy,
			MigrateStrategyAppend:    migrateStrategy,
		})
}

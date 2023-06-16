package client

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

// var migrateStrategy = destination.MigrateStrategy{
// 	AddColumn:           specs.MigrateModeSafe,
// 	AddColumnNotNull:    specs.MigrateModeForced,
// 	RemoveColumn:        specs.MigrateModeSafe,
// 	RemoveColumnNotNull: specs.MigrateModeForced,
// 	ChangeColumn:        specs.MigrateModeForced,
// }

func getTestConnection() string {
	testConn := os.Getenv("CQ_DEST_MONGODB_TEST_CONN")
	if testConn == "" {
		return "mongodb://localhost:27017"
	}
	return testConn
}

func TestPlugin(t *testing.T) {
	ctx := context.Background()
	p := plugin.NewPlugin("mongodb", "development", New)
	s := &Spec{
		ConnectionString: getTestConnection(),
		Database:         "destination_mongodb_test",
	}
	b, err := json.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	if err := p.Init(ctx, b); err != nil {
		t.Fatal(err)
	}
	plugin.TestWriterSuiteRunner(t,
		p,
		plugin.PluginTestSuiteTests{
			SkipMigrate: true,
			// SkipMigrateOverwriteForce: true,
			// SkipMigrateAppendForce:    true,

			// MigrateStrategyOverwrite: migrateStrategy,
			// MigrateStrategyAppend:    migrateStrategy,
		},
		// plugin.WithTestSourceTimePrecision(time.Millisecond),
	)
}

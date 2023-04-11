package client

import (
	"os"
	"runtime"
	"strconv"
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

const (
	defaultGremlinEndpoint = "ws://localhost:8182"
)

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func TestPlugin(t *testing.T) {
	defaultInsecure := "false"
	if runtime.GOOS == "darwin" {
		defaultInsecure = "true" // x509 issue with Neptune and macOS: https://github.com/golang/go/issues/51991
	}
	insecure, _ := strconv.ParseBool(getenv("CQ_DEST_GREMLIN_INSECURE", defaultInsecure))

	destination.PluginTestSuiteRunner(t,
		func() *destination.Plugin {
			return destination.NewPlugin("gremlin", "development", New, destination.WithManagedWriter())
		},
		specs.Destination{
			Spec: &Spec{
				Endpoint: getenv("CQ_DEST_GREMLIN_ENDPOINT", defaultGremlinEndpoint),
				Insecure: insecure,
				Username: os.Getenv("CQ_DEST_GREMLIN_USERNAME"),
				Password: os.Getenv("CQ_DEST_GREMLIN_PASSWORD"),
			},
		},
		destination.PluginTestSuiteTests{
			SkipMigrateOverwriteForce: true,
			SkipMigrateAppendForce:    true,

			MigrateStrategyOverwrite: migrateStrategy,
			MigrateStrategyAppend:    migrateStrategy,
		})
}

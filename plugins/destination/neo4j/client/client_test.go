package client

import (
	"context"
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/stretchr/testify/require"
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
	ctx := context.Background()
	p := plugin.NewPlugin("neo4j", "development", New)
	s := &Spec{
		Username:         getenv("CQ_DEST_NEO4J_USERNAME", defaultUsername),
		Password:         getenv("CQ_DEST_NEO4J_PASSWORD", defaultPassword),
		ConnectionString: getenv("CQ_DEST_NEO4J_CONNECTION_STRING", defaultConnectionString),
	}
	s.SetDefaults()
	require.NoError(t, s.Validate())
	b, err := json.Marshal(s)
	require.NoError(t, err)

	err = p.Init(ctx, b, plugin.NewClientOptions{})
	require.NoError(t, err)

	plugin.TestWriterSuiteRunner(t,
		p,
		plugin.WriterTestSuiteTests{
			SkipMigrate: true,
			SafeMigrations: plugin.SafeMigrations{
				AddColumn:    true,
				RemoveColumn: true,
			},
		},
		plugin.WithTestIgnoreNullsInLists(),
		plugin.WithTestDataOptions(schema.TestSourceOptions{
			TimePrecision: time.Microsecond,
			SkipLists:     true,
		}),
	)
}

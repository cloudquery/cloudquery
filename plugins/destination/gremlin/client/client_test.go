package client

import (
	"context"
	"encoding/json"
	"os"
	"runtime"
	"strconv"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/stretchr/testify/require"
)

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

	ctx := context.Background()
	p := plugin.NewPlugin("gremlin", "development", New)
	s := &Spec{
		Endpoint: getenv("CQ_DEST_GREMLIN_ENDPOINT", defaultGremlinEndpoint),
		Insecure: insecure,
		Username: os.Getenv("CQ_DEST_GREMLIN_USERNAME"),
		Password: os.Getenv("CQ_DEST_GREMLIN_PASSWORD"),
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
			SkipInsert:  true, // we do "no PKs = all columns are PKs" in this destination
			SkipMigrate: true,
			SafeMigrations: plugin.SafeMigrations{
				AddColumn:    true,
				RemoveColumn: true,
			},
		},
	)
}

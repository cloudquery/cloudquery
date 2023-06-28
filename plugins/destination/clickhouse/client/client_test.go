package client

import (
	"context"
	"net/url"
	"os"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/require"
)

func getTestConnection() string {
	if testConn := os.Getenv("CQ_DEST_CH_TEST_CONN"); len(testConn) > 0 {
		return testConn
	}

	return (&url.URL{
		User: url.UserPassword("cq", "test"),
		Host: "localhost:9000",
		Path: "cloudquery", // database
	}).String()
}

func TestPlugin(t *testing.T) {
	ctx := context.Background()
	p := plugin.NewPlugin("neo4j", "development", New)
	s := &Spec{ConnectionString: getTestConnection()}
	b, err := json.Marshal(s)
	require.NoError(t, err)
	require.NoError(t, p.Init(ctx, b))

	plugin.TestWriterSuiteRunner(t,
		p,
		plugin.WriterTestSuiteTests{
			SkipMigrate: true,
			SafeMigrations: plugin.SafeMigrations{
				AddColumn:    true,
				RemoveColumn: true,
			},
		},
		// destination.WithTestSourceAllowNull(typeconv.CanBeNullable),
	)
}

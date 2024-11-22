package client

import (
	"context"
	"net/url"
	"os"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/client/spec"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/clickhouse/resources/plugin"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/typeconv/ch/types"
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
	p := plugin.NewPlugin("clickhouse",
		internalPlugin.Version,
		New,
		plugin.WithJSONSchema(spec.JSONSchema),
	)
	s := &spec.Spec{ConnectionString: getTestConnection()}
	b, err := json.Marshal(s)
	require.NoError(t, err)
	require.NoError(t, p.Init(ctx, b, plugin.NewClientOptions{}))

	plugin.TestWriterSuiteRunner(t,
		p,
		plugin.WriterTestSuiteTests{
			SkipUpsert:       true,
			SkipDeleteStale:  true,
			SkipDeleteRecord: true,
			SafeMigrations: plugin.SafeMigrations{
				AddColumn:    true,
				RemoveColumn: true,
				// MovePKToCQOnly- is only a change to the underlying PKs, and because clickhouse only supports append only mode this is not a factor
				MovePKToCQOnly: true,
			},
			SkipSpecificMigrations: plugin.Migrations{
				RemoveUniqueConstraint: true,
				MovePKToCQOnly:         true,
			},
		},
		plugin.WithTestSourceAllowNull(types.CanBeNullable),
	)
}

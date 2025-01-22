//go:debug x509negativeserial=1
package client

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/require"
)

func getTestConnection() string {
	if testConn := os.Getenv("CQ_DEST_MS_SQL_TEST_CONN"); len(testConn) > 0 {
		return testConn
	}

	return fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		"localhost", "SA", "yourStrongP@ssword", 1433, "cloudquery",
	)
}

func TestPlugin(t *testing.T) {
	ctx := context.Background()
	p := plugin.NewPlugin("mssql", "development", New)
	s := &Spec{ConnectionString: getTestConnection()}
	b, err := json.Marshal(s)
	require.NoError(t, err)
	require.NoError(t, p.Init(ctx, b, plugin.NewClientOptions{}))

	plugin.TestWriterSuiteRunner(t,
		p,
		plugin.WriterTestSuiteTests{
			SkipDeleteRecord: true,
			SafeMigrations: plugin.SafeMigrations{
				AddColumn:    true,
				RemoveColumn: true,
			},
			SkipSpecificMigrations: plugin.Migrations{
				RemoveUniqueConstraint: true,
			},
			SkipSpecificWriteTests: plugin.WriteTests{
				DuplicatePK: true,
			},
		},
	)
}

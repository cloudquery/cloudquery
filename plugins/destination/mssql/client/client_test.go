package client

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/destination/mssql/resources/plugin"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
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
	p := destination.NewPlugin("mssql", plugin.Version, New, destination.WithManagedWriter())
	destination.PluginTestSuiteRunner(t, p,
		Spec{ConnectionString: getTestConnection()},
		destination.PluginTestSuiteTests{},
	)
}

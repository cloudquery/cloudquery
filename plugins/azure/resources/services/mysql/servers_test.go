//go:build integration

package mysql

import (
	"testing"

	"github.com/cloudquery/cq-provider-azure/client"
)

func TestIntegrationMySQLServers(t *testing.T) {
	client.AzureTestHelper(t, MySQLServers(),
		client.SnapshotsDirPath)
}

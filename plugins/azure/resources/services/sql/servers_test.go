//go:build integration

package sql

import (
	"testing"

	"github.com/cloudquery/cq-provider-azure/client"
)

func TestIntegrationSecuritySettings(t *testing.T) {
	client.AzureTestHelper(t, SQLServers(),
		client.SnapshotsDirPath)
}

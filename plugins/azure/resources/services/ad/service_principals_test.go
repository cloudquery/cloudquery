//go:build integration

package ad

import (
	"testing"

	"github.com/cloudquery/cq-provider-azure/client"
)

func TestIntegrationAdServicePrincipals(t *testing.T) {
	client.AzureTestHelper(t, ServicePrincipals(),
		client.SnapshotsDirPath)
}

//go:build integration

package network

import (
	"testing"

	"github.com/cloudquery/cq-provider-azure/client"
)

func TestIntegrationNetworkPublicIPAddresses(t *testing.T) {
	client.AzureTestHelper(t, NetworkPublicIPAddresses(),
		client.SnapshotsDirPath)
}

//go:build integration

package network

import (
	"testing"

	"github.com/cloudquery/cq-provider-azure/client"
)

func TestIntegrationNetworkWatchers(t *testing.T) {
	client.AzureTestHelper(t, NetworkWatchers(),
		client.SnapshotsDirPath)
}

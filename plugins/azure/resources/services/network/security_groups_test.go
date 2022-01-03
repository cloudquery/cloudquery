//go:build integration

package network

import (
	"testing"

	"github.com/cloudquery/cq-provider-azure/client"
)

func TestIntegrationNetworkSecurityGroups(t *testing.T) {
	client.AzureTestHelper(t, NetworkSecurityGroups())
}

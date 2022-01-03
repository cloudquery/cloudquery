//go:build integration

package keyvault

import (
	"testing"

	"github.com/cloudquery/cq-provider-azure/client"
)

func TestIntegrationKeyvaultVaults(t *testing.T) {
	client.AzureTestHelper(t, KeyvaultVaults())
}

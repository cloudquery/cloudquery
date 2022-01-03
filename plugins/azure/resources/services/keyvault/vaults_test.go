//go:build integration

package keyvault

import (
	"testing"

	"github.com/cloudquery/cq-provider-azure/client"
)

func TestIntegrationKeyvaultVaults(t *testing.T) {
	t.Skip("skip untill issue is fixed https://github.com/cloudquery/cq-provider-azure/issues/107")
	client.AzureTestHelper(t, KeyvaultVaults())
}

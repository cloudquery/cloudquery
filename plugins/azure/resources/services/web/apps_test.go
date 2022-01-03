//go:build integration

package web

import (
	"testing"

	"github.com/cloudquery/cq-provider-azure/client"
)

func TestIntegrationSubscriptionSubscriptions(t *testing.T) {
	client.AzureTestHelper(t, WebApps())
}

//go:build integration

package subscription

import (
	"testing"

	"github.com/cloudquery/cq-provider-azure/client"
)

func TestIntegrationSubscriptionSubscriptions(t *testing.T) {
	client.AzureTestHelper(t, SubscriptionSubscriptions(),
		client.SnapshotsDirPath)
}

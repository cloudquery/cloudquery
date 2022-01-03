//go:build integration

package resources

import (
	"testing"

	"github.com/cloudquery/cq-provider-azure/client"
)

func TestIntegrationResourcesGroups(t *testing.T) {
	client.AzureTestHelper(t, ResourcesGroups())
}

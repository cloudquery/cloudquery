//go:build integration

package resources

import (
	"testing"

	"github.com/cloudquery/cq-provider-azure/client"
)

func TestIntegrationResourcesPolicyAssignments(t *testing.T) {
	client.AzureTestHelper(t, ResourcesPolicyAssignments(),
		client.SnapshotsDirPath)
}

//go:build integration

package container

import (
	"testing"

	"github.com/cloudquery/cq-provider-azure/client"
)

func TestIntegrationManagedClusters(t *testing.T) {
	client.AzureTestHelper(t, ContainerManagedClusters(),
		client.SnapshotsDirPath)
}

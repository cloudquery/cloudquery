//go:build integration

package compute

import (
	"testing"

	"github.com/cloudquery/cq-provider-azure/client"
)

func TestIntegrationDisks(t *testing.T) {
	client.AzureTestHelper(t, ComputeDisks(),
		client.SnapshotsDirPath)
}

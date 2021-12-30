//go:build integration

package compute

import (
	"testing"

	"github.com/cloudquery/cq-provider-azure/client"
)

func TestIntegrationVirtualMachines(t *testing.T) {
	client.AzureTestHelper(t, ComputeVirtualMachines(),
		client.SnapshotsDirPath)
}

//go:build integration
// +build integration

package compute

import (
	"testing"

	"github.com/cloudquery/cq-provider-gcp/client"
)

func TestIntegrationComputeDisks(t *testing.T) {
	client.GcpTestHelper(t, ComputeDisks())
}

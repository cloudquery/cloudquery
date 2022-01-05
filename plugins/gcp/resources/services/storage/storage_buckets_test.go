//go:build integration
// +build integration

package storage

import (
	"testing"

	"github.com/cloudquery/cq-provider-gcp/client"
)

func TestIntegrationStorageBuckets(t *testing.T) {
	client.GcpTestHelper(t, StorageBuckets())
}

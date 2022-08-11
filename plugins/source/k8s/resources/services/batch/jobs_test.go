//go:build integration
// +build integration

package batch

import (
	"testing"

	"github.com/cloudquery/cq-provider-k8s/client"
)

func TestIntegrationJobs(t *testing.T) {
	client.K8sTestHelper(t, Jobs(), "./snapshots")
}

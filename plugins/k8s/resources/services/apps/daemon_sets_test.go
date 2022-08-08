//go:build integration
// +build integration

package apps

import (
	"testing"

	"github.com/cloudquery/cq-provider-k8s/client"
)

func TestIntegrationDaemonSets(t *testing.T) {
	client.K8sTestHelper(t, DaemonSets(), "./snapshots")
}

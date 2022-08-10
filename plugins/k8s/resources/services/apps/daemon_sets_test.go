//go:build integration
// +build integration

package apps

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/k8s/client"
)

func TestIntegrationDaemonSets(t *testing.T) {
	client.K8sTestHelper(t, DaemonSets(), "./snapshots")
}

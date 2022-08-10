//go:build integration
// +build integration

package core

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/k8s/client"
)

func TestIntegrationEndpoints(t *testing.T) {
	client.K8sTestHelper(t, Endpoints(), "./snapshots")
}

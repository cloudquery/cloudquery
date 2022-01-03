//go:build integration
// +build integration

package core

import (
	"testing"

	"github.com/cloudquery/cq-provider-k8s/client"
)

func TestIntegrationServices(t *testing.T) {
	client.K8sTestHelper(t, Services(), "./snapshots")
}

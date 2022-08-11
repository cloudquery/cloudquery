//go:build integration
// +build integration

package rbac

import (
	"testing"

	"github.com/cloudquery/cq-provider-k8s/client"
)

func TestIntegrationRoles(t *testing.T) {
	client.K8sTestHelper(t, Roles(), "./snapshots")
}

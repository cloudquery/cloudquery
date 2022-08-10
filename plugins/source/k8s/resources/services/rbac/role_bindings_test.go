//go:build integration
// +build integration

package rbac

import (
	"testing"

	"github.com/cloudquery/cq-provider-k8s/client"
)

func TestIntegrationRoleBindings(t *testing.T) {
	client.K8sTestHelper(t, RoleBindings(), "./snapshots")
}

//go:build integration
// +build integration

package rbac

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
)

func TestIntegrationRoleBindings(t *testing.T) {
	client.K8sTestHelper(t, RoleBindings(), "./snapshots")
}

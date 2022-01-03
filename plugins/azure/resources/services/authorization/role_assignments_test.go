//go:build integration

package authorization

import (
	"testing"

	"github.com/cloudquery/cq-provider-azure/client"
)

func TestIntegrationRoleAssignments(t *testing.T) {
	client.AzureTestHelper(t, AuthorizationRoleAssignments())
}

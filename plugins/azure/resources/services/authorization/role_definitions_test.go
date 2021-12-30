//go:build integration

package authorization

import (
	"testing"

	"github.com/cloudquery/cq-provider-azure/client"
)

func TestIntegrationRoleDefinitions(t *testing.T) {
	client.AzureTestHelper(t, AuthorizationRoleDefinitions(),
		client.SnapshotsDirPath)
}

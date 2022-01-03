//go:build integration

package security

import (
	"testing"

	"github.com/cloudquery/cq-provider-azure/client"
)

func TestIntegrationSecurityContacts(t *testing.T) {
	client.AzureTestHelper(t, SecurityContacts())
}

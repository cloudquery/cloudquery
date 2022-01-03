//go:build integration

package ad

import (
	"testing"

	"github.com/cloudquery/cq-provider-azure/client"
)

func TestIntegrationAdGroups(t *testing.T) {
	client.AzureTestHelper(t, Groups())
}

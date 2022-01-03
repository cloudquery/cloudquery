//go:build integration

package storage

import (
	"testing"

	"github.com/cloudquery/cq-provider-azure/client"
)

func TestIntegrationStorageAccounts(t *testing.T) {
	client.AzureTestHelper(t, StorageAccounts())
}

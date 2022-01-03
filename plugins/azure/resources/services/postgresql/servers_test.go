//go:build integration

package postgresql

import (
	"testing"

	"github.com/cloudquery/cq-provider-azure/client"
)

func TestIntegrationPostgresqlServers(t *testing.T) {
	client.AzureTestHelper(t, PostgresqlServers())
}

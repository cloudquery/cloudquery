//go:build integration
// +build integration

package sql

import (
	"testing"

	"github.com/cloudquery/cq-provider-gcp/client"
)

func TestIntegrationSQLInstances(t *testing.T) {
	client.GcpTestHelper(t, SQLInstances())
}

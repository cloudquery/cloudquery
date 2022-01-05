//go:build integration
// +build integration

package dns

import (
	"testing"

	"github.com/cloudquery/cq-provider-gcp/client"
)

func TestIntegrationDnsManagedZones(t *testing.T) {
	client.GcpTestHelper(t, DNSManagedZones())
}

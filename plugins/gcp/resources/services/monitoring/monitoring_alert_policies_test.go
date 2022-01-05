//go:build integration
// +build integration

package monitoring

import (
	"testing"

	"github.com/cloudquery/cq-provider-gcp/client"
)

func TestIntegrationMonitoringAlertPolicies(t *testing.T) {
	client.GcpTestHelper(t, MonitoringAlertPolicies())
}

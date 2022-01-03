//go:build integration

package monitor

import (
	"testing"

	"github.com/cloudquery/cq-provider-azure/client"
)

func TestIntegrationMonitorDiagnosticSettings(t *testing.T) {
	t.Skip("Skipping untill https://github.com/cloudquery/cq-provider-azure/issues/105 fixed")
	client.AzureTestHelper(t, MonitorDiagnosticSettings())
}

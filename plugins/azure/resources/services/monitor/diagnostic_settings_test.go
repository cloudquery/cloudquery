//go:build integration

package monitor

import (
	"testing"

	"github.com/cloudquery/cq-provider-azure/client"
)

func TestIntegrationMonitorDiagnosticSettings(t *testing.T) {
	client.AzureTestHelper(t, MonitorDiagnosticSettings())
}

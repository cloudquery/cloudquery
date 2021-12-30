//go:build integration

package monitor

import (
	"testing"

	"github.com/cloudquery/cq-provider-azure/client"
)

func TestIntegrationMonitorActivityLogAlerts(t *testing.T) {
	client.AzureTestHelper(t, MonitorActivityLogAlerts(),
		client.SnapshotsDirPath)
}

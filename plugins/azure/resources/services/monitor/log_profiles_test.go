//go:build integration

package monitor

import (
	"testing"

	"github.com/cloudquery/cq-provider-azure/client"
)

func TestIntegrationMonitorLogProfiles(t *testing.T) {
	client.AzureTestHelper(t, MonitorLogProfiles(),
		client.SnapshotsDirPath)
}

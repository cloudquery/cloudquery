package maintenance_windows

import (
	"testing"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
)

func buildMockHttpClient() *client.MockHttpClient {
	mockHttpClient := client.MockHttpClient{}

	response := pagerduty.ListMaintenanceWindowsResponse{}

	if err := faker.FakeObject(&response); err != nil {
		panic(err)
	}
	if err := client.FakeStringTimestamps(&response.MaintenanceWindows[0]); err != nil {
		panic(err)
	}

	response.More = false

	mockHttpClient.AddMockResponse("/maintenance_windows", response)

	return &mockHttpClient
}

func TestMaintenanceWindows(t *testing.T) {
	client.PagerdutyMockTestHelper(t, MaintenanceWindows(), buildMockHttpClient)
}

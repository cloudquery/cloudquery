// Auto generated code - DO NOT EDIT.

package monitor

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2019-11-01-preview/insights"
)

func TestMonitorActivityLogAlerts(t *testing.T) {
	client.MockTestHelper(t, ActivityLogAlerts(), createActivityLogAlertsMock)
}

func createActivityLogAlertsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockMonitorActivityLogAlertsClient(ctrl)
	s := services.Services{
		Monitor: services.MonitorClient{
			ActivityLogAlerts: mockClient,
		},
	}

	data := insights.ActivityLogAlertResource{}
	require.Nil(t, faker.FakeObject(&data))
	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/" + *data.ID
	data.ID = &id

	result := insights.ActivityLogAlertList{Value: &[]insights.ActivityLogAlertResource{data}}

	mockClient.EXPECT().ListBySubscriptionID(gomock.Any()).Return(result, nil)
	return s
}

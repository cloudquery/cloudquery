package resources_test

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2019-11-01-preview/insights"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/cq-provider-azure/resources"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildActivityLogAlertsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	alertsSvc := mocks.NewMockActivityLogAlertsClient(ctrl)
	s := services.Services{
		Monitor: services.MonitorClient{
			ActivityLogAlerts: alertsSvc,
		},
	}
	alert := insights.ActivityLogAlertResource{}
	err := faker.FakeData(&alert)
	if err != nil {
		t.Errorf("failed building mock %s", err)
	}

	page := insights.ActivityLogAlertList{Value: &[]insights.ActivityLogAlertResource{alert}}
	alertsSvc.EXPECT().ListBySubscriptionID(gomock.Any()).Return(page, nil)

	return s
}

func TestActivityLogAlerts(t *testing.T) {
	azureTestHelper(t, resources.MonitorActivityLogAlerts(), buildActivityLogAlertsMock)
}

package monitor

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2021-07-01-preview/insights"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildMonitorLogProfiles(t *testing.T, ctrl *gomock.Controller) services.Services {
	m := mocks.NewMockLogProfilesClient(ctrl)

	var r insights.LogProfileResource
	if err := faker.FakeData(&r); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().List(gomock.Any()).Return(
		insights.LogProfileCollection{Value: &[]insights.LogProfileResource{r}},
		nil,
	)

	return services.Services{
		Monitor: services.MonitorClient{LogProfiles: m},
	}
}

func TestMonitorLogProfiles(t *testing.T) {
	client.AzureMockTestHelper(t, MonitorLogProfiles(), buildMonitorLogProfiles, client.TestOptions{})
}

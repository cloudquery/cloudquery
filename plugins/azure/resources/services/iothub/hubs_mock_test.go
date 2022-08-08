package iothub

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/iothub/mgmt/2021-07-02/devices"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildIotHubHubsClientMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	m := mocks.NewMockIotHubClient(ctrl)
	var iothub devices.IotHubDescription
	if err := faker.FakeData(&iothub); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListBySubscription(gomock.Any()).Return(
		devices.NewIotHubDescriptionListResultPage(
			devices.IotHubDescriptionListResult{Value: &[]devices.IotHubDescription{iothub}},
			func(c context.Context, lr devices.IotHubDescriptionListResult) (devices.IotHubDescriptionListResult, error) {
				return devices.IotHubDescriptionListResult{}, nil
			},
		),
		nil,
	)

	return services.Services{IotHub: m}
}

func TestIotHubHubsServices(t *testing.T) {
	client.AzureMockTestHelper(t, IothubHubs(), buildIotHubHubsClientMock, client.TestOptions{})
}

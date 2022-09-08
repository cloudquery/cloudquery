// Auto generated code - DO NOT EDIT.

package iothub

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/go-faker/faker/v4"
	fakerOptions "github.com/go-faker/faker/v4/pkg/options"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/iothub/mgmt/2021-07-02/devices"
)

func TestIotHubDevices(t *testing.T) {
	client.MockTestHelper(t, Devices(), createDevicesMock)
}

func createDevicesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockIotHubDevicesClient(ctrl)
	s := services.Services{
		IotHub: services.IotHubClient{
			Devices: mockClient,
		},
	}

	data := devices.IotHubDescription{}
	fieldsToIgnore := []string{"Response"}
	require.Nil(t, faker.FakeData(&data, fakerOptions.WithIgnoreInterface(true), fakerOptions.WithFieldsToIgnore(fieldsToIgnore...), fakerOptions.WithRandomMapAndSliceMinSize(1), fakerOptions.WithRandomMapAndSliceMaxSize(1)))

	result := devices.NewIotHubDescriptionListResultPage(devices.IotHubDescriptionListResult{Value: &[]devices.IotHubDescription{data}}, func(ctx context.Context, result devices.IotHubDescriptionListResult) (devices.IotHubDescriptionListResult, error) {
		return devices.IotHubDescriptionListResult{}, nil
	})

	mockClient.EXPECT().ListBySubscription(gomock.Any()).Return(result, nil)
	return s
}

// Auto generated code - DO NOT EDIT.

package servicebus

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/preview/servicebus/mgmt/2021-06-01-preview/servicebus"
)

func createAccessKeysMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockServicebusAccessKeysClient(ctrl)
	s := services.Services{
		Servicebus: services.ServicebusClient{
			AccessKeys: mockClient,
		},
	}

	data := servicebus.AccessKeys{}
	require.Nil(t, faker.FakeObject(&data))

	result := servicebus.NewAccessKeysListResultPage(servicebus.AccessKeysListResult{Value: &[]servicebus.AccessKeys{data}}, func(ctx context.Context, result servicebus.AccessKeysListResult) (servicebus.AccessKeysListResult, error) {
		return servicebus.AccessKeysListResult{}, nil
	})

	mockClient.EXPECT().ListKeys(gomock.Any(), "test", "test", "test", "test").Return(result, nil)
	return s
}

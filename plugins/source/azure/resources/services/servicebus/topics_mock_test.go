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

func createTopicsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockServicebusTopicsClient(ctrl)
	s := services.Services{
		Servicebus: services.ServicebusClient{
			Topics: mockClient,
		},
	}

	data := servicebus.SBTopic{}
	require.Nil(t, faker.FakeObject(&data))

	// Ensure name and ID are consistent so we can reference it in other mock
	name := "test"
	data.Name = &name

	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := servicebus.NewSBTopicListResultPage(servicebus.SBTopicListResult{Value: &[]servicebus.SBTopic{data}}, func(ctx context.Context, result servicebus.SBTopicListResult) (servicebus.SBTopicListResult, error) {
		return servicebus.SBTopicListResult{}, nil
	})

	mockClient.EXPECT().ListByNamespace(gomock.Any(), "test", "test", nil, nil).Return(result, nil)
	return s
}

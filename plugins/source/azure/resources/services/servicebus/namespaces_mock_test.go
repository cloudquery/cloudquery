// Auto generated code - DO NOT EDIT.

package servicebus

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/preview/servicebus/mgmt/2021-06-01-preview/servicebus"
)

func TestServicebusNamespaces(t *testing.T) {
	client.MockTestHelper(t, Namespaces(), createNamespacesMock)
}

func createNamespacesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockServicebusNamespacesClient(ctrl)
	s := services.Services{
		Servicebus: services.ServicebusClient{
			Namespaces:         mockClient,
			Topics:             createTopicsMock(t, ctrl).Servicebus.Topics,
			AuthorizationRules: createAuthorizationRulesMock(t, ctrl).Servicebus.AuthorizationRules,
			AccessKeys:         createAccessKeysMock(t, ctrl).Servicebus.AccessKeys,
		},
	}

	data := servicebus.SBNamespace{}
	require.Nil(t, faker.FakeObject(&data))

	// Ensure name and ID are consistent so we can reference it in other mock
	name := "test"
	data.Name = &name

	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := servicebus.NewSBNamespaceListResultPage(servicebus.SBNamespaceListResult{Value: &[]servicebus.SBNamespace{data}}, func(ctx context.Context, result servicebus.SBNamespaceListResult) (servicebus.SBNamespaceListResult, error) {
		return servicebus.SBNamespaceListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).Return(result, nil)
	return s
}

// Auto generated code - DO NOT EDIT.

package eventhub

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/preview/eventhub/mgmt/2018-01-01-preview/eventhub"
)

func TestEventHubNamespaces(t *testing.T) {
	client.MockTestHelper(t, Namespaces(), createNamespacesMock)
}

func createNamespacesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockEventHubNamespacesClient(ctrl)
	s := services.Services{
		EventHub: services.EventHubClient{
			Namespaces: mockClient,
		},
	}

	data := eventhub.EHNamespace{}
	require.Nil(t, faker.FakeObject(&data))

	result := eventhub.NewEHNamespaceListResultPage(eventhub.EHNamespaceListResult{Value: &[]eventhub.EHNamespace{data}}, func(ctx context.Context, result eventhub.EHNamespaceListResult) (eventhub.EHNamespaceListResult, error) {
		return eventhub.EHNamespaceListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).Return(result, nil)
	return s
}

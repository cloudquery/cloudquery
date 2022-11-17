// Code generated by codegen; DO NOT EDIT.

package servicebus

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	api "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	mocks "github.com/cloudquery/cloudquery/plugins/source/azure/client/mocks/servicebus"
	service "github.com/cloudquery/cloudquery/plugins/source/azure/client/services/servicebus"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildNamespaces(t *testing.T, ctrl *gomock.Controller) *client.Services {
	mockNamespacesClient := mocks.NewMockNamespacesClient(ctrl)

	var response api.NamespacesClientListResponse
	require.NoError(t, faker.FakeObject(&response))
	// Use correct Azure ID format
	const id = "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	response.Value[0].ID = to.Ptr(id)

	mockNamespacesClient.EXPECT().NewListPager(gomock.Any()).
		Return(client.CreatePager(response)).MinTimes(1)

	servicebusClient := &service.ServicebusClient{
		NamespacesClient: mockNamespacesClient,
	}

	c := &client.Services{Servicebus: servicebusClient}

	buildTopics(t, ctrl, c)

	return c
}

func TestNamespaces(t *testing.T) {
	client.MockTestHelper(t, Namespaces(), buildNamespaces)
}

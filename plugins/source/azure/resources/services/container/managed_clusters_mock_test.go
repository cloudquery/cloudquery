// Code generated by codegen; DO NOT EDIT.

package container

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	api "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	mocks "github.com/cloudquery/cloudquery/plugins/source/azure/client/mocks/containerservice"
	service "github.com/cloudquery/cloudquery/plugins/source/azure/client/services/containerservice"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildManagedClusters(t *testing.T, ctrl *gomock.Controller) *client.Services {
	mockManagedClustersClient := mocks.NewMockManagedClustersClient(ctrl)

	var response api.ManagedClustersClientListResponse
	require.NoError(t, faker.FakeObject(&response))
	// Use correct Azure ID format
	const id = "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	response.Value[0].ID = to.Ptr(id)

	mockManagedClustersClient.EXPECT().NewListPager(gomock.Any()).
		Return(client.CreatePager(response)).MinTimes(1)

	containerserviceClient := &service.ContainerserviceClient{
		ManagedClustersClient: mockManagedClustersClient,
	}

	c := &client.Services{Containerservice: containerserviceClient}

	return c
}

func TestManagedClusters(t *testing.T) {
	client.MockTestHelper(t, ManagedClusters(), buildManagedClusters)
}

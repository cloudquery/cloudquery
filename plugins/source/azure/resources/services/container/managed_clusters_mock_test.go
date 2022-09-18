// Auto generated code - DO NOT EDIT.

package container

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2021-03-01/containerservice"
)

func TestContainerManagedClusters(t *testing.T) {
	client.MockTestHelper(t, ManagedClusters(), createManagedClustersMock)
}

func createManagedClustersMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockContainerManagedClustersClient(ctrl)
	s := services.Services{
		Container: services.ContainerClient{
			ManagedClusters: mockClient,
		},
	}

	data := containerservice.ManagedCluster{}
	require.Nil(t, faker.FakeObject(&data))
	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/" + *data.ID
	data.ID = &id

	result := containerservice.NewManagedClusterListResultPage(containerservice.ManagedClusterListResult{Value: &[]containerservice.ManagedCluster{data}}, func(ctx context.Context, result containerservice.ManagedClusterListResult) (containerservice.ManagedClusterListResult, error) {
		return containerservice.ManagedClusterListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).Return(result, nil)
	return s
}

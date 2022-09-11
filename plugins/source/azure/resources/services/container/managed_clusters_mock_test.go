// Auto generated code - DO NOT EDIT.

package container

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
	fieldsToIgnore := []string{"Response"}
	require.Nil(t, faker.FakeData(&data, fakerOptions.WithIgnoreInterface(true), fakerOptions.WithRecursionMaxDepth(2), fakerOptions.WithFieldsToIgnore(fieldsToIgnore...), fakerOptions.WithRandomMapAndSliceMinSize(1), fakerOptions.WithRandomMapAndSliceMaxSize(1)))

	result := containerservice.NewManagedClusterListResultPage(containerservice.ManagedClusterListResult{Value: &[]containerservice.ManagedCluster{data}}, func(ctx context.Context, result containerservice.ManagedClusterListResult) (containerservice.ManagedClusterListResult, error) {
		return containerservice.ManagedClusterListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).Return(result, nil)
	return s
}

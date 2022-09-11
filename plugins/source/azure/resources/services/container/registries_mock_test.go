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

	"github.com/Azure/azure-sdk-for-go/services/containerregistry/mgmt/2019-05-01/containerregistry"
)

func TestContainerRegistries(t *testing.T) {
	client.MockTestHelper(t, Registries(), createRegistriesMock)
}

func createRegistriesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockContainerRegistriesClient(ctrl)
	s := services.Services{
		Container: services.ContainerClient{
			Registries: mockClient,
		},
	}

	data := containerregistry.Registry{}
	require.Nil(t, faker.FakeObject(&data))
	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/" + *data.ID
	data.ID = &id

	result := containerregistry.NewRegistryListResultPage(containerregistry.RegistryListResult{Value: &[]containerregistry.Registry{data}}, func(ctx context.Context, result containerregistry.RegistryListResult) (containerregistry.RegistryListResult, error) {
		return containerregistry.RegistryListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).Return(result, nil)
	return s
}

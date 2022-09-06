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

	"github.com/Azure/azure-sdk-for-go/services/containerregistry/mgmt/2019-05-01/containerregistry"
)

func TestContainerRegistries(t *testing.T) {
	client.AzureMockTestHelper(t, Registries(), createRegistriesMock, client.TestOptions{})
}

func createRegistriesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockContainerRegistriesClient(ctrl)
	s := services.Services{
		Container: services.ContainerClient{
			Registries: mockClient,
		},
	}

	data := containerregistry.Registry{}
	fieldsToIgnore := []string{"Response"}
	require.Nil(t, faker.FakeData(&data, fakerOptions.WithIgnoreInterface(true), fakerOptions.WithFieldsToIgnore(fieldsToIgnore...), fakerOptions.WithRandomMapAndSliceMinSize(1), fakerOptions.WithRandomMapAndSliceMaxSize(1)))

	result := containerregistry.NewRegistryListResultPage(containerregistry.RegistryListResult{Value: &[]containerregistry.Registry{data}}, func(ctx context.Context, result containerregistry.RegistryListResult) (containerregistry.RegistryListResult, error) {
		return containerregistry.RegistryListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).Return(result, nil)
	return s
}

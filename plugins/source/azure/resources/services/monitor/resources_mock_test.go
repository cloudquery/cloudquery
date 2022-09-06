// Auto generated code - DO NOT EDIT.

package monitor

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/resources"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/go-faker/faker/v4"
	fakerOptions "github.com/go-faker/faker/v4/pkg/options"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestMonitorResources(t *testing.T) {
	client.AzureMockTestHelper(t, Resources(), createResourcesMock, client.TestOptions{})
}

func createResourcesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockMonitorResourcesClient(ctrl)
	s := services.Services{
		Monitor: services.MonitorClient{
			Resources: mockClient,
		},
	}

	data := resources.GenericResourceExpanded{}
	fieldsToIgnore := []string{"Response", "Properties"}
	require.Nil(t, faker.FakeData(&data, fakerOptions.WithIgnoreInterface(true), fakerOptions.WithFieldsToIgnore(fieldsToIgnore...), fakerOptions.WithRandomMapAndSliceMinSize(1), fakerOptions.WithRandomMapAndSliceMaxSize(1)))

	result := resources.NewListResultPage(resources.ListResult{Value: &[]resources.GenericResourceExpanded{data}}, func(ctx context.Context, result resources.ListResult) (resources.ListResult, error) {
		return resources.ListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any(), "", "", nil).Return(result, nil)
	return s
}

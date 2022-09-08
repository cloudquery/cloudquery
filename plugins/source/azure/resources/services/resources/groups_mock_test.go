// Auto generated code - DO NOT EDIT.

package resources

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

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2020-10-01/resources"
)

func TestResourcesGroups(t *testing.T) {
	client.MockTestHelper(t, Groups(), createGroupsMock)
}

func createGroupsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockResourcesGroupsClient(ctrl)
	s := services.Services{
		Resources: services.ResourcesClient{
			Groups: mockClient,
		},
	}

	data := resources.Group{}
	fieldsToIgnore := []string{"Response"}
	require.Nil(t, faker.FakeData(&data, fakerOptions.WithIgnoreInterface(true), fakerOptions.WithFieldsToIgnore(fieldsToIgnore...), fakerOptions.WithRandomMapAndSliceMinSize(1), fakerOptions.WithRandomMapAndSliceMaxSize(1)))

	result := resources.NewGroupListResultPage(resources.GroupListResult{Value: &[]resources.Group{data}}, func(ctx context.Context, result resources.GroupListResult) (resources.GroupListResult, error) {
		return resources.GroupListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any(), "", nil).Return(result, nil)
	return s
}

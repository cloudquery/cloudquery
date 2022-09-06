// Auto generated code - DO NOT EDIT.

package frontdoor

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

	"github.com/Azure/azure-sdk-for-go/services/frontdoor/mgmt/2020-11-01/frontdoor"
)

func TestFrontDoorDoors(t *testing.T) {
	client.AzureMockTestHelper(t, Doors(), createDoorsMock, client.TestOptions{})
}

func createDoorsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockFrontDoorDoorsClient(ctrl)
	s := services.Services{
		FrontDoor: services.FrontDoorClient{
			Doors: mockClient,
		},
	}

	data := frontdoor.FrontDoor{}
	fieldsToIgnore := []string{"Response", "Properties"}
	require.Nil(t, faker.FakeData(&data, fakerOptions.WithIgnoreInterface(true), fakerOptions.WithFieldsToIgnore(fieldsToIgnore...), fakerOptions.WithRandomMapAndSliceMinSize(1), fakerOptions.WithRandomMapAndSliceMaxSize(1)))

	result := frontdoor.NewListResultPage(frontdoor.ListResult{Value: &[]frontdoor.FrontDoor{data}}, func(ctx context.Context, result frontdoor.ListResult) (frontdoor.ListResult, error) {
		return frontdoor.ListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).Return(result, nil)
	return s
}

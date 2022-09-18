// Auto generated code - DO NOT EDIT.

package frontdoor

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/frontdoor/mgmt/2020-11-01/frontdoor"
)

func TestFrontDoorDoors(t *testing.T) {
	client.MockTestHelper(t, Doors(), createDoorsMock)
}

func createDoorsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockFrontDoorDoorsClient(ctrl)
	s := services.Services{
		FrontDoor: services.FrontDoorClient{
			Doors: mockClient,
		},
	}

	data := frontdoor.FrontDoor{}
	require.Nil(t, faker.FakeObject(&data))
	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/" + *data.ID
	data.ID = &id

	result := frontdoor.NewListResultPage(frontdoor.ListResult{Value: &[]frontdoor.FrontDoor{data}}, func(ctx context.Context, result frontdoor.ListResult) (frontdoor.ListResult, error) {
		return frontdoor.ListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).Return(result, nil)
	return s
}

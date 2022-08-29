package frontdoor

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/frontdoor/mgmt/2020-11-01/frontdoor"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/golang/mock/gomock"
)

func buildFrontDoorsServices(t *testing.T, ctrl *gomock.Controller) services.Services {
	m := mocks.NewMockFrontDoorDoorsClient(ctrl)

	m.EXPECT().List(gomock.Any()).Return(
		frontdoor.NewListResultPage(
			frontdoor.ListResult{Value: &[]frontdoor.FrontDoor{fakeFrontDoor(t)}},
			func(c context.Context, lr frontdoor.ListResult) (frontdoor.ListResult, error) {
				return frontdoor.ListResult{}, nil
			},
		),
		nil,
	)

	return services.Services{FrontDoor: services.FrontDoorClient{Doors: m}}
}

func TestFrontDoors(t *testing.T) {
	table := FrontDoors()
	client.AzureMockTestHelper(t, table, buildFrontDoorsServices, client.TestOptions{})
}

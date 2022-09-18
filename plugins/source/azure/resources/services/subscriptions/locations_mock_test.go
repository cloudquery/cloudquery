// Auto generated code - DO NOT EDIT.

package subscriptions

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
)

func TestSubscriptionsLocations(t *testing.T) {
	client.MockTestHelper(t, Locations(), createLocationsMock)
}

func createLocationsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSubscriptionsLocationsClient(ctrl)
	s := services.Services{
		Subscriptions: services.SubscriptionsClient{
			Locations: mockClient,
		},
	}

	data := armsubscriptions.Location{}
	require.Nil(t, faker.FakeObject(&data))
	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/" + *data.ID
	data.ID = &id

	pager := runtime.NewPager(runtime.PagingHandler[armsubscriptions.ClientListLocationsResponse]{
		More: func(page armsubscriptions.ClientListLocationsResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *armsubscriptions.ClientListLocationsResponse) (armsubscriptions.ClientListLocationsResponse, error) {
			return armsubscriptions.ClientListLocationsResponse{
				LocationListResult: armsubscriptions.LocationListResult{
					Value: []*armsubscriptions.Location{&data},
				},
			}, nil
		},
	})

	mockClient.EXPECT().NewListLocationsPager(gomock.Any(), gomock.Any()).Return(
		pager,
	)
	return s
}

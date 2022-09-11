// Auto generated code - DO NOT EDIT.

package subscriptions

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
	fieldsToIgnore := []string{"Response"}
	require.Nil(t, faker.FakeData(&data, fakerOptions.WithIgnoreInterface(true), fakerOptions.WithRecursionMaxDepth(2), fakerOptions.WithFieldsToIgnore(fieldsToIgnore...), fakerOptions.WithRandomMapAndSliceMinSize(1), fakerOptions.WithRandomMapAndSliceMaxSize(1)))

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

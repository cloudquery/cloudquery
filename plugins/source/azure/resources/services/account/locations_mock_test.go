package account

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildAccountLocationsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	m := mocks.NewMockSubscriptionsClient(ctrl)

	var subscriptionID string
	if err := faker.FakeData(&subscriptionID); err != nil {
		t.Fatal(err)
	}

	var model armsubscriptions.Location
	if err := faker.FakeData(&model); err != nil {
		t.Fatal(err)
	}
	pager := runtime.NewPager[armsubscriptions.ClientListLocationsResponse](runtime.PagingHandler[armsubscriptions.ClientListLocationsResponse]{
		More: func(page armsubscriptions.ClientListLocationsResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *armsubscriptions.ClientListLocationsResponse) (armsubscriptions.ClientListLocationsResponse, error) {
			return armsubscriptions.ClientListLocationsResponse{
				LocationListResult: armsubscriptions.LocationListResult{
					Value: []*armsubscriptions.Location{&model},
				},
			}, nil
		},
	})
	m.EXPECT().NewListLocationsPager(gomock.Any(), gomock.Any()).Return(
		pager,
	)

	return services.Services{
		Subscriptions: services.Subscriptions{
			SubscriptionID: subscriptionID,
			Subscriptions:  m,
		},
	}
}

func TestComputeAccountLocations(t *testing.T) {
	client.AzureMockTestHelper(t, Locations(), buildAccountLocationsMock, client.TestOptions{})
}

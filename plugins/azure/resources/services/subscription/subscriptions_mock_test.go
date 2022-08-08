package subscription

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

func buildSubscriptionsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	m := mocks.NewMockSubscriptionsClient(ctrl)

	var subscriptionID string
	if err := faker.FakeData(&subscriptionID); err != nil {
		t.Fatal(err)
	}

	var model armsubscriptions.Subscription
	if err := faker.FakeData(&model); err != nil {
		t.Fatal(err)
	}
	pager := runtime.NewPager[armsubscriptions.ClientListResponse](runtime.PagingHandler[armsubscriptions.ClientListResponse]{
		More: func(page armsubscriptions.ClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *armsubscriptions.ClientListResponse) (armsubscriptions.ClientListResponse, error) {
			return armsubscriptions.ClientListResponse{
				SubscriptionListResult: armsubscriptions.SubscriptionListResult{
					NextLink: nil,
					Value:    []*armsubscriptions.Subscription{&model},
				},
			}, nil
		},
	})
	m.EXPECT().NewListPager(gomock.Any()).Return(
		pager,
	)

	return services.Services{
		Subscriptions: services.Subscriptions{
			SubscriptionID: subscriptionID,
			Subscriptions:  m,
		},
	}
}

func TestSubscriptions(t *testing.T) {
	client.AzureMockTestHelper(t, Subscriptions(), buildSubscriptionsMock, client.TestOptions{})
}

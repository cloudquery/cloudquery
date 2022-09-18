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

func TestSubscriptionsSubscriptions(t *testing.T) {
	client.MockTestHelper(t, Subscriptions(), createSubscriptionsMock)
}

func createSubscriptionsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSubscriptionsSubscriptionsClient(ctrl)
	s := services.Services{
		Subscriptions: services.SubscriptionsClient{
			Subscriptions: mockClient,
		},
	}

	data := armsubscriptions.Subscription{}
	require.Nil(t, faker.FakeObject(&data))
	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/" + *data.ID
	data.ID = &id

	pager := runtime.NewPager(runtime.PagingHandler[armsubscriptions.ClientListResponse]{
		More: func(page armsubscriptions.ClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *armsubscriptions.ClientListResponse) (armsubscriptions.ClientListResponse, error) {
			return armsubscriptions.ClientListResponse{
				SubscriptionListResult: armsubscriptions.SubscriptionListResult{
					Value: []*armsubscriptions.Subscription{&data},
				},
			}, nil
		},
	})

	mockClient.EXPECT().NewListPager(gomock.Any()).Return(
		pager,
	)
	return s
}

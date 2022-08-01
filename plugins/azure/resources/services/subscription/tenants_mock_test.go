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

func buildTenantsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	m := mocks.NewMockTenantsClient(ctrl)

	var subscriptionID string
	if err := faker.FakeData(&subscriptionID); err != nil {
		t.Fatal(err)
	}

	var model armsubscriptions.TenantIDDescription
	if err := faker.FakeData(&model); err != nil {
		t.Fatal(err)
	}
	pager := runtime.NewPager[armsubscriptions.TenantsClientListResponse](runtime.PagingHandler[armsubscriptions.TenantsClientListResponse]{
		More: func(page armsubscriptions.TenantsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *armsubscriptions.TenantsClientListResponse) (armsubscriptions.TenantsClientListResponse, error) {
			return armsubscriptions.TenantsClientListResponse{
				TenantListResult: armsubscriptions.TenantListResult{
					NextLink: nil,
					Value:    []*armsubscriptions.TenantIDDescription{&model},
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
			Tenants:        m,
		},
	}
}

func TestTenants(t *testing.T) {
	client.AzureMockTestHelper(t, Tenants(), buildTenantsMock, client.TestOptions{})
}

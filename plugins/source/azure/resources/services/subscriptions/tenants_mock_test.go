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

func TestSubscriptionsTenants(t *testing.T) {
	client.MockTestHelper(t, Tenants(), createTenantsMock)
}

func createTenantsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSubscriptionsTenantsClient(ctrl)
	s := services.Services{
		Subscriptions: services.SubscriptionsClient{
			Tenants: mockClient,
		},
	}

	data := armsubscriptions.TenantIDDescription{}
	fieldsToIgnore := []string{"Response"}
	require.Nil(t, faker.FakeData(&data, fakerOptions.WithIgnoreInterface(true), fakerOptions.WithRecursionMaxDepth(2), fakerOptions.WithFieldsToIgnore(fieldsToIgnore...), fakerOptions.WithRandomMapAndSliceMinSize(1), fakerOptions.WithRandomMapAndSliceMaxSize(1)))

	pager := runtime.NewPager(runtime.PagingHandler[armsubscriptions.TenantsClientListResponse]{
		More: func(page armsubscriptions.TenantsClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *armsubscriptions.TenantsClientListResponse) (armsubscriptions.TenantsClientListResponse, error) {
			return armsubscriptions.TenantsClientListResponse{
				TenantListResult: armsubscriptions.TenantListResult{
					Value: []*armsubscriptions.TenantIDDescription{&data},
				},
			}, nil
		},
	})

	mockClient.EXPECT().NewListPager(gomock.Any()).Return(
		pager,
	)
	return s
}

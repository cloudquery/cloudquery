// Auto generated code - DO NOT EDIT.

package search

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/search/mgmt/2020-08-01/search"
)

func TestSearchServices(t *testing.T) {
	client.MockTestHelper(t, Services(), createServicesMock)
}

func createServicesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSearchServicesClient(ctrl)
	s := services.Services{
		Search: services.SearchClient{
			Services: mockClient,
		},
	}

	data := search.Service{}
	require.Nil(t, faker.FakeObject(&data))

	result := search.NewServiceListResultPage(search.ServiceListResult{Value: &[]search.Service{data}}, func(ctx context.Context, result search.ServiceListResult) (search.ServiceListResult, error) {
		return search.ServiceListResult{}, nil
	})

	mockClient.EXPECT().ListBySubscription(gomock.Any(), nil).Return(result, nil)
	return s
}

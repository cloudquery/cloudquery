// Auto generated code - DO NOT EDIT.

package cdn

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/cdn/mgmt/2020-09-01/cdn"
)

func createRoutesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockCDNRoutesClient(ctrl)
	s := services.Services{
		CDN: services.CDNClient{
			Routes: mockClient,
		},
	}

	data := cdn.Route{}
	require.Nil(t, faker.FakeObject(&data))

	result := cdn.NewRouteListResultPage(cdn.RouteListResult{Value: &[]cdn.Route{data}}, func(ctx context.Context, result cdn.RouteListResult) (cdn.RouteListResult, error) {
		return cdn.RouteListResult{}, nil
	})

	mockClient.EXPECT().ListByEndpoint(gomock.Any(), "test", "test", "test").Return(result, nil)
	return s
}

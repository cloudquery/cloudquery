// Auto generated code - DO NOT EDIT.

package web

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web"
)

func createVnetConnectionsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockWebVnetConnectionsClient(ctrl)
	s := services.Services{
		Web: services.WebClient{
			VnetConnections: mockClient,
		},
	}

	data := web.VnetInfo{}
	require.Nil(t, faker.FakeObject(&data))

	result := web.NewVnetInfoListResultPage(web.VnetInfoListResult{Value: &[]web.VnetInfo{data}}, func(ctx context.Context, result web.VnetInfoListResult) (web.VnetInfoListResult, error) {
		return web.VnetInfoListResult{}, nil
	})

	mockClient.EXPECT().GetVnetConnection(gomock.Any(), "test", "test").Return(result, nil)
	return s
}

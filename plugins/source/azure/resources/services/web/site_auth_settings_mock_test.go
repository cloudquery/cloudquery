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

func createSiteAuthSettingsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockWebSiteAuthSettingsClient(ctrl)
	s := services.Services{
		Web: services.WebClient{
			SiteAuthSettings: mockClient,
		},
	}

	data := web.SiteAuthSettings{}
	require.Nil(t, faker.FakeObject(&data))

	result := web.NewSiteAuthSettingsListResultPage(web.SiteAuthSettingsListResult{Value: &[]web.SiteAuthSettings{data}}, func(ctx context.Context, result web.SiteAuthSettingsListResult) (web.SiteAuthSettingsListResult, error) {
		return web.SiteAuthSettingsListResult{}, nil
	})

	mockClient.EXPECT().GetAuthSettings(gomock.Any(), "test", "test").Return(result, nil)
	return s
}

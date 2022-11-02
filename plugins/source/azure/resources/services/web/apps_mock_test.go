// Auto generated code - DO NOT EDIT.

package web

import (
	"context"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web"
)

func TestWebApps(t *testing.T) {
	client.MockTestHelper(t, Apps(), createAppsMock)
}

func createAppsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockWebAppsClient(ctrl)
	s := services.Services{
		Web: services.WebClient{
			Apps:               mockClient,
			SiteAuthSettings:   createSiteAuthSettingsMock(t, ctrl).Web.SiteAuthSettings,
			VnetConnections:    createVnetConnectionsMock(t, ctrl).Web.VnetConnections,
			PublishingProfiles: createPublishingProfilesMock(t, ctrl).Web.PublishingProfiles,
			SiteAuthSettingsV2: createSiteAuthSettingsV2Mock(t, ctrl).Web.SiteAuthSettingsV2,
			Functions:          createFunctionsMock(t, ctrl).Web.Functions,
		},
	}

	data := web.Site{}
	require.Nil(t, faker.FakeObject(&data))

	// Ensure name and ID are consistent so we can reference it in other mock
	name := "test"
	data.Name = &name

	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := web.NewAppCollectionPage(web.AppCollection{Value: &[]web.Site{data}}, func(ctx context.Context, result web.AppCollection) (web.AppCollection, error) {
		return web.AppCollection{}, nil
	})

	vnetName := "test"
	result.Values()[0].SiteConfig.VnetName = &vnetName
	resourceGroup := "test"
	result.Values()[0].ResourceGroup = &resourceGroup
	mockClient.EXPECT().List(gomock.Any()).Return(result, nil)
	return s
}

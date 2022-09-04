package web

import (
	"bytes"
	"context"
	"encoding/xml"
	"io"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web"
	"github.com/Azure/go-autorest/autorest"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildWebAppsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	siteAuthSettings := mocks.NewMockWebSiteAuthSettingsClient(ctrl)
	apps := mocks.NewMockWebAppsClient(ctrl)
	s := services.Services{
		Web: services.WebClient{
			Apps:             apps,
			SiteAuthSettings: siteAuthSettings,
		},
	}

	site := web.Site{}
	err := faker.FakeData(&site)
	if err != nil {
		t.Errorf("failed building mock %s", err)
	}

	ip := faker.IPv4()
	(*site.SiteProperties.SiteConfig.ScmIPSecurityRestrictions)[0].SubnetMask = &ip
	page := web.NewAppCollectionPage(web.AppCollection{Value: &[]web.Site{site}}, func(ctx context.Context, collection web.AppCollection) (web.AppCollection, error) {
		return web.AppCollection{}, nil
	})
	apps.EXPECT().List(gomock.Any()).Return(page, nil)
	var vi web.VnetInfo
	require.NoError(t, faker.FakeDataSkipFields(&vi, []string{"Routes"}))
	apps.EXPECT().GetVnetConnection(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(vi, nil)

	pp := PublishData{
		PublishData: []PublishProfile{
			{
				PublishUrl: "test",
				UserName:   "test",
				UserPWD:    "test"},
		},
	}

	data, err := xml.Marshal([]PublishData{pp})
	if err != nil {
		t.Errorf("failed building xml %s", err)
	}

	value := io.NopCloser(bytes.NewReader(data)) // r type is io.ReadCloser
	response := web.ReadCloser{Response: autorest.Response{Response: &http.Response{Body: value}}}
	apps.EXPECT().ListPublishingProfileXMLWithSecrets(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(response, nil)

	auth := web.SiteAuthSettings{}
	err = faker.FakeData(&auth)
	if err != nil {
		t.Errorf("failed building mock %s", err)
	}
	siteAuthSettings.EXPECT().GetAuthSettings(gomock.Any(), gomock.Any(), gomock.Any()).Return(auth, nil)

	return s
}

func TestWebApps(t *testing.T) {
	client.AzureMockTestHelper(t, WebApps(), buildWebAppsMock, client.TestOptions{})
}

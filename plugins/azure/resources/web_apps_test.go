package resources_test

import (
	"bytes"
	"context"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web"
	"github.com/Azure/go-autorest/autorest"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/cq-provider-azure/resources"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildWebAppsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	apps := mocks.NewMockAppsClient(ctrl)
	s := services.Services{
		Web: services.WebClient{
			Apps: apps,
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

	pp := resources.PublishData{
		PublishData: []resources.PublishProfile{
			{
				PublishUrl: "test",
				UserName:   "test",
				UserPWD:    "test"},
		},
	}

	data, err := xml.Marshal([]resources.PublishData{pp})
	if err != nil {
		t.Errorf("failed building xml %s", err)
	}

	value := ioutil.NopCloser(bytes.NewReader(data)) // r type is io.ReadCloser
	response := web.ReadCloser{Response: autorest.Response{Response: &http.Response{Body: value}}}
	apps.EXPECT().ListPublishingProfileXMLWithSecrets(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(response, nil)

	auth := web.SiteAuthSettings{}
	err = faker.FakeData(&auth)
	if err != nil {
		t.Errorf("failed building mock %s", err)
	}
	apps.EXPECT().GetAuthSettings(gomock.Any(), gomock.Any(), gomock.Any()).Return(auth, nil)

	return s
}

func TestWebApps(t *testing.T) {
	azureTestHelper(t, resources.WebApps(), buildWebAppsMock)
}

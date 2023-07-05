package cdns

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/digitalocean/godo"
	"github.com/golang/mock/gomock"
)

func createCdn(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCdnService(ctrl)

	var data []godo.CDN
	if err := faker.FakeObject(&data); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().List(gomock.Any(), gomock.Any()).Return(data, &godo.Response{}, nil)

	return client.Services{
		Cdn: m,
	}
}

func TestCdn(t *testing.T) {
	client.MockTestHelper(t, Cdns(), createCdn, client.TestOptions{})
}

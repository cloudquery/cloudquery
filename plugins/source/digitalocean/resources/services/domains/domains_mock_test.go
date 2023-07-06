package domains

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/digitalocean/godo"
	"github.com/golang/mock/gomock"
)

func createDomains(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDomainsService(ctrl)

	var data []godo.Domain
	if err := faker.FakeObject(&data); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().List(gomock.Any(), gomock.Any()).Return(data, &godo.Response{}, nil)

	createRecords(t, m)

	return client.Services{
		Domains: m,
	}
}

func TestDomains(t *testing.T) {
	client.MockTestHelper(t, Domains(), createDomains, client.TestOptions{})
}

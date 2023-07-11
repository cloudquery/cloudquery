package droplets

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/digitalocean/godo"
	"github.com/golang/mock/gomock"
)

func createDroplets(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDropletsService(ctrl)

	var data []godo.Droplet
	if err := faker.FakeObject(&data); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().List(gomock.Any(), gomock.Any()).Return(data, &godo.Response{}, nil)

	createNeighbors(t, m)

	return client.Services{
		Droplets: m,
	}
}

func TestDroplets(t *testing.T) {
	client.MockTestHelper(t, Droplets(), createDroplets, client.TestOptions{})
}

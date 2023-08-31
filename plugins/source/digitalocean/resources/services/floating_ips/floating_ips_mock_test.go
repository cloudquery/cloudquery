package floating_ips

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/digitalocean/godo"
	"github.com/golang/mock/gomock"
)

func createFloatingIps(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockFloatingIpsService(ctrl)

	var data []godo.FloatingIP
	if err := faker.FakeObject(&data); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().List(gomock.Any(), gomock.Any()).Return(data, &godo.Response{}, nil)

	return client.Services{
		FloatingIps: m,
	}
}

func TestFloatingIps(t *testing.T) {
	client.MockTestHelper(t, FloatingIps(), createFloatingIps, client.TestOptions{})
}

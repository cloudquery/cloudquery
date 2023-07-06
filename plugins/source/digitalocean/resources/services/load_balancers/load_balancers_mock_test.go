package load_balancers

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/digitalocean/godo"
	"github.com/golang/mock/gomock"
)

func createLoadBalancers(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockLoadBalancersService(ctrl)

	var data []godo.LoadBalancer
	if err := faker.FakeObject(&data); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().List(gomock.Any(), gomock.Any()).Return(data, &godo.Response{}, nil)

	return client.Services{
		LoadBalancers: m,
	}
}

func TestLoadBalancers(t *testing.T) {
	client.MockTestHelper(t, LoadBalancers(), createLoadBalancers, client.TestOptions{})
}

package vpcs

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/digitalocean/godo"
	"github.com/golang/mock/gomock"
)

func createVpcs(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockVpcsService(ctrl)

	var data []*godo.VPC
	if err := faker.FakeObject(&data); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().List(gomock.Any(), gomock.Any()).Return(data, &godo.Response{}, nil)

	createMembers(t, m)

	return client.Services{
		Vpcs: m,
	}
}

func TestVpcs(t *testing.T) {
	client.MockTestHelper(t, Vpcs(), createVpcs, client.TestOptions{})
}

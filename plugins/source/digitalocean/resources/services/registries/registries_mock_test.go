package registries

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/digitalocean/godo"
	"github.com/golang/mock/gomock"
)

func createRegistry(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRegistryService(ctrl)

	var data godo.Registry
	if err := faker.FakeObject(&data); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().Get(gomock.Any()).Return(&data, nil, nil)

	// add children mocks
	createRepositories(t, m)

	return client.Services{
		Registry: m,
	}
}

func TestRegistry(t *testing.T) {
	client.MockTestHelper(t, Registries(), createRegistry, client.TestOptions{})
}

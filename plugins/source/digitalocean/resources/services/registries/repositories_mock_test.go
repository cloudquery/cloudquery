package registries

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/digitalocean/godo"
	"github.com/golang/mock/gomock"
)

func createRepositories(t *testing.T, m *mocks.MockRegistryService) {
	var data []*godo.Repository
	if err := faker.FakeData(&data); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListRepositories(gomock.Any(), gomock.Any(), gomock.Any()).Return(data, &godo.Response{}, nil)
}

package droplets

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/digitalocean/godo"
	"github.com/golang/mock/gomock"
)

func createNeighbors(t *testing.T, m *mocks.MockDropletsService) {
	var data []godo.Droplet
	if err := faker.FakeObject(&data); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().Neighbors(gomock.Any(), gomock.Any()).Return(data, &godo.Response{}, nil)
}

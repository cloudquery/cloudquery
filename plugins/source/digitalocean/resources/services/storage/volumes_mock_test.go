package storage

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/digitalocean/godo"
	"github.com/golang/mock/gomock"
)

func createVolumes(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockStorageService(ctrl)

	var data []godo.Volume
	if err := faker.FakeObject(&data); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListVolumes(gomock.Any(), gomock.Any()).Return(data, &godo.Response{}, nil)

	return client.Services{
		Storage: m,
	}
}

func TestVolumes(t *testing.T) {
	client.MockTestHelper(t, Volumes(), createVolumes, client.TestOptions{})
}

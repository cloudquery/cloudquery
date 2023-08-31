package images

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/digitalocean/godo"
	"github.com/golang/mock/gomock"
)

func createImages(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockImagesService(ctrl)

	var data []godo.Image
	if err := faker.FakeObject(&data); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().List(gomock.Any(), gomock.Any()).Return(data, &godo.Response{}, nil)

	return client.Services{
		Images: m,
	}
}

func TestImages(t *testing.T) {
	client.MockTestHelper(t, Images(), createImages, client.TestOptions{})
}

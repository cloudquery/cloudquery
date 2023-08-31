package sizes

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/digitalocean/godo"
	"github.com/golang/mock/gomock"
)

func createSizes(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSizesService(ctrl)

	var data []godo.Size
	if err := faker.FakeObject(&data); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().List(gomock.Any(), gomock.Any()).Return(data, &godo.Response{}, nil)

	return client.Services{
		Sizes: m,
	}
}

func TestSizes(t *testing.T) {
	client.MockTestHelper(t, Sizes(), createSizes, client.TestOptions{})
}

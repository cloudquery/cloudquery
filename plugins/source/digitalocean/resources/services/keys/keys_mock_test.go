package keys

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/digitalocean/godo"
	"github.com/golang/mock/gomock"
)

func createKeys(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockKeysService(ctrl)

	var data []godo.Key
	if err := faker.FakeObject(&data); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().List(gomock.Any(), gomock.Any()).Return(data, &godo.Response{}, nil)

	return client.Services{
		Keys: m,
	}
}

func TestKeys(t *testing.T) {
	client.MockTestHelper(t, Keys(), createKeys, client.TestOptions{})
}

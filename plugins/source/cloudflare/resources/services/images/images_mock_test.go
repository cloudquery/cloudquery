package images

import (
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildImages(t *testing.T, ctrl *gomock.Controller) client.Clients {
	mock := mocks.NewMockApi(ctrl)

	faker.SetIgnoreInterface(true)

	var image cloudflare.Image
	if err := faker.FakeData(&image); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListImages(
		gomock.Any(),
		client.TestAccountID,
		cloudflare.PaginationOptions{},
	).Return(
		[]cloudflare.Image{image},
		nil,
	)

	return client.Clients{
		client.TestAccountID: mock,
	}
}

func TestImages(t *testing.T) {
	client.MockTestHelper(t, Images(), buildImages)
}

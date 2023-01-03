package items

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/hackernews/client"
	"github.com/cloudquery/cloudquery/plugins/source/hackernews/client/mocks"
	"github.com/cloudquery/cloudquery/plugins/source/hackernews/client/services"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/hermanschaaf/hackernews"
)

func buildItemsMock(t *testing.T, ctrl *gomock.Controller) services.HackernewsClient {
	m := mocks.NewMockHackernewsClient(ctrl)
	f := hackernews.Item{}
	err := faker.FakeObject(&f)
	if err != nil {
		t.Fatal(err)
	}
	maxID := 1
	m.EXPECT().MaxItemID(gomock.Any()).Return(maxID, nil)
	m.EXPECT().GetItem(gomock.Any(), gomock.Any()).Return(f, nil)
	return m
}

func TestItems(t *testing.T) {
	client.MockTestHelper(t, Items(), buildItemsMock, client.TestOptions{})
}

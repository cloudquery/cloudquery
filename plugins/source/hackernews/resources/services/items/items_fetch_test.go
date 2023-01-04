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

func buildItemsMockNoCursor(t *testing.T, ctrl *gomock.Controller) services.HackernewsClient {
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

// In this test, we mock the state backend to return a cursor value of "", meaning no cursor. The max item ID will be 1.
// We then expect one item to be fetched, and the cursor should be set to "1".
func TestItems_NoCursor(t *testing.T) {
	ctrl := gomock.NewController(t)
	mbe := mocks.NewMockBackend(ctrl)
	mbe.EXPECT().Get(gomock.Any(), gomock.Any(), gomock.Any()).Return("", nil)
	mbe.EXPECT().Set(gomock.Any(), "hackernews_items", "id", "1").Return(nil)
	client.MockTestHelper(t, Items(), buildItemsMockNoCursor, client.TestOptions{
		Backend: mbe,
	})
}

func buildItemsMockWithCursor(t *testing.T, ctrl *gomock.Controller) services.HackernewsClient {
	m := mocks.NewMockHackernewsClient(ctrl)
	f := hackernews.Item{}
	err := faker.FakeObject(&f)
	if err != nil {
		t.Fatal(err)
	}
	maxID := 10
	m.EXPECT().MaxItemID(gomock.Any()).Return(maxID, nil)
	m.EXPECT().GetItem(gomock.Any(), gomock.Any()).Return(f, nil)
	return m
}

// In this test, we mock the state backend to return a cursor value of "9", and the max item ID will be 10.
// We then expect that one item will be fetched, and the cursor will be set to "10".
func TestItems_WithCursor(t *testing.T) {
	ctrl := gomock.NewController(t)
	mbe := mocks.NewMockBackend(ctrl)
	mbe.EXPECT().Get(gomock.Any(), gomock.Any(), gomock.Any()).Return("9", nil)
	mbe.EXPECT().Set(gomock.Any(), "hackernews_items", "id", "10").Return(nil)
	client.MockTestHelper(t, Items(), buildItemsMockWithCursor, client.TestOptions{
		Backend: mbe,
	})
}

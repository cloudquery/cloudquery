package items

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/hackernews/v3/client"
	"github.com/cloudquery/cloudquery/plugins/source/hackernews/v3/client/mocks"
	"github.com/cloudquery/cloudquery/plugins/source/hackernews/v3/client/services"
	"github.com/cloudquery/plugin-sdk/v4/faker"
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
	maxID := 5
	m.EXPECT().MaxItemID(gomock.Any()).Return(maxID, nil)
	m.EXPECT().GetItem(gomock.Any(), gomock.Any()).Times(5).Return(f, nil)
	return m
}

// In this test, we mock the state backend to return a cursor value of "", meaning no cursor. The max item ID will be 5.
// We then expect one item to be fetched, and the cursor should be set to "5".
func TestItems_NoCursor(t *testing.T) {
	ctrl := gomock.NewController(t)
	mbe := mocks.NewMockBackendClient(ctrl)
	mbe.EXPECT().GetKey(gomock.Any(), "hackernews_items").Return("", nil)
	mbe.EXPECT().SetKey(gomock.Any(), "hackernews_items", "5").Times(1).Return(nil)
	mbe.EXPECT().Flush(gomock.Any()).MinTimes(1).Return(nil)
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
	m.EXPECT().GetItem(gomock.Any(), gomock.Any()).Times(5).Return(f, nil)
	return m
}

// In this test, we mock the state backend to return a cursor value of "5", and the max item ID will be 10.
// We then expect that one item will be fetched, and the cursor will be set to "10".
func TestItems_WithCursor(t *testing.T) {
	ctrl := gomock.NewController(t)
	mbe := mocks.NewMockBackendClient(ctrl)
	mbe.EXPECT().GetKey(gomock.Any(), gomock.Any()).Return("5", nil)
	mbe.EXPECT().SetKey(gomock.Any(), "hackernews_items", "10").Times(1).Return(nil)
	mbe.EXPECT().Flush(gomock.Any()).MinTimes(1).Return(nil)
	client.MockTestHelper(t, Items(), buildItemsMockWithCursor, client.TestOptions{
		Backend: mbe,
	})
}

func buildItemsMockWithStartTime(t *testing.T, ctrl *gomock.Controller) services.HackernewsClient {
	m := mocks.NewMockHackernewsClient(ctrl)
	maxID := 5
	start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	times := make([]int, 5)
	for i := 1; i <= maxID; i++ {
		times[i-1] = int(start.AddDate(0, 0, i).Unix())
	}
	m.EXPECT().MaxItemID(gomock.Any()).AnyTimes().Return(maxID, nil)
	m.EXPECT().GetItem(gomock.Any(), gomock.Any()).AnyTimes().DoAndReturn(
		func(ctx context.Context, id int) (hackernews.Item, error) {
			f := hackernews.Item{}
			err := faker.FakeObject(&f)
			if err != nil {
				return f, errors.New("failed to fake object")
			}
			if id <= len(times) {
				f.ID = id
				f.Time = times[id-1]
				return f, nil
			}
			return hackernews.Item{}, errors.New("not found")
		})
	return m
}

// In this test, we request a certain start time and expect that the sync will start from the first post after that start time.
func TestItems_WithStartTime(t *testing.T) {
	ctrl := gomock.NewController(t)
	mbe := mocks.NewMockBackendClient(ctrl)
	mbe.EXPECT().GetKey(gomock.Any(), gomock.Any()).Return("", nil)
	mbe.EXPECT().SetKey(gomock.Any(), "hackernews_items", "5").Return(nil)
	mbe.EXPECT().Flush(gomock.Any()).MinTimes(1).Return(nil)
	client.MockTestHelper(t, Items(), buildItemsMockWithStartTime, client.TestOptions{
		Backend:   mbe,
		StartTime: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
	})
}

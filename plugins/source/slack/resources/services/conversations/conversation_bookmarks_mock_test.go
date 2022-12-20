package conversations

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/slack/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/slack-go/slack"
)

func buildConversationBookmarksMock(t *testing.T, m *mocks.MockSlackClient) {
	bookmarks := make([]slack.Bookmark, 1)
	err := faker.FakeObject(&bookmarks)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListBookmarksContext(gomock.Any(), gomock.Any()).AnyTimes().Return(bookmarks, nil)
}

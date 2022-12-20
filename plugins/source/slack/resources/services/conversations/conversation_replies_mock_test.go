package conversations

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/slack/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/slack-go/slack"
)

func buildConversationRepliesMock(t *testing.T, m *mocks.MockSlackClient) {
	d := make([]slack.Message, 0, 1)
	err := faker.FakeObject(&d)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetConversationRepliesContext(gomock.Any(), gomock.Any()).Times(1).Return(d, true, "cursor1", nil)
	m.EXPECT().GetConversationRepliesContext(gomock.Any(), gomock.Any()).AnyTimes().Return(d, false, "", nil)
}

package conversations

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/cloudquery/plugins/source/slack/client/mocks"
	"github.com/cloudquery/cloudquery/plugins/source/slack/client/services"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/slack-go/slack"
)

func buildConversationsMock(t *testing.T, ctrl *gomock.Controller) services.SlackClient {
	m := mocks.NewMockSlackClient(ctrl)
	d := make([]slack.Channel, 0, 1)
	err := faker.FakeObject(&d)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetConversationsContext(gomock.Any(), gomock.Any()).Times(1).Return(d, "cursor1", nil)
	m.EXPECT().GetConversationsContext(gomock.Any(), gomock.Any()).Times(1).Return(d, "", nil)
	return m
}

func TestConversations(t *testing.T) {
	client.MockTestHelper(t, Conversations(), buildConversationsMock, client.TestOptions{})
}

package conversations

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/slack/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/slack-go/slack"
)

func buildConversationHistoriesMock(t *testing.T, m *mocks.MockSlackClient) {
	d := &slack.GetConversationHistoryResponse{}
	err := faker.FakeObject(&d)
	if err != nil {
		t.Fatal(err)
	}
	d.HasMore = false
	msgs := make([]slack.Message, 1)
	err = faker.FakeObject(&msgs)
	if err != nil {
		t.Fatal(err)
	}
	d.Messages = msgs

	m.EXPECT().GetConversationHistoryContext(gomock.Any(), gomock.Any()).AnyTimes().Return(d, nil)
}

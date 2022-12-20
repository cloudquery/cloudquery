package conversations

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/slack/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildConversationMembersMock(t *testing.T, m *mocks.MockSlackClient) {
	users := make([]string, 1)
	err := faker.FakeObject(&users)
	if err != nil {
		t.Fatal(err)
	}
	nextCursor := "testCursor"
	m.EXPECT().GetUsersInConversationContext(gomock.Any(), gomock.Any()).Times(1).Return(users, nextCursor, nil)
	m.EXPECT().GetUsersInConversationContext(gomock.Any(), gomock.Any()).AnyTimes().Return(users, "", nil)
}

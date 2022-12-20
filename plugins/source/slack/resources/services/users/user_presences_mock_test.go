package users

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/slack/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/slack-go/slack"
)

func addUserPresencesMock(t *testing.T, m *mocks.MockSlackClient) {
	d := &slack.UserPresence{}
	err := faker.FakeObject(&d)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetUserPresenceContext(gomock.Any(), gomock.Any()).Return(d, nil)
}

package users

import (
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/cloudquery/plugins/source/slack/client/mocks"
	"github.com/cloudquery/cloudquery/plugins/source/slack/client/services"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/slack-go/slack"
)

func buildUsersMock(t *testing.T, ctrl *gomock.Controller) services.SlackClient {
	m := mocks.NewMockSlackClient(ctrl)
	d := make([]slack.User, 0, 1)
	err := faker.FakeObject(&d)
	if err != nil {
		t.Fatal(err)
	}
	d[0].Updated = slack.JSONTime(time.Now().Unix())
	addUserPresencesMock(t, m)
	m.EXPECT().GetUsersContext(gomock.Any(), gomock.Any()).Return(d, nil)
	return m
}

func TestUsers(t *testing.T) {
	client.MockTestHelper(t, Users(), buildUsersMock, client.TestOptions{})
}

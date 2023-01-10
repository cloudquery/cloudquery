package teams

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/cloudquery/plugins/source/slack/client/mocks"
	"github.com/cloudquery/cloudquery/plugins/source/slack/client/services"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/slack-go/slack"
)

func buildTeamsMock(t *testing.T, ctrl *gomock.Controller) services.SlackClient {
	m := mocks.NewMockSlackClient(ctrl)
	var d *slack.TeamInfo
	err := faker.FakeObject(&d)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetTeamInfoContext(gomock.Any()).Return(d, nil)
	return m
}

func TestTeams(t *testing.T) {
	client.MockTestHelper(t, Teams(), buildTeamsMock, client.TestOptions{})
}

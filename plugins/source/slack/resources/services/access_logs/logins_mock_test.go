package access_logs

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/cloudquery/plugins/source/slack/client/mocks"
	"github.com/cloudquery/cloudquery/plugins/source/slack/client/services"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/slack-go/slack"
)

func buildLoginsMock(t *testing.T, ctrl *gomock.Controller) services.SlackClient {
	m := mocks.NewMockSlackClient(ctrl)
	d := make([]slack.Login, 0, 1)
	err := faker.FakeObject(&d)
	if err != nil {
		t.Fatal(err)
	}
	p1 := &slack.Paging{
		Count: 1000,
		Total: 1000,
		Page:  1,
		Pages: 2,
	}
	p2 := &slack.Paging{
		Count: 1000,
		Total: 20,
		Page:  2,
		Pages: 2,
	}
	m.EXPECT().GetAccessLogsContext(gomock.Any(), gomock.Any()).Times(1).Return(d, p1, nil)
	m.EXPECT().GetAccessLogsContext(gomock.Any(), gomock.Any()).Times(1).Return(d, p2, nil)

	return m
}

func TestLogins(t *testing.T) {
	client.MockTestHelper(t, Logins(), buildLoginsMock, client.TestOptions{})
}

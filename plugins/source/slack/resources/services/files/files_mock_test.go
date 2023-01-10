package files

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/cloudquery/plugins/source/slack/client/mocks"
	"github.com/cloudquery/cloudquery/plugins/source/slack/client/services"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/slack-go/slack"
)

func buildFilesMock(t *testing.T, ctrl *gomock.Controller) services.SlackClient {
	m := mocks.NewMockSlackClient(ctrl)
	d := make([]slack.File, 0, 1)
	err := faker.FakeObject(&d)
	if err != nil {
		t.Fatal(err)
	}
	p := &slack.Paging{
		Count: 1,
		Total: 2,
		Page:  1,
		Pages: 2,
	}
	m.EXPECT().GetFilesContext(gomock.Any(), gomock.Any()).Times(1).Return(d, p, nil)
	p2 := &slack.Paging{
		Count: 1,
		Total: 2,
		Page:  2,
		Pages: 2,
	}
	m.EXPECT().GetFilesContext(gomock.Any(), gomock.Any()).Times(1).Return(d, p2, nil)
	return m
}

func TestFiles(t *testing.T) {
	client.MockTestHelper(t, Files(), buildFilesMock, client.TestOptions{})
}

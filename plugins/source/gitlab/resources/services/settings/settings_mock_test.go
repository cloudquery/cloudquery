package settings

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/xanzy/go-gitlab"
)

func buildSettings(t *testing.T, ctrl *gomock.Controller) client.Services {
	settingMock := mocks.NewMockSettingsClient(ctrl)

	var settings *gitlab.Settings
	if err := faker.FakeObject(&settings); err != nil {
		t.Fatal(err)
	}

	settingMock.EXPECT().GetSettings(gomock.Any()).Return(settings, &gitlab.Response{}, nil)

	return client.Services{
		Settings: settingMock,
	}
}

func TestSettings(t *testing.T) {
	client.GitlabMockTestHelper(t, Settings(), buildSettings, client.TestOptions{})
}

package external

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/cloudquery/plugins/source/github/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v49/github"
)

func buildExternalGroups(t *testing.T, ctrl *gomock.Controller) client.GithubServices {
	mock := mocks.NewMockTeamsService(ctrl)

	var cs *github.ExternalGroupList
	if err := faker.FakeObject(&cs); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListExternalGroups(gomock.Any(), "testorg", gomock.Any()).Return(cs, &github.Response{}, nil)
	return client.GithubServices{Teams: mock}
}

func TestExternalGroups(t *testing.T) {
	client.GithubMockTestHelper(t, Groups(), buildExternalGroups, client.TestOptions{})
}

package resources

import (
	"testing"

	"github.com/cloudquery/cq-provider-github/client"
	"github.com/cloudquery/cq-provider-github/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v45/github"
)

func buildExternalGroups(t *testing.T, ctrl *gomock.Controller) client.GithubServices {
	mock := mocks.NewMockTeamsService(ctrl)

	var cs *github.ExternalGroupList
	if err := faker.FakeData(&cs); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListExternalGroups(gomock.Any(), "testorg", gomock.Any()).Return(cs, &github.Response{}, nil)
	return client.GithubServices{Teams: mock}
}

func TestExternalGroups(t *testing.T) {
	client.GithubMockTestHelper(t, ExternalGroups(), buildExternalGroups, client.TestOptions{})
}

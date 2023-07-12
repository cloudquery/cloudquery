package issues

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/cloudquery/plugins/source/github/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v49/github"
)

func buildIssues(t *testing.T, ctrl *gomock.Controller) client.GithubServices {
	mock := mocks.NewMockIssuesService(ctrl)

	var cs github.Issue
	if err := faker.FakeObject(&cs); err != nil {
		t.Fatal(err)
	}
	someId := int64(5555555)
	cs.Repository = &github.Repository{ID: &someId}

	mock.EXPECT().ListByRepo(gomock.Any(), "testorg", gomock.Any(), gomock.Any()).Return(
		[]*github.Issue{&cs}, &github.Response{}, nil)

	return client.GithubServices{Issues: mock}
}

func TestIssues(t *testing.T) {
	client.GithubMockTestHelper(t, Issues(), buildIssues, client.TestOptions{})
}

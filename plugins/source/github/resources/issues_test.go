package resources

import (
	"testing"

	"github.com/cloudquery/cq-provider-github/client"
	"github.com/cloudquery/cq-provider-github/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v45/github"
)

func buildIssues(t *testing.T, ctrl *gomock.Controller) client.GithubServices {
	mock := mocks.NewMockIssuesService(ctrl)

	var cs github.Issue
	if err := faker.FakeDataSkipFields(&cs, []string{"Repository"}); err != nil {
		t.Fatal(err)
	}
	someId := int64(5555555)
	cs.Repository = &github.Repository{ID: &someId}

	mock.EXPECT().ListByOrg(gomock.Any(), "testorg", gomock.Any()).Return(
		[]*github.Issue{&cs}, &github.Response{}, nil)

	return client.GithubServices{Issues: mock}
}

func TestIssues(t *testing.T) {
	client.GithubMockTestHelper(t, Issues(), buildIssues, client.TestOptions{})
}

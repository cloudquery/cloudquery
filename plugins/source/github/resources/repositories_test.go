package resources

import (
	"testing"

	"github.com/cloudquery/cq-provider-github/client"
	"github.com/cloudquery/cq-provider-github/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v45/github"
)

func buildRepositiories(t *testing.T, ctrl *gomock.Controller) client.GithubServices {
	mock := mocks.NewMockRepositoriesService(ctrl)

	var cs github.Repository
	if err := faker.FakeDataSkipFields(&cs, []string{"Parent", "Source", "TemplateRepository"}); err != nil {
		t.Fatal(err)
	}
	someId := int64(5555555)
	cs.Parent = &github.Repository{ID: &someId}
	cs.TemplateRepository = &github.Repository{ID: &someId}
	cs.Source = &github.Repository{ID: &someId}

	mock.EXPECT().ListByOrg(gomock.Any(), "testorg", gomock.Any()).Return(
		[]*github.Repository{&cs}, &github.Response{}, nil)

	return client.GithubServices{Repositories: mock}
}

func TestRepos(t *testing.T) {
	client.GithubMockTestHelper(t, Repositories(), buildRepositiories, client.TestOptions{})
}

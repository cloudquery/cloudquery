package repositories

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/cloudquery/plugins/source/github/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v48/github"
)

func buildRepositiories(t *testing.T, ctrl *gomock.Controller) client.GithubServices {
	mock := mocks.NewMockRepositoriesService(ctrl)

	var cs github.Repository
	if err := faker.FakeObject(&cs); err != nil {
		t.Fatal(err)
	}
	someId := int64(5555555)
	cs.Parent = &github.Repository{ID: &someId}
	cs.TemplateRepository = &github.Repository{ID: &someId}
	cs.Source = &github.Repository{ID: &someId}

	mock.EXPECT().ListByOrg(gomock.Any(), "testorg", gomock.Any()).Return(
		[]*github.Repository{&cs}, &github.Response{}, nil)

	var release github.RepositoryRelease
	if err := faker.FakeObject(&release); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().ListReleases(gomock.Any(), "testorg", gomock.Any(), gomock.Any()).Return(
		[]*github.RepositoryRelease{&release}, &github.Response{}, nil)

	var releaseAsset github.ReleaseAsset
	if err := faker.FakeObject(&releaseAsset); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().ListReleaseAssets(gomock.Any(), "testorg", gomock.Any(), gomock.Any(), gomock.Any()).Return(
		[]*github.ReleaseAsset{&releaseAsset}, &github.Response{}, nil)

	dependabot := buildDependabot(t, ctrl)

	return client.GithubServices{
		Dependabot:   dependabot,
		Repositories: mock,
	}
}

func TestRepos(t *testing.T) {
	client.GithubMockTestHelper(t, Repositories(), buildRepositiories, client.TestOptions{})
}

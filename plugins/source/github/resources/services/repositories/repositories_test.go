package repositories

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/cloudquery/plugins/source/github/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v49/github"
)

func buildRepositories(t *testing.T, ctrl *gomock.Controller) client.GithubServices {
	mock := mocks.NewMockRepositoriesService(ctrl)

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

	var branch github.Branch
	if err := faker.FakeObject(&branch); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().ListBranches(gomock.Any(), "testorg", gomock.Any(), gomock.Any()).Return(
		[]*github.Branch{&branch}, &github.Response{}, nil)

	var protection github.Protection
	if err := faker.FakeObject(&protection); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().GetBranchProtection(gomock.Any(), "testorg", gomock.Any(), gomock.Any()).Return(
		&protection, &github.Response{}, nil)

	dependabot := buildDependabot(t, ctrl)

	return client.GithubServices{
		Dependabot:   dependabot,
		Repositories: mock,
	}
}

func TestRepos(t *testing.T) {
	client.GithubMockTestHelper(t, Repositories(), buildRepositories, client.TestOptions{})
}

package repositories

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/cloudquery/plugins/source/github/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v59/github"
	"github.com/stretchr/testify/require"
)

func buildRepositories(t *testing.T, ctrl *gomock.Controller) client.GithubServices {
	mock := mocks.NewMockRepositoriesService(ctrl)
	dependencyGraph := mocks.NewMockDependencyGraphService(ctrl)

	var sbom github.SBOM
	require.NoError(t, faker.FakeObject(&sbom))
	dependencyGraph.EXPECT().GetSBOM(gomock.Any(), "testorg", gomock.Any()).Return(
		&sbom, &github.Response{}, nil)

	var release github.RepositoryRelease
	require.NoError(t, faker.FakeObject(&release))

	mock.EXPECT().ListReleases(gomock.Any(), "testorg", gomock.Any(), gomock.Any()).Return(
		[]*github.RepositoryRelease{&release}, &github.Response{}, nil)

	var releaseAsset github.ReleaseAsset
	require.NoError(t, faker.FakeObject(&releaseAsset))

	mock.EXPECT().ListReleaseAssets(gomock.Any(), "testorg", gomock.Any(), gomock.Any(), gomock.Any()).Return(
		[]*github.ReleaseAsset{&releaseAsset}, &github.Response{}, nil)

	var branch github.Branch
	require.NoError(t, faker.FakeObject(&branch))

	mock.EXPECT().ListBranches(gomock.Any(), "testorg", gomock.Any(), gomock.Any()).Return(
		[]*github.Branch{&branch}, &github.Response{}, nil)

	var protection github.Protection
	require.NoError(t, faker.FakeObject(&protection))

	mock.EXPECT().GetBranchProtection(gomock.Any(), "testorg", gomock.Any(), gomock.Any()).Return(
		&protection, &github.Response{}, nil)

	dependabot := buildDependabot(t, ctrl)

	var key github.Key
	require.NoError(t, faker.FakeObject(&key))

	mock.EXPECT().ListKeys(gomock.Any(), "testorg", gomock.Any(), gomock.Any()).Return(
		[]*github.Key{&key}, &github.Response{}, nil)

	return client.GithubServices{
		Dependabot:      dependabot,
		Repositories:    mock,
		DependencyGraph: dependencyGraph,
	}
}

func TestRepos(t *testing.T) {
	client.GithubMockTestHelper(t, Repositories(), buildRepositories, client.TestOptions{})
}

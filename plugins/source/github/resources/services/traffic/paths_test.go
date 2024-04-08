package traffic

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/cloudquery/plugins/source/github/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v59/github"
	"github.com/stretchr/testify/require"
)

func buildPaths(t *testing.T, ctrl *gomock.Controller) client.GithubServices {
	repositoriesMock := mocks.NewMockRepositoriesService(ctrl)

	var paths []*github.TrafficPath
	require.NoError(t, faker.FakeObject(&paths))

	repositoriesMock.EXPECT().ListTrafficPaths(gomock.Any(), "test string", "test string").Return(paths, nil, nil)
	return client.GithubServices{Repositories: repositoriesMock}
}

func TestPaths(t *testing.T) {
	client.GithubMockTestHelper(t, Paths(), buildPaths, client.TestOptions{})
}

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

func buildClones(t *testing.T, ctrl *gomock.Controller) client.GithubServices {
	repositoriesMock := mocks.NewMockRepositoriesService(ctrl)

	var clones *github.TrafficClones
	require.NoError(t, faker.FakeObject(&clones))

	opts := github.TrafficBreakdownOptions{}

	repositoriesMock.EXPECT().ListTrafficClones(gomock.Any(), "test string", "test string", &opts).Return(clones, nil, nil)
	return client.GithubServices{Repositories: repositoriesMock}
}

func TestClones(t *testing.T) {
	client.GithubMockTestHelper(t, Clones(), buildClones, client.TestOptions{})
}

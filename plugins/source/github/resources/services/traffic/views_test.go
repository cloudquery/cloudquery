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

func buildViews(t *testing.T, ctrl *gomock.Controller) client.GithubServices {
	repositoriesMock := mocks.NewMockRepositoriesService(ctrl)

	var views *github.TrafficViews
	require.NoError(t, faker.FakeObject(&views))

	opts := github.TrafficBreakdownOptions{}

	repositoriesMock.EXPECT().ListTrafficViews(gomock.Any(), "test string", "test string", &opts).Return(views, nil, nil)
	return client.GithubServices{Repositories: repositoriesMock}
}

func TestViews(t *testing.T) {
	client.GithubMockTestHelper(t, Views(), buildViews, client.TestOptions{})
}

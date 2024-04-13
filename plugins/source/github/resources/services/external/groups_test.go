package external

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/cloudquery/plugins/source/github/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v59/github"
	"github.com/stretchr/testify/require"
)

func buildExternalGroups(t *testing.T, ctrl *gomock.Controller) client.GithubServices {
	mock := mocks.NewMockTeamsService(ctrl)

	var cs *github.ExternalGroupList
	require.NoError(t, faker.FakeObject(&cs))
	mock.EXPECT().ListExternalGroups(gomock.Any(), "testorg", gomock.Any()).Return(cs, &github.Response{}, nil)
	return client.GithubServices{Teams: mock}
}

func TestExternalGroups(t *testing.T) {
	client.GithubMockTestHelper(t, Groups(), buildExternalGroups, client.TestOptions{})
}

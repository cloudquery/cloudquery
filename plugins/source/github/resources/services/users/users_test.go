package users

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/cloudquery/plugins/source/github/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v59/github"
	"github.com/stretchr/testify/require"
)

func buildUsers(t *testing.T, ctrl *gomock.Controller) client.GithubServices {
	mock := mocks.NewMockUsersService(ctrl)

	var u github.User
	require.NoError(t, faker.FakeObject(&u))

	mock.EXPECT().Get(gomock.Any(), "").Return(
		&u, &github.Response{}, nil)

	var key github.Key
	require.NoError(t, faker.FakeObject(&key))

	mock.EXPECT().ListKeys(gomock.Any(), "", gomock.Any()).Return(
		[]*github.Key{&key}, &github.Response{}, nil)

	return client.GithubServices{Users: mock}
}

func TestTeams(t *testing.T) {
	client.GithubMockTestHelper(t, Users(), buildUsers, client.TestOptions{})
}

package organizations

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/cloudquery/plugins/source/github/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v59/github"
	"github.com/stretchr/testify/require"
)

func buildOrganizations(t *testing.T, ctrl *gomock.Controller) client.GithubServices {
	mock := mocks.NewMockOrganizationsService(ctrl)

	var cs *github.Organization
	require.NoError(t, faker.FakeObject(&cs))
	mock.EXPECT().Get(gomock.Any(), gomock.Any()).Return(cs, &github.Response{}, nil)

	var u github.User
	require.NoError(t, faker.FakeObject(&u))
	mock.EXPECT().ListMembers(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		[]*github.User{&u}, &github.Response{}, nil)

	var m github.Membership
	require.NoError(t, faker.FakeObject(&m))
	mock.EXPECT().GetOrgMembership(gomock.Any(), *u.Login, gomock.Any()).Return(
		&m, &github.Response{}, nil)

	dependabot := buildDependabot(t, ctrl)

	return client.GithubServices{
		Dependabot:    dependabot,
		Organizations: mock,
	}
}

func TestOrganizations(t *testing.T) {
	client.GithubMockTestHelper(t, Organizations(), buildOrganizations, client.TestOptions{})
}

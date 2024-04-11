package teams

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/cloudquery/plugins/source/github/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v59/github"
	"github.com/stretchr/testify/require"
)

func buildTeams(t *testing.T, ctrl *gomock.Controller) client.GithubServices {
	mock := mocks.NewMockTeamsService(ctrl)

	var cs github.Team
	require.NoError(t, faker.FakeObject(&cs))
	someId := int64(5555555)
	someURL := "https://github.com/orgs/1/teams/test"
	cs.MembersURL = &someURL
	cs.Parent = &github.Team{ID: &someId}

	mock.EXPECT().ListTeams(gomock.Any(), "testorg", gomock.Any()).Return(
		[]*github.Team{&cs}, &github.Response{}, nil)

	var u github.User
	require.NoError(t, faker.FakeObject(&u))
	mock.EXPECT().ListTeamMembersByID(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(
		[]*github.User{&u}, &github.Response{}, nil)

	var r github.Repository
	require.NoError(t, faker.FakeObject(&r))
	r.Parent = &github.Repository{ID: &someId}
	r.TemplateRepository = &github.Repository{ID: &someId}
	r.Source = &github.Repository{ID: &someId}

	mock.EXPECT().ListTeamReposByID(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(
		[]*github.Repository{&r}, &github.Response{}, nil)

	var m github.Membership
	require.NoError(t, faker.FakeObject(&m))

	mock.EXPECT().GetTeamMembershipBySlug(gomock.Any(), "testorg", *cs.Slug, *u.Login).Return(
		&m, &github.Response{}, nil)

	return client.GithubServices{Teams: mock}
}

func TestTeams(t *testing.T) {
	client.GithubMockTestHelper(t, Teams(), buildTeams, client.TestOptions{})
}

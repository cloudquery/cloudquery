package teams

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/cloudquery/plugins/source/github/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v49/github"
)

func buildTeams(t *testing.T, ctrl *gomock.Controller) client.GithubServices {
	mock := mocks.NewMockTeamsService(ctrl)

	var cs github.Team
	if err := faker.FakeObject(&cs); err != nil {
		t.Fatal(err)
	}
	someId := int64(5555555)
	someURL := "https://github.com/orgs/1/teams/test"
	cs.MembersURL = &someURL
	cs.Parent = &github.Team{ID: &someId}

	mock.EXPECT().ListTeams(gomock.Any(), "testorg", gomock.Any()).Return(
		[]*github.Team{&cs}, &github.Response{}, nil)

	var u github.User
	if err := faker.FakeObject(&u); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListTeamMembersByID(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(
		[]*github.User{&u}, &github.Response{}, nil)

	var r github.Repository
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}
	r.Parent = &github.Repository{ID: &someId}
	r.TemplateRepository = &github.Repository{ID: &someId}
	r.Source = &github.Repository{ID: &someId}

	mock.EXPECT().ListTeamReposByID(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(
		[]*github.Repository{&r}, &github.Response{}, nil)

	var m github.Membership
	if err := faker.FakeObject(&m); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().GetTeamMembershipBySlug(gomock.Any(), "testorg", *cs.Slug, *u.Login).Return(
		&m, &github.Response{}, nil)

	return client.GithubServices{Teams: mock}
}

func TestTeams(t *testing.T) {
	client.GithubMockTestHelper(t, Teams(), buildTeams, client.TestOptions{})
}

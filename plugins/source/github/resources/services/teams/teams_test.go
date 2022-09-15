package teams

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/cloudquery/plugins/source/github/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v45/github"
)

func buildTeams(t *testing.T, ctrl *gomock.Controller) client.GithubServices {
	mock := mocks.NewMockTeamsService(ctrl)

	var cs github.Team
	if err := faker.FakeDataSkipFields(&cs, []string{"Parent", "MembersURL"}); err != nil {
		t.Fatal(err)
	}
	someId := int64(5555555)
	someURL := "https://github.com/orgs/1/teams/test"
	cs.MembersURL = &someURL
	cs.Parent = &github.Team{ID: &someId}

	mock.EXPECT().ListTeams(gomock.Any(), "testorg", gomock.Any()).Return(
		[]*github.Team{&cs}, &github.Response{}, nil)

	var u github.User
	if err := faker.FakeDataSkipFields(&u, []string{"Parent"}); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListTeamMembersByID(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(
		[]*github.User{&u}, &github.Response{}, nil)

	var r github.Repository
	if err := faker.FakeDataSkipFields(&r, []string{"Parent", "Source", "TemplateRepository"}); err != nil {
		t.Fatal(err)
	}
	r.Parent = &github.Repository{ID: &someId}
	r.TemplateRepository = &github.Repository{ID: &someId}
	r.Source = &github.Repository{ID: &someId}

	mock.EXPECT().ListTeamReposByID(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(
		[]*github.Repository{&r}, &github.Response{}, nil)

	var m github.Membership
	if err := faker.FakeDataSkipFields(&m, []string{}); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().GetTeamMembershipBySlug(gomock.Any(), "testorg", *cs.Slug, *u.Login).Return(
		&m, &github.Response{}, nil)

	return client.GithubServices{Teams: mock}
}

func TestTeams(t *testing.T) {
	client.GithubMockTestHelper(t, Teams(), buildTeams, client.TestOptions{})
}

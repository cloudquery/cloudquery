package organizations

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/cloudquery/plugins/source/github/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v49/github"
)

func buildOrganizations(t *testing.T, ctrl *gomock.Controller) client.GithubServices {
	mock := mocks.NewMockOrganizationsService(ctrl)

	var cs *github.Organization
	if err := faker.FakeObject(&cs); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().Get(gomock.Any(), gomock.Any()).Return(cs, &github.Response{}, nil)

	var u github.User
	if err := faker.FakeObject(&u); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListMembers(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		[]*github.User{&u}, &github.Response{}, nil)

	var m github.Membership
	if err := faker.FakeObject(&m); err != nil {
		t.Fatal(err)
	}
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

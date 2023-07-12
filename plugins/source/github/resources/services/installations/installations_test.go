package installations

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/cloudquery/plugins/source/github/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v49/github"
)

func buildInstallations(t *testing.T, ctrl *gomock.Controller) client.GithubServices {
	mock := mocks.NewMockOrganizationsService(ctrl)

	var cs github.Installation
	if err := faker.FakeObject(&cs); err != nil {
		t.Fatal(err)
	}
	total := 1
	mock.EXPECT().ListInstallations(gomock.Any(), "testorg", gomock.Any()).Return(
		&github.OrganizationInstallations{TotalCount: &total, Installations: []*github.Installation{&cs}}, &github.Response{}, nil)

	return client.GithubServices{Organizations: mock}
}

func TestInstallations(t *testing.T) {
	client.GithubMockTestHelper(t, Installations(), buildInstallations, client.TestOptions{})
}

package resources

import (
	"testing"

	"github.com/cloudquery/cq-provider-github/client"
	"github.com/cloudquery/cq-provider-github/client/mocks"
	"github.com/google/go-github/v45/github"

	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildInstallations(t *testing.T, ctrl *gomock.Controller) client.GithubServices {
	mock := mocks.NewMockOrganizationsService(ctrl)

	var cs github.Installation
	if err := faker.FakeData(&cs); err != nil {
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

package resources

import (
	"testing"

	"github.com/cloudquery/cq-provider-github/client"
	"github.com/cloudquery/cq-provider-github/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v45/github"
)

func buildPackageBilling(t *testing.T, ctrl *gomock.Controller) client.GithubServices {
	mock := mocks.NewMockBillingService(ctrl)

	var cs *github.PackageBilling
	if err := faker.FakeData(&cs); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().GetPackagesBillingOrg(gomock.Any(), "testorg").Return(cs, &github.Response{}, nil)
	return client.GithubServices{Billing: mock}
}

func TestPackageBillings(t *testing.T) {
	client.GithubMockTestHelper(t, PackageBillings(), buildPackageBilling, client.TestOptions{})
}

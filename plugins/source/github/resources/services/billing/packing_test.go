package billing

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/cloudquery/plugins/source/github/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v49/github"
)

func buildPackage(t *testing.T, ctrl *gomock.Controller) client.GithubServices {
	mock := mocks.NewMockBillingService(ctrl)

	var cs *github.PackageBilling
	if err := faker.FakeObject(&cs); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().GetPackagesBillingOrg(gomock.Any(), "testorg").Return(cs, &github.Response{}, nil)
	return client.GithubServices{Billing: mock}
}

func TestPackageBillings(t *testing.T) {
	client.GithubMockTestHelper(t, Package(), buildPackage, client.TestOptions{})
}

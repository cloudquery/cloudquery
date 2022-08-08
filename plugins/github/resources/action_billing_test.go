package resources

import (
	"testing"

	"github.com/cloudquery/cq-provider-github/client"
	"github.com/cloudquery/cq-provider-github/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v45/github"
)

func buildActionBilling(t *testing.T, ctrl *gomock.Controller) client.GithubServices {
	mock := mocks.NewMockBillingService(ctrl)

	var cs *github.ActionBilling
	if err := faker.FakeData(&cs); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().GetActionsBillingOrg(gomock.Any(), "testorg").Return(cs, &github.Response{}, nil)
	return client.GithubServices{Billing: mock}
}

func TestActionBillings(t *testing.T) {
	client.GithubMockTestHelper(t, ActionBillings(), buildActionBilling, client.TestOptions{})
}

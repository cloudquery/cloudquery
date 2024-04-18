package billing

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/cloudquery/plugins/source/github/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v59/github"
	"github.com/stretchr/testify/require"
)

func buildAction(t *testing.T, ctrl *gomock.Controller) client.GithubServices {
	mock := mocks.NewMockBillingService(ctrl)

	var cs *github.ActionBilling
	require.NoError(t, faker.FakeObject(&cs))
	mock.EXPECT().GetActionsBillingOrg(gomock.Any(), "testorg").Return(cs, &github.Response{}, nil)
	return client.GithubServices{Billing: mock}
}

func TestActionBillings(t *testing.T) {
	client.GithubMockTestHelper(t, Action(), buildAction, client.TestOptions{})
}

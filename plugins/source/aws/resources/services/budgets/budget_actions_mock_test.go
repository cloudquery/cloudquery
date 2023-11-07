package budgets

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/budgets"
	"github.com/aws/aws-sdk-go-v2/service/budgets/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildBudgetActionsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockBudgetsClient(ctrl)

	var action types.Action
	require.NoError(t, faker.FakeObject(&action))

	m.EXPECT().DescribeBudgetActionsForAccount(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&budgets.DescribeBudgetActionsForAccountOutput{Actions: []types.Action{action}}, nil,
	)

	return client.Services{Budgets: m}
}
func TestBudgetActions(t *testing.T) {
	client.AwsMockTestHelper(t, BudgetActions(), buildBudgetActionsMock, client.TestOptions{})
}

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

func buildBudgetsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockBudgetsClient(ctrl)

	var budget types.Budget
	require.NoError(t, faker.FakeObject(&budget))

	m.EXPECT().DescribeBudgets(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&budgets.DescribeBudgetsOutput{Budgets: []types.Budget{budget}}, nil,
	)

	return client.Services{Budgets: m}
}
func TestBudgets(t *testing.T) {
	client.AwsMockTestHelper(t, Budgets(), buildBudgetsMock, client.TestOptions{})
}

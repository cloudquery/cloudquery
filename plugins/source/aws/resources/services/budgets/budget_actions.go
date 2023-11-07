package budgets

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/budgets"
	"github.com/aws/aws-sdk-go-v2/service/budgets/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func BudgetActions() *schema.Table {
	tableName := "aws_budget_actions"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_budgets_Action.html`,
		Resolver:    fetchBudgetActions,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "budgets"),
		Transform:   transformers.TransformWithStruct(&types.Action{}, transformers.WithPrimaryKeys("ActionId")),
		Columns:     schema.ColumnList{client.DefaultAccountIDColumn(true), client.DefaultRegionColumn(true)},
	}
}

func fetchBudgetActions(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceBudgets).Budgets
	input := &budgets.DescribeBudgetActionsForAccountInput{AccountId: aws.String(cl.AccountID)}

	p := budgets.NewDescribeBudgetActionsForAccountPaginator(svc, input, func(options *budgets.DescribeBudgetActionsForAccountPaginatorOptions) {
		options.StopOnDuplicateToken = true
	})

	for p.HasMorePages() {
		page, err := p.NextPage(ctx, func(options *budgets.Options) { options.Region = cl.Region })
		if err != nil {
			return err
		}
		res <- page.Actions
	}

	return nil
}

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

func Budgets() *schema.Table {
	tableName := "aws_budgets"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_budgets_Budget.html`,
		Resolver:    fetchBudgets,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "budgets"),
		Transform:   transformers.TransformWithStruct(&types.Budget{}, transformers.WithPrimaryKeys("BudgetName")),
		Columns:     schema.ColumnList{client.DefaultAccountIDColumn(true), client.DefaultRegionColumn(true)},
	}
}

func fetchBudgets(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceBudgets).Budgets
	input := &budgets.DescribeBudgetsInput{AccountId: aws.String(cl.AccountID)}

	p := budgets.NewDescribeBudgetsPaginator(svc, input, func(options *budgets.DescribeBudgetsPaginatorOptions) {
		options.StopOnDuplicateToken = true
	})

	for p.HasMorePages() {
		page, err := p.NextPage(ctx, func(options *budgets.Options) { options.Region = cl.Region })
		if err != nil {
			return err
		}
		res <- page.Budgets
	}

	return nil
}

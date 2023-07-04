package billing

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/billing/apiv1/billingpb"
	"cloud.google.com/go/billing/budgets/apiv1/budgetspb"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	budgets "cloud.google.com/go/billing/budgets/apiv1"
)

func fetchBudgets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &budgetspb.ListBudgetsRequest{
		Parent: parent.Item.(*pb.BillingAccount).Name,
	}
	gcpClient, err := budgets.NewBudgetClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListBudgets(ctx, req, c.CallOptions...)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		res <- resp
	}
	return nil
}

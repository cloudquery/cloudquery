package billing

import (
	pb "cloud.google.com/go/billing/budgets/apiv1/budgetspb"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Budgets() *schema.Table {
	return &schema.Table{
		Name:        "gcp_billing_budgets",
		Description: `https://cloud.google.com/billing/docs/reference/budget/rest/v1/billingAccounts.budgets#Budget`,
		Resolver:    fetchBudgets,
		Multiplex:   client.ProjectMultiplex,
		Transform:   client.TransformWithStruct(&pb.Budget{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:       "project_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveProject,
				PrimaryKey: true,
			},
		},
	}
}

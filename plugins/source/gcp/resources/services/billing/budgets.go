package billing

import (
	pb "cloud.google.com/go/billing/budgets/apiv1/budgetspb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Budgets() *schema.Table {
	return &schema.Table{
		Name:        "gcp_billing_budgets",
		Description: `https://cloud.google.com/billing/docs/reference/budget/rest/v1/billingAccounts.budgets#Budget`,
		Resolver:    fetchBudgets,
		Multiplex:   client.ProjectMultiplex,
		Transform:   transformers.TransformWithStruct(&pb.Budget{}, client.Options()...),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

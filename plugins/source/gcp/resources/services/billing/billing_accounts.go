package billing

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/billing/apiv1/billingpb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"

	billing "cloud.google.com/go/billing/apiv1"
)

func BillingAccounts() *schema.Table {
	return &schema.Table{
		Name:        "gcp_billing_billing_accounts",
		Description: `https://cloud.google.com/billing/docs/reference/rest/v1/billingAccounts#BillingAccount`,
		Resolver:    fetchBillingAccounts,
		Multiplex:   client.ProjectMultiplexEnabledServices("cloudbilling.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.BillingAccount{}, client.Options()...),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
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
		Relations: []*schema.Table{
			Budgets(),
		},
	}
}

func fetchBillingAccounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListBillingAccountsRequest{}
	gcpClient, err := billing.NewCloudBillingClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListBillingAccounts(ctx, req, c.CallOptions...)
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

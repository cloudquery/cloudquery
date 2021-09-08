package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-digitalocean/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/digitalocean/godo"
)

func BillingHistory() *schema.Table {
	return &schema.Table{
		Name:         "digitalocean_billing_history",
		Resolver:     fetchBillingHistory,
		DeleteFilter: client.DeleteFilter,
		Columns: []schema.Column{
			{
				Name:        "description",
				Description: "Description of the billing history entry.",
				Type:        schema.TypeString,
			},
			{
				Name:        "amount",
				Description: "Amount of the billing history entry.",
				Type:        schema.TypeString,
			},
			{
				Name:        "invoice_id",
				Description: "ID of the invoice associated with the billing history entry, if  applicable.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("InvoiceID"),
			},
			{
				Name:        "invoice_uuid",
				Description: "UUID of the invoice associated with the billing history entry, if  applicable.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("InvoiceUUID"),
			},
			{
				Name:        "date",
				Description: "Time the billing history entry occurred.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "type",
				Description: "Type of billing history entry.",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchBillingHistory(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client)
	// create options. initially, these will be blank
	opt := &godo.ListOptions{}
	for {
		billingHistory, resp, err := svc.DoClient.BillingHistory.List(ctx, opt)
		if err != nil {
			return err
		}
		// pass the current page's project to our result channel
		res <- billingHistory.BillingHistory
		// if we are at the last page, break out the for loop
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			return err
		}
		// set the page we want for the next request
		opt.Page = page + 1
	}
	return nil
}

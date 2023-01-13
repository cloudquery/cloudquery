package invoices

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func InvoiceItems() *schema.Table {
	return &schema.Table{
		Name:        "stripe_invoice_items",
		Description: `https://stripe.com/docs/api/invoiceitems`,
		Transform:   transformers.TransformWithStruct(&stripe.InvoiceItem{}, client.SharedTransformers(transformers.WithSkipFields("APIResource", "ID"), transformers.WithIgnoreInTestsTransformer(client.CreateIgnoreInTestsTransformer("Plan")))...),
		Resolver:    fetchInvoiceItems,

		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchInvoiceItems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	lp := &stripe.InvoiceItemListParams{}

	it := cl.Services.InvoiceItems.List(lp)
	for it.Next() {
		res <- it.InvoiceItem()
	}

	return it.Err()
}

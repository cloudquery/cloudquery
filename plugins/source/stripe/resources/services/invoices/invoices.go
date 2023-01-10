package invoices

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func Invoices() *schema.Table {
	return &schema.Table{
		Name:        "stripe_invoices",
		Description: `https://stripe.com/docs/api/invoices`,
		Transform:   transformers.TransformWithStruct(&stripe.Invoice{}, transformers.WithSkipFields("APIResource", "ID"), transformers.WithIgnoreInTestsTransformer(client.CreateIgnoreInTestsTransformer("DefaultSource"))),
		Resolver:    fetchInvoices,

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

func fetchInvoices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	it := cl.Services.Invoices.List(&stripe.InvoiceListParams{})
	for it.Next() {
		res <- it.Invoice()
	}
	return it.Err()
}

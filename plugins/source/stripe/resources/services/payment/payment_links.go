package payment

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func PaymentLinks() *schema.Table {
	return &schema.Table{
		Name:        "stripe_payment_links",
		Description: `https://stripe.com/docs/api/payment_links`,
		Transform:   transformers.TransformWithStruct(&stripe.PaymentLink{}, client.SharedTransformers(transformers.WithSkipFields("APIResource", "ID"))...),
		Resolver:    fetchPaymentLinks,

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

func fetchPaymentLinks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	lp := &stripe.PaymentLinkListParams{}

	it := cl.Services.PaymentLinks.List(lp)
	for it.Next() {
		res <- it.PaymentLink()
	}

	return it.Err()
}

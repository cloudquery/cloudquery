package payment

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func PaymentIntents() *schema.Table {
	return &schema.Table{
		Name:        "stripe_payment_intents",
		Description: `https://stripe.com/docs/api/payment_intents`,
		Transform:   transformers.TransformWithStruct(&stripe.PaymentIntent{}, transformers.WithSkipFields("APIResource", "ID"), transformers.WithIgnoreInTestsTransformer(client.CreateIgnoreInTestsTransformer("Source"))),
		Resolver:    fetchPaymentIntents,

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

func fetchPaymentIntents(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	it := cl.Services.PaymentIntents.List(&stripe.PaymentIntentListParams{})
	for it.Next() {
		res <- it.PaymentIntent()
	}
	return it.Err()
}

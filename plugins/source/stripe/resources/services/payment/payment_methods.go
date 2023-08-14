package payment

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/stripe/stripe-go/v74"
)

func PaymentMethods() *schema.Table {
	return &schema.Table{
		Name:        "stripe_payment_methods",
		Description: `https://stripe.com/docs/api/payment_methods`,
		Transform:   client.TransformWithStruct(&stripe.PaymentMethod{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchPaymentMethods,

		Columns: []schema.Column{
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchPaymentMethods(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	lp := &stripe.PaymentMethodListParams{}

	it := cl.Services.PaymentMethods.List(lp)
	for it.Next() {
		res <- it.PaymentMethod()
	}

	return it.Err()
}

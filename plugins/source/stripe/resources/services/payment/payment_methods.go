package payment

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func PaymentMethods() *schema.Table {
	return &schema.Table{
		Name:        "stripe_payment_methods",
		Description: `https://stripe.com/docs/api/payment_methods`,
		Transform:   transformers.TransformWithStruct(&stripe.PaymentMethod{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchPaymentMethods("payment_methods"),

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

func fetchPaymentMethods(tableName string) schema.TableResolver {
	return func(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
		cl := meta.(*client.Client)

		lp := &stripe.PaymentMethodListParams{}

		it := cl.Services.PaymentMethods.List(lp)
		for it.Next() {

			res <- it.PaymentMethod()

		}

		return it.Err()

	}
}

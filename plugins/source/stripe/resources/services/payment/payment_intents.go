package payment

import (
	"context"

	"fmt"
	"strconv"

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
		Resolver:    fetchPaymentIntents("payment_intents"),

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

func fetchPaymentIntents(tableName string) schema.TableResolver {
	return func(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
		cl := meta.(*client.Client)

		lp := &stripe.PaymentIntentListParams{}

		if cl.Backend != nil {
			value, err := cl.Backend.Get(ctx, tableName, cl.ID())
			if err != nil {
				return fmt.Errorf("failed to retrieve state from backend: %w", err)
			}
			if value != "" {
				vi, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					return fmt.Errorf("retrieved invalid state backend: %q %w", value, err)
				}
				lp.Created = &vi
			}
		}

		it := cl.Services.PaymentIntents.List(lp)
		for it.Next() {
			res <- it.PaymentIntent()
		}
		return it.Err()
	}
}

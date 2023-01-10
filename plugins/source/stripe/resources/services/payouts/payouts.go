package payouts

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func Payouts() *schema.Table {
	return &schema.Table{
		Name:        "stripe_payouts",
		Description: `https://stripe.com/docs/api/payouts`,
		Transform:   transformers.TransformWithStruct(&stripe.Payout{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchPayouts,

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

func fetchPayouts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	it := cl.Services.Payouts.List(&stripe.PayoutListParams{})
	for it.Next() {
		res <- it.Payout()
	}
	return it.Err()
}

package application_fees

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func FeeRefunds() *schema.Table {
	return &schema.Table{
		Name:        "stripe_fee_refunds",
		Description: `https://stripe.com/docs/api/fee_refunds`,
		Transform:   transformers.TransformWithStruct(&stripe.FeeRefund{}, client.SharedTransformers(transformers.WithSkipFields("APIResource", "ID"))...),
		Resolver:    fetchFeeRefunds,

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

func fetchFeeRefunds(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	p := parent.Item.(*stripe.ApplicationFee)

	lp := &stripe.FeeRefundListParams{
		ID: stripe.String(p.ID),
	}

	it := cl.Services.FeeRefunds.List(lp)
	for it.Next() {
		res <- it.FeeRefund()
	}

	return it.Err()
}

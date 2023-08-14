package application_fees

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/stripe/stripe-go/v74"
)

func FeeRefunds() *schema.Table {
	return &schema.Table{
		Name:        "stripe_fee_refunds",
		Description: `https://stripe.com/docs/api/fee_refunds`,
		Transform:   client.TransformWithStruct(&stripe.FeeRefund{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchFeeRefunds,

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

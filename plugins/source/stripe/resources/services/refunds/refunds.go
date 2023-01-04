// Code generated by codegen; DO NOT EDIT.

package refunds

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func Refunds() *schema.Table {
	return &schema.Table{
		Name:      "stripe_refunds",
		Transform: transformers.TransformWithStruct(&stripe.Refund{}, transformers.WithSkipFields([]string{"ID", "APIResource"})),
		Resolver:  fetchRefunds,
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

func fetchRefunds(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	it := cl.Services.Refunds.List(&stripe.RefundListParams{})
	for it.Next() {
		res <- it.Refund()
	}
	return it.Err()
}

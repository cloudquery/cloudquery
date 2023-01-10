package application_fees

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func ApplicationFees() *schema.Table {
	return &schema.Table{
		Name:        "stripe_application_fees",
		Description: `https://stripe.com/docs/api/application_fees`,
		Transform:   transformers.TransformWithStruct(&stripe.ApplicationFee{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchApplicationFees,

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

		Relations: []*schema.Table{
			FeeRefunds(),
		},
	}
}

func fetchApplicationFees(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	it := cl.Services.ApplicationFees.List(&stripe.ApplicationFeeListParams{})
	for it.Next() {
		res <- it.ApplicationFee()
	}
	return it.Err()
}

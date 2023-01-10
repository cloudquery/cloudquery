package issuing

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func IssuingCardholders() *schema.Table {
	return &schema.Table{
		Name:        "stripe_issuing_cardholders",
		Description: `https://stripe.com/docs/api/issuing_cardholders`,
		Transform:   transformers.TransformWithStruct(&stripe.IssuingCardholder{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchIssuingCardholders,

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

func fetchIssuingCardholders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	it := cl.Services.IssuingCardholders.List(&stripe.IssuingCardholderListParams{})
	for it.Next() {
		res <- it.IssuingCardholder()
	}
	return it.Err()
}

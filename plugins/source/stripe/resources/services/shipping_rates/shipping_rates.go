package shipping_rates

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func ShippingRates() *schema.Table {
	return &schema.Table{
		Name:        "stripe_shipping_rates",
		Description: `https://stripe.com/docs/api/shipping_rates`,
		Transform:   transformers.TransformWithStruct(&stripe.ShippingRate{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchShippingRates,

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

func fetchShippingRates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	it := cl.Services.ShippingRates.List(&stripe.ShippingRateListParams{})
	for it.Next() {
		res <- it.ShippingRate()
	}
	return it.Err()
}

package tax

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func TaxRates() *schema.Table {
	return &schema.Table{
		Name:        "stripe_tax_rates",
		Description: `https://stripe.com/docs/api/tax_rates`,
		Transform:   transformers.TransformWithStruct(&stripe.TaxRate{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchTaxRates,

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

func fetchTaxRates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	it := cl.Services.TaxRates.List(&stripe.TaxRateListParams{})
	for it.Next() {
		res <- it.TaxRate()
	}
	return it.Err()
}

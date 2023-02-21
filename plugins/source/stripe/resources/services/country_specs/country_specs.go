package country_specs

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func CountrySpecs() *schema.Table {
	return &schema.Table{
		Name:        "stripe_country_specs",
		Description: `https://stripe.com/docs/api/country_specs`,
		Transform:   transformers.TransformWithStruct(&stripe.CountrySpec{}, client.SharedTransformers(transformers.WithSkipFields("APIResource", "ID"))...),
		Resolver:    fetchCountrySpecs,

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

func fetchCountrySpecs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	lp := &stripe.CountrySpecListParams{}

	it := cl.Services.CountrySpecs.List(lp)
	for it.Next() {
		res <- it.CountrySpec()
	}

	return it.Err()
}

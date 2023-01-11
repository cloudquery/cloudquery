package apple_pay_domains

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func ApplePayDomains() *schema.Table {
	return &schema.Table{
		Name:        "stripe_apple_pay_domains",
		Description: `https://stripe.com/docs/api/apple_pay_domains`,
		Transform:   transformers.TransformWithStruct(&stripe.ApplePayDomain{}, client.SharedTransformers(transformers.WithSkipFields("APIResource", "ID"))...),
		Resolver:    fetchApplePayDomains,

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

func fetchApplePayDomains(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	lp := &stripe.ApplePayDomainListParams{}

	it := cl.Services.ApplePayDomains.List(lp)
	for it.Next() {
		res <- it.ApplePayDomain()
	}

	return it.Err()
}

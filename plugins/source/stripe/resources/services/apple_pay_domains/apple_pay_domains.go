package apple_pay_domains

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/stripe/stripe-go/v74"
)

func ApplePayDomains() *schema.Table {
	return &schema.Table{
		Name:        "stripe_apple_pay_domains",
		Description: `https://stripe.com/docs/api`,
		Transform:   client.TransformWithStruct(&stripe.ApplePayDomain{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchApplePayDomains,

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

func fetchApplePayDomains(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	lp := &stripe.ApplePayDomainListParams{}

	it := cl.Services.ApplePayDomains.List(lp)
	for it.Next() {
		res <- it.ApplePayDomain()
	}

	return it.Err()
}

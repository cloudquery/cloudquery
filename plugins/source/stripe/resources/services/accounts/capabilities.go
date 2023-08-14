package accounts

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/stripe/stripe-go/v74"
)

func Capabilities() *schema.Table {
	return &schema.Table{
		Name:        "stripe_capabilities",
		Description: `https://stripe.com/docs/api/capabilities`,
		Transform:   client.TransformWithStruct(&stripe.Capability{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchCapabilities,

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

func fetchCapabilities(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	p := parent.Item.(*stripe.Account)

	lp := &stripe.CapabilityListParams{
		Account: stripe.String(p.ID),
	}

	it := cl.Services.Capabilities.List(lp)
	for it.Next() {
		res <- it.Capability()
	}

	return it.Err()
}

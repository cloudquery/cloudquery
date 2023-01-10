package subscription

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func Subscriptions() *schema.Table {
	return &schema.Table{
		Name:        "stripe_subscriptions",
		Description: `https://stripe.com/docs/api/subscriptions`,
		Transform:   transformers.TransformWithStruct(&stripe.Subscription{}, transformers.WithSkipFields("APIResource", "ID"), transformers.WithIgnoreInTestsTransformer(client.CreateIgnoreInTestsTransformer("DefaultSource"))),
		Resolver:    fetchSubscriptions,

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

func fetchSubscriptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	it := cl.Services.Subscriptions.List(&stripe.SubscriptionListParams{})
	for it.Next() {
		res <- it.Subscription()
	}
	return it.Err()
}

package webhook_endpoints

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/stripe/stripe-go/v74"
)

func WebhookEndpoints() *schema.Table {
	return &schema.Table{
		Name:        "stripe_webhook_endpoints",
		Description: `https://stripe.com/docs/api/webhook_endpoints`,
		Transform:   client.TransformWithStruct(&stripe.WebhookEndpoint{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchWebhookEndpoints,

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

func fetchWebhookEndpoints(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	lp := &stripe.WebhookEndpointListParams{}

	it := cl.Services.WebhookEndpoints.List(lp)
	for it.Next() {
		res <- it.WebhookEndpoint()
	}

	return it.Err()
}

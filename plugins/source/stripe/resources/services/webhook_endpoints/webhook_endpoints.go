package webhook_endpoints

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func WebhookEndpoints() *schema.Table {
	return &schema.Table{
		Name:        "stripe_webhook_endpoints",
		Description: `https://stripe.com/docs/api/webhook_endpoints`,
		Transform:   transformers.TransformWithStruct(&stripe.WebhookEndpoint{}, client.SharedTransformers(transformers.WithSkipFields("APIResource", "ID"))...),
		Resolver:    fetchWebhookEndpoints,

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

func fetchWebhookEndpoints(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	lp := &stripe.WebhookEndpointListParams{}

	it := cl.Services.WebhookEndpoints.List(lp)
	for it.Next() {
		res <- it.WebhookEndpoint()
	}

	return it.Err()
}

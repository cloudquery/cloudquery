package sigma

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func SigmaScheduledQueryRuns() *schema.Table {
	return &schema.Table{
		Name:        "stripe_sigma_scheduled_query_runs",
		Description: `https://stripe.com/docs/api/sigma_scheduled_query_runs`,
		Transform:   transformers.TransformWithStruct(&stripe.SigmaScheduledQueryRun{}, client.SharedTransformers(transformers.WithSkipFields("APIResource", "ID"))...),
		Resolver:    fetchSigmaScheduledQueryRuns,

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

func fetchSigmaScheduledQueryRuns(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	lp := &stripe.SigmaScheduledQueryRunListParams{}

	it := cl.Services.SigmaScheduledQueryRuns.List(lp)
	for it.Next() {
		res <- it.SigmaScheduledQueryRun()
	}

	return it.Err()
}

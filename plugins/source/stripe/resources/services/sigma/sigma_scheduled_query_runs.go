package sigma

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/stripe/stripe-go/v74"
)

func SigmaScheduledQueryRuns() *schema.Table {
	return &schema.Table{
		Name:        "stripe_sigma_scheduled_query_runs",
		Description: `https://stripe.com/docs/api/sigma/scheduled_queries`,
		Transform:   client.TransformWithStruct(&stripe.SigmaScheduledQueryRun{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchSigmaScheduledQueryRuns,

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

func fetchSigmaScheduledQueryRuns(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	lp := &stripe.SigmaScheduledQueryRunListParams{}

	it := cl.Services.SigmaScheduledQueryRuns.List(lp)
	for it.Next() {
		res <- it.SigmaScheduledQueryRun()
	}

	return it.Err()
}

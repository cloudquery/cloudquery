package plans

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func Plans() *schema.Table {
	return &schema.Table{
		Name:        "stripe_plans",
		Description: `https://stripe.com/docs/api/plans`,
		Transform:   transformers.TransformWithStruct(&stripe.Plan{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchPlans,

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

func fetchPlans(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	it := cl.Services.Plans.List(&stripe.PlanListParams{})
	for it.Next() {
		res <- it.Plan()
	}
	return it.Err()
}

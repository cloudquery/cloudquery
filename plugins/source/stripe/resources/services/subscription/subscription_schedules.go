package subscription

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func SubscriptionSchedules() *schema.Table {
	return &schema.Table{
		Name:        "stripe_subscription_schedules",
		Description: `https://stripe.com/docs/api/subscription_schedules`,
		Transform:   transformers.TransformWithStruct(&stripe.SubscriptionSchedule{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchSubscriptionSchedules,

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

func fetchSubscriptionSchedules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	it := cl.Services.SubscriptionSchedules.List(&stripe.SubscriptionScheduleListParams{})
	for it.Next() {
		res <- it.SubscriptionSchedule()
	}
	return it.Err()
}

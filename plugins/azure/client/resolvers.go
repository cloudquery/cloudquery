package client

import (
	"context"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ResolveAzureSubscription(_ context.Context, meta schema.ClientMeta, r *schema.Resource, _ schema.Column) error {
	cl := meta.(*Client)
	return r.Set("subscription_id", cl.SubscriptionId)
}

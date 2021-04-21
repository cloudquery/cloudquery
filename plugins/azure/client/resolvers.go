package client

import (
	"context"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ResolveAzureSubscription(_ context.Context, meta schema.ClientMeta, r *schema.Resource, _ schema.Column) error {
	client := meta.(*Client)
	r.Set("subscription_id", client.SubscriptionId)
	return nil
}

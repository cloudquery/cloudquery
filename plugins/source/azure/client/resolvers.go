package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
)

func ResolveAzureSubscription(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	return r.Set(c.Name, meta.(*Client).SubscriptionId)
}

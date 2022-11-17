package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
)

func SubscriptionIDResolver(_ context.Context, meta schema.ClientMeta, r *schema.Resource, col schema.Column) error {
	cl := meta.(*Client)
	return r.Set(col.Name, cl.SubscriptionID)
}

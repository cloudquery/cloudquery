package client

import (
	"context"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ResolveAccountId(_ context.Context, meta schema.ClientMeta, r *schema.Resource, col schema.Column) error {
	client := meta.(*Client)
	return r.Set(col.Name, client.AccountId)
}

func ResolveZoneId(_ context.Context, meta schema.ClientMeta, r *schema.Resource, col schema.Column) error {
	client := meta.(*Client)
	return r.Set(col.Name, client.ZoneId)
}

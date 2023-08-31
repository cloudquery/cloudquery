package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func ResolveAccountID(_ context.Context, meta schema.ClientMeta, r *schema.Resource, col schema.Column) error {
	client := meta.(*Client)
	return r.Set(col.Name, client.AccountId)
}

func ResolveZoneID(_ context.Context, meta schema.ClientMeta, r *schema.Resource, col schema.Column) error {
	client := meta.(*Client)
	return r.Set(col.Name, client.ZoneId)
}

package client

import (
	"context"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/thoas/go-funk"
)

func ResolveAccount(_ context.Context, meta schema.ClientMeta, r *schema.Resource, _ schema.Column) error {
	client := meta.(*Client)
	return r.Set("account_id", client.AccountID)
}

func ResolveRegion(_ context.Context, meta schema.ClientMeta, r *schema.Resource, _ schema.Column) error {
	client := meta.(*Client)
	return r.Set("region", client.Region)
}

func TimestampResolver(layout, path string) schema.ColumnResolver {
	return func(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		s := funk.Get(r.Item, path, funk.WithAllowZero()).(string)
		if s == "" {
			return r.Set(c.Name, nil)
		}
		t, err := time.Parse(layout, s)
		if err != nil {
			return err
		}
		return r.Set(c.Name, t)
	}
}

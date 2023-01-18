package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/thoas/go-funk"
)

func ResolveOrg(_ context.Context, meta schema.ClientMeta, r *schema.Resource, _ schema.Column) error {
	client := meta.(*Client)
	return r.Set("org", client.Org)
}

func ResolveParentColumn(field string) schema.ColumnResolver {
	return func(_ context.Context, _ schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		return r.Set(c.Name, funk.Get(r.Parent.Item, field))
	}
}

func ResolveGrandParentColumn(field string) schema.ColumnResolver {
	return func(_ context.Context, _ schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		return r.Set(c.Name, funk.Get(r.Parent.Parent.Item, field))
	}
}

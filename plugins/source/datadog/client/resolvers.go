package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/thoas/go-funk"
)

func ResolveParentColumn(field string) schema.ColumnResolver {
	return func(_ context.Context, _ schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		return r.Set(c.Name, funk.Get(r.Parent.Item, field))
	}
}

func ResolveAccountName(_ context.Context, meta schema.ClientMeta, r *schema.Resource, col schema.Column) error {
	client := meta.(*Client)
	return r.Set(col.Name, client.multiplexedAccount.Name)
}

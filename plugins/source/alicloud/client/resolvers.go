package client

import (
	"context"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/thoas/go-funk"
	"time"
)

func ResolveParentColumn(field string) schema.ColumnResolver {
	return func(_ context.Context, _ schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		return r.Set(c.Name, funk.Get(r.Parent.Item, field))
	}
}

func ResolveUpdateDate() schema.ColumnResolver {
	return func(_ context.Context, _ schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		return r.Set(c.Name, time.Now().AddDate(0, 0, -1).Format("2006-01-02"))
	}
}

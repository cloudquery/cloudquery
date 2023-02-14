package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
)

func ResolveOrg(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	return r.Set(c.Name, meta.(*Client).Org)
}

func ResolveRepositoryID(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	return r.Set(c.Name, meta.(*Client).Repository.ID)
}

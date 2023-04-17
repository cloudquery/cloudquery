package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func ResolveTailnet(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	return resource.Set(c.Name, meta.(*Client).pluginSpec.Tailnet)
}

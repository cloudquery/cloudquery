package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
)

func ResolveSiteID(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	return resource.Set("site_id", meta.(*Client).PluginSpec.SiteId)
}

package stats

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/fastly/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchStatsRegions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	f := func() error {
		r, err := c.Fastly.GetRegions()
		if err != nil {
			return err
		}
		res <- r.Data
		return nil
	}
	return c.RetryOnError(ctx, "fastly_stats_regions", f)
}

func setRegionName(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	return resource.Set(c.Name, resource.Item.(string))
}

package analytics

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/homebrew/client"
	"github.com/cloudquery/cloudquery/plugins/source/homebrew/internal/homebrew"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Installs365Days() *schema.Table {
	return &schema.Table{
		Name:          "homebrew_analytics_installs_365d",
		Description:   `https://formulae.brew.sh/analytics/install/365d/`,
		Resolver:      fetchInstalls365Days,
		IsIncremental: true,
		Transform: transformers.TransformWithStruct(
			&homebrew.InstallItem{},
			transformers.WithPrimaryKeys("Formula"),
		),
	}
}

func fetchInstalls365Days(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	installs, err := c.Homebrew.GetInstalls(ctx, homebrew.Days365)
	if err != nil {
		return err
	}
	res <- installs.Items
	return nil
}

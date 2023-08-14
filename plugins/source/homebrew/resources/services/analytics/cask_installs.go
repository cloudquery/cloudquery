package analytics

import (
	"context"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/source/homebrew/client"
	"github.com/cloudquery/cloudquery/plugins/source/homebrew/internal/homebrew"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func CaskInstalls(days homebrew.Days) *schema.Table {
	return &schema.Table{
		Name:        fmt.Sprintf("homebrew_analytics_cask_installs_%s", days),
		Title:       fmt.Sprintf("Homebrew Analytics Cask Installs (%d days)", days.Number()),
		Description: fmt.Sprintf(`https://formulae.brew.sh/analytics/cask-install/%s/`, days),
		Resolver:    fetchCaskInstalls(days),
		Transform: transformers.TransformWithStruct(
			&homebrew.CaskInstallItem{},
			transformers.WithPrimaryKeys("Number"),
		),
	}
}

func fetchCaskInstalls(days homebrew.Days) func(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	return func(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
		c := meta.(*client.Client)
		installs, err := c.Homebrew.GetCaskInstalls(ctx, days)
		if err != nil {
			return err
		}
		res <- installs.Items
		return nil
	}
}

package analytics

import (
	"context"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/source/homebrew/client"
	"github.com/cloudquery/cloudquery/plugins/source/homebrew/internal/homebrew"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Installs(days homebrew.Days) *schema.Table {
	return &schema.Table{
		Name:        fmt.Sprintf("homebrew_analytics_installs_%s", days),
		Title:       fmt.Sprintf("Homebrew Analytics Installs (%d days)", days.Number()),
		Description: fmt.Sprintf(`https://formulae.brew.sh/analytics/install/%s/`, days),
		Resolver:    fetchInstalls(days),
		Transform: transformers.TransformWithStruct(
			&homebrew.InstallItem{},
			transformers.WithPrimaryKeys("Number"),
		),
	}
}

func fetchInstalls(days homebrew.Days) func(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	return func(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
		c := meta.(*client.Client)
		installs, err := c.Homebrew.GetInstalls(ctx, days)
		if err != nil {
			return err
		}
		res <- installs.Items
		return nil
	}
}

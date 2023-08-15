package analytics

import (
	"context"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/source/homebrew/client"
	"github.com/cloudquery/cloudquery/plugins/source/homebrew/internal/homebrew"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func BuildErrors(days homebrew.Days) *schema.Table {
	return &schema.Table{
		Name:        fmt.Sprintf("homebrew_analytics_build_errors_%s", days),
		Title:       fmt.Sprintf("Homebrew Analytics Build Errors (%d days)", days.Number()),
		Description: fmt.Sprintf(`https://formulae.brew.sh/analytics/build-error/%s/`, days),
		Resolver:    fetchBuildErrors(days),
		Transform: transformers.TransformWithStruct(
			&homebrew.BuildErrorItem{},
			transformers.WithPrimaryKeys("Number"),
		),
	}
}

func fetchBuildErrors(days homebrew.Days) func(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	return func(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
		c := meta.(*client.Client)
		errors, err := c.Homebrew.GetBuildErrors(ctx, days)
		if err != nil {
			return err
		}
		res <- errors.Items
		return nil
	}
}

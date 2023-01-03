package repositories

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/go-github/v48/github"
)

func fetchAlerts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	repo := parent.Item.(*github.Repository)

	alerts, _, err := c.Github.Dependabot.ListRepoAlerts(ctx, c.Org, *repo.Name, nil)
	if err != nil {
		return err
	}

	res <- alerts

	return nil
}

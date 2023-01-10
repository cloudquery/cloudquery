package repositories

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/go-github/v48/github"
)

func fetchSecrets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	repo := parent.Item.(*github.Repository)

	secrets, _, err := c.Github.Dependabot.ListRepoSecrets(ctx, c.Org, *repo.Name, nil)
	if err != nil {
		return err
	}

	res <- secrets.Secrets

	return nil
}

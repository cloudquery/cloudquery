package repositories

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/google/go-github/v59/github"
)

func secrets() *schema.Table {
	return &schema.Table{
		Name:      "github_repository_dependabot_secrets",
		Resolver:  fetchSecrets,
		Transform: client.TransformWithStruct(&github.Secret{}, transformers.WithPrimaryKeys("Name")),
		Columns:   []schema.Column{client.OrgColumn, client.RepositoryIDColumn},
	}
}

func fetchSecrets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	repo := parent.Item.(*github.Repository)

	opts := &github.ListOptions{PerPage: 100}
	for {
		secrets, resp, err := c.Github.Dependabot.ListRepoSecrets(ctx, c.Org, *repo.Name, opts)
		if err != nil {
			return err
		}
		res <- secrets.Secrets

		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	return nil
}

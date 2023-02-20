package repositories

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/google/go-github/v48/github"
)

func secrets() *schema.Table {
	return &schema.Table{
		Name:     "github_repository_dependabot_secrets",
		Resolver: fetchSecrets,
		Transform: transformers.TransformWithStruct(&github.Secret{},
			append(client.SharedTransformers(), transformers.WithPrimaryKeys("Name"))...),
		Columns: []schema.Column{client.OrgColumn, client.RepositoryIDColumn},
	}
}

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

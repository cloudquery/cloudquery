package repositories

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/google/go-github/v49/github"
)

func keys() *schema.Table {
	return &schema.Table{
		Name:      "github_repository_keys",
		Resolver:  fetchKeys,
		Transform: client.TransformWithStruct(&github.Key{}, transformers.WithPrimaryKeys("ID")),
		Columns: []schema.Column{
			client.OrgColumn,
			client.RepositoryIDColumn,
		},
	}
}

func fetchKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	repo := parent.Item.(*github.Repository)
	opts := &github.ListOptions{PerPage: 100}
	for {
		keysList, resp, err := c.Github.Repositories.ListKeys(ctx, c.Org, *repo.Name, opts)
		if err != nil {
			return err
		}
		res <- keysList
		opts.Page = resp.NextPage
		if opts.Page == resp.LastPage {
			break
		}
	}

	return nil
}

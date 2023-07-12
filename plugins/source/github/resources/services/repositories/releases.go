package repositories

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/google/go-github/v49/github"
)

func releases() *schema.Table {
	return &schema.Table{
		Name:      "github_releases",
		Resolver:  fetchReleases,
		Transform: client.TransformWithStruct(&github.RepositoryRelease{}, transformers.WithPrimaryKeys("ID")),
		Columns:   []schema.Column{client.OrgColumn, client.RepositoryIDColumn},
		Relations: []*schema.Table{assets()},
	}
}

func fetchReleases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	repo := parent.Item.(*github.Repository)
	opts := &github.ListOptions{PerPage: 1000}
	for {
		releases, resp, err := c.Github.Repositories.ListReleases(ctx, c.Org, *repo.Name, opts)
		if err != nil {
			return err
		}
		res <- releases

		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}
	return nil
}

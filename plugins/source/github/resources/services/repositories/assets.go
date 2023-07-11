package repositories

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/google/go-github/v49/github"
)

func assets() *schema.Table {
	return &schema.Table{
		Name:      "github_release_assets",
		Resolver:  fetchAssets,
		Transform: client.TransformWithStruct(&github.ReleaseAsset{}, transformers.WithPrimaryKeys("ID")),
		Columns:   []schema.Column{client.OrgColumn, client.RepositoryIDColumn},
	}
}

func fetchAssets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	release := parent.Item.(*github.RepositoryRelease)
	repo := parent.Parent.Item.(*github.Repository)
	opts := &github.ListOptions{PerPage: 100}
	for {
		releases, resp, err := c.Github.Repositories.ListReleaseAssets(ctx, c.Org, *repo.Name, *release.ID, opts)
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

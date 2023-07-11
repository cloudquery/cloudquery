package hooks

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/google/go-github/v49/github"
)

func Hooks() *schema.Table {
	return &schema.Table{
		Name:      "github_hooks",
		Resolver:  fetchHooks,
		Multiplex: client.OrgMultiplex,
		Transform: client.TransformWithStruct(&github.Hook{}, transformers.WithPrimaryKeys("ID")),
		Columns:   []schema.Column{client.OrgColumn},
		Relations: []*schema.Table{deliveries()},
	}
}

func fetchHooks(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	opts := &github.ListOptions{PerPage: 100}
	for {
		hooks, resp, err := c.Github.Organizations.ListHooks(ctx, c.Org, opts)
		if err != nil {
			return err
		}
		res <- hooks

		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}
	return nil
}

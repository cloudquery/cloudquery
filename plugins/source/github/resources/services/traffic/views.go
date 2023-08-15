package traffic

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/google/go-github/v49/github"
)

func Views() *schema.Table {
	return &schema.Table{
		Name:        "github_traffic_views",
		Description: "https://docs.github.com/en/rest/metrics/traffic?apiVersion=2022-11-28#get-page-views",
		Resolver:    fetchViews,
		Multiplex:   client.OrgRepositoryMultiplex,
		Transform:   client.TransformWithStruct(&github.TrafficViews{}),
		Columns:     []schema.Column{client.OrgColumn, client.RepositoryIDColumn},
	}
}

func fetchViews(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	repo := c.Repository
	opts := &github.TrafficBreakdownOptions{}
	views, _, err := c.Github.Repositories.ListTrafficViews(ctx, *repo.Owner.Login, *c.Repository.Name, opts)
	if err != nil {
		return err
	}
	res <- views
	return nil
}

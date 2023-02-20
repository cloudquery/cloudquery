package traffic

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/google/go-github/v48/github"
)

func Paths() *schema.Table {
	return &schema.Table{
		Name:        "github_traffic_paths",
		Description: "https://docs.github.com/en/rest/metrics/traffic?apiVersion=2022-11-28#get-top-referral-paths",
		Resolver:    fetchPaths,
		Multiplex:   client.OrgRepositoryMultiplex,
		Transform:   transformers.TransformWithStruct(&github.TrafficPath{}, append(client.SharedTransformers(), transformers.WithPrimaryKeys("Path"))...),
		Columns:     []schema.Column{client.OrgColumn, client.RepositoryIDColumn},
	}
}

func fetchPaths(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	repo := c.Repository
	paths, _, err := c.Github.Repositories.ListTrafficPaths(ctx, *repo.Owner.Login, *c.Repository.Name)
	if err != nil {
		return err
	}
	res <- paths
	return nil
}

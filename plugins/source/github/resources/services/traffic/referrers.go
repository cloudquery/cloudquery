package traffic

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/google/go-github/v49/github"
)

func Referrers() *schema.Table {
	return &schema.Table{
		Name:        "github_traffic_referrers",
		Description: "https://docs.github.com/en/rest/metrics/traffic?apiVersion=2022-11-28#get-top-referral-sources",
		Resolver:    fetchReferrers,
		Multiplex:   client.OrgRepositoryMultiplex,
		Transform:   client.TransformWithStruct(&github.TrafficReferrer{}, transformers.WithPrimaryKeys("Referrer")),
		Columns:     []schema.Column{client.OrgColumn, client.RepositoryIDColumn},
	}
}

func fetchReferrers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	repo := c.Repository
	referrers, _, err := c.Github.Repositories.ListTrafficReferrers(ctx, *repo.Owner.Login, *c.Repository.Name)
	if err != nil {
		return err
	}
	res <- referrers
	return nil
}

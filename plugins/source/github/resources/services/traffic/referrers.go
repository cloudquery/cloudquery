package traffic

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/google/go-github/v48/github"
)

func Referrers() *schema.Table {
	return &schema.Table{
		Name:        "github_traffic_referrers",
		Description: "https://docs.github.com/en/rest/metrics/traffic?apiVersion=2022-11-28#get-top-referral-sources",
		Resolver:    fetchReferrers,
		Multiplex:   client.OrgRepositoryMultiplex,
		Transform:   transformers.TransformWithStruct(&github.TrafficReferrer{}, append(client.SharedTransformers(), transformers.WithPrimaryKeys("Referrer"))...),
		Columns: []schema.Column{
			{
				Name:        "org",
				Type:        schema.TypeString,
				Resolver:    client.ResolveOrg,
				Description: `The Github Organization of the resource.`,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "repository_id",
				Type:     schema.TypeInt,
				Resolver: client.ResolveRepositoryId,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
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

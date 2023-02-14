package traffic

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/google/go-github/v48/github"
)

func Clones() *schema.Table {
	return &schema.Table{
		Name:        "github_traffic_clones",
		Description: "https://docs.github.com/en/rest/metrics/traffic?apiVersion=2022-11-28#get-repository-clones",
		Resolver:    fetchClones,
		Multiplex:   client.OrgRepositoryMultiplex,
		Transform:   transformers.TransformWithStruct(&github.TrafficClones{}, client.SharedTransformers()...),
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

func fetchClones(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	repo := c.Repository
	opts := &github.TrafficBreakdownOptions{}
	clones, _, err := c.Github.Repositories.ListTrafficClones(ctx, *repo.Owner.Login, *c.Repository.Name, opts)
	if err != nil {
		return err
	}
	res <- clones
	return nil
}

package organizations

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/google/go-github/v48/github"
)

func Organizations() *schema.Table {
	return &schema.Table{
		Name:      "github_organizations",
		Resolver:  fetchOrganizations,
		Multiplex: client.OrgMultiplex,
		Transform: transformers.TransformWithStruct(&github.Organization{},
			append(client.SharedTransformers(), transformers.WithPrimaryKeys("ID"))...),
		Columns:   []schema.Column{client.OrgColumn},
		Relations: []*schema.Table{alerts(), members(), secrets()},
	}
}

func fetchOrganizations(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	org, _, err := c.Github.Organizations.Get(ctx, c.Org)
	if err != nil {
		return err
	}
	res <- org
	return nil
}

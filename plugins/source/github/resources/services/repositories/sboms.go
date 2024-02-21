package repositories

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/google/go-github/v59/github"
)

func sboms() *schema.Table {
	return &schema.Table{
		Name:      "github_repository_sboms",
		Resolver:  fetchSboms,
		Multiplex: client.OrgRepositoryMultiplex,
		Transform: client.TransformWithStruct(&github.SBOMInfo{}),
		Columns: []schema.Column{
			client.OrgColumn,
			client.RepositoryIDColumn,
		},
	}
}

func fetchSboms(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	repo := parent.Item.(*github.Repository)

	sbom, _, err := c.Github.DependencyGraph.GetSBOM(ctx, c.Org, *repo.Name)
	if err != nil {
		return err
	}
	res <- sbom.SBOM

	return nil
}

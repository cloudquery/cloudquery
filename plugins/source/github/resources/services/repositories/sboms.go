package repositories

import (
	"context"
	"fmt"
	"net/http"

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

	sbom, rawResp, err := c.Github.DependencyGraph.GetSBOM(ctx, c.Org, *repo.Name)
	if err != nil {
		if rawResp != nil && rawResp.StatusCode == http.StatusNotFound {
			return fmt.Errorf("sbom not found for repository %s/%s. You might need to enable dependency graph under the repository insights tab", c.Org, *repo.Name)
		}
		return err
	}
	res <- sbom.SBOM

	return nil
}

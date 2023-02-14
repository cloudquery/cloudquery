package repositories

import (
	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/google/go-github/v48/github"
)

func releases() *schema.Table {
	return &schema.Table{
		Name:      "github_releases",
		Resolver:  fetchReleases,
		Multiplex: client.OrgMultiplex,
		Transform: transformers.TransformWithStruct(&github.RepositoryRelease{},
			append(client.SharedTransformers(), transformers.WithPrimaryKeys("ID"))...),
		Columns:   []schema.Column{client.OrgColumn, repoIDColumn},
		Relations: []*schema.Table{assets()},
	}
}

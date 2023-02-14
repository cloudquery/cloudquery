package teams

import (
	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/google/go-github/v48/github"
)

func repositories() *schema.Table {
	return &schema.Table{
		Name:     "github_team_repositories",
		Resolver: fetchRepositories,
		Transform: transformers.TransformWithStruct(&github.Repository{},
			append(client.SharedTransformers(), transformers.WithPrimaryKeys("ID"))...),
		Columns: []schema.Column{client.OrgColumn, teamIDColumn},
	}
}

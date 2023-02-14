package repositories

import (
	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/google/go-github/v48/github"
)

func Repositories() *schema.Table {
	return &schema.Table{
		Name:      "github_repositories",
		Resolver:  fetchRepositories,
		Multiplex: client.OrgRepositoryMultiplex,
		Transform: transformers.TransformWithStruct(&github.Repository{},
			append(client.SharedTransformers(), transformers.WithPrimaryKeys("ID"))...),
		Columns:   []schema.Column{client.OrgColumn},
		Relations: []*schema.Table{alerts(), releases(), secrets()},
	}
}

var repoIDColumn = schema.Column{
	Name:            "repository_id",
	Type:            schema.TypeInt,
	Resolver:        client.ResolveParentColumn("ID"),
	CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
}

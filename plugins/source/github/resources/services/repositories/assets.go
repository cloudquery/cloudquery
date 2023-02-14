package repositories

import (
	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/google/go-github/v48/github"
)

func assets() *schema.Table {
	return &schema.Table{
		Name:      "github_release_assets",
		Resolver:  fetchAssets,
		Multiplex: client.OrgMultiplex,
		Transform: transformers.TransformWithStruct(&github.ReleaseAsset{},
			append(client.SharedTransformers(), transformers.WithPrimaryKeys("ID"))...),
		Columns: []schema.Column{client.OrgColumn,
			{
				Name:            "repository_id",
				Type:            schema.TypeInt,
				Resolver:        client.ResolveGrandParentColumn("ID"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
	}
}
